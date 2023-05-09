package go_fmt_print_linter_test

import (
	"testing"

	"github.com/olibaa/go_fmt_print_linter"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, go_fmt_print_linter.Analyzer, "test_case_a", "test_case_b")
}
