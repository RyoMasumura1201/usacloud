// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ProductLicenseListCompleteFlags(ctx command.Context, params *params.ListProductLicenseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["ProductLicense"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProductLicense"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["ProductLicense"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["ProductLicense"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["ProductLicense"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProductLicenseReadCompleteFlags(ctx command.Context, params *params.ReadProductLicenseParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["ProductLicense"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}