package multierrfinder

import (
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "multierrfinder finds multierr"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "multierrfinder",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			// ["fmt","strconv",...]
			for _, spec := range n.Specs {
				s, _ := spec.(*ast.ImportSpec)
				path, _ := strconv.Unquote(s.Path.Value)
				if path == "go.uber.org/multierr" {
					pass.Reportf(s.Pos(), "hogehoge")
				}
			}

			// case *ast.Ident:
			// 	if n.Name == "gopher" {
			// 		pass.Reportf(n.Pos(), "identifier is gopher")
			// 	}
		}
	})

	return nil, nil
}
