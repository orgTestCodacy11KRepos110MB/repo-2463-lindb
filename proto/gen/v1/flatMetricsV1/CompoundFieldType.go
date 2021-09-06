// Licensed to LinDB under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. LinDB licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatMetricsV1

import "strconv"

type CompoundFieldType int8

const (
	CompoundFieldTypeUnSpecified         CompoundFieldType = 0
	CompoundFieldTypeDeltaHistogram      CompoundFieldType = 1
	CompoundFieldTypeCumultaiveHistogram CompoundFieldType = 2
)

var EnumNamesCompoundFieldType = map[CompoundFieldType]string{
	CompoundFieldTypeUnSpecified:         "UnSpecified",
	CompoundFieldTypeDeltaHistogram:      "DeltaHistogram",
	CompoundFieldTypeCumultaiveHistogram: "CumultaiveHistogram",
}

var EnumValuesCompoundFieldType = map[string]CompoundFieldType{
	"UnSpecified":         CompoundFieldTypeUnSpecified,
	"DeltaHistogram":      CompoundFieldTypeDeltaHistogram,
	"CumultaiveHistogram": CompoundFieldTypeCumultaiveHistogram,
}

func (v CompoundFieldType) String() string {
	if s, ok := EnumNamesCompoundFieldType[v]; ok {
		return s
	}
	return "CompoundFieldType(" + strconv.FormatInt(int64(v), 10) + ")"
}