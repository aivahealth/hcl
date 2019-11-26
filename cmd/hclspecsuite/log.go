package main

import (
	"github.com/aivahealth/hcl"
)

type LogBeginCallback func(testName string, testFile *TestFile)
type LogProblemsCallback func(testName string, testFile *TestFile, diags hcl.Diagnostics)
