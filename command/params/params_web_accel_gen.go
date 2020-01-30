// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListWebAccelParam is input parameters for the sacloud API
type ListWebAccelParam struct {
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

// NewListWebAccelParam return new ListWebAccelParam
func NewListWebAccelParam() *ListWebAccelParam {
	return &ListWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListWebAccelParam) FillValueToSkeleton() {
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
func (p *ListWebAccelParam) Validate() []error {
	errors := []error{}

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

func (p *ListWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *ListWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *ListWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *ListWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *ListWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}

// ReadWebAccelParam is input parameters for the sacloud API
type ReadWebAccelParam struct {
	Selector          []string `json:"selector"`
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

// NewReadWebAccelParam return new ReadWebAccelParam
func NewReadWebAccelParam() *ReadWebAccelParam {
	return &ReadWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadWebAccelParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
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
func (p *ReadWebAccelParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
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

func (p *ReadWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *ReadWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *ReadWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *ReadWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *ReadWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *ReadWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *ReadWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadWebAccelParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadWebAccelParam) GetId() int64 {
	return p.Id
}

// CertificateInfoWebAccelParam is input parameters for the sacloud API
type CertificateInfoWebAccelParam struct {
	Selector          []string `json:"selector"`
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

// NewCertificateInfoWebAccelParam return new CertificateInfoWebAccelParam
func NewCertificateInfoWebAccelParam() *CertificateInfoWebAccelParam {
	return &CertificateInfoWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *CertificateInfoWebAccelParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
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
func (p *CertificateInfoWebAccelParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
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

func (p *CertificateInfoWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *CertificateInfoWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["certificate-info"]
}

func (p *CertificateInfoWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CertificateInfoWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CertificateInfoWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CertificateInfoWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CertificateInfoWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *CertificateInfoWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *CertificateInfoWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CertificateInfoWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CertificateInfoWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CertificateInfoWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CertificateInfoWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CertificateInfoWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CertificateInfoWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CertificateInfoWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *CertificateInfoWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CertificateInfoWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *CertificateInfoWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CertificateInfoWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CertificateInfoWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *CertificateInfoWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *CertificateInfoWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CertificateInfoWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CertificateInfoWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *CertificateInfoWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *CertificateInfoWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CertificateInfoWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *CertificateInfoWebAccelParam) SetId(v int64) {
	p.Id = v
}

func (p *CertificateInfoWebAccelParam) GetId() int64 {
	return p.Id
}

// CertificateNewWebAccelParam is input parameters for the sacloud API
type CertificateNewWebAccelParam struct {
	Cert              string   `json:"cert"`
	Key               string   `json:"key"`
	CertContent       string   `json:"cert-content"`
	KeyContent        string   `json:"key-content"`
	Selector          []string `json:"selector"`
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

// NewCertificateNewWebAccelParam return new CertificateNewWebAccelParam
func NewCertificateNewWebAccelParam() *CertificateNewWebAccelParam {
	return &CertificateNewWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *CertificateNewWebAccelParam) FillValueToSkeleton() {
	if isEmpty(p.Cert) {
		p.Cert = ""
	}
	if isEmpty(p.Key) {
		p.Key = ""
	}
	if isEmpty(p.CertContent) {
		p.CertContent = ""
	}
	if isEmpty(p.KeyContent) {
		p.KeyContent = ""
	}
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
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
func (p *CertificateNewWebAccelParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["WebAccel"].Commands["certificate-new"].Params["cert"].ValidateFunc
		errs := validator("--cert", p.Cert)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["WebAccel"].Commands["certificate-new"].Params["key"].ValidateFunc
		errs := validator("--key", p.Key)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--cert-content", p.CertContent, map[string]interface{}{

			"--cert": p.Cert,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--key-content", p.KeyContent, map[string]interface{}{

			"--key": p.Key,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
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

func (p *CertificateNewWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *CertificateNewWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["certificate-new"]
}

func (p *CertificateNewWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CertificateNewWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CertificateNewWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CertificateNewWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CertificateNewWebAccelParam) SetCert(v string) {
	p.Cert = v
}

func (p *CertificateNewWebAccelParam) GetCert() string {
	return p.Cert
}
func (p *CertificateNewWebAccelParam) SetKey(v string) {
	p.Key = v
}

func (p *CertificateNewWebAccelParam) GetKey() string {
	return p.Key
}
func (p *CertificateNewWebAccelParam) SetCertContent(v string) {
	p.CertContent = v
}

func (p *CertificateNewWebAccelParam) GetCertContent() string {
	return p.CertContent
}
func (p *CertificateNewWebAccelParam) SetKeyContent(v string) {
	p.KeyContent = v
}

func (p *CertificateNewWebAccelParam) GetKeyContent() string {
	return p.KeyContent
}
func (p *CertificateNewWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *CertificateNewWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *CertificateNewWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CertificateNewWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CertificateNewWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CertificateNewWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CertificateNewWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CertificateNewWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CertificateNewWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CertificateNewWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CertificateNewWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CertificateNewWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *CertificateNewWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CertificateNewWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *CertificateNewWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CertificateNewWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CertificateNewWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *CertificateNewWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *CertificateNewWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CertificateNewWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CertificateNewWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *CertificateNewWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *CertificateNewWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CertificateNewWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *CertificateNewWebAccelParam) SetId(v int64) {
	p.Id = v
}

func (p *CertificateNewWebAccelParam) GetId() int64 {
	return p.Id
}

// CertificateUpdateWebAccelParam is input parameters for the sacloud API
type CertificateUpdateWebAccelParam struct {
	Cert              string   `json:"cert"`
	Key               string   `json:"key"`
	CertContent       string   `json:"cert-content"`
	KeyContent        string   `json:"key-content"`
	Selector          []string `json:"selector"`
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

// NewCertificateUpdateWebAccelParam return new CertificateUpdateWebAccelParam
func NewCertificateUpdateWebAccelParam() *CertificateUpdateWebAccelParam {
	return &CertificateUpdateWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *CertificateUpdateWebAccelParam) FillValueToSkeleton() {
	if isEmpty(p.Cert) {
		p.Cert = ""
	}
	if isEmpty(p.Key) {
		p.Key = ""
	}
	if isEmpty(p.CertContent) {
		p.CertContent = ""
	}
	if isEmpty(p.KeyContent) {
		p.KeyContent = ""
	}
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
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
func (p *CertificateUpdateWebAccelParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["WebAccel"].Commands["certificate-update"].Params["cert"].ValidateFunc
		errs := validator("--cert", p.Cert)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["WebAccel"].Commands["certificate-update"].Params["key"].ValidateFunc
		errs := validator("--key", p.Key)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--cert-content", p.CertContent, map[string]interface{}{

			"--cert": p.Cert,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--key-content", p.KeyContent, map[string]interface{}{

			"--key": p.Key,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
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

func (p *CertificateUpdateWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *CertificateUpdateWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["certificate-update"]
}

func (p *CertificateUpdateWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CertificateUpdateWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CertificateUpdateWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CertificateUpdateWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CertificateUpdateWebAccelParam) SetCert(v string) {
	p.Cert = v
}

func (p *CertificateUpdateWebAccelParam) GetCert() string {
	return p.Cert
}
func (p *CertificateUpdateWebAccelParam) SetKey(v string) {
	p.Key = v
}

func (p *CertificateUpdateWebAccelParam) GetKey() string {
	return p.Key
}
func (p *CertificateUpdateWebAccelParam) SetCertContent(v string) {
	p.CertContent = v
}

func (p *CertificateUpdateWebAccelParam) GetCertContent() string {
	return p.CertContent
}
func (p *CertificateUpdateWebAccelParam) SetKeyContent(v string) {
	p.KeyContent = v
}

func (p *CertificateUpdateWebAccelParam) GetKeyContent() string {
	return p.KeyContent
}
func (p *CertificateUpdateWebAccelParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *CertificateUpdateWebAccelParam) GetSelector() []string {
	return p.Selector
}
func (p *CertificateUpdateWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CertificateUpdateWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CertificateUpdateWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CertificateUpdateWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CertificateUpdateWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CertificateUpdateWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CertificateUpdateWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CertificateUpdateWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CertificateUpdateWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CertificateUpdateWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *CertificateUpdateWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CertificateUpdateWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *CertificateUpdateWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CertificateUpdateWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CertificateUpdateWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *CertificateUpdateWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *CertificateUpdateWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CertificateUpdateWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *CertificateUpdateWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *CertificateUpdateWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *CertificateUpdateWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *CertificateUpdateWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *CertificateUpdateWebAccelParam) SetId(v int64) {
	p.Id = v
}

func (p *CertificateUpdateWebAccelParam) GetId() int64 {
	return p.Id
}

// DeleteCacheWebAccelParam is input parameters for the sacloud API
type DeleteCacheWebAccelParam struct {
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
}

// NewDeleteCacheWebAccelParam return new DeleteCacheWebAccelParam
func NewDeleteCacheWebAccelParam() *DeleteCacheWebAccelParam {
	return &DeleteCacheWebAccelParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteCacheWebAccelParam) FillValueToSkeleton() {
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

}

// Validate checks current values in model
func (p *DeleteCacheWebAccelParam) Validate() []error {
	errors := []error{}

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

func (p *DeleteCacheWebAccelParam) GetResourceDef() *schema.Resource {
	return define.Resources["WebAccel"]
}

func (p *DeleteCacheWebAccelParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete-cache"]
}

func (p *DeleteCacheWebAccelParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteCacheWebAccelParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteCacheWebAccelParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteCacheWebAccelParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteCacheWebAccelParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteCacheWebAccelParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteCacheWebAccelParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteCacheWebAccelParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteCacheWebAccelParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteCacheWebAccelParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteCacheWebAccelParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteCacheWebAccelParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteCacheWebAccelParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteCacheWebAccelParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteCacheWebAccelParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteCacheWebAccelParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteCacheWebAccelParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteCacheWebAccelParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteCacheWebAccelParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteCacheWebAccelParam) GetFormat() string {
	return p.Format
}
func (p *DeleteCacheWebAccelParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteCacheWebAccelParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteCacheWebAccelParam) SetQuery(v string) {
	p.Query = v
}

func (p *DeleteCacheWebAccelParam) GetQuery() string {
	return p.Query
}
func (p *DeleteCacheWebAccelParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *DeleteCacheWebAccelParam) GetQueryFile() string {
	return p.QueryFile
}
