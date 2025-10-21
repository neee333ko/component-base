package v1

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/neee333ko/component-base/pkg/json"
)

type Extend map[string]interface{}

func (ext Extend) String() string {
	bytes, _ := json.Marshal(ext)

	return string(bytes)
}

func (ext Extend) Merge(extendShadow string) Extend {
	var newExt Extend

	json.Unmarshal([]byte(extendShadow), &newExt)

	for k, v := range newExt {
		if _, ok := ext[k]; !ok {
			ext[k] = v
		}
	}

	return ext
}

type TypeMeta struct {
	Type       string `json:"type,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
}

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}

type ObjectMeta struct {
	ID         uint      `json:"id,omitempty" gorm:"primaryKey;column:id;autoIncrement"`
	InstanceID string    `json:"instanceID,omitempty" gorm:"unique;not null;type:varchar(32);column:instance_id"`
	Name       string    `json:"name,omitempty" gorm:"column:name;type:varchar(64);not null" validate:"name"`
	Ext        Extend    `json:"extend,omitempty" gorm:"-" validate:"omitempty"`
	ExtShadow  string    `json:"-" gorm:"column:ext_shadow;type:text" validate:"omitempty"`
	CreatedAt  time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}

func (object *ObjectMeta) BeforeCreate(tx *gorm.DB) error {
	object.ExtShadow = object.Ext.String()

	return nil
}

func (object *ObjectMeta) BeforeUpdate(tx *gorm.DB) error {
	object.ExtShadow = object.Ext.String()

	return nil
}

func (object *ObjectMeta) AfterFind(tx *gorm.DB) error {
	object.Ext.Merge(object.ExtShadow)

	return nil
}

type ListOptions struct {
	TypeMeta       `json:",inline"`
	LabelSelector  string `json:"labelSelector,omitempty" form:"labelSelector"`
	FieldSelector  string `json:"fieldSelector,omitempty" form:"fieldSelector"`
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
	Limit          *int64 `json:"limit,omitempty" form:"limit"`
	Offset         *int64 `json:"offset,omitempty" form:"offset"`
}

type ExportOptions struct {
	TypeMeta `json:",inline"`

	// Should this value be exported.  Export strips fields that a user can not specify.
	// Deprecated. Planned for removal in 1.18.
	Export bool `json:"export"`
	// Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'.
	// Deprecated. Planned for removal in 1.18.
	Exact bool `json:"exact"`
}

// GetOptions is the standard query options to the standard REST get call.
type GetOptions struct {
	TypeMeta `json:",inline"`
}

// DeleteOptions may be provided when deleting an API object.
type DeleteOptions struct {
	TypeMeta `json:",inline"`

	// +optional
	Unscoped bool `json:"unscoped"`
}

// CreateOptions may be provided when creating an API object.
type CreateOptions struct {
	TypeMeta `json:",inline"`

	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`
}

// PatchOptions may be provided when patching an API object.
// PatchOptions is meant to be a superset of UpdateOptions.
type PatchOptions struct {
	TypeMeta `json:",inline"`

	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`

	// Force is going to "force" Apply requests. It means user will
	// re-acquire conflicting fields owned by other people. Force
	// flag must be unset for non-apply patch requests.
	// +optional
	Force bool `json:"force,omitempty"`
}

// UpdateOptions may be provided when updating an API object.
// All fields in UpdateOptions should also be present in PatchOptions.
type UpdateOptions struct {
	TypeMeta `json:",inline"`

	// When present, indicates that modifications should not be
	// persisted. An invalid or unrecognized dryRun directive will
	// result in an error response and no further processing of the
	// request. Valid values are:
	// - All: all dry run stages will be processed
	// +optional
	DryRun []string `json:"dryRun,omitempty"`
}

// AuthorizeOptions may be provided when authorize an API object.
type AuthorizeOptions struct {
	TypeMeta `json:",inline"`
}

// TableOptions are used when a Table is requested by the caller.
type TableOptions struct {
	TypeMeta `json:",inline"`

	// NoHeaders is only exposed for internal callers. It is not included in our OpenAPI definitions
	// and may be removed as a field in a future release.
	NoHeaders bool `json:"-"`
}
