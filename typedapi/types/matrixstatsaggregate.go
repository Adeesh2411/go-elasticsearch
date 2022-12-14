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

// MatrixStatsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/_types/aggregations/Aggregate.ts#L748-L752
type MatrixStatsAggregate struct {
	DocCount int64                  `json:"doc_count"`
	Fields   []MatrixStatsFields    `json:"fields"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

// NewMatrixStatsAggregate returns a MatrixStatsAggregate.
func NewMatrixStatsAggregate() *MatrixStatsAggregate {
	r := &MatrixStatsAggregate{}

	return r
}