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

package proxylb

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.Example, "example", "", p.Example, "Output example parameters with JSON format")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format options: [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "Query for JSON output")
	fs.StringVarP(&p.QueryDriver, "query-driver", "", p.QueryDriver, "Name of the driver that handles queries to JSON output options: [jmespath/jq]")
	fs.StringVarP(&p.Name, "name", "", p.Name, "(*required) ")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.StringVarP(&p.Plan, "plan", "", p.Plan, "(*required) options: [100/500/1000/5000/10000/50000/100000]")
	fs.StringVarP(&p.HealthCheck.Protocol, "health-check-protocol", "", p.HealthCheck.Protocol, "(*required) ")
	fs.StringVarP(&p.HealthCheck.Path, "health-check-path", "", p.HealthCheck.Path, "")
	fs.StringVarP(&p.HealthCheck.Host, "health-check-host", "", p.HealthCheck.Host, "")
	fs.IntVarP(&p.HealthCheck.DelayLoop, "health-check-delay-loop", "", p.HealthCheck.DelayLoop, "(*required) ")
	fs.StringVarP(&p.SorryServer.IPAddress, "sorry-server-ip-address", "", p.SorryServer.IPAddress, "(aliases: --ipaddress)")
	fs.IntVarP(&p.SorryServer.Port, "sorry-server-port", "", p.SorryServer.Port, "")
	fs.StringVarP(&p.LetsEncrypt.CommonName, "lets-encrypt-common-name", "", p.LetsEncrypt.CommonName, "")
	fs.BoolVarP(&p.LetsEncrypt.Enabled, "lets-encrypt-enabled", "", p.LetsEncrypt.Enabled, "")
	fs.BoolVarP(&p.LetsEncrypt.AcceptTOS, "lets-encrypt-accept-tos", "", p.LetsEncrypt.AcceptTOS, "The flag to accept the current Let's Encrypt terms of service(see: https://letsencrypt.org/repository/)")
	fs.StringVarP(&p.StickySession.Method, "sticky-session-method", "", p.StickySession.Method, "")
	fs.BoolVarP(&p.StickySession.Enabled, "sticky-session-enabled", "", p.StickySession.Enabled, "")
	fs.IntVarP(&p.Timeout.InactiveSec, "inactive-sec", "", p.Timeout.InactiveSec, "")
	fs.BoolVarP(&p.UseVIPFailover, "vip-fail-over", "", p.UseVIPFailover, "")
	fs.StringVarP(&p.Region, "region", "", p.Region, "(*required) options: [tk1/is1/anycast]")
	fs.StringVarP(&p.BindPortsData, "bind-ports", "", p.BindPortsData, "")
	fs.StringVarP(&p.ServersData, "servers", "", p.ServersData, "")
	fs.StringVarP(&p.RulesData, "rules", "", p.RulesData, "")
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
	case "ipaddress":
		name = "sorry-server-ip-address"
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
		fs = pflag.NewFlagSet("proxy-lb", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("bind-ports"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-delay-loop"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-host"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-path"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-protocol"))
		fs.AddFlag(cmd.LocalFlags().Lookup("inactive-sec"))
		fs.AddFlag(cmd.LocalFlags().Lookup("lets-encrypt-accept-tos"))
		fs.AddFlag(cmd.LocalFlags().Lookup("lets-encrypt-common-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("lets-encrypt-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("plan"))
		fs.AddFlag(cmd.LocalFlags().Lookup("region"))
		fs.AddFlag(cmd.LocalFlags().Lookup("rules"))
		fs.AddFlag(cmd.LocalFlags().Lookup("servers"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sorry-server-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sorry-server-port"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sticky-session-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sticky-session-method"))
		fs.AddFlag(cmd.LocalFlags().Lookup("vip-fail-over"))
		sets = append(sets, &core.FlagSet{
			Title: "Proxy-Lb-specific options",
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

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("plan", util.FlagCompletionFunc("100", "500", "1000", "5000", "10000", "50000", "100000"))
	cmd.RegisterFlagCompletionFunc("region", util.FlagCompletionFunc("tk1", "is1", "anycast"))

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
