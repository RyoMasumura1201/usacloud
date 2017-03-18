// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func DiskEditCompleteFlags(ctx Context, params *EditDiskParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "default-route", "gateway":
		comp = define.Resources["Disk"].Commands["edit"].Params["default-route"].CompleteFunc
	case "disable-password-auth", "disable-pw-auth":
		comp = define.Resources["Disk"].Commands["edit"].Params["disable-password-auth"].CompleteFunc
	case "hostname":
		comp = define.Resources["Disk"].Commands["edit"].Params["hostname"].CompleteFunc
	case "id":
		comp = define.Resources["Disk"].Commands["edit"].Params["id"].CompleteFunc
	case "ipaddress", "ip":
		comp = define.Resources["Disk"].Commands["edit"].Params["ipaddress"].CompleteFunc
	case "nw-masklen", "network-masklen":
		comp = define.Resources["Disk"].Commands["edit"].Params["nw-masklen"].CompleteFunc
	case "password":
		comp = define.Resources["Disk"].Commands["edit"].Params["password"].CompleteFunc
	case "ssh-key-ids":
		comp = define.Resources["Disk"].Commands["edit"].Params["ssh-key-ids"].CompleteFunc
	case "startup-script-ids", "note-ids":
		comp = define.Resources["Disk"].Commands["edit"].Params["startup-script-ids"].CompleteFunc
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