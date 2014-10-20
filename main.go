package main

import (
	"fmt"
	"github.com/kevin1024/cadvisor_statsd/cadvisor"
	"time"
)


func main() {
	for {
		output, _ := cadvisor.GetAllSubcontainers()
		fmt.Printf("%v", output)
		time.Sleep(5*time.Second)
	}
}
