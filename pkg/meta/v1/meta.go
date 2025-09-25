package v1

import (
	"time"

	"github.com/neee333ko/component-base/pkg/scheme"
)

type ObjectAccessor interface {
	GetObject() Object
}

type Object interface {
	GetID() uint
	SetID(uint)
	GetName() string
	SetName(string)
	GetCreatedAt() time.Time
	SetCreatedAt(time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(time.Time)
}

var _ Object = &ObjectMeta{}
var _ ObjectAccessor = &ObjectMeta{}

func (o *ObjectMeta) GetID() uint              { return o.ID }
func (o *ObjectMeta) SetID(id uint)            { o.ID = id }
func (o *ObjectMeta) GetName() string          { return o.Name }
func (o *ObjectMeta) SetName(name string)      { o.Name = name }
func (o *ObjectMeta) GetCreatedAt() time.Time  { return o.CreatedAt }
func (o *ObjectMeta) SetCreatedAt(t time.Time) { o.CreatedAt = t }
func (o *ObjectMeta) GetUpdatedAt() time.Time  { return o.UpdatedAt }
func (o *ObjectMeta) SetUpdatedAt(t time.Time) { o.UpdatedAt = t }

func (o *ObjectMeta) GetObject() Object { return o }

type List interface {
	SetTotalCount(int64)
	GetTotalCount() int64
}

var _ List = &ListMeta{}

func (l *ListMeta) SetTotalCount(c int64) { l.TotalCount = c }
func (l *ListMeta) GetTotalCount() int64  { return l.TotalCount }

type Type interface {
	GetVersion() string
	SetVersion(string)
	GetKind() string
	SetKind(string)
}

var _ Type = &TypeMeta{}
var _ scheme.ObjectKind = &TypeMeta{}

func (t *TypeMeta) GetVersion() string  { return t.ApiVersion }
func (t *TypeMeta) SetVersion(v string) { t.ApiVersion = v }
func (t *TypeMeta) GetKind() string     { return t.Type }
func (t *TypeMeta) SetKind(k string)    { t.Type = k }

func (t *TypeMeta) GetGroupVersionKind() *scheme.GroupVersionKind {
	return scheme.FromAPIVersionAndKind(t.ApiVersion, t.Type)
}

func (t *TypeMeta) SetGroupVersionKind(gvk *scheme.GroupVersionKind) {
	t.ApiVersion, t.Type = gvk.ToAPIVersionAndKind()
}

func (t *TypeMeta) GetObjectKind() scheme.ObjectKind { return t }
