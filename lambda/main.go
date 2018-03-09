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

package lambda

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Func contains all the information about the lambda
type Func struct {
	ProjectID  string
	Parameters []*ssm.Parameter
	SSM        *ssm.SSM
}

func (l *Func) getParameters() ([]*ssm.Parameter, error) {
	var err error

	l.Parameters = make([]*ssm.Parameter, 0)
	l.Parameters, err = l.getSSMParameters(true, true, nil)
	if err != nil {
		return nil, err
	}

	return l.Parameters, err
}

func (l *Func) getSSMParameters(recursive bool, withDecryption bool, nextToken *string) ([]*ssm.Parameter, error) {
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
