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
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockedGetParametersByPathOutput struct {
	ssmiface.SSMAPI
	Resp ssm.GetParametersByPathOutput
}

func (m mockedGetParametersByPathOutput) GetParametersByPath(in *ssm.GetParametersByPathInput) (*ssm.GetParametersByPathOutput, error) {
	return &m.Resp, nil // only return mocked resp
}

func TestGetSSMParameters(t *testing.T) {
	cases := []struct {
		Resp     ssm.GetParametersByPathOutput
		Expected []*ssm.Parameter
	}{
		{
			// Case 1, expected parameters to be returned, non encrypted
			Resp: ssm.GetParametersByPathOutput{
				NextToken: nil,
				Parameters: []*ssm.Parameter{
					{
						Name:    aws.String("/test_service/test_parameter"),
						Value:   aws.String("test_value"),
						Type:    aws.String(ssm.ParameterTypeString),
						Version: aws.Int64(1),
					},
				},
			},
			Expected: []*ssm.Parameter{
				{
					Name:    aws.String("/test_service/test_parameter"),
					Value:   aws.String("test_value"),
					Type:    aws.String(ssm.ParameterTypeString),
					Version: aws.Int64(1),
				},
			},
		},
	}

	for i, c := range cases {
		s := SSMStore{
			ServiceID: "test_service",
			SSM:       mockedGetParametersByPathOutput{Resp: c.Resp},
		}
		parameters, err := s.GetParameters()
		if err != nil {
			t.Fatalf("%d, unexpected error, %v", i, err)
		}
		if a, e := len(parameters), len(c.Expected); a != e {
			t.Fatalf("%d, expected %d messages, got %d", i, e, a)
		}
		for j, parameter := range parameters {
			if a, e := parameter, c.Expected[j]; aws.StringValue(a.Name) != aws.StringValue(e.Name) {
				t.Errorf("%d, expected %v message, got %v", i, e, a)
			}
		}
	}
}
