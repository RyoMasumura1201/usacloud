// Copyright 2017-2020 The Usacloud Authors
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

package vdef

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sacloud/libsacloud/v2/pkg/mapconv"
)

// ConverterFilters mapconvでの変換時に利用されるフィルターの定義、definitionsに登録したものは実行時に動的に追加される
var ConverterFilters = map[string]mapconv.FilterFunc{
	"rfc3339":         strToTime,
	"path_to_reader":  pathToReader,
	"path_or_content": pathOrContent,
}

func strToTime(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid time value: %v", v)
	}
	if s == "" {
		return time.Time{}, nil
	}

	allowDatetimeFormatList := []string{
		time.RFC3339,
	}
	for _, format := range allowDatetimeFormatList {
		d, err := time.Parse(format, s)
		if err == nil {
			// success
			return d, nil
		}
	}
	return nil, fmt.Errorf("invalid time format: %v", v)
}

// pathToReader ファイルパスからio.Reader(実体は*os.File)を返す
//
// Note: ファイルはここではクローズされないため、このフィルタを適用する先のリクエストでCloseを適切に呼ぶようにする
// libsacloud serviceの場合はservice内でcloseが呼ばれる
func pathToReader(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid filepath value: %v", v)
	}
	if s == "" {
		return nil, nil
	}

	file, err := os.Open(s)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// pathOrContent 値がファイルだった場合はファイルの内容を、そうでない場合は値をそのまま返す
func pathOrContent(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	s, ok := v.(string)
	if !ok {
		return nil, fmt.Errorf("invalid filepath value: %v", v)
	}
	if s == "" {
		return nil, nil
	}

	file, err := os.Open(s)
	if err != nil {
		if os.IsNotExist(err) {
			return s, nil // そのまま返す
		}
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return string(data), nil
}