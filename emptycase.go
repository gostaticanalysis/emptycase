package emptycase

import (
	"github.com/gostaticanalysis/comment"
	"go/ast"
	"strings"

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

var analyzerTest = false

func run(pass *analysis.Pass) (interface{}, error) {
	commentMap := comment.New(pass.Fset, pass.Files)

	// Report a case with no body.
	// However, if a comment is attached to it, it is considered an intentional empty case and not reported.
	var emptyCaseChecker = func(caseNode ast.Node, body []ast.Stmt) {
		if len(body) != 0 {
			return
		}

		commentGroup := commentMap.Comments(caseNode)
		if len(commentGroup) == 0 {
			pass.Reportf(caseNode.Pos(), "empty case")
			return
		}

		if analyzerTest {
			// Comments that indicate the expected value of the analyzer test are not considered comments.
			if len(commentGroup) == 1 && strings.HasPrefix(strings.TrimLeft(commentGroup[0].List[0].Text[2:], " "), "want") {
				pass.Reportf(caseNode.Pos(), "empty case")
			}
		}
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CaseClause)(nil),
		(*ast.CommClause)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CaseClause:
			emptyCaseChecker(n, n.Body)
		case *ast.CommClause:
			emptyCaseChecker(n, n.Body)
		}
	})

	return nil, nil
}
