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