package fuzzconfig

import (
	"github.com/aivahealth/hcl/json"
)

func Fuzz(data []byte) int {
	_, diags := json.Parse(data, "<fuzz-conf>")

	if diags.HasErrors() {
		return 0
	}

	return 1
}
