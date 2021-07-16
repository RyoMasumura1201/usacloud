// Copyright 2017-2021 The Usacloud Authors
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

package vpcrouter

import (
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/usacloud/pkg/cli"
	"github.com/sacloud/usacloud/pkg/util"
)

type routerSettingUpdate struct {
	InternetConnectionEnabled *bool

	StaticNATData string                         `cli:"static-nat" mapconv:"-"`
	StaticNAT     *[]*sacloud.VPCRouterStaticNAT `cli:"-" mapconv:",omitempty"`

	PortForwardingData string                              `cli:"port-forwarding" mapconv:"-"`
	PortForwarding     *[]*sacloud.VPCRouterPortForwarding `cli:"-" mapconv:",omitempty"`

	FirewallData string                        `cli:"firewall" mapconv:"-"`
	Firewall     *[]*sacloud.VPCRouterFirewall `cli:"-" mapconv:",omitempty"`

	DHCPServerData string                          `cli:"dhcp-server" mapconv:"-"`
	DHCPServer     *[]*sacloud.VPCRouterDHCPServer `cli:"-" mapconv:",omitempty"`

	DHCPStaticMappingData string                                 `cli:"dhcp-static-mapping" mapconv:"-"`
	DHCPStaticMapping     *[]*sacloud.VPCRouterDHCPStaticMapping `cli:"-" mapconv:",omitempty"`

	PPTPServerData string                       `cli:"pptp" mapconv:"-"`
	PPTPServer     *sacloud.VPCRouterPPTPServer `cli:"-" mapconv:",omitempty"`

	L2TPIPsecServerData string                            `cli:"l2tp" mapconv:"-"`
	L2TPIPsecServer     *sacloud.VPCRouterL2TPIPsecServer `cli:"-" mapconv:",omitempty"`

	WireGuardData string                      `cli:"wireguard" mapconv:"-"`
	WireGuard     *sacloud.VPCRouterWireGuard `cli:"-" mapconv:",omitempty"`

	RemoteAccessUsersData string                                `cli:"users" mapconv:"-"`
	RemoteAccessUsers     *[]*sacloud.VPCRouterRemoteAccessUser `cli:"-" mapconv:",omitempty"`

	SiteToSiteIPsecVPNData string                                  `cli:"site-to-site-vpn" mapconv:"-"`
	SiteToSiteIPsecVPN     *[]*sacloud.VPCRouterSiteToSiteIPsecVPN `cli:"-" mapconv:",omitempty"`

	StaticRouteData string                           `cli:"static-route" mapconv:"-"`
	StaticRoute     *[]*sacloud.VPCRouterStaticRoute `cli:"-" mapconv:",omitempty"`

	SyslogHost *string
}

func (r *routerSettingUpdate) Customize(_ cli.Context) error {
	if r.StaticNATData != "" {
		var staticNat []*sacloud.VPCRouterStaticNAT
		if err := util.MarshalJSONFromPathOrContent(r.StaticNATData, &staticNat); err != nil {
			return err
		}
		if r.StaticNAT == nil {
			r.StaticNAT = &[]*sacloud.VPCRouterStaticNAT{}
		}
		*r.StaticNAT = append(*r.StaticNAT, staticNat...)
	}

	if r.PortForwardingData != "" {
		var portForwarding []*sacloud.VPCRouterPortForwarding
		if err := util.MarshalJSONFromPathOrContent(r.PortForwardingData, &portForwarding); err != nil {
			return err
		}
		if r.PortForwarding == nil {
			r.PortForwarding = &[]*sacloud.VPCRouterPortForwarding{}
		}
		*r.PortForwarding = append(*r.PortForwarding, portForwarding...)
	}

	if r.FirewallData != "" {
		var firewall []*sacloud.VPCRouterFirewall
		if err := util.MarshalJSONFromPathOrContent(r.FirewallData, &firewall); err != nil {
			return err
		}
		if r.Firewall == nil {
			r.Firewall = &[]*sacloud.VPCRouterFirewall{}
		}
		*r.Firewall = append(*r.Firewall, firewall...)
	}

	if r.DHCPServerData != "" {
		var dhcpServer []*sacloud.VPCRouterDHCPServer
		if err := util.MarshalJSONFromPathOrContent(r.DHCPServerData, &dhcpServer); err != nil {
			return err
		}
		if r.DHCPServer == nil {
			r.DHCPServer = &[]*sacloud.VPCRouterDHCPServer{}
		}
		*r.DHCPServer = append(*r.DHCPServer, dhcpServer...)
	}

	if r.DHCPStaticMappingData != "" {
		var dhcpStaticMapping []*sacloud.VPCRouterDHCPStaticMapping
		if err := util.MarshalJSONFromPathOrContent(r.DHCPStaticMappingData, &dhcpStaticMapping); err != nil {
			return err
		}
		if r.DHCPStaticMapping == nil {
			r.DHCPStaticMapping = &[]*sacloud.VPCRouterDHCPStaticMapping{}
		}
		*r.DHCPStaticMapping = append(*r.DHCPStaticMapping, dhcpStaticMapping...)
	}

	if r.PPTPServerData != "" {
		var pptp sacloud.VPCRouterPPTPServer
		if err := util.MarshalJSONFromPathOrContent(r.PPTPServerData, &pptp); err != nil {
			return err
		}
		*r.PPTPServer = pptp
	}

	if r.L2TPIPsecServerData != "" {
		var l2tp sacloud.VPCRouterL2TPIPsecServer
		if err := util.MarshalJSONFromPathOrContent(r.L2TPIPsecServerData, &l2tp); err != nil {
			return err
		}
		*r.L2TPIPsecServer = l2tp
	}
	if r.WireGuardData != "" {
		var wireGuard sacloud.VPCRouterWireGuard
		if err := util.MarshalJSONFromPathOrContent(r.WireGuardData, &wireGuard); err != nil {
			return err
		}
		*r.WireGuard = wireGuard
	}

	if r.RemoteAccessUsersData != "" {
		var users []*sacloud.VPCRouterRemoteAccessUser
		if err := util.MarshalJSONFromPathOrContent(r.RemoteAccessUsersData, &users); err != nil {
			return err
		}
		if r.RemoteAccessUsers == nil {
			r.RemoteAccessUsers = &[]*sacloud.VPCRouterRemoteAccessUser{}
		}
		*r.RemoteAccessUsers = append(*r.RemoteAccessUsers, users...)
	}

	if r.SiteToSiteIPsecVPNData != "" {
		var s2s []*sacloud.VPCRouterSiteToSiteIPsecVPN
		if err := util.MarshalJSONFromPathOrContent(r.SiteToSiteIPsecVPNData, &s2s); err != nil {
			return err
		}
		if r.SiteToSiteIPsecVPN == nil {
			r.SiteToSiteIPsecVPN = &[]*sacloud.VPCRouterSiteToSiteIPsecVPN{}
		}
		*r.SiteToSiteIPsecVPN = append(*r.SiteToSiteIPsecVPN, s2s...)
	}

	if r.StaticRouteData != "" {
		var staticRoutes []*sacloud.VPCRouterStaticRoute
		if err := util.MarshalJSONFromPathOrContent(r.StaticRouteData, &staticRoutes); err != nil {
			return err
		}
		if r.StaticRoute == nil {
			r.StaticRoute = &[]*sacloud.VPCRouterStaticRoute{}
		}
		*r.StaticRoute = append(*r.StaticRoute, staticRoutes...)
	}

	return nil
}
