// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e


package types

// EqlFeatures type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/xpack/usage/types.ts#L95-L103
type EqlFeatures struct {
	Event     uint                 `json:"event"`
	Join      uint                 `json:"join"`
	Joins     EqlFeaturesJoin      `json:"joins"`
	Keys      EqlFeaturesKeys      `json:"keys"`
	Pipes     EqlFeaturesPipes     `json:"pipes"`
	Sequence  uint                 `json:"sequence"`
	Sequences EqlFeaturesSequences `json:"sequences"`
}

// NewEqlFeatures returns a EqlFeatures.
func NewEqlFeatures() *EqlFeatures {
	r := &EqlFeatures{}

	return r
}