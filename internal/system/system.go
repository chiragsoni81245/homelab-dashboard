// This package contains the logic to fetch system stats
package system

import (
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
)

type DiskUsage struct {
	Name   string `json:"name"`
	Used   uint64 `json:"used"`
	Total  uint64 `json:"total"`
	Usage  float64 `json:"usage"`
}

type SystemStatus struct {
	CPU    struct {
		Usage float64 `json:"usage"`
		Tempreture float64 `json:"tempreture"`
	} `json:"cpu"`
	Memory struct {
		Used  uint64 `json:"used"`
		Total uint64 `json:"total"`
		Usage float64 `json:"usage"`
	} `json:"memory"`
	Disks []DiskUsage `json:"disks"`
}

func GetSystemStatus() (*SystemStatus, error) {
	status := &SystemStatus{}

	// CPU %
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}
	if len(cpuPercent) > 0 {
		status.CPU.Usage = cpuPercent[0]
	}

	// Temperature
	hostTempretures, err := host.SensorsTemperatures()
	for _, t := range hostTempretures {
		// filter only CPU sensors (varies by system: "coretemp", "cpu", etc.)
		if t.SensorKey == "" || (t.SensorKey[:4] != "core" && t.SensorKey[:3] != "cpu") {
			continue
		}
		status.CPU.Tempreture += t.Temperature
	}
	status.CPU.Tempreture = status.CPU.Tempreture / float64(len(hostTempretures))

	// Memory
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	status.Memory.Used = vmStat.Used
	status.Memory.Total = vmStat.Total
	status.Memory.Usage = vmStat.UsedPercent

	// Disks
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	diskMap := make(map[string]*DiskUsage)

	for _, p := range partitions {
		usageStat, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		// derive parent disk name (e.g. /dev/sda1 -> sda)
		base := filepath.Base(p.Device) // e.g. "sda1"
		diskName := strings.TrimRight(base, "0123456789") // crude strip of trailing digits

		if _, ok := diskMap[diskName]; !ok {
			diskMap[diskName] = &DiskUsage{Name: diskName}
		}

		d := diskMap[diskName]
		d.Used += usageStat.Used
		d.Total += usageStat.Total
	}
	for _, d := range diskMap {
		if d.Name == "loop" { continue }

		if d.Total > 0 {
			d.Used = uint64(d.Used) / 1024 / 1024 / 1024
			d.Total = uint64(d.Total) / 1024 / 1024 / 1024
			d.Usage = (float64(d.Used) / float64(d.Total)) * 100
		}
		status.Disks = append(status.Disks, *d)
	}

	return status, nil
}
