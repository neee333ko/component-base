package scheme

type ObjectKind interface {
	SetGroupVersionKind(*GroupVersionKind)
	GetGroupVersionKind() *GroupVersionKind
}

var EmptyObjectKind = &emptyObjectKind{}

type emptyObjectKind struct{}

func (e *emptyObjectKind) SetGroupVersionKind(gvk *GroupVersionKind) {}
func (e *emptyObjectKind) GetGroupVersionKind() *GroupVersionKind    { return nil }
