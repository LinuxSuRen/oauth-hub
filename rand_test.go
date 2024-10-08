/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package oauth

import (
	"strings"
	"testing"
)

const (
	maxRangeTestCount = 500
	testStringLength  = 32
)

func TestString(t *testing.T) {
	valid := "bcdfghjklmnpqrstvwxz2456789"
	for _, l := range []int{0, 1, 2, 10, 123} {
		s := String(l)
		if len(s) != l {
			t.Errorf("expected string of size %d, got %q", l, s)
		}
		for _, c := range s {
			if !strings.ContainsRune(valid, c) {
				t.Errorf("expected valid characters, got %v", c)
			}
		}
	}
}

func BenchmarkRandomStringGeneration(b *testing.B) {
	b.ResetTimer()
	var s string
	for i := 0; i < b.N; i++ {
		s = String(testStringLength)
	}
	b.StopTimer()
	if len(s) == 0 {
		b.Fatal(s)
	}
}
