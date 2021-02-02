// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gojson

package protocol

import (
	"testing"

	json "github.com/goccy/go-json"
)

func TestPublishDiagnosticsParams(t *testing.T) {
	testPublishDiagnosticsParams(t, json.MarshalNoEscape, json.UnmarshalNoEscape)
}
