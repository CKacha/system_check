# system_check
CLI app that can check certain system resources because HWiNFO is too damn confusing

built for [c-li](c-li.hackclub.com)

## download

grab a prebuilt binary from the [latest release](https://github.com/CKacha/system_check/releases/latest) — windows, macOS, and linux (amd64).

## run

```
go run .
```

## build it

```
go build -o system_check.exe .
./system_check.exe
```

## things system check shows you:

- OS: hostname, platform, kernel, uptime
- CPU: model, physical/logical cores, current usage
- Memory: total/used, usage
- Disks: per-volume used/total, usage

Thanks to [gopsutil](https://github.com/shirou/gopsutil)!

