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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	s "github.com/axelspringer/vodka-aws/store"
)

const (
	defaultMaxRetries = 5
)

// New returns a new wrapper for the Lambda functionality
func New(serviceID string) *Func {

	ssmSession := session.Must(session.NewSession())
	ssm := ssm.New(ssmSession, &aws.Config{
		MaxRetries: aws.Int(defaultMaxRetries),
	})

	store := &s.SSMStore{
		ServiceID: serviceID,
		SSM:       ssm,
	}

	return &Func{
		ProjectID: serviceID,
		Store:     store,
	}
}
