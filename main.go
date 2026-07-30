package main

import (
	"fmt"
	"os"

	"github.com/ckacha/system_check/internal/display"
	"github.com/ckacha/system_check/internal/report"
)

func main() {
	r, err := report.Gather()
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning:", err)
	}
	display.Print(r)
}
