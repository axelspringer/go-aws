// Copyright 2018 Axel Springer SE
// Copyright 2018 Spring KG
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

// CloudWatchCodePipelineEvent represents a CodePipeline event which digested through CloudWatch
type CloudWatchCodePipelineEvent struct {
	Version    string                                `json:"version"`
	ID         string                                `json:"id"`
	DetailType string                                `json:"detail-type"`
	Source     string                                `json:"source"`
	Account    string                                `json:"account"`
	Time       string                                `json:"time"`
	Region     string                                `json:"region"`
	Resources  []CloudWatchCodePipelineEventResource `json:"resources"`
	Detail     CloudWatchCodePipelineEventDetails    `json:"detail"`
}

// CloudWatchCodePipelineEventDetails represents the CodePipeline events details as digestes by CloudWatch
type CloudWatchCodePipelineEventDetails struct {
	Pipeline    string  `json:"pipeline"`
	ExecutionID string  `json:"execution-id"`
	State       string  `json:"state"`
	Version     float64 `json:"version"`
}

// CloudWatchCodePipelineEventResource represents a resource in a CodePipeline Event in CloudWatch
type CloudWatchCodePipelineEventResource string
