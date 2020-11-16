package controller

import (
	"context"
	"fmt"
	"github.com/alperhankendi/devnot-workshop/internal/movies"
	"github.com/alperhankendi/devnot-workshop/pkg/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	resource struct {
		service *movies.Service

		//Ver3
		queue chan *movies.Movie
	}
)

func NewController(srv *movies.Service) *resource {

	r := &resource{
		service: srv,
		//Ver3
		queue: make(chan *movies.Movie, 10000),
	}
	//Ver3
	go r.startProcessor()

	//Ver4
	JobQueue = make(chan Job, MaxQueue)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run(r)
	return r
}

func (receiver *resource) CreateV1(c echo.Context) error {

	item := new(movies.Movie)
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse request. Error:%s", err.Error()))
	}
	err := receiver.service.Create(c.Request().Context(), item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("failed to create item. Error:%s", err.Error()))
	}

	return c.JSON(http.StatusCreated, "")
}
func (receiver *resource) CreateV2(c echo.Context) error {

	item := new(movies.Movie)
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse request. Error:%s", err.Error()))
	}
	go receiver.service.Create(context.Background(), item)

	return c.JSON(http.StatusCreated, "")
}
func (receiver *resource) CreateV3(c echo.Context) error {

	item := new(movies.Movie)
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse request. Error:%s", err.Error()))
	}
	receiver.queue <- item
	return c.JSON(http.StatusCreated, "")
}
func (receiver *resource) CreateV4(c echo.Context) error {

	item := new(movies.Movie)
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("failed to parse request. Error:%s", err.Error()))
	}

	work := Job{Payload: item}
	JobQueue <- work

	return c.JSON(http.StatusCreated, "")
}
func (receiver *resource) startProcessor() {

	for {
		select {
		case item := <-receiver.queue:
			receiver.service.Create(context.Background(), item)
			//fmt.Printf("processed.\n")
		}
	}
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (receiver *resource) GetV1(c echo.Context) error {

	id := c.Param("id")
	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, "id can not be null or empty")
	}
	item, err := receiver.service.Get(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, _ := json.Marshal(&item)
	return c.String(http.StatusOK, string(result))
}

var (
	MaxWorker = 3 * 2
	MaxQueue  = 20
)

type Job struct {
	Payload *movies.Movie
	r       *resource
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}
func (w Worker) Start(r *resource) {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := r.service.Create(context.Background(), job.Payload); err != nil {
					log.Logger.Printf("failed to create item: %s", err.Error())
				}
			case <-w.quit:
				return
			}
		}
	}()
}
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
