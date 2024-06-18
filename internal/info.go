package internal

import (
	"metricsExtractor/services"
)

type SystemInfoProvider interface {
	CollectSystemInfo() (*services.SystemInfo, error)
}

type Info struct {
	systemService *services.SystemService
}

func NewInfo(systemService *services.SystemService) *Info {
	return &Info{
		systemService: systemService,
	}
}

func (i *Info) CollectSystemInfo() (*services.SystemInfo, error) {
	cpuUsage, err := i.systemService.GetCPUUsage()
	if err != nil {
		return nil, err
	}

	memoryUsage, err := i.systemService.GetMemoryUsage()
	if err != nil {
		return nil, err
	}

	diskUsage, err := i.systemService.GetDiskUsage()
	if err != nil {
		return nil, err
	}

	cpuTemperature, err := i.systemService.GetCPUTemperature()
	if err != nil {
		return nil, err
	}

	systemLoad, err := i.systemService.GetSystemLoad()
	if err != nil {
		return nil, err
	}

	return &services.SystemInfo{
		CPUUsage:       cpuUsage,
		MemoryUsage:    memoryUsage,
		DiskUsage:      diskUsage,
		CPUTemperature: cpuTemperature,
		SystemLoad:     systemLoad,
	}, nil
}
