// Copyright 2017 The Bazel Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"strings"
	"testing"
)

func TestReverseYAML(t *testing.T) {
	input := `apiVersion: v1
kind: Service
metadata:
  name: test-svc
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-deploy
`
	expected := `apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-deploy
---
apiVersion: v1
kind: Service
metadata:
  name: test-svc
`
	gotBytes, err := reverseYAML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("reverseYAML failed: %v", err)
	}
	got := string(gotBytes)
	if strings.TrimSpace(got) != strings.TrimSpace(expected) {
		t.Errorf("got:\n%s\nwant:\n%s", got, expected)
	}
}

func TestReverseYAML_Single(t *testing.T) {
	input := `apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
`
	gotBytes, err := reverseYAML(strings.NewReader(input))
	if err != nil {
		t.Fatalf("reverseYAML failed: %v", err)
	}
	got := string(gotBytes)
	if strings.TrimSpace(got) != strings.TrimSpace(input) {
		t.Errorf("got:\n%s\nwant:\n%s", got, input)
	}
}

func TestReverseYAML_Empty(t *testing.T) {
	gotBytes, err := reverseYAML(strings.NewReader(""))
	if err != nil {
		t.Fatalf("reverseYAML failed on empty input: %v", err)
	}
	if len(gotBytes) != 0 {
		t.Errorf("expected empty output, got: %s", string(gotBytes))
	}
}
