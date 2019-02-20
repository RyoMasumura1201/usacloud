// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListProductServerParam is input parameters for the sacloud API
type ListProductServerParam struct {
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

// NewListProductServerParam return new ListProductServerParam
func NewListProductServerParam() *ListProductServerParam {
	return &ListProductServerParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListProductServerParam) FillValueToSkeleton() {
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
func (p *ListProductServerParam) Validate() []error {
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
		validator := define.Resources["ProductServer"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListProductServerParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductServer"]
}

func (p *ListProductServerParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListProductServerParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListProductServerParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListProductServerParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListProductServerParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListProductServerParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductServerParam) GetName() []string {
	return p.Name
}
func (p *ListProductServerParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListProductServerParam) GetId() []int64 {
	return p.Id
}
func (p *ListProductServerParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductServerParam) GetFrom() int {
	return p.From
}
func (p *ListProductServerParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductServerParam) GetMax() int {
	return p.Max
}
func (p *ListProductServerParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductServerParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductServerParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListProductServerParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListProductServerParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListProductServerParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListProductServerParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListProductServerParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListProductServerParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductServerParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductServerParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductServerParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductServerParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductServerParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductServerParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductServerParam) GetFormat() string {
	return p.Format
}
func (p *ListProductServerParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListProductServerParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListProductServerParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListProductServerParam) GetQuery() string {
	return p.Query
}
func (p *ListProductServerParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListProductServerParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadProductServerParam is input parameters for the sacloud API
type ReadProductServerParam struct {
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

// NewReadProductServerParam return new ReadProductServerParam
func NewReadProductServerParam() *ReadProductServerParam {
	return &ReadProductServerParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadProductServerParam) FillValueToSkeleton() {
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
func (p *ReadProductServerParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductServer"].Commands["read"].Params["id"].ValidateFunc
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

func (p *ReadProductServerParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductServer"]
}

func (p *ReadProductServerParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadProductServerParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadProductServerParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadProductServerParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadProductServerParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadProductServerParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadProductServerParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadProductServerParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadProductServerParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadProductServerParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadProductServerParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadProductServerParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadProductServerParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadProductServerParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductServerParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductServerParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductServerParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductServerParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductServerParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductServerParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductServerParam) GetFormat() string {
	return p.Format
}
func (p *ReadProductServerParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadProductServerParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadProductServerParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadProductServerParam) GetQuery() string {
	return p.Query
}
func (p *ReadProductServerParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadProductServerParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadProductServerParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadProductServerParam) GetId() int64 {
	return p.Id
}
