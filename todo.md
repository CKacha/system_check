mmm want to make smjth but what could it be

maybe cli thing to get system info?

what does HWiNFO do again
https://www.hwinfo.com/about-software/supported-components/
https://www.hwinfo.com/sdk/ 
https://sourceforge.net/projects/hwinfo/


ok gonna use gopsutil for cross-platform system info
    https://github.com/shirou/gopsutil
    https://pkg.go.dev/github.com/shirou/gopsutil/v4
stuff i can probably use:
    report.Gather() 
        returns partial data + wrapped error if any single source fails
    (disk.Partitions(false))
        disk section only lists physical partitions, skips virtual/pseudo mounts

thing to show 
    live view, (probably a loop with a clear screen)
    gpu info
    colorized output
    flags? 

added a menu: pick a section (os/cpu/mem/disk) for that data + a plain-english
"what this means" explainer, or "everything" for the brief overview, which then
offers one drill-down into a section. number+enter input via bufio.Scanner,
no TUI lib for now.