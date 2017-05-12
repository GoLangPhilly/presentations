package service

import (
	"github.com/docker/docker/api/types"
	"os"
)


type config struct {
	serverHostPost string
}

func getConfig() *config {

	envConfig := config{}

	envConfig.serverHostPost = os.Getenv("GRPC_HOST_PORT")
	if envConfig.serverHostPost == "" {
		envConfig.serverHostPost = "localhost:9090"
	}

	return &envConfig
}



func convert(containerStats *types.StatsJSON) *ContainerStats {

	var (
		previousCPU            uint64
		previousSystem         uint64
		memPercent, cpuPercent float64
	)

	// MemoryStats.Limit will never be 0 unless the container is not running and we haven't
	// got any data from cgroup
	memLimit := float64(containerStats.MemoryStats.Limit)
	if memLimit != 0 {
		memPercent = float64(containerStats.MemoryStats.Usage) / float64(memLimit) * 100.0
	}
	previousCPU = containerStats.PreCPUStats.CPUUsage.TotalUsage
	previousSystem = containerStats.PreCPUStats.SystemUsage
	cpuPercent = calculateCPUPercent(previousCPU, previousSystem, containerStats)

	cs := &ContainerStats{
		Container:  containerStats.Name,
		CpuPercentage: cpuPercent,
		MemoryPercentage: memPercent,
		MemoryLimit: convertMemoryLimit(memLimit),
		MemorySizeType: convertMemorySize(memLimit),


	}

	return cs


}

func calculateCPUPercent(previousCPU, previousSystem uint64, v *types.StatsJSON) float64 {
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

type MemSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB MemSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func convertMemoryLimit(ml float64) float64 {
	memoryLimit := MemSize(ml)

	retVal := MemSize(memoryLimit)

	switch {
	case memoryLimit >= YB:
		retVal = memoryLimit/YB
	case memoryLimit >= ZB:
		retVal = memoryLimit/ZB
	case memoryLimit >= EB:
		retVal = memoryLimit/EB
	case memoryLimit >= PB:
		retVal = memoryLimit/PB
	case memoryLimit >= TB:
		retVal = memoryLimit/TB
	case memoryLimit >= GB:
		retVal = memoryLimit/GB
	case memoryLimit >= MB:
		retVal = memoryLimit/MB
	case memoryLimit >= KB:
		retVal = memoryLimit/KB
	}
	return float64(retVal)
}

func convertMemorySize(ml float64) ContainerStats_MemorySize {

	memoryLimit := MemSize(ml)
	switch {
	case memoryLimit >= YB:
		return ContainerStats_YB
	case memoryLimit >= ZB:
		return ContainerStats_ZB
	case memoryLimit >= EB:
		return ContainerStats_EB
	case memoryLimit >= PB:
		return ContainerStats_PB
	case memoryLimit >= TB:
		return ContainerStats_TB
	case memoryLimit >= GB:
		return ContainerStats_GB
	case memoryLimit >= MB:
		return ContainerStats_MB
	case memoryLimit >= KB:
		return ContainerStats_KB
	}
	return ContainerStats_B

}


