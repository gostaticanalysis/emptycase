package emptycase

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "emptycase find empty cases"

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "emptycase",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CaseClause)(nil),
		(*ast.CommClause)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CaseClause:
			if len(n.Body) == 0 {
				pass.Reportf(n.Pos(), "empty case")
			}
		case *ast.CommClause:
			if len(n.Body) == 0 {
				pass.Reportf(n.Pos(), "empty case")
			}
		}
	})

	return nil, nil
}
