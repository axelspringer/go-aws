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

// CloudWatchEvent represents an CloudWatch Event
type CloudWatchEvent struct {
	Version    string                    `json:"version"`
	ID         string                    `json:"id"`
	DetailType string                    `json:"detail-type"`
	Source     string                    `json:"source"`
	Account    string                    `json:"account"`
	Time       string                    `json:"time"`
	Region     string                    `json:"region"`
	Resources  []CloudWatchEventResource `json:"resources"`
	Detail     string                    `json:"detail"`
}

// CloudWatchEventResource represents associated resources of a CloudWatchEvent
type CloudWatchEventResource string
