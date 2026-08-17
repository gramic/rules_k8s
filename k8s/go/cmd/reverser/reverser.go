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
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type yamlDoc struct {
	vMap   yaml.MapSlice
	vList  []interface{}
	isInt  bool
	vInt   int
	isBool bool
	vBool  bool
	isStr  bool
	vStr   string
}

func (y *yamlDoc) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(&y.vMap); err == nil {
		return nil
	}
	if err := unmarshal(&y.vList); err == nil {
		return nil
	}
	if err := unmarshal(&y.vInt); err == nil {
		y.isInt = true
		return nil
	}
	if err := unmarshal(&y.vBool); err == nil {
		y.isBool = true
		return nil
	}
	if err := unmarshal(&y.vStr); err == nil {
		y.isStr = true
		return nil
	}
	return fmt.Errorf("unable to parse given blob as a YAML map, list, string, integer or boolean")
}

func (y *yamlDoc) val() interface{} {
	if y.vMap != nil {
		return y.vMap
	}
	if y.vList != nil {
		return y.vList
	}
	if y.isInt {
		return y.vInt
	}
	if y.isBool {
		return y.vBool
	}
	if y.isStr {
		return y.vStr
	}
	return nil
}

func reverseYAML(r io.Reader) ([]byte, error) {
	d := yaml.NewDecoder(r)
	var docs []interface{}
	for {
		y := yamlDoc{}
		if err := d.Decode(&y); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("error decoding YAML: %w", err)
		}
		docs = append(docs, y.val())
	}

	if len(docs) == 0 {
		return nil, nil
	}

	buf := &bytes.Buffer{}
	e := yaml.NewEncoder(buf)
	for i := len(docs) - 1; i >= 0; i-- {
		if err := e.Encode(docs[i]); err != nil {
			return nil, fmt.Errorf("error encoding YAML: %w", err)
		}
	}
	if err := e.Close(); err != nil {
		return nil, fmt.Errorf("error closing YAML encoder: %w", err)
	}

	return buf.Bytes(), nil
}

func main() {
	res, err := reverseYAML(os.Stdin)
	if err != nil {
		log.Fatalf("reverser error: %v", err)
	}
	if len(res) > 0 {
		fmt.Print(string(res))
	}
}
