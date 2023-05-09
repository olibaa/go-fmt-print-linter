package main

import (
	"github.com/olibaa/gofmtprintlinter"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(
		gofmtprintlinter.Analyzer,
	)
}
