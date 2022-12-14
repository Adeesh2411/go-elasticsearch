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

// DateHistogramGrouping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d63a0e35ee85d84c83d0612ff3c0641a7a1e7e2e/specification/rollup/_types/Groupings.ts#L30-L38
type DateHistogramGrouping struct {
	CalendarInterval *Duration `json:"calendar_interval,omitempty"`
	Delay            *Duration `json:"delay,omitempty"`
	Field            string    `json:"field"`
	FixedInterval    *Duration `json:"fixed_interval,omitempty"`
	Format           *string   `json:"format,omitempty"`
	Interval         *Duration `json:"interval,omitempty"`
	TimeZone         *string   `json:"time_zone,omitempty"`
}

// NewDateHistogramGrouping returns a DateHistogramGrouping.
func NewDateHistogramGrouping() *DateHistogramGrouping {
	r := &DateHistogramGrouping{}

	return r
}