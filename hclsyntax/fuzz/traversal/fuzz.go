package fuzztraversal

import (
	"github.com/aivahealth/hcl"
	"github.com/aivahealth/hcl/hclsyntax"
)

func Fuzz(data []byte) int {
	_, diags := hclsyntax.ParseTraversalAbs(data, "<fuzz-trav>", hcl.Pos{Line: 1, Column: 1})

	if diags.HasErrors() {
		return 0
	}

	return 1
}
