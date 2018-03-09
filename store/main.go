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

package store

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// ensure SSMStore confirms to Store interface
var _ Store = &SSMStore{}

// New returns a new wrapper for the Lambda functionality
func New(projectID string) *SSMStore {
	return &SSMStore{
		ProjectID: projectID,
	}
}

// GetParameters fetches the SSM parameters for the configured ProjectID
// and also returns the parameters
func (l *SSMStore) GetParameters() ([]*ssm.Parameter, error) {
	var err error

	l.Parameters = make([]*ssm.Parameter, 0)
	l.Parameters, err = l.getSSMParameters(true, true, nil)
	if err != nil {
		return nil, err
	}

	return l.Parameters, err
}

// SetEnv is setting the available env variables to os
func (l *SSMStore) SetEnv() error {
	var err error

	return err
}

// GetEnv returns an environemnt of parameter name and value
func (l *SSMStore) GetEnv() (map[string]string, error) {
	var err error

	env := make(map[string]string)

	if l.Parameters == nil {
		_, err = l.GetParameters()
		if err != nil {
			return nil, err
		}
	}

	for _, parameter := range l.Parameters {
		// should be improved
		env[strings.Split(aws.StringValue(parameter.Name), "/")[1]] = aws.StringValue(parameter.Value)
	}

	return env, err
}

func (l *SSMStore) getSSMParameters(recursive bool, withDecryption bool, nextToken *string) ([]*ssm.Parameter, error) {
	var err error

	params := &ssm.GetParametersByPathInput{
		Path:           aws.String(fmt.Sprintf("/%s", l.ProjectID)),
		WithDecryption: aws.Bool(withDecryption),
	}

	output, err := l.SSM.GetParametersByPath(params)
	if err != nil {
		return l.Parameters, err
	}

	l.Parameters = append(l.Parameters, output.Parameters...)

	if nextToken != nil {
		l.getSSMParameters(recursive, withDecryption, nextToken)
	}

	return l.Parameters, err
}
