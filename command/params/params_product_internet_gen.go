// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListProductInternetParam is input parameters for the sacloud API
type ListProductInternetParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
	From              int      `json:"from"`
	Max               int      `json:"max"`
	Sort              []string `json:"sort"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewListProductInternetParam return new ListProductInternetParam
func NewListProductInternetParam() *ListProductInternetParam {
	return &ListProductInternetParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListProductInternetParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *ListProductInternetParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductInternet"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListProductInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductInternet"]
}

func (p *ListProductInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListProductInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListProductInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListProductInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListProductInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListProductInternetParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductInternetParam) GetName() []string {
	return p.Name
}
func (p *ListProductInternetParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListProductInternetParam) GetId() []int64 {
	return p.Id
}
func (p *ListProductInternetParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductInternetParam) GetFrom() int {
	return p.From
}
func (p *ListProductInternetParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductInternetParam) GetMax() int {
	return p.Max
}
func (p *ListProductInternetParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductInternetParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListProductInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListProductInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListProductInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListProductInternetParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListProductInternetParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListProductInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductInternetParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductInternetParam) GetFormat() string {
	return p.Format
}
func (p *ListProductInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListProductInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListProductInternetParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListProductInternetParam) GetQuery() string {
	return p.Query
}
func (p *ListProductInternetParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListProductInternetParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadProductInternetParam is input parameters for the sacloud API
type ReadProductInternetParam struct {
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
	Id                int64    `json:"id"`
}

// NewReadProductInternetParam return new ReadProductInternetParam
func NewReadProductInternetParam() *ReadProductInternetParam {
	return &ReadProductInternetParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadProductInternetParam) FillValueToSkeleton() {
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *ReadProductInternetParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductInternet"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadProductInternetParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductInternet"]
}

func (p *ReadProductInternetParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadProductInternetParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadProductInternetParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadProductInternetParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadProductInternetParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadProductInternetParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadProductInternetParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadProductInternetParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadProductInternetParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadProductInternetParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadProductInternetParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadProductInternetParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadProductInternetParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadProductInternetParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductInternetParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductInternetParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductInternetParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductInternetParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductInternetParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductInternetParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductInternetParam) GetFormat() string {
	return p.Format
}
func (p *ReadProductInternetParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadProductInternetParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadProductInternetParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadProductInternetParam) GetQuery() string {
	return p.Query
}
func (p *ReadProductInternetParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadProductInternetParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadProductInternetParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadProductInternetParam) GetId() int64 {
	return p.Id
}
