// Code generated by 'github.com/sacloud/usacloud/tools/gen-cli-commands'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/schema"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func init() {
	createParam := NewCreateISOImageParam()
	deleteParam := NewDeleteISOImageParam()
	downloadParam := NewDownloadISOImageParam()
	ftpCloseParam := NewFtpCloseISOImageParam()
	ftpOpenParam := NewFtpOpenISOImageParam()
	listParam := NewListISOImageParam()
	readParam := NewReadISOImageParam()
	updateParam := NewUpdateISOImageParam()
	uploadParam := NewUploadISOImageParam()

	cliCommand := &cli.Command{
		Name:  "iso-image",
		Usage: "A manage commands of ISOImage",
		Subcommands: []*cli.Command{
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create ISOImage",
				Flags: []cli.Flag{
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
						Name:        "iso-file",
						Usage:       "[Required] set iso image file",
						Destination: &createParam.IsoFile,
					},
					&cli.StringFlag{
						Name:        "name",
						Usage:       "[Required] set resource display name",
						Destination: &createParam.Name,
					},
					&cli.IntFlag{
						Name:        "size",
						Usage:       "[Required] set iso size(GB)",
						Value:       5,
						Destination: &createParam.Size,
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
								ISOImageCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
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
										ISOImageCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageCreateCompleteFlags(ctx, createParam, name, cur)
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
							ISOImageCreateCompleteArgs(ctx, createParam, cur, prev, commandName)
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
					return ISOImageCreate(ctx, createParam)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"d", "rm"},
				Usage:     "Delete ISOImage",
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
								ISOImageDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
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
										ISOImageDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageDeleteCompleteFlags(ctx, deleteParam, name, cur)
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
							ISOImageDeleteCompleteArgs(ctx, deleteParam, cur, prev, commandName)
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
					return ISOImageDelete(ctx, deleteParam)
				},
			},
			{
				Name:      "download",
				Usage:     "Download ISOImage",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "file-destination",
						Usage:       "[Required] set file destination path",
						Destination: &downloadParam.FileDestination,
					},
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &downloadParam.Id,
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
					ctx := NewContext(c, realArgs, downloadParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								ISOImageDownloadCompleteArgs(ctx, downloadParam, cur, prev, commandName)
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
										ISOImageDownloadCompleteArgs(ctx, downloadParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageDownloadCompleteFlags(ctx, downloadParam, name, cur)
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
							ISOImageDownloadCompleteArgs(ctx, downloadParam, cur, prev, commandName)
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
					if errors := downloadParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), downloadParam)

					// Run command with params
					return ISOImageDownload(ctx, downloadParam)
				},
			},
			{
				Name:      "ftp-close",
				Usage:     "FtpClose ISOImage",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &ftpCloseParam.Id,
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
					ctx := NewContext(c, realArgs, ftpCloseParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								ISOImageFtpCloseCompleteArgs(ctx, ftpCloseParam, cur, prev, commandName)
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
										ISOImageFtpCloseCompleteArgs(ctx, ftpCloseParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageFtpCloseCompleteFlags(ctx, ftpCloseParam, name, cur)
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
							ISOImageFtpCloseCompleteArgs(ctx, ftpCloseParam, cur, prev, commandName)
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
					if errors := ftpCloseParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), ftpCloseParam)

					// Run command with params
					return ISOImageFtpClose(ctx, ftpCloseParam)
				},
			},
			{
				Name:      "ftp-open",
				Usage:     "FtpOpen ISOImage",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &ftpOpenParam.Id,
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
					ctx := NewContext(c, realArgs, ftpOpenParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								ISOImageFtpOpenCompleteArgs(ctx, ftpOpenParam, cur, prev, commandName)
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
										ISOImageFtpOpenCompleteArgs(ctx, ftpOpenParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageFtpOpenCompleteFlags(ctx, ftpOpenParam, name, cur)
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
							ISOImageFtpOpenCompleteArgs(ctx, ftpOpenParam, cur, prev, commandName)
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
					if errors := ftpOpenParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), ftpOpenParam)

					// Run command with params
					return ISOImageFtpOpen(ctx, ftpOpenParam)
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l", "ls", "find"},
				Usage:   "List ISOImage",
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
					&cli.StringFlag{
						Name:        "scope",
						Usage:       "set filter by scope('user' or 'shared')",
						Destination: &listParam.Scope,
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
								ISOImageListCompleteArgs(ctx, listParam, cur, prev, commandName)
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
										ISOImageListCompleteArgs(ctx, listParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageListCompleteFlags(ctx, listParam, name, cur)
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
							ISOImageListCompleteArgs(ctx, listParam, cur, prev, commandName)
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
					return ISOImageList(ctx, listParam)
				},
			},
			{
				Name:      "read",
				Aliases:   []string{"r"},
				Usage:     "Read ISOImage",
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
								ISOImageReadCompleteArgs(ctx, readParam, cur, prev, commandName)
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
										ISOImageReadCompleteArgs(ctx, readParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageReadCompleteFlags(ctx, readParam, name, cur)
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
							ISOImageReadCompleteArgs(ctx, readParam, cur, prev, commandName)
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
					return ISOImageRead(ctx, readParam)
				},
			},
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     "Update ISOImage",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
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
								ISOImageUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
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
										ISOImageUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageUpdateCompleteFlags(ctx, updateParam, name, cur)
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
							ISOImageUpdateCompleteArgs(ctx, updateParam, cur, prev, commandName)
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
					return ISOImageUpdate(ctx, updateParam)
				},
			},
			{
				Name:      "upload",
				Usage:     "Upload ISOImage",
				ArgsUsage: "[ResourceID]",
				Flags: []cli.Flag{
					&cli.Int64Flag{
						Name:        "id",
						Usage:       "[Required] set resource ID",
						Destination: &uploadParam.Id,
					},
					&cli.StringFlag{
						Name:        "iso-file",
						Usage:       "[Required] set iso image file",
						Destination: &uploadParam.IsoFile,
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
					ctx := NewContext(c, realArgs, uploadParam)

					if strings.HasPrefix(prev, "-") {
						// prev if flag , is values setted?
						if strings.Contains(prev, "=") {
							if strings.HasPrefix(cur, "-") {
								completionFlagNames(c, commandName)
								return
							} else {
								ISOImageUploadCompleteArgs(ctx, uploadParam, cur, prev, commandName)
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
										ISOImageUploadCompleteArgs(ctx, uploadParam, cur, prev, commandName)
										return
									}
								} else {
									// prev is flag , call completion func of each flags
									ISOImageUploadCompleteFlags(ctx, uploadParam, name, cur)
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
							ISOImageUploadCompleteArgs(ctx, uploadParam, cur, prev, commandName)
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
					if errors := uploadParam.Validate(); len(errors) > 0 {
						return flattenErrorsWithPrefix(errors, "Options")
					}

					// create command context
					ctx := NewContext(c, c.Args().Slice(), uploadParam)

					// Run command with params
					return ISOImageUpload(ctx, uploadParam)
				},
			},
		},
	}

	// build Category-Resource mapping
	appendResourceCategoryMap("iso-image", &schema.Category{
		Key:         "storage",
		DisplayName: "Storage",
		Order:       20,
	})

	// build Category-Command mapping

	appendCommandCategoryMap("iso-image", "create", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "delete", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "download", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "ftp-close", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "ftp-open", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "list", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "read", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "update", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})
	appendCommandCategoryMap("iso-image", "upload", &schema.Category{
		Key:         "default",
		DisplayName: "",
		Order:       2147483647,
	})

	// build Category-Param mapping

	appendFlagCategoryMap("iso-image", "create", "description", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "create", "icon-id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "create", "iso-file", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "create", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "create", "size", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "create", "tags", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "delete", "force", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "delete", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "download", "file-destination", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "download", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "ftp-close", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "ftp-open", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "list", "from", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "list", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "list", "max", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "list", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "list", "scope", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "list", "sort", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "read", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "update", "description", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "update", "icon-id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "update", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "update", "name", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "update", "tags", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "upload", "id", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})
	appendFlagCategoryMap("iso-image", "upload", "iso-file", &schema.Category{
		Key:         "default",
		DisplayName: "Other options",
		Order:       2147483647,
	})

	// append command to GlobalContext
	Commands = append(Commands, cliCommand)
}
