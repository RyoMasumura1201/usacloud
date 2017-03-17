// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	createParam := NewCreateInternetParam()
	deleteParam := NewDeleteInternetParam()
	listParam := NewListInternetParam()
	monitorParam := NewMonitorInternetParam()
	readParam := NewReadInternetParam()
	updateParam := NewUpdateInternetParam()
	updateBandwidthParam := NewUpdateBandwidthInternetParam()

	cliCommand := &cli.Command{
		Name:  "internet",
		Usage: "A manage commands of Internet",
		Subcommands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create Internet",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "band-width",
						Usage:       "[Required] set band-width(Mbpm)",
						Value:       100,
						Destination: &createParam.BandWidth,
					},
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &createParam.Description,
					},
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &createParam.IconId,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
					},
					&cli.IntFlag{
						Name:        "nw-masklen",
						Aliases:     []string{"network-masklen"},
						Usage:       "[Required] set Global-IPAddress prefix",
						Value:       28,
						Destination: &createParam.NwMasklen,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, createParam)

					// Set option values for slice
					createParam.Tags = c.StringSlice("tags")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetCreateCompleteFlags(ctx, createParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					createParam.Tags = c.StringSlice("tags")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := createParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), createParam)

					// Run command with params
					return InternetCreate(ctx, createParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete Internet",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "force",
						Aliases:     []string{"f"},
						Destination: &deleteParam.Force,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &deleteParam.Id,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, deleteParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetDeleteCompleteFlags(ctx, deleteParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := deleteParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), deleteParam)

					// confirm
					if !deleteParam.Force && !confirmContinue("delete this") {
						return nil
					}

					// Run command with params
					return InternetDelete(ctx, deleteParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List Internet",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "from",
						Usage:       "set offset",
						Destination: &listParam.From,
					},
					&cli.Int64SliceFlag{
						Name:  "id",
						Usage: "set filter by id(s)",
					},
					&cli.IntFlag{
						Name:        "max",
						Usage:       "set limit",
						Destination: &listParam.Max,
					},
					&cli.StringSliceFlag{
						Name:  "name",
						Usage: "set filter by name(s)",
					},
					&cli.StringSliceFlag{
						Name:  "sort",
						Usage: "set field(s) for sort",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, listParam)

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetListCompleteArgs(ctx, listParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetListCompleteFlags(ctx, listParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetListCompleteArgs(ctx, listParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					listParam.Id = c.Int64Slice("id")
					listParam.Name = c.StringSlice("name")
					listParam.Sort = c.StringSlice("sort")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Validate specific for each command params
					if errors := listParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), listParam)

					// Run command with params
					return InternetList(ctx, listParam)
				},
			},
			{
				Name:      "monitor",
				Usage:     "Monitor Internet",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "end",
						Usage:       "set end-time",
						Destination: &monitorParam.End,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &monitorParam.Id,
					},
					&cli.StringFlag{
						Name:        "key-format",
						Usage:       "[Required] set monitoring value key-format",
						Value:       "sakuracloud.{{.ID}}.internet",
						Destination: &monitorParam.KeyFormat,
					},
					&cli.StringFlag{
						Name:        "start",
						Usage:       "set start-time",
						Destination: &monitorParam.Start,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, monitorParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetMonitorCompleteArgs(ctx, monitorParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetMonitorCompleteArgs(ctx, monitorParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetMonitorCompleteFlags(ctx, monitorParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetMonitorCompleteArgs(ctx, monitorParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := monitorParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), monitorParam)

					// Run command with params
					return InternetMonitor(ctx, monitorParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read Internet",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &readParam.Id,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, readParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetReadCompleteArgs(ctx, readParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetReadCompleteArgs(ctx, readParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetReadCompleteFlags(ctx, readParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetReadCompleteArgs(ctx, readParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := readParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), readParam)

					// Run command with params
					return InternetRead(ctx, readParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update Internet",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "band-width",
						Usage:       "set band-width(Mbpm)",
						Destination: &updateParam.BandWidth,
					},
					&cli.StringFlag{
						Name:        "description",
						Aliases:     []string{"desc"},
						Usage:       "set resource description",
						Destination: &updateParam.Description,
					},
					&cli.Int64Flag{
						Name:        "icon-id",
						Usage:       "set Icon ID",
						Destination: &updateParam.IconId,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &updateParam.Id,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "set resource display name",
						Destination: &updateParam.Name,
					},
					&cli.StringSliceFlag{
						Name:  "tags",
						Usage: "set resource tags",
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, updateParam)

					// Set option values for slice
					updateParam.Tags = c.StringSlice("tags")

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetUpdateCompleteFlags(ctx, updateParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Set option values for slice
					updateParam.Tags = c.StringSlice("tags")

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := updateParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), updateParam)

					// Run command with params
					return InternetUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "update-bandwidth",
				Usage:     "UpdateBandwidth Internet",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "band-width",
						Usage:       "[Required] set band-width(Mbpm)",
						Value:       100,
						Destination: &updateBandwidthParam.BandWidth,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &updateBandwidthParam.Id,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					// c.Args() == arg1 arg2 arg3 -- [cur] [prev] [commandName]
					args := c.Args().Slice()
					commandName := args[c.NArg()-1]
					prev := args[c.NArg()-2]
					cur := args[c.NArg()-3]

					// set real args
					realArgs := args[0 : c.NArg()-3]

					// Validate global params
					GlobalOption.Validate(false)

					// build command context
					ctx := NewContext(c, realArgs, updateBandwidthParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								InternetUpdateBandwidthCompleteArgs(ctx, updateBandwidthParam, cur, prev, commandName)
								return
							}
						}

						// cleanup flag name
						name := prev
						for {
							if !strings.HasPrefix(name, "-") {
								break
							}
							name = strings.Replace(name, "-", "", 1)
						}

						// flag is exists? , is BoolFlag?
						exists := false
						for _, flag := range c.App.Command(commandName).Flags {

							for _, n := range flag.Names() {
								if n == name {
									exists = true
									break
								}
							}

							if exists {
								if _, ok := flag.(*cli.BoolFlag); ok {
									if strings.HasPrefix(cur, "-") {
										completionFlagNames(c, commandName)
										return
									} else {
										InternetUpdateBandwidthCompleteArgs(ctx, updateBandwidthParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									InternetUpdateBandwidthCompleteFlags(ctx, updateBandwidthParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completionFlagNames(c, commandName)
							return
						} else {
							InternetUpdateBandwidthCompleteArgs(ctx, updateBandwidthParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					// Validate global params
					if errors := GlobalOption.Validate(false); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// id is can set from option or args(first)
					if c.NArg() > 0 {
						c.Set("id", c.Args().First())
					}

					// Validate specific for each command params
					if errors := updateBandwidthParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), updateBandwidthParam)

					// Run command with params
					return InternetUpdateBandwidth(ctx, updateBandwidthParam)
				},
			},
		},
	}

	// build Category-Resource mapping
	appendResourceCategoryMap("internet", &schema.Category{
		Key:         "networking",
		DisplayName: "Networking",
		Order:       30,
	})

	// build Category-Command mapping

	appendCommandCategoryMap("internet", "create", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("internet", "delete", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("internet", "list", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("internet", "monitor", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("internet", "read", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("internet", "update", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("internet", "update-bandwidth", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	appendFlagCategoryMap("internet", "create", "band-width", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "create", "description", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "create", "icon-id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "create", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "create", "nw-masklen", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "create", "tags", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "delete", "force", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "delete", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "list", "from", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "list", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "list", "max", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "list", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "list", "sort", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "monitor", "end", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "monitor", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "monitor", "key-format", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "monitor", "start", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "read", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update", "band-width", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update", "description", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update", "icon-id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update", "tags", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update-bandwidth", "band-width", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("internet", "update-bandwidth", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
