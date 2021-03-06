package ast

import (
	"testing"
)

func TestPredefinedExpr(t *testing.T) {
	nodes := map[string]Node{
		`0x33d6e08 <col:30> 'const char [25]' lvalue __PRETTY_FUNCTION__`: &PredefinedExpr{
			Address:  "0x33d6e08",
			Position: "col:30",
			Type:     "const char [25]",
			Lvalue:   true,
			Name:     "__PRETTY_FUNCTION__",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
