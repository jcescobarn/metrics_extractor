package services

import (
	"os"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	CPUUsage       float64
	MemoryUsage    float64
	DiskUsage      float64
	CPUTemperature float64
	SystemLoad     float64
}

type SystemService struct{}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (ss *SystemService) GetSystemLoad() (float64, error) {
	avg, err := load.Avg()
	if err != nil {
		return 0.0, err

	}
	return avg.Load1, nil
}

func (ss *SystemService) GetCPUTemperature() (float64, error) {
	bytes, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}
	tempStr := strings.TrimSpace(string(bytes))
	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		return 0, err
	}
	return temp / 1000, nil
}

func (ss *SystemService) GetCPUUsage() (float64, error) {
	cpu_use_percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0, err
	}
	return cpu_use_percent[0], nil
}

func (ss *SystemService) GetMemoryUsage() (float64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, err
	}
	return v.UsedPercent, nil
}

func (ss *SystemService) GetDiskUsage() (float64, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return 0.0, err

	}
	return usage.UsedPercent, nil
}
