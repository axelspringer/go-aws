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
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

// Store is the interface to a SSM Store
type Store interface {
	GetParameters() ([]*ssm.Parameter, error)
	GetEnv() (map[string]string, error)
	SetEnv() error
}

// SSMStore contains the parameters, secrets, etc. stored in SSM
type SSMStore struct {
	ServiceID  string
	Parameters []*ssm.Parameter
	SSM        ssmiface.SSMAPI
}
