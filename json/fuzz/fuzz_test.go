// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fuzzjson

import (
	"testing"

	"github.com/ryan-jan/hcl/json"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		_, diags := json.Parse(data, "<fuzz-conf>")

		if diags.HasErrors() {
			t.Logf("Error when parsing JSON %v", data)
			for _, diag := range diags {
				t.Logf("- %s", diag.Error())
			}
		}
	})
}
