// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protojson_test

import (
	"testing"

	"github.com/thoohv5/protobuf-go/encoding/protojson"

	"github.com/thoohv5/protobuf-go/types/known/durationpb"
)

func BenchmarkUnmarshal_Duration(b *testing.B) {
	input := []byte(`"-123456789.123456789s"`)

	for i := 0; i < b.N; i++ {
		err := protojson.Unmarshal(input, &durationpb.Duration{})
		if err != nil {
			b.Fatal(err)
		}
	}
}
