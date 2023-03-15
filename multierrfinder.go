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

	// まずimportだけを見てmultierrが使われていることを確認する
	multierrIsUsed := false
	inspect.Preorder([]ast.Node{(*ast.GenDecl)(nil)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.GenDecl:
			// ["fmt","strconv",...]
			for _, spec := range n.Specs {
				// spec を *ast.ImportSpecにキャストする
				// 参考 : https://golang.hateblo.jp/entry/golang-interface-type-assertion
				s, _ := spec.(*ast.ImportSpec)

				// s.Path.Valueは"hoge"のような形で文字列が入っているので、
				// ""を文字列から削除する
				// 参考　: https://tenntenn.dev/ja/posts/quote/
				path, _ := strconv.Unquote(s.Path.Value)
				
				if path == "go.uber.org/multierr" {
					pass.Reportf(s.Pos(), "hogehoge")
					multierrIsUsed = true
				}
			}
		}
	})

	if !multierrIsUsed {
		return nil, nil
	}

	// 次にコード内でmultierr.Errors(err)が使われていることを確認する
	inspect.Preorder([]ast.Node{(*ast.SelectorExpr)(nil)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.SelectorExpr:
			i := n.X.(*ast.Ident)
			if i.Name == "multierr"{
				if n.Sel.Name == "Errors" {
					pass.Reportf(i.Pos(), "multierr found")
				}
			}
		}
	})

	return nil, nil
}
