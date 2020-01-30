// Copyright 2017-2020 The Usacloud Authors
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

package funcs

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func ServerInterfaceAddForRouter(ctx command.Context, params *params.InterfaceAddForRouterServerParam) error {

	client := ctx.GetAPIClient()
	api := client.GetServerAPI()
	diskAPI := client.GetDiskAPI()
	switchAPI := client.GetSwitchAPI()
	interfaceAPI := client.GetInterfaceAPI()

	p, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", e)
	}

	if len(p.GetInterfaces()) > 0 {
		return fmt.Errorf("Interface to connect to Router+Switch must be the first interface of server")
	}

	// is router?
	sw, err := switchAPI.Read(params.SwitchId)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", e)
	}
	if sw.Internet == nil {
		return fmt.Errorf("SwitchID must be ID of Router+Switch")
	}

	// validate connected disk if need DiskEdit
	if !params.WithoutDiskEdit {
		disks := p.GetDisks()
		if len(disks) == 0 {
			return fmt.Errorf("Server haven't any disks. Can't call EditDisk API without server connected disks")
		}
		// EditDisk API supported?
		res, err := diskAPI.CanEditDisk(disks[0].ID)
		if err != nil {
			return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", e)
		}
		if !res {
			return fmt.Errorf("Can't call EditDisk API with disk(%d)", disks[0].ID)

		}

	}

	if p.IsUp() {
		return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", "server is running")
	}

	// call manipurate functions
	nic, err := interfaceAPI.CreateAndConnectToServer(params.Id)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", err)
	}

	_, err = interfaceAPI.ConnectToSwitch(nic.ID, params.SwitchId)
	if err != nil {
		return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", err)
	}

	if !params.WithoutDiskEdit {
		// disk edit
		editParam := diskAPI.NewCondig()
		editParam.SetBackground(true)

		if params.Ipaddress != "" {
			editParam.SetUserIPAddress(params.Ipaddress)
		}
		if params.NwMasklen > 0 {
			editParam.SetNetworkMaskLen(fmt.Sprintf("%d", params.NwMasklen))
		}
		if params.DefaultRoute != "" {
			editParam.SetDefaultRoute(params.DefaultRoute)
		}

		_, err := diskAPI.Config(p.GetDisks()[0].ID, editParam)
		if err != nil {
			return fmt.Errorf("ServerInterfaceAddForRouter is failed: %s", err)
		}
		if err := diskAPI.SleepWhileCopying(p.GetDisks()[0].ID, client.DefaultTimeoutDuration); err != nil {
			return fmt.Errorf("ServerInterfaceAddForInternet is failed: %s", err)
		}
	}

	return nil

}
