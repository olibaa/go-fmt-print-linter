package go_fmt_print_linter

import (
	"flag"
	"go/ast"
	"strings"

	"github.com/gostaticanalysis/comment"
	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
)

const (
	name = "go_fmt_print_linter"
	doc  = "go_fmt_print_linter is a static analysis tool to detect `fmt.Print` or `fmt.Println` or `fmt.Printf`"
)

var Analyzer = &analysis.Analyzer{
	Name: name,
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		commentmap.Analyzer,
	},
}

var (
	generatedComment    = flag.String("generated-comment", "Code generated by", "Please set a comment that can identify the file generated.")
	ignoreGeneratedFile = flag.Bool("ignore-gen-file", false, "ignore generated file flag")
)

func run(pass *analysis.Pass) (interface{}, error) {
	flag.Parse()
	cmaps := pass.ResultOf[commentmap.Analyzer].(comment.Maps)

	for _, file := range pass.Files {

		// generated file is skip if ignore-gen-file flag is true
		if *ignoreGeneratedFile && isGeneratedFile(file) {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.CallExpr:

				fun, ok := x.Fun.(*ast.SelectorExpr)
				if !ok {
					return true
				}

				ident, ok := fun.X.(*ast.Ident)
				if !ok {
					return true
				}

				if ident.Name == "fmt" && (fun.Sel.Name == "Print" || fun.Sel.Name == "Println" || fun.Sel.Name == "Printf") {
					if cmaps.IgnorePos(x.Pos(), "go_fmt_print_linter") {
						return true
					}
					pass.Reportf(x.Pos(), "Did you forget to delete \"%s.%s\"? If you add it, please add an ignore comment.", ident.Name, fun.Sel.Name)
				}
			}
			return true
		})
	}
	return nil, nil
}

func isGeneratedFile(file *ast.File) bool {
	for _, c := range file.Comments {
		if strings.Contains(c.Text(), *generatedComment) {
			return true
		}
	}
	return false
}
