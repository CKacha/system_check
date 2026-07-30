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

func PrintOS(o report.OS, explain bool) {
	fmt.Println("== OS ==")
	fmt.Printf("  hostname:  %s\n", o.Hostname)
	fmt.Printf("  platform:  %s %s\n", o.Platform, o.PlatformVer)
	fmt.Printf("  kernel:    %s\n", o.KernelVersion)
	fmt.Printf("  uptime:    %s\n", o.Uptime.Round(1e9))
	if explain {
		fmt.Println()
		fmt.Println("  what this means:")
		fmt.Println("  hostname is this computer's name on the network.")
		fmt.Println("  platform/kernel tell you which OS and build you're running.")
		fmt.Println("  uptime is how long it's been since the last restart - a long uptime")
		fmt.Println("  is probably not good, make sure to restart your computer occasionally!")
	}
	fmt.Println()
}

func PrintCPU(c report.CPU, explain bool) {
	fmt.Println("== CPU ==")
	fmt.Printf("  model:     %s\n", c.Model)
	fmt.Printf("  cores:     %d physical / %d logical\n", c.PhysicalCores, c.LogicalCores)
	fmt.Printf("  usage:     %s %.1f%%\n", bar(c.UsagePercent, 20), c.UsagePercent)
	if explain {
		fmt.Println()
		fmt.Println("  what this means:")
		fmt.Println("  physical cores are actual processor cores; logical cores also")
		fmt.Println("  count hyperthreads, so logical is often double physical.")
		fmt.Println("  usage is how busy the CPU is right now; spikes are normal,")
		fmt.Println("  but if it's stuck near 100% something is hogging the processor.")
	}
	fmt.Println()
}

func PrintMemory(m report.Memory, explain bool) {
	fmt.Println("== Memory ==")
	fmt.Printf("  total:     %.1f GB\n", bytesToGB(m.TotalBytes))
	fmt.Printf("  used:      %.1f GB\n", bytesToGB(m.UsedBytes))
	fmt.Printf("  usage:     %s %.1f%%\n", bar(m.UsedPercent, 20), m.UsedPercent)
	if explain {
		fmt.Println()
		fmt.Println("  what this means:")
		fmt.Println("  this is your RAM, the short-term memory your programs use while running.")
		fmt.Println("  high usage isn't automatically bad (unused RAM is wasted RAM),")
		fmt.Println("  but if it's maxed out and things feel slow, you may need to close")
		fmt.Println("  some programs or consider more RAM.")
	}
	fmt.Println()
}

func PrintDisks(disks []report.DiskVolume, explain bool) {
	fmt.Println("== Disks ==")
	for _, d := range disks {
		fmt.Printf("  %s (%s)\n", d.MountPoint, d.Fstype)
		fmt.Printf("    %.1f GB / %.1f GB used  %s %.1f%%\n",
			bytesToGB(d.UsedBytes), bytesToGB(d.TotalBytes), bar(d.UsedPercent, 20), d.UsedPercent)
	}
	if explain {
		fmt.Println()
		fmt.Println("  what this means:")
		fmt.Println("  this is storage space(not RAM!), it's where your files, apps, and OS live.")
		fmt.Println("  once a drive gets close to full (usually 90%+), your computer can")
		fmt.Println("  slow down or apps can fail to save/update, so it's worth clearing space.")
	}
	fmt.Println()
}

func PrintAll(r report.Report) {
	fmt.Println("|| smol system check ||")
	fmt.Println()
	PrintOS(r.OS, false)
	PrintCPU(r.CPU, false)
	PrintMemory(r.Memory, false)
	PrintDisks(r.Disks, false)
}
