/**
 * Copyright (c) 2023 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cache

import (
	cloudmodel "github.com/deepflowio/deepflow/server/controller/cloud/model"
	"github.com/deepflowio/deepflow/server/controller/db/mysql"
	. "github.com/deepflowio/deepflow/server/controller/recorder/common"
)

func (b *DiffBaseDataSet) addVMPodNodeConnection(dbItem *mysql.VMPodNodeConnection, seq int) {
	b.VMPodNodeConnections[dbItem.Lcuuid] = &VMPodNodeConnection{
		DiffBase: DiffBase{
			Sequence: seq,
			Lcuuid:   dbItem.Lcuuid,
		},
		SubDomainLcuuid: dbItem.SubDomain,
	}
	b.GetLogFunc()(addDiffBase(RESOURCE_TYPE_VM_POD_NODE_CONNECTION_EN, b.VMPodNodeConnections[dbItem.Lcuuid]))
}

func (b *DiffBaseDataSet) deleteVMPodNodeConnection(lcuuid string) {
	delete(b.VMPodNodeConnections, lcuuid)
	log.Info(deleteDiffBase(RESOURCE_TYPE_VM_POD_NODE_CONNECTION_EN, lcuuid))
}

type VMPodNodeConnection struct {
	DiffBase
	SubDomainLcuuid string `json:"sub_domain_lcuuid"`
}

func (p *VMPodNodeConnection) Update(cloudItem *cloudmodel.VMPodNodeConnection) {
	p.SubDomainLcuuid = cloudItem.SubDomainLcuuid
	log.Info(updateDiffBase(RESOURCE_TYPE_VM_POD_NODE_CONNECTION_EN, p))
}