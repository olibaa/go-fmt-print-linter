package gofmtprintlinter_test

import (
	"testing"

	"github.com/olibaa/gofmtprintlinter"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gofmtprintlinter.Analyzer, "test_case_a", "test_case_b")
}
