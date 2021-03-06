package ast

import (
	"testing"
)

func TestImplicitValueInitExpr(t *testing.T) {
	nodes := map[string]Node{
		`0x7f8c3396fbd8 <<invalid sloc>> 'sqlite3StatValueType':'long long'`: &ImplicitValueInitExpr{
			Address:  "0x7f8c3396fbd8",
			Position: "<invalid sloc>",
			Type1:    "sqlite3StatValueType",
			Type2:    "long long",
			Children: []Node{},
		},
		`0x7feecb0d6af0 <<invalid sloc>> 'char'`: &ImplicitValueInitExpr{
			Address:  "0x7feecb0d6af0",
			Position: "<invalid sloc>",
			Type1:    "char",
			Type2:    "",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
