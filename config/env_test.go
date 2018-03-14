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

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvProjectID(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		isSuccess bool
		envName   string
		projectID string
	}{
		{true, "PROJECT_ID", "foo"},
		{true, "PROJECT_ID", "bar"},
		{false, "PROJECT_ID1", "xxx"},
		{false, "PROJECT_ID2", "xxx"},
		{false, "PROJECT_I", "xxx"},
	}

	defer os.Clearenv()

	for _, tt := range tests {
		target := fmt.Sprintf("%+v", tt)
		os.Clearenv()

		env, err := EnvProjectID()
		assert.Equal("", env, target)
		assert.EqualError(err, errNoProjectIDString)

		os.Setenv(tt.envName, tt.projectID)
		env, err = EnvProjectID()
		if !tt.isSuccess {
			assert.Equal("", env, target)
			assert.EqualError(err, errNoProjectIDString)
			return
		}

		if assert.NoError(err) {
			assert.Equal(tt.projectID, env, target)
		}

		os.Clearenv()
		env, err = EnvProjectID()
		assert.Equal("", env, target)
		assert.EqualError(err, errNoProjectIDString)
	}
}
