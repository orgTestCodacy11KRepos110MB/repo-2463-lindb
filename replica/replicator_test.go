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

package replica

import (
	"testing"

	"github.com/lindb/lindb/models"
	"github.com/lindb/lindb/pkg/timeutil"

	"github.com/stretchr/testify/assert"
)

func TestReplicator_String(t *testing.T) {
	time, _ := timeutil.ParseTimestamp("2019-12-12 10:11:10")
	r := replicator{channel: &ReplicatorChannel{
		State: &models.ReplicaState{
			Database:   "test",
			ShardID:    1,
			Leader:     1,
			Follower:   2,
			FamilyTime: time,
		},
	}}

	assert.Equal(t, "[database:test,shard:1,family:20191212101110,from(leader):1,to(follower):2]", r.String())
}
