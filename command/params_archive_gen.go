// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// CreateArchiveParam is input parameters for the sacloud API
type CreateArchiveParam struct {
	ArchiveFile     string
	Description     string
	IconId          int64
	Name            string
	Size            int
	SourceArchiveId int64
	SourceDiskId    int64
	Tags            []string
}

// NewCreateArchiveParam return new CreateArchiveParam
func NewCreateArchiveParam() *CreateArchiveParam {
	return &CreateArchiveParam{}
}

// Validate checks current values in model
func (p *CreateArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["archive-file"].ValidateFunc
		errs := validator("--archive-file", p.ArchiveFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--archive-file", p.ArchiveFile, map[string]interface{}{

			"--source-archive-id": p.SourceArchiveId,
			"--source-disk-id":    p.SourceDiskId,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["size"].ValidateFunc
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--size", p.Size, map[string]interface{}{

			"--source-archive-id": p.SourceArchiveId,
			"--source-disk-id":    p.SourceDiskId,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["source-archive-id"].ValidateFunc
		errs := validator("--source-archive-id", p.SourceArchiveId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--source-archive-id", p.SourceArchiveId, map[string]interface{}{

			"--archive-file":   p.ArchiveFile,
			"--size":           p.Size,
			"--source-disk-id": p.SourceDiskId,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["source-disk-id"].ValidateFunc
		errs := validator("--source-disk-id", p.SourceDiskId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--source-disk-id", p.SourceDiskId, map[string]interface{}{

			"--archive-file":      p.ArchiveFile,
			"--size":              p.Size,
			"--source-archive-id": p.SourceArchiveId,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *CreateArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateArchiveParam) SetArchiveFile(v string) {
	p.ArchiveFile = v
}

func (p *CreateArchiveParam) GetArchiveFile() string {
	return p.ArchiveFile
}
func (p *CreateArchiveParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateArchiveParam) GetDescription() string {
	return p.Description
}
func (p *CreateArchiveParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateArchiveParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateArchiveParam) SetName(v string) {
	p.Name = v
}

func (p *CreateArchiveParam) GetName() string {
	return p.Name
}
func (p *CreateArchiveParam) SetSize(v int) {
	p.Size = v
}

func (p *CreateArchiveParam) GetSize() int {
	return p.Size
}
func (p *CreateArchiveParam) SetSourceArchiveId(v int64) {
	p.SourceArchiveId = v
}

func (p *CreateArchiveParam) GetSourceArchiveId() int64 {
	return p.SourceArchiveId
}
func (p *CreateArchiveParam) SetSourceDiskId(v int64) {
	p.SourceDiskId = v
}

func (p *CreateArchiveParam) GetSourceDiskId() int64 {
	return p.SourceDiskId
}
func (p *CreateArchiveParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateArchiveParam) GetTags() []string {
	return p.Tags
}

// DeleteArchiveParam is input parameters for the sacloud API
type DeleteArchiveParam struct {
	Force bool
	Id    int64
}

// NewDeleteArchiveParam return new DeleteArchiveParam
func NewDeleteArchiveParam() *DeleteArchiveParam {
	return &DeleteArchiveParam{}
}

// Validate checks current values in model
func (p *DeleteArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *DeleteArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteArchiveParam) SetForce(v bool) {
	p.Force = v
}

func (p *DeleteArchiveParam) GetForce() bool {
	return p.Force
}
func (p *DeleteArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteArchiveParam) GetId() int64 {
	return p.Id
}

// DownloadArchiveParam is input parameters for the sacloud API
type DownloadArchiveParam struct {
	FileDestination string
	Id              int64
}

// NewDownloadArchiveParam return new DownloadArchiveParam
func NewDownloadArchiveParam() *DownloadArchiveParam {
	return &DownloadArchiveParam{}
}

// Validate checks current values in model
func (p *DownloadArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--file-destination", p.FileDestination)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["download"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DownloadArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *DownloadArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["download"]
}

func (p *DownloadArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DownloadArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DownloadArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DownloadArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DownloadArchiveParam) SetFileDestination(v string) {
	p.FileDestination = v
}

func (p *DownloadArchiveParam) GetFileDestination() string {
	return p.FileDestination
}
func (p *DownloadArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *DownloadArchiveParam) GetId() int64 {
	return p.Id
}

// FtpCloseArchiveParam is input parameters for the sacloud API
type FtpCloseArchiveParam struct {
	Id int64
}

// NewFtpCloseArchiveParam return new FtpCloseArchiveParam
func NewFtpCloseArchiveParam() *FtpCloseArchiveParam {
	return &FtpCloseArchiveParam{}
}

// Validate checks current values in model
func (p *FtpCloseArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["ftp-close"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *FtpCloseArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *FtpCloseArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["ftp-close"]
}

func (p *FtpCloseArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *FtpCloseArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *FtpCloseArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *FtpCloseArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *FtpCloseArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *FtpCloseArchiveParam) GetId() int64 {
	return p.Id
}

// FtpOpenArchiveParam is input parameters for the sacloud API
type FtpOpenArchiveParam struct {
	Id int64
}

// NewFtpOpenArchiveParam return new FtpOpenArchiveParam
func NewFtpOpenArchiveParam() *FtpOpenArchiveParam {
	return &FtpOpenArchiveParam{}
}

// Validate checks current values in model
func (p *FtpOpenArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["ftp-open"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *FtpOpenArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *FtpOpenArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["ftp-open"]
}

func (p *FtpOpenArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *FtpOpenArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *FtpOpenArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *FtpOpenArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *FtpOpenArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *FtpOpenArchiveParam) GetId() int64 {
	return p.Id
}

// ListArchiveParam is input parameters for the sacloud API
type ListArchiveParam struct {
	From  int
	Id    []int64
	Max   int
	Name  []string
	Scope string
	Sort  []string
}

// NewListArchiveParam return new ListArchiveParam
func NewListArchiveParam() *ListArchiveParam {
	return &ListArchiveParam{}
}

// Validate checks current values in model
func (p *ListArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Archive"].Commands["list"].Params["id"].ValidateFunc
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
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *ListArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListArchiveParam) SetFrom(v int) {
	p.From = v
}

func (p *ListArchiveParam) GetFrom() int {
	return p.From
}
func (p *ListArchiveParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListArchiveParam) GetId() []int64 {
	return p.Id
}
func (p *ListArchiveParam) SetMax(v int) {
	p.Max = v
}

func (p *ListArchiveParam) GetMax() int {
	return p.Max
}
func (p *ListArchiveParam) SetName(v []string) {
	p.Name = v
}

func (p *ListArchiveParam) GetName() []string {
	return p.Name
}
func (p *ListArchiveParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListArchiveParam) GetScope() string {
	return p.Scope
}
func (p *ListArchiveParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListArchiveParam) GetSort() []string {
	return p.Sort
}

// ReadArchiveParam is input parameters for the sacloud API
type ReadArchiveParam struct {
	Id int64
}

// NewReadArchiveParam return new ReadArchiveParam
func NewReadArchiveParam() *ReadArchiveParam {
	return &ReadArchiveParam{}
}

// Validate checks current values in model
func (p *ReadArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *ReadArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadArchiveParam) GetId() int64 {
	return p.Id
}

// UpdateArchiveParam is input parameters for the sacloud API
type UpdateArchiveParam struct {
	Description string
	IconId      int64
	Id          int64
	Name        string
	Tags        []string
}

// NewUpdateArchiveParam return new UpdateArchiveParam
func NewUpdateArchiveParam() *UpdateArchiveParam {
	return &UpdateArchiveParam{}
}

// Validate checks current values in model
func (p *UpdateArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Archive"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *UpdateArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateArchiveParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateArchiveParam) GetDescription() string {
	return p.Description
}
func (p *UpdateArchiveParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateArchiveParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateArchiveParam) GetId() int64 {
	return p.Id
}
func (p *UpdateArchiveParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateArchiveParam) GetName() string {
	return p.Name
}
func (p *UpdateArchiveParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateArchiveParam) GetTags() []string {
	return p.Tags
}

// UploadArchiveParam is input parameters for the sacloud API
type UploadArchiveParam struct {
	ArchiveFile string
	Id          int64
}

// NewUploadArchiveParam return new UploadArchiveParam
func NewUploadArchiveParam() *UploadArchiveParam {
	return &UploadArchiveParam{}
}

// Validate checks current values in model
func (p *UploadArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--archive-file", p.ArchiveFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["upload"].Params["archive-file"].ValidateFunc
		errs := validator("--archive-file", p.ArchiveFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["upload"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UploadArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *UploadArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["upload"]
}

func (p *UploadArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UploadArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UploadArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UploadArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UploadArchiveParam) SetArchiveFile(v string) {
	p.ArchiveFile = v
}

func (p *UploadArchiveParam) GetArchiveFile() string {
	return p.ArchiveFile
}
func (p *UploadArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *UploadArchiveParam) GetId() int64 {
	return p.Id
}

// WaitForCopyArchiveParam is input parameters for the sacloud API
type WaitForCopyArchiveParam struct {
	Id int64
}

// NewWaitForCopyArchiveParam return new WaitForCopyArchiveParam
func NewWaitForCopyArchiveParam() *WaitForCopyArchiveParam {
	return &WaitForCopyArchiveParam{}
}

// Validate checks current values in model
func (p *WaitForCopyArchiveParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Archive"].Commands["wait-for-copy"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *WaitForCopyArchiveParam) getResourceDef() *schema.Resource {
	return define.Resources["Archive"]
}

func (p *WaitForCopyArchiveParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["wait-for-copy"]
}

func (p *WaitForCopyArchiveParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *WaitForCopyArchiveParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *WaitForCopyArchiveParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *WaitForCopyArchiveParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *WaitForCopyArchiveParam) SetId(v int64) {
	p.Id = v
}

func (p *WaitForCopyArchiveParam) GetId() int64 {
	return p.Id
}
