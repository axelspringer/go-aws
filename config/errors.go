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

package config

const (
	errNoProjectIDString = "no valid PROJECT_ID found in os environment"
)

// NewErrNoProjectID returns a new error of
// non-existing PROJECT_ID environment variable
func NewErrNoProjectID() error {
	return &errNoProjectID{errNoProjectIDString}
}

type errNoProjectID struct {
	s string
}

func (e *errNoProjectID) Error() string {
	return e.s
}
