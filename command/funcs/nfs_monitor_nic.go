package funcs

import (
	"bytes"
	"fmt"
	"sort"
	"text/template"
	"time"

	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
)

func NFSMonitorNic(ctx command.Context, params *params.MonitorNicNFSParam) error {

	client := ctx.GetAPIClient()
	api := client.GetNFSAPI()

	_, e := api.Read(params.Id)
	if e != nil {
		return fmt.Errorf("NFSMonitorNic is failed: %s", e)
	}

	end := parseDateTimeString(params.End)
	start := end.Add(-1 * time.Hour)
	if params.Start != "" {
		start = parseDateTimeString(params.Start)
	}

	// validate start <= end
	if !(start.Unix() <= end.Unix()) {
		return fmt.Errorf("Invalid Parameter : start(%s) or end(%s) is invalid", start, end)
	}

	req := sacloud.NewResourceMonitorRequest(&start, &end)

	list := MonitorValues{}

	res, err := api.MonitorInterface(params.Id, req)
	if err != nil {
		return fmt.Errorf("NFSMonitorNic is failed: %s", err)
	}

	// collect values
	receiveValues, err := res.FlattenPacketReceiveValue()
	if err != nil {
		return fmt.Errorf("NFSMonitorNic is failed: %s", err)
	}
	sendValues, err := res.FlattenPacketSendValue()
	if err != nil {
		return fmt.Errorf("NFSMonitorNic is failed: %s", err)
	}

	// sort
	sort.Slice(receiveValues, func(i, j int) bool { return receiveValues[i].Time.Before(receiveValues[j].Time) })
	sort.Slice(sendValues, func(i, j int) bool { return sendValues[i].Time.Before(sendValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("NFSMonitorNic is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	for i := range receiveValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", receiveValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", receiveValues[i].Time.Unix()),
			"Send":      fmt.Sprintf("%f", receiveValues[i].Value),
			"Receive":   fmt.Sprintf("%f", sendValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}
