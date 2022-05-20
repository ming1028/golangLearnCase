package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net"
	"time"
)

func main() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("cpu info err:%#v\n", err)
		return
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	fmt.Println("======cpu负载======")
	info, err := load.Avg()
	fmt.Printf("cpu负载：%v\n", info)
	fmt.Println("======cpu负载======")
	for {
		percent, _ := cpu.Percent(time.Second, true)
		fmt.Printf("cpu pencent:%v\n", percent)
		break
	}
	fmt.Println("======mem info======")
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info%v\n", memInfo)
	fmt.Println("======mem info======")

	fmt.Println("======host info======")
	hostInfo, _ := host.Info()
	fmt.Printf("host info%v uptime:%v boottime:%v\n", hostInfo, hostInfo.Uptime, hostInfo.BootTime)
	fmt.Println("======host info======")

	fmt.Println("======net======")
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		fmt.Println(ipAddr.IP.String())
	}
	fmt.Println("======net======")
}
