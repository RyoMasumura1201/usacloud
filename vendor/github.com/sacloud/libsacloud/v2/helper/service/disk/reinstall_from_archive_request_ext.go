// Copyright 2016-2020 The Libsacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package disk

import (
	"context"

	"github.com/sacloud/libsacloud/v2/helper/query"
	"github.com/sacloud/libsacloud/v2/sacloud"
)

func (req *ReinstallFromArchiveRequest) toRequestParameter(ctx context.Context, caller sacloud.APICaller) (*sacloud.DiskInstallRequest, error) {
	param := &sacloud.DiskInstallRequest{
		SourceArchiveID: req.SourceArchiveID,
	}
	if req.SourceArchiveID.IsEmpty() {
		// find archive by ostype
		archive, err := query.FindArchiveByOSType(ctx, sacloud.NewArchiveOp(caller), req.Zone, req.OSType)
		if err != nil {
			return nil, err
		}
		param.SourceArchiveID = archive.ID
	}
	return param, nil
}