package main

import (
	"fmt"
	"strings"
	//"strconv"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/cadvisor/info"
	"github.com/google/cadvisor/client"
	//"github.com/cactus/go-statsd-client/statsd"
	"time"
)

func StatsDCounter(containerName string, statName string, statValue uint64) string {
	containerName = strings.TrimPrefix(strings.Replace(containerName, "/", ".", -1), ".")
	return fmt.Sprintf("%v.%v: %v|g)", containerName, statName, statValue)
}

func ContainerStatsToStatsDStrings(name string, startStats *info.ContainerStats, endStats *info.ContainerStats, intervalNs uint64) []string {
	return []string{
		StatsDCounter(name, "network.rxbytes", endStats.Network.RxBytes - startStats.Network.RxBytes),
		StatsDCounter(name, "network.rxpackets", endStats.Network.RxPackets - startStats.Network.RxPackets),
		StatsDCounter(name, "network.rxerrors", endStats.Network.RxErrors - startStats.Network.RxErrors),
		StatsDCounter(name, "network.rxdropped", endStats.Network.RxDropped - startStats.Network.RxDropped),
		StatsDCounter(name, "network.txbytes", endStats.Network.TxBytes - startStats.Network.TxBytes),
		StatsDCounter(name, "network.txpackets", endStats.Network.TxPackets - startStats.Network.TxPackets),
		StatsDCounter(name, "network.txerrors", endStats.Network.TxErrors - startStats.Network.TxErrors),
		StatsDCounter(name, "network.txdropped", endStats.Network.TxDropped - startStats.Network.TxDropped),
		StatsDCounter(name, "cpu.total", (((endStats.Cpu.Usage.Total - startStats.Cpu.Usage.Total) * 100) / intervalNs)),
	}
}

func main() {
	for {
		client, _ := client.NewClient("http://192.168.59.103:8080/")
		//mInfo, _ := client.MachineInfo()
		request := info.ContainerInfoRequest{2}
		sInfo, _ := client.ContainerInfo("/docker/1718df726f98e33de03524667dafaf414e9e3d1afcc17f1c2cb20068d48d721c", &request)
		first := sInfo.Stats[0]
		last := sInfo.Stats[len(sInfo.Stats)-1]

		intervalNs := uint64((last.Timestamp.Sub(first.Timestamp)/time.Nanosecond))
		fmt.Printf("time elpased: %v\n", intervalNs)
	        spew.Dump(ContainerStatsToStatsDStrings(sInfo.ContainerReference.Name, first, last, intervalNs))


		//spew.Dump(sInfo.Stats)
		//for _, containerStats := range sInfo.Stats {
		        //spew.Dump(containerStats)
			//spew.Dump(ContainerStatsToStatsDStrings(sInfo.ContainerReference.Name, containerStats))
		//}
		time.Sleep(10 * time.Second)
	}
}
