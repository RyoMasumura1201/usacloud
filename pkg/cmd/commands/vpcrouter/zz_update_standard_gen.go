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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package vpcrouter

import (
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *updateStandardParameter) CleanupEmptyValue(fs *pflag.FlagSet) {
	if !fs.Changed("name") {
		p.Name = nil
	}
	if !fs.Changed("description") {
		p.Description = nil
	}
	if !fs.Changed("tags") {
		p.Tags = nil
	}
	if !fs.Changed("icon-id") {
		p.IconID = nil
	}
	if !fs.Changed("internet-connection-enabled") {
		p.RouterSetting.InternetConnectionEnabled = nil
	}
	if !fs.Changed("syslog-host") {
		p.RouterSetting.SyslogHost = nil
	}
}

func (p *updateStandardParameter) buildFlags(fs *pflag.FlagSet) {
	if p.Name == nil {
		p.Name = pointer.NewString("")
	}
	if p.Description == nil {
		p.Description = pointer.NewString("")
	}
	if p.Tags == nil {
		p.Tags = pointer.NewStringSlice([]string{})
	}
	if p.IconID == nil {
		p.IconID = pointer.NewID(types.ID(0))
	}
	if p.RouterSetting.InternetConnectionEnabled == nil {
		p.RouterSetting.InternetConnectionEnabled = pointer.NewBool(false)
	}
	if p.RouterSetting.SyslogHost == nil {
		p.RouterSetting.SyslogHost = pointer.NewString("")
	}
	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "(*required) ")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format options: [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "Query for JSON output")
	fs.StringVarP(&p.QueryDriver, "query-driver", "", p.QueryDriver, "Name of the driver that handles queries to JSON output options: [jmespath/jq]")
	fs.StringVarP(p.Name, "name", "", "", "")
	fs.StringVarP(p.Description, "description", "", "", "")
	fs.StringSliceVarP(p.Tags, "tags", "", nil, "")
	fs.VarP(core.NewIDFlag(p.IconID, p.IconID), "icon-id", "", "")
	fs.StringVarP(&p.PrivateNetworkInterfacesData, "private-network-interfaces", "", p.PrivateNetworkInterfacesData, "")
	fs.BoolVarP(p.RouterSetting.InternetConnectionEnabled, "internet-connection-enabled", "", false, "")
	fs.StringVarP(&p.RouterSetting.StaticNATData, "static-nat", "", p.RouterSetting.StaticNATData, "")
	fs.StringVarP(&p.RouterSetting.PortForwardingData, "port-forwarding", "", p.RouterSetting.PortForwardingData, "")
	fs.StringVarP(&p.RouterSetting.FirewallData, "firewall", "", p.RouterSetting.FirewallData, "")
	fs.StringVarP(&p.RouterSetting.DHCPServerData, "dhcp-server", "", p.RouterSetting.DHCPServerData, "")
	fs.StringVarP(&p.RouterSetting.DHCPStaticMappingData, "dhcp-static-mapping", "", p.RouterSetting.DHCPStaticMappingData, "")
	fs.StringVarP(&p.RouterSetting.PPTPServerData, "pptp", "", p.RouterSetting.PPTPServerData, "")
	fs.StringVarP(&p.RouterSetting.L2TPIPsecServerData, "l2tp", "", p.RouterSetting.L2TPIPsecServerData, "")
	fs.StringVarP(&p.RouterSetting.WireGuardData, "wireguard", "", p.RouterSetting.WireGuardData, "")
	fs.StringVarP(&p.RouterSetting.RemoteAccessUsersData, "users", "", p.RouterSetting.RemoteAccessUsersData, "")
	fs.StringVarP(&p.RouterSetting.SiteToSiteIPsecVPNData, "site-to-site-vpn", "", p.RouterSetting.SiteToSiteIPsecVPNData, "")
	fs.StringVarP(&p.RouterSetting.StaticRouteData, "static-route", "", p.RouterSetting.StaticRouteData, "")
	fs.StringVarP(p.RouterSetting.SyslogHost, "syslog-host", "", "", "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *updateStandardParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	}
	return pflag.NormalizedName(name)
}

func (p *updateStandardParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("common", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Common options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("vpc-router", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("dhcp-server"))
		fs.AddFlag(cmd.LocalFlags().Lookup("dhcp-static-mapping"))
		fs.AddFlag(cmd.LocalFlags().Lookup("firewall"))
		fs.AddFlag(cmd.LocalFlags().Lookup("internet-connection-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("l2tp"))
		fs.AddFlag(cmd.LocalFlags().Lookup("port-forwarding"))
		fs.AddFlag(cmd.LocalFlags().Lookup("pptp"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-network-interfaces"))
		fs.AddFlag(cmd.LocalFlags().Lookup("site-to-site-vpn"))
		fs.AddFlag(cmd.LocalFlags().Lookup("static-nat"))
		fs.AddFlag(cmd.LocalFlags().Lookup("static-route"))
		fs.AddFlag(cmd.LocalFlags().Lookup("syslog-host"))
		fs.AddFlag(cmd.LocalFlags().Lookup("users"))
		fs.AddFlag(cmd.LocalFlags().Lookup("wireguard"))
		sets = append(sets, &core.FlagSet{
			Title: "Vpc-Router-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("zone", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("zone"))
		sets = append(sets, &core.FlagSet{
			Title: "Zone options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("wait", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("no-wait"))
		sets = append(sets, &core.FlagSet{
			Title: "Wait options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("input", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-driver"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("example", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("example"))
		sets = append(sets, &core.FlagSet{
			Title: "Parameter example",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *updateStandardParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *updateStandardParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
