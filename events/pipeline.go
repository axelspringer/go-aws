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

const (
	// CodePipelineStarted represents a started CodePipeline
	CodePipelineStarted = "STARTED"
	// CodePipelineSucceeded represents a succeeded CodePipeline
	CodePipelineSucceeded = "SUCCEEDED"
	// CodePipelineResumed represents a resumed CodePipeline
	CodePipelineResumed = "RESUMED"
	// CodePipelineFailed represents a failed CodePipeline
	CodePipelineFailed = "FAILED"
	// CodePipelineCanceled represents a canceled CodePipeline
	CodePipelineCanceled = "Failed"
	// CodePipelineSuperseded represents a superseded CodePipeline
	CodePipelineSuperseded = "SUPERSEDED"
)

// CodePipelineEvent represents a CodePipeline event which digested through CloudWatch
type CodePipelineEvent struct {
	CloudWatchEvent
	Detail CodePipelineEventDetails `json:"detail"`
}

// CodePipelineEventDetails represents the CodePipeline events details as digestes by CloudWatch
type CodePipelineEventDetails struct {
	Pipeline    string  `json:"pipeline"`
	ExecutionID string  `json:"execution-id"`
	State       string  `json:"state"`
	Version     float64 `json:"version"`
}
