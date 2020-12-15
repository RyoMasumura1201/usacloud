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

// Code generated by 'github.com/sacloud/usacloud/tools/gen-commands'; DO NOT EDIT

package mobilegateway

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "(*required) ")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.Name, "name", "", p.Name, "(*required) ")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.BoolVarP(&p.InternetConnectionEnabled, "internet-connection-enabled", "", p.InternetConnectionEnabled, "")
	fs.BoolVarP(&p.InterDeviceCommunicationEnabled, "inter-device-communication-enabled", "", p.InterDeviceCommunicationEnabled, "")
	fs.StringVarP(&p.SIMsData, "sims", "", p.SIMsData, "")
	fs.StringVarP(&p.SIMRoutesData, "sim-routes", "", p.SIMRoutesData, "")
	fs.StringVarP(&p.StaticRoutesData, "static-routes", "", p.StaticRoutesData, "")
	fs.VarP(core.NewIDFlag(&p.PrivateInterface.SwitchID, &p.PrivateInterface.SwitchID), "private-interface-switch-id", "", "")
	fs.StringVarP(&p.PrivateInterface.IPAddress, "private-interface-ip-address", "", p.PrivateInterface.IPAddress, "(*required) ")
	fs.IntVarP(&p.PrivateInterface.NetworkMaskLen, "private-interface-network-mask-len", "", p.PrivateInterface.NetworkMaskLen, "")
	fs.StringVarP(&p.DNS.DNS1, "dns1", "", p.DNS.DNS1, "(*required) ")
	fs.StringVarP(&p.DNS.DNS2, "dns2", "", p.DNS.DNS2, "(*required) ")
	fs.IntVarP(&p.TrafficConfig.TrafficQuotaInMB, "traffic-config-traffic-quota-in-mb", "", p.TrafficConfig.TrafficQuotaInMB, "")
	fs.IntVarP(&p.TrafficConfig.BandWidthLimitInKbps, "traffic-config-band-width-limit-in-kbps", "", p.TrafficConfig.BandWidthLimitInKbps, "")
	fs.BoolVarP(&p.TrafficConfig.EmailNotifyEnabled, "traffic-config-email-notify-enabled", "", p.TrafficConfig.EmailNotifyEnabled, "")
	fs.BoolVarP(&p.TrafficConfig.SlackNotifyEnabled, "traffic-config-slack-notify-enabled", "", p.TrafficConfig.SlackNotifyEnabled, "")
	fs.StringVarP(&p.TrafficConfig.SlackNotifyWebhooksURL, "traffic-config-slack-notify-webhooks-url", "", p.TrafficConfig.SlackNotifyWebhooksURL, "")
	fs.BoolVarP(&p.TrafficConfig.AutoTrafficShaping, "traffic-config-auto-traffic-shaping", "", p.TrafficConfig.AutoTrafficShaping, "")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
	fs.BoolVarP(&p.BootAfterCreate, "boot-after-create", "", p.BootAfterCreate, "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *createParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
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

func (p *createParameter) buildFlagsUsage(cmd *cobra.Command) {
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
		fs = pflag.NewFlagSet("mobile-gateway", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("boot-after-create"))
		fs.AddFlag(cmd.LocalFlags().Lookup("dns1"))
		fs.AddFlag(cmd.LocalFlags().Lookup("dns2"))
		fs.AddFlag(cmd.LocalFlags().Lookup("inter-device-communication-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("internet-connection-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sim-routes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sims"))
		fs.AddFlag(cmd.LocalFlags().Lookup("static-routes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-auto-traffic-shaping"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-band-width-limit-in-kbps"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-email-notify-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-slack-notify-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-slack-notify-webhooks-url"))
		fs.AddFlag(cmd.LocalFlags().Lookup("traffic-config-traffic-quota-in-mb"))
		sets = append(sets, &core.FlagSet{
			Title: "Mobile-Gateway-specific options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("network", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("private-interface-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-interface-network-mask-len"))
		fs.AddFlag(cmd.LocalFlags().Lookup("private-interface-switch-id"))
		sets = append(sets, &core.FlagSet{
			Title: "Network options",
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

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
