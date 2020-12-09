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

package simplemonitor

import (
	"github.com/sacloud/libsacloud/v2/sacloud/pointer"
	"github.com/sacloud/libsacloud/v2/sacloud/types"
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *updateParameter) CleanupEmptyValue(fs *pflag.FlagSet) {
	if !fs.Changed("description") {
		p.Description = nil
	}
	if !fs.Changed("tags") {
		p.Tags = nil
	}
	if !fs.Changed("icon-id") {
		p.IconID = nil
	}
	if !fs.Changed("delay-loop") {
		p.DelayLoop = nil
	}
	if !fs.Changed("enabled") {
		p.Enabled = nil
	}
	if !fs.Changed("health-check-protocol") {
		p.HealthCheck.Protocol = nil
	}
	if !fs.Changed("health-check-port") {
		p.HealthCheck.Port = nil
	}
	if !fs.Changed("health-check-path") {
		p.HealthCheck.Path = nil
	}
	if !fs.Changed("health-check-status") {
		p.HealthCheck.Status = nil
	}
	if !fs.Changed("health-check-sni") {
		p.HealthCheck.SNI = nil
	}
	if !fs.Changed("health-check-host") {
		p.HealthCheck.Host = nil
	}
	if !fs.Changed("health-check-basic-auth-username") {
		p.HealthCheck.BasicAuthUsername = nil
	}
	if !fs.Changed("health-check-basic-auth-password") {
		p.HealthCheck.BasicAuthPassword = nil
	}
	if !fs.Changed("health-check-q-name") {
		p.HealthCheck.QName = nil
	}
	if !fs.Changed("health-check-expected-data") {
		p.HealthCheck.ExpectedData = nil
	}
	if !fs.Changed("health-check-community") {
		p.HealthCheck.Community = nil
	}
	if !fs.Changed("health-check-snmp-version") {
		p.HealthCheck.SNMPVersion = nil
	}
	if !fs.Changed("health-check-oid") {
		p.HealthCheck.OID = nil
	}
	if !fs.Changed("health-check-remaining-days") {
		p.HealthCheck.RemainingDays = nil
	}
	if !fs.Changed("notify-email-enabled") {
		p.NotifyEmailEnabled = nil
	}
	if !fs.Changed("notify-email-html") {
		p.NotifyEmailHTML = nil
	}
	if !fs.Changed("notify-slack-enabled") {
		p.NotifySlackEnabled = nil
	}
	if !fs.Changed("slack-webhooks-url") {
		p.SlackWebhooksURL = nil
	}
	if !fs.Changed("notify-interval") {
		p.NotifyInterval = nil
	}
}

func (p *updateParameter) buildFlags(fs *pflag.FlagSet) {
	if p.Description == nil {
		p.Description = pointer.NewString("")
	}
	if p.Tags == nil {
		p.Tags = pointer.NewStringSlice([]string{})
	}
	if p.IconID == nil {
		p.IconID = pointer.NewID(types.ID(0))
	}
	if p.DelayLoop == nil {
		p.DelayLoop = pointer.NewInt(0)
	}
	if p.Enabled == nil {
		p.Enabled = pointer.NewBool(false)
	}
	if p.HealthCheck.Protocol == nil {
		p.HealthCheck.Protocol = pointer.NewString("")
	}
	if p.HealthCheck.Port == nil {
		p.HealthCheck.Port = pointer.NewInt(0)
	}
	if p.HealthCheck.Path == nil {
		p.HealthCheck.Path = pointer.NewString("")
	}
	if p.HealthCheck.Status == nil {
		p.HealthCheck.Status = pointer.NewInt(0)
	}
	if p.HealthCheck.SNI == nil {
		p.HealthCheck.SNI = pointer.NewBool(false)
	}
	if p.HealthCheck.Host == nil {
		p.HealthCheck.Host = pointer.NewString("")
	}
	if p.HealthCheck.BasicAuthUsername == nil {
		p.HealthCheck.BasicAuthUsername = pointer.NewString("")
	}
	if p.HealthCheck.BasicAuthPassword == nil {
		p.HealthCheck.BasicAuthPassword = pointer.NewString("")
	}
	if p.HealthCheck.QName == nil {
		p.HealthCheck.QName = pointer.NewString("")
	}
	if p.HealthCheck.ExpectedData == nil {
		p.HealthCheck.ExpectedData = pointer.NewString("")
	}
	if p.HealthCheck.Community == nil {
		p.HealthCheck.Community = pointer.NewString("")
	}
	if p.HealthCheck.SNMPVersion == nil {
		p.HealthCheck.SNMPVersion = pointer.NewString("")
	}
	if p.HealthCheck.OID == nil {
		p.HealthCheck.OID = pointer.NewString("")
	}
	if p.HealthCheck.RemainingDays == nil {
		p.HealthCheck.RemainingDays = pointer.NewInt(0)
	}
	if p.NotifyEmailEnabled == nil {
		p.NotifyEmailEnabled = pointer.NewBool(false)
	}
	if p.NotifyEmailHTML == nil {
		p.NotifyEmailHTML = pointer.NewBool(false)
	}
	if p.NotifySlackEnabled == nil {
		p.NotifySlackEnabled = pointer.NewString("")
	}
	if p.SlackWebhooksURL == nil {
		p.SlackWebhooksURL = pointer.NewString("")
	}
	if p.NotifyInterval == nil {
		p.NotifyInterval = pointer.NewInt(0)
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
	fs.StringVarP(p.Description, "description", "", "", "")
	fs.StringSliceVarP(p.Tags, "tags", "", nil, "")
	fs.VarP(core.NewIDFlag(p.IconID, p.IconID), "icon-id", "", "")
	fs.IntVarP(p.DelayLoop, "delay-loop", "", 0, "")
	fs.BoolVarP(p.Enabled, "enabled", "", false, "")
	fs.StringVarP(p.HealthCheck.Protocol, "health-check-protocol", "", "", "options: [http/https/ping/tcp/dns/ssh/smtp/pop3/snmp/sslcertificate]")
	fs.IntVarP(p.HealthCheck.Port, "health-check-port", "", 0, "")
	fs.StringVarP(p.HealthCheck.Path, "health-check-path", "", "", "")
	fs.IntVarP(p.HealthCheck.Status, "health-check-status", "", 0, "")
	fs.BoolVarP(p.HealthCheck.SNI, "health-check-sni", "", false, "")
	fs.StringVarP(p.HealthCheck.Host, "health-check-host", "", "", "")
	fs.StringVarP(p.HealthCheck.BasicAuthUsername, "health-check-basic-auth-username", "", "", "")
	fs.StringVarP(p.HealthCheck.BasicAuthPassword, "health-check-basic-auth-password", "", "", "")
	fs.StringVarP(p.HealthCheck.QName, "health-check-q-name", "", "", "")
	fs.StringVarP(p.HealthCheck.ExpectedData, "health-check-expected-data", "", "", "")
	fs.StringVarP(p.HealthCheck.Community, "health-check-community", "", "", "")
	fs.StringVarP(p.HealthCheck.SNMPVersion, "health-check-snmp-version", "", "", "")
	fs.StringVarP(p.HealthCheck.OID, "health-check-oid", "", "", "")
	fs.IntVarP(p.HealthCheck.RemainingDays, "health-check-remaining-days", "", 0, "")
	fs.BoolVarP(p.NotifyEmailEnabled, "notify-email-enabled", "", false, "")
	fs.BoolVarP(p.NotifyEmailHTML, "notify-email-html", "", false, "")
	fs.StringVarP(p.NotifySlackEnabled, "notify-slack-enabled", "", "", "")
	fs.StringVarP(p.SlackWebhooksURL, "slack-webhooks-url", "", "", "")
	fs.IntVarP(p.NotifyInterval, "notify-interval", "", 0, "")
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
		fs = pflag.NewFlagSet("simple-monitor", pflag.ContinueOnError)
		fs.AddFlag(cmd.LocalFlags().Lookup("description"))
		fs.AddFlag(cmd.LocalFlags().Lookup("tags"))
		fs.AddFlag(cmd.LocalFlags().Lookup("icon-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("delay-loop"))
		fs.AddFlag(cmd.LocalFlags().Lookup("enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-protocol"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-port"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-path"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-status"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-sni"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-host"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-basic-auth-username"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-basic-auth-password"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-q-name"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-expected-data"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-community"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-snmp-version"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-oid"))
		fs.AddFlag(cmd.LocalFlags().Lookup("health-check-remaining-days"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-email-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-email-html"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-slack-enabled"))
		fs.AddFlag(cmd.LocalFlags().Lookup("slack-webhooks-url"))
		fs.AddFlag(cmd.LocalFlags().Lookup("notify-interval"))
		sets = append(sets, &core.FlagSet{
			Title: "Simple-Monitor options",
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
	cmd.RegisterFlagCompletionFunc("health-check-protocol", util.FlagCompletionFunc("http", "https", "ping", "tcp", "dns", "ssh", "smtp", "pop3", "snmp", "sslcertificate"))

}

func (p *updateParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
