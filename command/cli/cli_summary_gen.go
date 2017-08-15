// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/completion"
	"github.com/sacloud/usacloud/command/funcs"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	showParam := params.NewShowSummaryParam()

	cliCommand := &cli.Command{
		Name:  "summary",
		Usage: "Show summary of resource usage",
		Action: func(c *cli.Context) error {
			comm := c.App.Command("show")
			if comm != nil {
				return comm.Action(c)
			}
			return cli.ShowSubcommandHelp(c)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "param-template",
				Usage: "Set input parameter from string(JSON)",
			},
			&cli.StringFlag{
				Name:  "param-template-file",
				Usage: "Set input parameter from file",
			},
			&cli.BoolFlag{
				Name:  "generate-skeleton",
				Usage: "Output skelton of parameter JSON",
			},
			&cli.StringFlag{
				Name:    "output-type",
				Aliases: []string{"out"},
				Usage:   "Output type [json/csv/tsv]",
			},
			&cli.StringSliceFlag{
				Name:    "column",
				Aliases: []string{"col"},
				Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Only display IDs",
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"fmt"},
				Usage:   "Output format(see text/template package document for detail)",
			},
			&cli.StringFlag{
				Name:  "format-file",
				Usage: "Output format from file(see text/template package document for detail)",
			},
			&cli.BoolFlag{
				Name:    "paid-resources-only",
				Aliases: []string{"paid"},
				Usage:   "Show paid-resource only",
				Value:   false,
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "show",
				Usage: "Show Summary (default)",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "param-template",
						Usage: "Set input parameter from string(JSON)",
					},
					&cli.StringFlag{
						Name:  "param-template-file",
						Usage: "Set input parameter from file",
					},
					&cli.BoolFlag{
						Name:  "generate-skeleton",
						Usage: "Output skelton of parameter JSON",
					},
					&cli.StringFlag{
						Name:    "output-type",
						Aliases: []string{"out"},
						Usage:   "Output type [json/csv/tsv]",
					},
					&cli.StringSliceFlag{
						Name:    "column",
						Aliases: []string{"col"},
						Usage:   "Output columns(using when '--output-type' is in [csv/tsv] only)",
					},
					&cli.BoolFlag{
						Name:    "quiet",
						Aliases: []string{"q"},
						Usage:   "Only display IDs",
					},
					&cli.StringFlag{
						Name:    "format",
						Aliases: []string{"fmt"},
						Usage:   "Output format(see text/template package document for detail)",
					},
					&cli.StringFlag{
						Name:  "format-file",
						Usage: "Output format from file(see text/template package document for detail)",
					},
					&cli.BoolFlag{
						Name:    "paid-resources-only",
						Aliases: []string{"paid"},
						Usage:   "Show paid-resource only",
						Value:   false,
					},
				},
				ShellComplete: func(c *cli.Context) {

					if c.NArg() < 3 { // invalid args
						return
					}

					if err := checkConfigVersion(); err != nil {
						return
					}
					if err := applyConfigFromFile(c); err != nil {
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
					command.GlobalOption.Validate(false)

					// build command context
					ctx := command.NewContext(c, realArgs, showParam)

					// Set option values
					if c.IsSet("param-template") {
						showParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						showParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						showParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						showParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						showParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						showParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						showParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						showParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("paid-resources-only") {
						showParam.PaidResourcesOnly = c.Bool("paid-resources-only")
					}

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completion.FlagNames(c, commandName)
								return
							} else {
								completion.SummaryShowCompleteArgs(ctx, showParam, cur, prev, commandName)
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
										completion.FlagNames(c, commandName)
										return
									} else {
										completion.SummaryShowCompleteArgs(ctx, showParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									completion.SummaryShowCompleteFlags(ctx, showParam, name, cur)
									return
								}
							}
						}
						// here, prev is wrong, so noop.
					} else {
						if strings.HasPrefix(cur, "-") {
							completion.FlagNames(c, commandName)
							return
						} else {
							completion.SummaryShowCompleteArgs(ctx, showParam, cur, prev, commandName)
							return
						}
					}
				},
				Action: func(c *cli.Context) error {

					if err := checkConfigVersion(); err != nil {
						return err
					}
					if err := applyConfigFromFile(c); err != nil {
						return err
					}

					showParam.ParamTemplate = c.String("param-template")
					showParam.ParamTemplateFile = c.String("param-template-file")
					strInput, err := command.GetParamTemplateValue(showParam)
					if err != nil {
						return err
					}
					if strInput != "" {
						p := params.NewShowSummaryParam()
						err := json.Unmarshal([]byte(strInput), p)
						if err != nil {
							return fmt.Errorf("Failed to parse JSON: %s", err)
						}
						mergo.MergeWithOverwrite(showParam, p)
					}

					// Set option values
					if c.IsSet("param-template") {
						showParam.ParamTemplate = c.String("param-template")
					}
					if c.IsSet("param-template-file") {
						showParam.ParamTemplateFile = c.String("param-template-file")
					}
					if c.IsSet("generate-skeleton") {
						showParam.GenerateSkeleton = c.Bool("generate-skeleton")
					}
					if c.IsSet("output-type") {
						showParam.OutputType = c.String("output-type")
					}
					if c.IsSet("column") {
						showParam.Column = c.StringSlice("column")
					}
					if c.IsSet("quiet") {
						showParam.Quiet = c.Bool("quiet")
					}
					if c.IsSet("format") {
						showParam.Format = c.String("format")
					}
					if c.IsSet("format-file") {
						showParam.FormatFile = c.String("format-file")
					}
					if c.IsSet("paid-resources-only") {
						showParam.PaidResourcesOnly = c.Bool("paid-resources-only")
					}

					// Validate global params
					if errors := command.GlobalOption.Validate(false); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "GlobalOptions")
					}

					// Generate skeleton
					if showParam.GenerateSkeleton {
						showParam.GenerateSkeleton = false
						showParam.FillValueToSkeleton()
						d, err := json.MarshalIndent(showParam, "", "\t")
						if err != nil {
							return fmt.Errorf("Failed to Marshal JSON: %s", err)
						}
						fmt.Fprintln(command.GlobalOption.Out, string(d))
						return nil
					}

					// Validate specific for each command params
					if errors := showParam.Validate(); len(errors) > 0 {
						return command.FlattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := command.NewContext(c, c.Args().Slice(), showParam)

					// Run command with params
					return funcs.SummaryShow(ctx, showParam)

				},
			},
		},
	}

	// build Category-Resource mapping
	AppendResourceCategoryMap("summary", &schema.Category{
		Key:         "summary",
		DisplayName: "Summary",
		Order:       100,
	})

	// build Category-Command mapping

	AppendCommandCategoryMap("summary", "show", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	AppendFlagCategoryMap("summary", "show", "column", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("summary", "show", "format", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("summary", "show", "format-file", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("summary", "show", "generate-skeleton", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("summary", "show", "output-type", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})
	AppendFlagCategoryMap("summary", "show", "paid-resources-only", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	AppendFlagCategoryMap("summary", "show", "param-template", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("summary", "show", "param-template-file", &schema.Category{
		Key:         "Input",
		DisplayName: "Input options",
		Order:       2147483627,
	})
	AppendFlagCategoryMap("summary", "show", "quiet", &schema.Category{
		Key:         "output",
		DisplayName: "Output options",
		Order:       2147483637,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
