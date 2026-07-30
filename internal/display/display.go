package display

import (
	"fmt"
	"strings"

	"github.com/ckacha/system_check/internal/report"
)

func bytesToGB(b uint64) float64 {
	return float64(b) / (1024 * 1024 * 1024)
}

func bar(percent float64, width int) string {
	filled := min(int(percent/100*float64(width)), width)
	return "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
}

func Print(r report.Report) {
	fmt.Println("|| smol system check ||")
	fmt.Println()

	fmt.Println("== OS ==")
	fmt.Printf("  hostname:  %s\n", r.OS.Hostname)
	fmt.Printf("  platform:  %s %s\n", r.OS.Platform, r.OS.PlatformVer)
	fmt.Printf("  kernel:    %s\n", r.OS.KernelVersion)
	fmt.Printf("  uptime:    %s\n", r.OS.Uptime.Round(1e9))
	fmt.Println()

	fmt.Println("== CPU ==")
	fmt.Printf("  model:     %s\n", r.CPU.Model)
	fmt.Printf("  cores:     %d physical / %d logical\n", r.CPU.PhysicalCores, r.CPU.LogicalCores)
	fmt.Printf("  usage:     %s %.1f%%\n", bar(r.CPU.UsagePercent, 20), r.CPU.UsagePercent)
	fmt.Println()

	fmt.Println("== Memory ==")
	fmt.Printf("  total:     %.1f GB\n", bytesToGB(r.Memory.TotalBytes))
	fmt.Printf("  used:      %.1f GB\n", bytesToGB(r.Memory.UsedBytes))
	fmt.Printf("  usage:     %s %.1f%%\n", bar(r.Memory.UsedPercent, 20), r.Memory.UsedPercent)
	fmt.Println()

	fmt.Println("== Disks ==")
	for _, d := range r.Disks {
		fmt.Printf("  %s (%s)\n", d.MountPoint, d.Fstype)
		fmt.Printf("    %.1f GB / %.1f GB used  %s %.1f%%\n",
			bytesToGB(d.UsedBytes), bytesToGB(d.TotalBytes), bar(d.UsedPercent, 20), d.UsedPercent)
	}
}
