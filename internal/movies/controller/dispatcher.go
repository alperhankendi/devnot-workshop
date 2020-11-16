/*

We have decided to utilize a common pattern when using Go channels, in order to create a 2-tier channel system,
one for queuing jobs and another to control how many workers operate on the JobQueue concurrently.
The idea was to parallelize the creating movie to mongo to a somewhat sustainable rate,
one that would not cripple the machine nor start generating connections errors from mongo.
So we have opted for creating a Job/Worker pattern.

For those that are familiar with Java, C#, etc, think about this as the Golang way of implementing
a Worker Thread-Pool utilizing channels instead.
*/

package controller

import (
	"fmt"
	"runtime"
)

/*
  örnek..
  1.WORKERPool (1xgr) ==> { 20x worker } (20xgr)
  2.WORKERPool (1xgr) ==> { 20x worker } (20xgr)
  ..
  ..
  6.WORKERPool (1xgr) ==> { 20x worker } (20xgr)
*/

type Dispatcher struct {
	maxWorkers int
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run(r *resource) {
	fmt.Printf("GOMAXPROCS is %d\n", runtime.GOMAXPROCS(0))
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start(r)
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	fmt.Println("Worker que dispatcher started...")
	for {

		select {
		case job := <-JobQueue:

			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}
