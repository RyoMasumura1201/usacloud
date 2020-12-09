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

package gslb

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
	if !fs.Changed("health-check-protocol") {
		p.HealthCheck.Protocol = nil
	}
	if !fs.Changed("health-check-host-header") {
		p.HealthCheck.HostHeader = nil
	}
	if !fs.Changed("health-check-path") {
		p.HealthCheck.Path = nil
	}
	if !fs.Changed("health-check-response-code") {
		p.HealthCheck.ResponseCode = nil
	}
	if !fs.Changed("health-check-port") {
		p.HealthCheck.Port = nil
	}
	if !fs.Changed("delay-loop") {
		p.DelayLoop = nil
	}
	if !fs.Changed("weighted") {
		p.Weighted = nil
	}
	if !fs.Changed("sorry-server") {
		p.SorryServer = nil
	}
	if !fs.Changed("servers") {
		p.ServersData = nil
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
	if p.HealthCheck.Protocol == nil {
		p.HealthCheck.Protocol = pointer.NewString("")
	}
	if p.HealthCheck.HostHeader == nil {
		p.HealthCheck.HostHeader = pointer.NewString("")
	}
	if p.HealthCheck.Path == nil {
		p.HealthCheck.Path = pointer.NewString("")
	}
	if p.HealthCheck.ResponseCode == nil {
		p.HealthCheck.ResponseCode = pointer.NewInt(0)
	}
	if p.HealthCheck.Port == nil {
		p.HealthCheck.Port = pointer.NewInt(0)
	}
	if p.DelayLoop == nil {
		p.DelayLoop = pointer.NewInt(0)
	}
	if p.Weighted == nil {
		p.Weighted = pointer.NewBool(false)
	}
	if p.SorryServer == nil {
		p.SorryServer = pointer.NewString("")
	}
	if p.ServersData == nil {
		p.ServersData = pointer.NewString("")
	}
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.FormatFile, "format-file", "", p.FormatFile, "Output format in Go templates(from file)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.QueryFile, "query-file", "", p.QueryFile, "JMESPath query(from file)")
	fs.StringVarP(p.Name, "name", "", "", "")
	fs.StringVarP(p.Description, "description", "", "", "")
	fs.StringSliceVarP(p.Tags, "tags", "", nil, "")
	fs.VarP(core.NewIDFlag(p.IconID, p.IconID), "icon-id", "", "")
	fs.StringVarP(p.HealthCheck.Protocol, "health-check-protocol", "", "", "")
	fs.StringVarP(p.HealthCheck.HostHeader, "health-check-host-header", "", "", "")
	fs.StringVarP(p.HealthCheck.Path, "health-check-path", "", "", "")
	fs.IntVarP(p.HealthCheck.ResponseCode, "health-check-response-code", "", 0, "")
	fs.IntVarP(p.HealthCheck.Port, "health-check-port", "", 0, "")
	fs.IntVarP(p.DelayLoop, "delay-loop", "", 0, "")
	fs.BoolVarP(p.Weighted, "weighted", "", false, "")
	fs.StringVarP(p.SorryServer, "sorry-server", "", "", "")
	fs.StringVarP(p.ServersData, "servers", "", "", "")
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
	}
	return pflag.NormalizedName(name)
}

func (p *updateParameter) buildFlagsUsage(cmd *cobra.Command) {
	var sets []*core.FlagSet
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("gslb", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-protocol"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-host-header"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-path"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-response-code"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-port"))
		fs.AddFlag(cmd.LocalFlags().Lookup("delay-loop"))
		fs.AddFlag(cmd.LocalFlags().Lookup("weighted"))
		fs.AddFlag(cmd.LocalFlags().Lookup("sorry-server"))
		fs.AddFlag(cmd.LocalFlags().Lookup("servers"))
		sets = append(sets, &core.FlagSet{
			Title: "Gslb options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("Input", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("parameters"))
		fs.AddFlag(cmd.LocalFlags().Lookup("generate-skeleton"))
		fs.AddFlag(cmd.LocalFlags().Lookup("assumeyes"))
		sets = append(sets, &core.FlagSet{
			Title: "Input options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("output", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format"))
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
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
