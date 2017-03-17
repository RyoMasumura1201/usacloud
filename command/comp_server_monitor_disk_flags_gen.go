// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package command

import (
	"fmt"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func ServerMonitorDiskCompleteFlags(ctx Context, params *MonitorDiskServerParam, flagName string, currentValue string) {
	var comp schema.SchemaCompletionFunc

	switch flagName {
	case "end":
		comp = define.Resources["Server"].Commands["monitor-disk"].Params["end"].CompleteFunc
	case "id":
		comp = define.Resources["Server"].Commands["monitor-disk"].Params["id"].CompleteFunc
	case "index":
		comp = define.Resources["Server"].Commands["monitor-disk"].Params["index"].CompleteFunc
	case "key-format":
		comp = define.Resources["Server"].Commands["monitor-disk"].Params["key-format"].CompleteFunc
	case "start":
		comp = define.Resources["Server"].Commands["monitor-disk"].Params["start"].CompleteFunc
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}
