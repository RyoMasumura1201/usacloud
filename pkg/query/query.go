// Copyright 2017-2021 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/itchyny/gojq"
	"github.com/jmespath/go-jmespath"
)

func ByJMESPath(v interface{}, query string) (result interface{}, err error) {
	defer func() {
		ret := recover()
		if ret != nil {
			err = fmt.Errorf("jmespath: failed to process query: %s", ret)
		}
	}()
	result, err = jmespath.Search(query, v)
	return
}

func ByGoJQ(input interface{}, query string) (result interface{}, err error) {
	// Note: queryによってはpanicしてスタックトレースを吐くケースがあるためrecoverしておく
	defer func() {
		ret := recover()
		if ret != nil {
			err = fmt.Errorf("gojq: failed to process query: %s", ret)
		}
	}()

	q, err := gojq.Parse(query)
	if err != nil {
		return nil, fmt.Errorf("gojq parse failed: %v", err)
	}

	// gojqにinputを渡す前に[]map[string]interface{}へ変換しておく
	// see: https://pkg.go.dev/github.com/itchyny/gojq#readme-usage-as-a-library
	// > the query input should have type []interface{} for an array and map[string]interface{} for a map
	iter := q.Run(convertInputToMap(input))

	var results []interface{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, fmt.Errorf("gojq: %s", err.Error())
		}
		results = append(results, v)
	}
	return results, err
}

func convertInputToMap(input interface{}) []interface{} {
	var inputs []interface{}
	ins, ok := input.([]interface{})
	if ok {
		inputs = ins
	} else {
		inputs = []interface{}{input}
	}

	for i, v := range inputs {
		if !structs.IsStruct(v) {
			continue
		}
		inputs[i] = structs.Map(v)
	}
	return inputs
}
