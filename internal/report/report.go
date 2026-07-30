package report

import (
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

type OS struct {
	Platform      string
	PlatformVer   string
	KernelVersion string
	Hostname      string
	Uptime        time.Duration
}

type CPU struct {
	Model         string
	PhysicalCores int
	LogicalCores  int
	UsagePercent  float64
}

type Memory struct {
	TotalBytes     uint64
	UsedBytes      uint64
	AvailableBytes uint64
	UsedPercent    float64
}

type DiskVolume struct {
	MountPoint  string
	Fstype      string
	TotalBytes  uint64
	UsedBytes   uint64
	UsedPercent float64
}

type Report struct {
	OS     OS
	CPU    CPU
	Memory Memory
	Disks  []DiskVolume
}

func Gather() (Report, error) {
	var r Report
	var firstErr error
	note := func(err error) {
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}

	if info, err := host.Info(); err == nil {
		r.OS = OS{
			Platform:      info.Platform,
			PlatformVer:   info.PlatformVersion,
			KernelVersion: info.KernelVersion,
			Hostname:      info.Hostname,
			Uptime:        time.Duration(info.Uptime) * time.Second,
		}
	} else {
		note(err)
	}

	if counts, err := cpu.Counts(false); err == nil {
		r.CPU.PhysicalCores = counts
	} else {
		note(err)
	}

	if counts, err := cpu.Counts(true); err == nil {
		r.CPU.LogicalCores = counts
	} else {
		note(err)
	}

	if infos, err := cpu.Info(); err == nil && len(infos) > 0 {
		r.CPU.Model = strings.TrimSpace(infos[0].ModelName)
	} else if err != nil {
		note(err)
	}

	if percents, err := cpu.Percent(200*time.Millisecond, false); err == nil && len(percents) > 0 {
		r.CPU.UsagePercent = percents[0]
	} else if err != nil {
		note(err)
	}

	if vm, err := mem.VirtualMemory(); err == nil {
		r.Memory = Memory{
			TotalBytes:     vm.Total,
			UsedBytes:      vm.Used,
			AvailableBytes: vm.Available,
			UsedPercent:    vm.UsedPercent,
		}
	} else {
		note(err)
	}

	if parts, err := disk.Partitions(false); err == nil {
		for _, p := range parts {
			usage, err := disk.Usage(p.Mountpoint)
			if err != nil {
				continue
			}

			r.Disks = append(r.Disks, DiskVolume{
				MountPoint:  p.Mountpoint,
				Fstype:      p.Fstype,
				TotalBytes:  usage.Total,
				UsedBytes:   usage.Used,
				UsedPercent: usage.UsedPercent,
			})
		}

	} else {
		note(err)
	}

	if firstErr != nil {
		return r, fmt.Errorf("some system info wasn't able to be read: %w", firstErr)
	}

	return r, nil

}
