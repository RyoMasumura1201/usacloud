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

func DatabaseMonitorBackupDiskSize(ctx command.Context, params *params.MonitorBackupDiskSizeDatabaseParam) error {

	client := ctx.GetAPIClient()
	api := client.GetDatabaseAPI()

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

	res, err := api.MonitorDatabase(params.Id, req)
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}

	// collect values
	totalValues, err := res.FlattenTotalDisk2SizeValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}
	usedValues, err := res.FlattenUsedDisk2SizeValue()
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}

	// sort
	sort.Slice(totalValues, func(i, j int) bool { return totalValues[i].Time.Before(totalValues[j].Time) })
	sort.Slice(usedValues, func(i, j int) bool { return usedValues[i].Time.Before(usedValues[j].Time) })

	// build key string
	key := ""
	buf := bytes.NewBufferString("")
	t := template.New("t")
	template.Must(t.Parse(params.KeyFormat))
	err = t.Execute(buf, map[string]interface{}{
		"ID": params.Id,
	})
	if err != nil {
		return fmt.Errorf("DatabaseMonitorBackupDiskSize is failed: %s", err)
	}
	key = buf.String()

	// build sortable struct
	list := MonitorValues{}
	for i := range totalValues {
		list = append(list, MonitorValue{
			"Key":       key,
			"TimeStamp": fmt.Sprintf("%s", totalValues[i].Time),
			"UnixTime":  fmt.Sprintf("%d", totalValues[i].Time.Unix()),
			"Used":      fmt.Sprintf("%f", usedValues[i].Value),
			"Total":     fmt.Sprintf("%f", totalValues[i].Value),
		})
	}
	sort.Sort(list)

	values := []interface{}{}
	for _, v := range list {
		values = append(values, v)
	}

	return ctx.GetOutput().Print(values...)

}
