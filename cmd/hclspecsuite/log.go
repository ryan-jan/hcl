// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/ryan-jan/hcl"
)

type LogBeginCallback func(testName string, testFile *TestFile)
type LogProblemsCallback func(testName string, testFile *TestFile, diags hcl.Diagnostics)
