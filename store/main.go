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
func New(serviceID string) *SSMStore {
	return &SSMStore{
		ServiceID: serviceID, // Setting the ServiceID to an id of a service in SSM
	}
}

// GetParameters fetches the SSM parameters for the configured ProjectID
// and also returns the parameters
func (s *SSMStore) GetParameters() ([]*ssm.Parameter, error) {
	var err error

	s.Parameters = make([]*ssm.Parameter, 0)
	s.Parameters, err = s.getSSMParameters(true, true, nil)
	if err != nil {
		return nil, err
	}

	return s.Parameters, err
}

// TestEnv is testing for the existence of required parameters
func (s *SSMStore) TestEnv(parameters []string) (bool, error) {
	var err error

	env, err := s.GetEnv()
	if err != nil {
		return false, err
	}

	for _, parameter := range parameters {
		if _, ok := env[parameter]; !ok {
			return false, fmt.Errorf("%v does not exists", parameter)
		}
	}

	return true, err
}

// SetEnv is setting the available env variables to os
func (s *SSMStore) SetEnv() error {
	var err error

	return err
}

// GetEnv returns an environment of parameter name and value
func (s *SSMStore) GetEnv() (map[string]string, error) {
	var err error

	env := make(map[string]string)

	if s.Parameters == nil {
		_, err = s.GetParameters()
		if err != nil {
			return nil, err
		}
	}

	for _, parameter := range s.Parameters {
		// TODO: Refactor to more general solution
		env[strings.Split(aws.StringValue(parameter.Name), "/")[2]] = aws.StringValue(parameter.Value)
	}

	return env, err
}

// getSSMParameters is wrapping the functionality to retrieve parameters from SSM
func (s *SSMStore) getSSMParameters(recursive bool, withDecryption bool, nextToken *string) ([]*ssm.Parameter, error) {
	var err error

	params := &ssm.GetParametersByPathInput{
		Path:           aws.String(fmt.Sprintf("/%s", s.ServiceID)),
		WithDecryption: aws.Bool(withDecryption),
	}

	output, err := s.SSM.GetParametersByPath(params)
	if err != nil {
		return s.Parameters, err
	}

	s.Parameters = append(s.Parameters, output.Parameters...)

	if nextToken != nil {
		s.getSSMParameters(recursive, withDecryption, nextToken)
	}

	return s.Parameters, err
}
