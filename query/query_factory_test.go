// Licensed to LinDB under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. LinDB licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package query

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/lindb/lindb/sql/stmt"
)

func TestExecutorFactory_NewExecutor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	factory := NewQueryFactory(nil, nil, nil, nil)
	assert.NotNil(t, factory.NewMetricQuery(
		context.Background(),
		"",
		""))
	assert.NotNil(t, factory.NewMetadataQuery(
		context.Background(),
		"",
		&stmt.Metadata{}))

	assert.NotNil(t, factory.newStorageMetricQuery(
		nil,
		nil,
		factory.newStorageQueryContext(nil, &stmt.Query{})))
	assert.NotNil(t, factory.newStorageMetadataQuery(
		nil, nil, nil))
}