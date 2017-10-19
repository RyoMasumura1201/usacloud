// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func GSLBListCompleteFlags(ctx command.Context, params *params.ListGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["GSLB"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["GSLB"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["GSLB"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["GSLB"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["GSLB"].Commands["list"].BuildedParams().Get("sort")
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

func GSLBServerInfoCompleteFlags(ctx command.Context, params *params.ServerInfoGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["GSLB"].Commands["server-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["server-info"].BuildedParams().Get("id")
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

func GSLBCreateCompleteFlags(ctx command.Context, params *params.CreateGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "protocol":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("protocol")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "host-header":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("host-header")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "path":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("path")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "response-code":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("response-code")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "delay-loop":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("delay-loop")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "weighted":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("weighted")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sorry-server":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("sorry-server")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["GSLB"].Commands["create"].BuildedParams().Get("icon-id")
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

func GSLBServerAddCompleteFlags(ctx command.Context, params *params.ServerAddGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "ipaddress":
		param := define.Resources["GSLB"].Commands["server-add"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disabled":
		param := define.Resources["GSLB"].Commands["server-add"].BuildedParams().Get("disabled")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "weight":
		param := define.Resources["GSLB"].Commands["server-add"].BuildedParams().Get("weight")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["GSLB"].Commands["server-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["server-add"].BuildedParams().Get("id")
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

func GSLBReadCompleteFlags(ctx command.Context, params *params.ReadGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["GSLB"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["read"].BuildedParams().Get("id")
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

func GSLBServerUpdateCompleteFlags(ctx command.Context, params *params.ServerUpdateGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["GSLB"].Commands["server-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ipaddress":
		param := define.Resources["GSLB"].Commands["server-update"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disabled":
		param := define.Resources["GSLB"].Commands["server-update"].BuildedParams().Get("disabled")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "weight":
		param := define.Resources["GSLB"].Commands["server-update"].BuildedParams().Get("weight")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["GSLB"].Commands["server-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["server-update"].BuildedParams().Get("id")
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

func GSLBServerDeleteCompleteFlags(ctx command.Context, params *params.ServerDeleteGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["GSLB"].Commands["server-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["GSLB"].Commands["server-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["server-delete"].BuildedParams().Get("id")
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

func GSLBUpdateCompleteFlags(ctx command.Context, params *params.UpdateGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "protocol":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("protocol")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "host-header":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("host-header")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "path":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("path")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "response-code":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("response-code")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "delay-loop":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("delay-loop")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "weighted":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("weighted")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sorry-server":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("sorry-server")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["update"].BuildedParams().Get("id")
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

func GSLBDeleteCompleteFlags(ctx command.Context, params *params.DeleteGSLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["GSLB"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["GSLB"].Commands["delete"].BuildedParams().Get("id")
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