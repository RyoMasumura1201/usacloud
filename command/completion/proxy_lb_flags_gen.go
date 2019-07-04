// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"

	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ProxyLBListCompleteFlags(ctx command.Context, params *params.ListProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["ProxyLB"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags", "selector":
		param := define.Resources["ProxyLB"].Commands["list"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["ProxyLB"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["ProxyLB"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["ProxyLB"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBCreateCompleteFlags(ctx command.Context, params *params.CreateProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "plan":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("plan")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "protocol":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("protocol")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "host-header":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("host-header")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "path":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("path")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "delay-loop":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("delay-loop")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sticky-session":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("sticky-session")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sorry-server-ipaddress":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("sorry-server-ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sorry-server-port":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("sorry-server-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["ProxyLB"].Commands["create"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBReadCompleteFlags(ctx command.Context, params *params.ReadProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["read"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBUpdateCompleteFlags(ctx command.Context, params *params.UpdateProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "protocol":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("protocol")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "host-header":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("host-header")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "path":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("path")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "delay-loop":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("delay-loop")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sticky-session":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("sticky-session")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sorry-server-ipaddress":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("sorry-server-ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sorry-server-port":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("sorry-server-port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "name":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "tags":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("tags")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "icon-id":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("icon-id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBDeleteCompleteFlags(ctx command.Context, params *params.DeleteProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBPlanChangeCompleteFlags(ctx command.Context, params *params.PlanChangeProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "plan":
		param := define.Resources["ProxyLB"].Commands["plan-change"].BuildedParams().Get("plan")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["plan-change"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["plan-change"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBBindPortInfoCompleteFlags(ctx command.Context, params *params.BindPortInfoProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["bind-port-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["bind-port-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBBindPortAddCompleteFlags(ctx command.Context, params *params.BindPortAddProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "mode":
		param := define.Resources["ProxyLB"].Commands["bind-port-add"].BuildedParams().Get("mode")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port":
		param := define.Resources["ProxyLB"].Commands["bind-port-add"].BuildedParams().Get("port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "redirect-to-https":
		param := define.Resources["ProxyLB"].Commands["bind-port-add"].BuildedParams().Get("redirect-to-https")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "support-http2":
		param := define.Resources["ProxyLB"].Commands["bind-port-add"].BuildedParams().Get("support-http2")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["bind-port-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["bind-port-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBBindPortUpdateCompleteFlags(ctx command.Context, params *params.BindPortUpdateProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "mode":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("mode")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "redirect-to-https":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("redirect-to-https")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "support-http2":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("support-http2")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["bind-port-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBBindPortDeleteCompleteFlags(ctx command.Context, params *params.BindPortDeleteProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["ProxyLB"].Commands["bind-port-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["bind-port-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["bind-port-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBResponseHeaderInfoCompleteFlags(ctx command.Context, params *params.ResponseHeaderInfoProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "port-index":
		param := define.Resources["ProxyLB"].Commands["response-header-info"].BuildedParams().Get("port-index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["response-header-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["response-header-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBResponseHeaderAddCompleteFlags(ctx command.Context, params *params.ResponseHeaderAddProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "port-index":
		param := define.Resources["ProxyLB"].Commands["response-header-add"].BuildedParams().Get("port-index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "header":
		param := define.Resources["ProxyLB"].Commands["response-header-add"].BuildedParams().Get("header")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "value":
		param := define.Resources["ProxyLB"].Commands["response-header-add"].BuildedParams().Get("value")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["response-header-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["response-header-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBResponseHeaderUpdateCompleteFlags(ctx command.Context, params *params.ResponseHeaderUpdateProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["ProxyLB"].Commands["response-header-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port-index":
		param := define.Resources["ProxyLB"].Commands["response-header-update"].BuildedParams().Get("port-index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "header":
		param := define.Resources["ProxyLB"].Commands["response-header-update"].BuildedParams().Get("header")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "value":
		param := define.Resources["ProxyLB"].Commands["response-header-update"].BuildedParams().Get("value")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["response-header-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["response-header-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBResponseHeaderDeleteCompleteFlags(ctx command.Context, params *params.ResponseHeaderDeleteProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["ProxyLB"].Commands["response-header-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port-index":
		param := define.Resources["ProxyLB"].Commands["response-header-delete"].BuildedParams().Get("port-index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["response-header-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["response-header-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBAcmeInfoCompleteFlags(ctx command.Context, params *params.AcmeInfoProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["acme-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["acme-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBAcmeSettingCompleteFlags(ctx command.Context, params *params.AcmeSettingProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "accept-tos":
		param := define.Resources["ProxyLB"].Commands["acme-setting"].BuildedParams().Get("accept-tos")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "common-name":
		param := define.Resources["ProxyLB"].Commands["acme-setting"].BuildedParams().Get("common-name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disable":
		param := define.Resources["ProxyLB"].Commands["acme-setting"].BuildedParams().Get("disable")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["acme-setting"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["acme-setting"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBAcmeRenewCompleteFlags(ctx command.Context, params *params.AcmeRenewProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["acme-renew"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["acme-renew"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBServerInfoCompleteFlags(ctx command.Context, params *params.ServerInfoProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["server-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["server-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBServerAddCompleteFlags(ctx command.Context, params *params.ServerAddProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "ipaddress":
		param := define.Resources["ProxyLB"].Commands["server-add"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disabled":
		param := define.Resources["ProxyLB"].Commands["server-add"].BuildedParams().Get("disabled")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port":
		param := define.Resources["ProxyLB"].Commands["server-add"].BuildedParams().Get("port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["server-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["server-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBServerUpdateCompleteFlags(ctx command.Context, params *params.ServerUpdateProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["ProxyLB"].Commands["server-update"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "ipaddress":
		param := define.Resources["ProxyLB"].Commands["server-update"].BuildedParams().Get("ipaddress")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "disabled":
		param := define.Resources["ProxyLB"].Commands["server-update"].BuildedParams().Get("disabled")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "port":
		param := define.Resources["ProxyLB"].Commands["server-update"].BuildedParams().Get("port")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["server-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["server-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBServerDeleteCompleteFlags(ctx command.Context, params *params.ServerDeleteProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "index":
		param := define.Resources["ProxyLB"].Commands["server-delete"].BuildedParams().Get("index")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["server-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["server-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBCertificateInfoCompleteFlags(ctx command.Context, params *params.CertificateInfoProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["certificate-info"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["certificate-info"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBCertificateAddCompleteFlags(ctx command.Context, params *params.CertificateAddProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "server-certificate", "server-cert":
		param := define.Resources["ProxyLB"].Commands["certificate-add"].BuildedParams().Get("server-certificate")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "intermediate-certificate", "issuer-cert":
		param := define.Resources["ProxyLB"].Commands["certificate-add"].BuildedParams().Get("intermediate-certificate")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "private-key":
		param := define.Resources["ProxyLB"].Commands["certificate-add"].BuildedParams().Get("private-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["certificate-add"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["certificate-add"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBCertificateUpdateCompleteFlags(ctx command.Context, params *params.CertificateUpdateProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "server-certificate", "server-cert":
		param := define.Resources["ProxyLB"].Commands["certificate-update"].BuildedParams().Get("server-certificate")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "intermediate-certificate", "issuer-cert":
		param := define.Resources["ProxyLB"].Commands["certificate-update"].BuildedParams().Get("intermediate-certificate")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "private-key":
		param := define.Resources["ProxyLB"].Commands["certificate-update"].BuildedParams().Get("private-key")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["certificate-update"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["certificate-update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBCertificateDeleteCompleteFlags(ctx command.Context, params *params.CertificateDeleteProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "selector":
		param := define.Resources["ProxyLB"].Commands["certificate-delete"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["certificate-delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func ProxyLBMonitorCompleteFlags(ctx command.Context, params *params.MonitorProxyLBParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "start":
		param := define.Resources["ProxyLB"].Commands["monitor"].BuildedParams().Get("start")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "end":
		param := define.Resources["ProxyLB"].Commands["monitor"].BuildedParams().Get("end")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "key-format":
		param := define.Resources["ProxyLB"].Commands["monitor"].BuildedParams().Get("key-format")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "selector":
		param := define.Resources["ProxyLB"].Commands["monitor"].BuildedParams().Get("selector")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["ProxyLB"].Commands["monitor"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out", "o":
		comp = schema.CompleteInStrValues("json", "yaml", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
