package main

import (
	"fmt"
	//"strconv"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/cadvisor/info"
	"github.com/kevin1024/cadvisor/client"
	//"github.com/cactus/go-statsd-client/statsd"
	"time"
)

func ContainerStatsToStatsDStrings(cs *info.ContainerStats) []string {
	fmt.Printf("---> %v\n", cs)
	return []string{
		//cs.Cpu.Usage.Total,
		"whatever",
		"stats",
	}
}

func main() {
	for {
		client, _ := client.NewClient("http://192.168.59.103:8080/")
		//mInfo, _ := client.MachineInfo()
		request := info.ContainerInfoRequest{2}
		sInfo, _ := client.ContainerInfo("/docker", &request)
		for _, containerStats := range sInfo.Stats {
		        spew.Dump(containerStats)
			spew.Dump(ContainerStatsToStatsDStrings(containerStats))
		}
		time.Sleep(5 * time.Second)
	}
}
