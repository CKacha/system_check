package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ckacha/system_check/internal/display"
	"github.com/ckacha/system_check/internal/report"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("== system_check ==")
		fmt.Println("what do you want to check?")
		fmt.Println("  1) OS")
		fmt.Println("  2) CPU")
		fmt.Println("  3) Memory")
		fmt.Println("  4) Disks")
		fmt.Println("  5) everything")
		fmt.Println("  q) quit")
		fmt.Print("> ")

		if !scanner.Scan() {
			return
		}
		choice := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if choice == "q" || choice == "quit" || choice == "0" {
			return
		}

		if !isValidChoice(choice) {
			fmt.Println("not a valid option, try again")
			continue
		}

		r, err := report.Gather()
		if err != nil {
			fmt.Fprintln(os.Stderr, "warning:", err)
		}

		fmt.Println()
		switch choice {
		case "1":
			display.PrintOS(r.OS, true)
		case "2":
			display.PrintCPU(r.CPU, true)
		case "3":
			display.PrintMemory(r.Memory, true)
		case "4":
			display.PrintDisks(r.Disks, true)
		case "5":
			display.PrintAll(r)
			drillDown(scanner, r)
		}
	}
}

func isValidChoice(choice string) bool {
	switch choice {
	case "1", "2", "3", "4", "5":
		return true
	}
	return false
}

func drillDown(scanner *bufio.Scanner, r report.Report) {
	fmt.Println("want more detail on something? (os/cpu/mem/disk/n)")
	fmt.Print("> ")

	if !scanner.Scan() {
		return
	}

	fmt.Println()
	switch strings.ToLower(strings.TrimSpace(scanner.Text())) {
	case "os":
		display.PrintOS(r.OS, true)
	case "cpu":
		display.PrintCPU(r.CPU, true)
	case "mem", "memory":
		display.PrintMemory(r.Memory, true)
	case "disk", "disks":
		display.PrintDisks(r.Disks, true)
	}
}
