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

package api

import (
	"github.com/gin-gonic/gin"

	"github.com/lindb/lindb/app/broker/api/admin"
	"github.com/lindb/lindb/app/broker/api/cluster"
	"github.com/lindb/lindb/app/broker/api/query"
	"github.com/lindb/lindb/app/broker/api/state"
	"github.com/lindb/lindb/app/broker/api/write"
	"github.com/lindb/lindb/app/broker/deps"
)

// API represents broker http api.
type API struct {
	master      *cluster.MasterAPI
	database    *admin.DatabaseAPI
	flusher     *admin.DatabaseFlusherAPI
	storage     *admin.StorageClusterAPI
	brokerState *state.BrokerAPI
	//storageState    *state.StorageAPI
	prometheus      *write.PrometheusWriter
	influxIngestion *write.InfluxWriter
	nativeIngestion *write.NativeWriter
	metric          *query.MetricAPI
	metadata        *query.MetadataAPI
}

// NewAPI creates broker http api.
func NewAPI(deps *deps.HTTPDeps) *API {
	return &API{
		master:      cluster.NewMasterAPI(deps),
		database:    admin.NewDatabaseAPI(deps),
		flusher:     admin.NewDatabaseFlusherAPI(deps),
		storage:     admin.NewStorageClusterAPI(deps),
		brokerState: state.NewBrokerAPI(deps),
		//storageState:    state.NewStorageAPI(deps),
		prometheus:      write.NewPrometheusWriter(deps),
		influxIngestion: write.NewInfluxWriter(deps),
		nativeIngestion: write.NewNativeWriter(deps),
		metric:          query.NewMetricAPI(deps),
		metadata:        query.NewMetadataAPI(deps),
	}
}

// RegisterRouter registers v1 http api router.
func (api *API) RegisterRouter(router *gin.RouterGroup) {
	api.master.Register(router)
	api.database.Register(router)
	api.flusher.Register(router)
	api.storage.Register(router)

	api.brokerState.Register(router)
	//TODO need add
	//api.storageState.Register(router)

	api.metadata.Register(router)
	api.metric.Register(router)
	api.influxIngestion.Register(router)
	api.nativeIngestion.Register(router)
	api.prometheus.Register(router)
}