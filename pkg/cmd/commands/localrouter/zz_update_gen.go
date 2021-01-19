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

package localrouter

import (
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *updateParameter) CleanupEmptyValue(fs *pflag.FlagSet) {
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
	if !fs.Changed("switch-code") {
		p.Switch.Code = nil
	}
	if !fs.Changed("switch-category") {
		p.Switch.Category = nil
	}
	if !fs.Changed("switch-zone-id") {
		p.Switch.ZoneID = nil
	}
	if !fs.Changed("virtual-ip-address") {
		p.Interface.VirtualIPAddress = nil
	}
	if !fs.Changed("ip-addresses") {
		p.Interface.IPAddress = nil
	}
	if !fs.Changed("netmask") {
		p.Interface.NetworkMaskLen = nil
	}
	if !fs.Changed("vrid") {
		p.Interface.VRID = nil
	}
	if !fs.Changed("peers") {
		p.PeersData = nil
	}
	if !fs.Changed("static-routes") {
		p.StaticRoutesData = nil
	}
}

func (p *updateParameter) buildFlags(fs *pflag.FlagSet) {
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
	if p.Switch.Code == nil {
		p.Switch.Code = pointer.NewString("")
	}
	if p.Switch.Category == nil {
		p.Switch.Category = pointer.NewString("")
	}
	if p.Switch.ZoneID == nil {
		p.Switch.ZoneID = pointer.NewString("")
	}
	if p.Interface.VirtualIPAddress == nil {
		p.Interface.VirtualIPAddress = pointer.NewString("")
	}
	if p.Interface.IPAddress == nil {
		p.Interface.IPAddress = pointer.NewStringSlice([]string{})
	}
	if p.Interface.NetworkMaskLen == nil {
		p.Interface.NetworkMaskLen = pointer.NewInt(0)
	}
	if p.Interface.VRID == nil {
		p.Interface.VRID = pointer.NewInt(0)
	}
	if p.PeersData == nil {
		p.PeersData = pointer.NewString("")
	}
	if p.StaticRoutesData == nil {
		p.StaticRoutesData = pointer.NewString("")
	}
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
	fs.StringVarP(p.Switch.Code, "switch-code", "", "", "")
	fs.StringVarP(p.Switch.Category, "switch-category", "", "", "")
	fs.StringVarP(p.Switch.ZoneID, "switch-zone-id", "", "", "")
	fs.StringVarP(p.Interface.VirtualIPAddress, "virtual-ip-address", "", "", "")
	fs.StringSliceVarP(p.Interface.IPAddress, "ip-addresses", "", nil, "")
	fs.IntVarP(p.Interface.NetworkMaskLen, "netmask", "", 0, "(aliases: --network-mask-len)")
	fs.IntVarP(p.Interface.VRID, "vrid", "", 0, "")
	fs.StringVarP(p.PeersData, "peers", "", "", "")
	fs.StringVarP(p.StaticRoutesData, "static-routes", "", "", "")
	fs.SetNormalizeFunc(p.normalizeFlagName)
}

func (p *updateParameter) normalizeFlagName(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "skeleton":
		name = "generate-skeleton"
	case "out":
		name = "output-type"
	case "fmt":
		name = "format"
	case "network-mask-len":
		name = "netmask"
	}
	return pflag.NormalizedName(name)
}

func (p *updateParameter) buildFlagsUsage(cmd *cobra.Command) {
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
		fs = pflag.NewFlagSet("local-router", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("ip-addresses"))
		fs.AddFlag(cmd.LocalFlags().Lookup("netmask"))
		fs.AddFlag(cmd.LocalFlags().Lookup("peers"))
		fs.AddFlag(cmd.LocalFlags().Lookup("static-routes"))
		fs.AddFlag(cmd.LocalFlags().Lookup("switch-category"))
		fs.AddFlag(cmd.LocalFlags().Lookup("switch-code"))
		fs.AddFlag(cmd.LocalFlags().Lookup("switch-zone-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("virtual-ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("vrid"))
		sets = append(sets, &core.FlagSet{
			Title: "Local-Router-specific options",
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

func (p *updateParameter) setCompletionFunc(cmd *cobra.Command) {

}

func (p *updateParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
