package main

import (
	"github.com/olibaa/go_fmt_print_linter"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(
		go_fmt_print_linter.Analyzer,
	)
}
