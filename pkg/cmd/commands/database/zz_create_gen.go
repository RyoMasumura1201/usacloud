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

package database

import (
	"github.com/sacloud/usacloud/pkg/cmd/core"
	"github.com/sacloud/usacloud/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func (p *createParameter) CleanupEmptyValue(fs *pflag.FlagSet) {

}

func (p *createParameter) buildFlags(fs *pflag.FlagSet) {

	fs.StringVarP(&p.Zone, "zone", "", p.Zone, "")
	fs.StringVarP(&p.Parameters, "parameters", "", p.Parameters, "Input parameters in JSON format")
	fs.BoolVarP(&p.GenerateSkeleton, "generate-skeleton", "", p.GenerateSkeleton, "Output skeleton of parameters with JSON format (aliases: --skeleton)")
	fs.BoolVarP(&p.AssumeYes, "assumeyes", "y", p.AssumeYes, "Assume that the answer to any question which would be asked is yes")
	fs.StringVarP(&p.OutputType, "output-type", "o", p.OutputType, "Output format: one of the following [table/json/yaml] (aliases: --out)")
	fs.BoolVarP(&p.Quiet, "quiet", "q", p.Quiet, "Output IDs only")
	fs.StringVarP(&p.Format, "format", "", p.Format, "Output format in Go templates (aliases: --fmt)")
	fs.StringVarP(&p.FormatFile, "format-file", "", p.FormatFile, "Output format in Go templates(from file)")
	fs.StringVarP(&p.Query, "query", "", p.Query, "JMESPath query")
	fs.StringVarP(&p.QueryFile, "query-file", "", p.QueryFile, "JMESPath query(from file)")
	fs.StringVarP(&p.Name, "name", "", p.Name, "")
	fs.StringVarP(&p.Description, "description", "", p.Description, "")
	fs.StringSliceVarP(&p.Tags, "tags", "", p.Tags, "")
	fs.VarP(core.NewIDFlag(&p.IconID, &p.IconID), "icon-id", "", "")
	fs.StringVarP(&p.DatabaseType, "database-type", "", p.DatabaseType, "options: [postgresql/postgres/mariadb]")
	fs.StringVarP(&p.PlanID, "plan", "", p.PlanID, "options: [10g/30g/90g/240g/500g/1t]")
	fs.VarP(core.NewIDFlag(&p.SwitchID, &p.SwitchID), "switch-id", "", "")
	fs.StringSliceVarP(&p.IPAddresses, "ip-address", "", p.IPAddresses, "(aliases: --ipaddress)")
	fs.IntVarP(&p.NetworkMaskLen, "network-mask-len", "", p.NetworkMaskLen, "")
	fs.StringVarP(&p.DefaultRoute, "default-route", "", p.DefaultRoute, "")
	fs.IntVarP(&p.Port, "port", "", p.Port, "")
	fs.StringSliceVarP(&p.SourceNetwork, "source-network", "", p.SourceNetwork, "")
	fs.StringVarP(&p.Username, "username", "", p.Username, "")
	fs.StringVarP(&p.Password, "password", "", p.Password, "")
	fs.BoolVarP(&p.EnableReplication, "enable-replication", "", p.EnableReplication, "")
	fs.StringVarP(&p.ReplicaUserPassword, "replica-user-password", "", p.ReplicaUserPassword, "")
	fs.BoolVarP(&p.EnableWebUI, "enable-web-ui", "", p.EnableWebUI, "")
	fs.BoolVarP(&p.EnableBackup, "enable-backup", "", p.EnableBackup, "")
	fs.StringSliceVarP(&p.BackupWeekdays, "backup-weekdays", "", p.BackupWeekdays, "options: [all/sun/mon/tue/wed/thu/fri/sat]")
	fs.IntVarP(&p.BackupStartTimeHour, "backup-start-time-hour", "", p.BackupStartTimeHour, "")
	fs.IntVarP(&p.BackupStartTimeMinute, "backup-start-time-minute", "", p.BackupStartTimeMinute, "options: [0/15/30/45]")
	fs.BoolVarP(&p.NoWait, "no-wait", "", p.NoWait, "")
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
		name = "ip-address"
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
		fs = pflag.NewFlagSet("plan", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("database-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("plan"))
		sets = append(sets, &core.FlagSet{
			Title: "Plan options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("WebUI", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("enable-web-ui"))
		sets = append(sets, &core.FlagSet{
			Title: "WebUI options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("backup", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("enable-backup"))
		fs.AddFlag(cmd.LocalFlags().Lookup("backup-weekdays"))
		fs.AddFlag(cmd.LocalFlags().Lookup("backup-start-time-hour"))
		fs.AddFlag(cmd.LocalFlags().Lookup("backup-start-time-minute"))
		sets = append(sets, &core.FlagSet{
			Title: "Backup options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("network", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("switch-id"))
		fs.AddFlag(cmd.LocalFlags().Lookup("ip-address"))
		fs.AddFlag(cmd.LocalFlags().Lookup("network-mask-len"))
		fs.AddFlag(cmd.LocalFlags().Lookup("default-route"))
		fs.AddFlag(cmd.LocalFlags().Lookup("port"))
		fs.AddFlag(cmd.LocalFlags().Lookup("source-network"))
		sets = append(sets, &core.FlagSet{
			Title: "Network options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("replication", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("enable-replication"))
		fs.AddFlag(cmd.LocalFlags().Lookup("replica-user-password"))
		sets = append(sets, &core.FlagSet{
			Title: "Replication options",
			Flags: fs,
		})
	}
	{
		var fs *pflag.FlagSet
		fs = pflag.NewFlagSet("user", pflag.ContinueOnError)
		fs.SortFlags = false
		fs.AddFlag(cmd.LocalFlags().Lookup("username"))
		fs.AddFlag(cmd.LocalFlags().Lookup("password"))
		sets = append(sets, &core.FlagSet{
			Title: "User options",
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
		fs.AddFlag(cmd.LocalFlags().Lookup("format-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("output-type"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query"))
		fs.AddFlag(cmd.LocalFlags().Lookup("query-file"))
		fs.AddFlag(cmd.LocalFlags().Lookup("quiet"))
		sets = append(sets, &core.FlagSet{
			Title: "Output options",
			Flags: fs,
		})
	}

	core.BuildFlagsUsage(cmd, sets)
}

func (p *createParameter) setCompletionFunc(cmd *cobra.Command) {
	cmd.RegisterFlagCompletionFunc("database-type", util.FlagCompletionFunc("postgresql", "postgres", "mariadb"))
	cmd.RegisterFlagCompletionFunc("plan", util.FlagCompletionFunc("10g", "30g", "90g", "240g", "500g", "1t"))
	cmd.RegisterFlagCompletionFunc("backup-weekdays", util.FlagCompletionFunc("all", "sun", "mon", "tue", "wed", "thu", "fri", "sat"))
	cmd.RegisterFlagCompletionFunc("backup-start-time-minute", util.FlagCompletionFunc("0", "15", "30", "45"))

}

func (p *createParameter) SetupCobraCommandFlags(cmd *cobra.Command) {
	p.buildFlags(cmd.Flags())
	p.buildFlagsUsage(cmd)
	p.setCompletionFunc(cmd)
}
