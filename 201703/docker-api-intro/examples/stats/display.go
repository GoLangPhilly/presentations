package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"time"
)

func displayContainerStats(containerStats *types.StatsJSON) {

	var (
		previousCPU            uint64
		previousSystem         uint64
		memPercent, cpuPercent float64
	)

	// MemoryStats.Limit will never be 0 unless the container is not running and we haven't
	// got any data from cgroup
	if containerStats.MemoryStats.Limit != 0 {
		memPercent = float64(containerStats.MemoryStats.Usage) / float64(containerStats.MemoryStats.Limit) * 100.0
	}
	previousCPU = containerStats.PreCPUStats.CPUUsage.TotalUsage
	previousSystem = containerStats.PreCPUStats.SystemUsage
	cpuPercent = calculateCPUPercentUnix(previousCPU, previousSystem, containerStats)

	fmt.Print("\u001B[2J\u001B[0;0f")
	fmt.Printf("%s%12s%12s\n", "CONTAINER", "CPU %", "MEM %")

	fmt.Print(containerStats.Name, " ")
	if cpuPercent > 30.0 {
		fmt.Print("\x1b[31;1m")
	} else {
		fmt.Print("\x1b[32;1m")
	}
	fmt.Printf("% 16.2f%c", cpuPercent, '%')
	fmt.Print("\x1b[0m")
	fmt.Printf("%11.2f%c", memPercent, '%')
	time.Sleep(1 * time.Second)
}

func calculateCPUPercentUnix(previousCPU, previousSystem uint64, v *types.StatsJSON) float64 {
	var (
		cpuPercent = 0.0
		// calculate the change for the cpu usage of the container in between readings
		cpuDelta = float64(v.CPUStats.CPUUsage.TotalUsage) - float64(previousCPU)
		// calculate the change for the entire system between readings
		systemDelta = float64(v.CPUStats.SystemUsage) - float64(previousSystem)
	)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(len(v.CPUStats.CPUUsage.PercpuUsage)) * 100.0
	}
	return cpuPercent * 1
}
