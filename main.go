package main

import (
	"fmt"
	"github.com/alperhankendi/devnot-workshop/cmd"
)

var (
	BuildVersion string = ""
)

func main() {
	fmt.Printf("Application is starting. Version:%s", BuildVersion)
	cmd.Execute()
}
