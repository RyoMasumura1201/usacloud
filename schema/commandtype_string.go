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

// Code generated by "stringer -type=CommandType command_type.go"; DO NOT EDIT

package schema

import "fmt"

const _CommandType_name = "CommandInvalidCommandListCommandCreateCommandReadCommandUpdateCommandDeleteCommandManipulateMultiCommandManipulateSingleCommandManipulateIDOnlyCommandCustom"

var _CommandType_index = [...]uint8{0, 14, 25, 38, 49, 62, 75, 97, 120, 143, 156}

func (i CommandType) String() string {
	if i < 0 || i >= CommandType(len(_CommandType_index)-1) {
		return fmt.Sprintf("CommandType(%d)", i)
	}
	return _CommandType_name[_CommandType_index[i]:_CommandType_index[i+1]]
}
