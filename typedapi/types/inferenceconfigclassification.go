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
// https://github.com/elastic/elasticsearch-specification/tree/899364a63e7415b60033ddd49d50a30369da26d7

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// InferenceConfigClassification type.
//
// https://github.com/elastic/elasticsearch-specification/blob/899364a63e7415b60033ddd49d50a30369da26d7/specification/ingest/_types/Processors.ts#L257-L263
type InferenceConfigClassification struct {
	NumTopClasses                 *int    `json:"num_top_classes,omitempty"`
	NumTopFeatureImportanceValues *int    `json:"num_top_feature_importance_values,omitempty"`
	PredictionFieldType           *string `json:"prediction_field_type,omitempty"`
	ResultsField                  *string `json:"results_field,omitempty"`
	TopClassesResultsField        *string `json:"top_classes_results_field,omitempty"`
}

func (s *InferenceConfigClassification) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "num_top_classes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumTopClasses = &value
			case float64:
				f := int(v)
				s.NumTopClasses = &f
			}

		case "num_top_feature_importance_values":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumTopFeatureImportanceValues = &value
			case float64:
				f := int(v)
				s.NumTopFeatureImportanceValues = &f
			}

		case "prediction_field_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.PredictionFieldType = &o

		case "results_field":
			if err := dec.Decode(&s.ResultsField); err != nil {
				return err
			}

		case "top_classes_results_field":
			if err := dec.Decode(&s.TopClassesResultsField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewInferenceConfigClassification returns a InferenceConfigClassification.
func NewInferenceConfigClassification() *InferenceConfigClassification {
	r := &InferenceConfigClassification{}

	return r
}
