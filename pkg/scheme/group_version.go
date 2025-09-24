package scheme

import (
	"fmt"
	"strings"
)

func ParseGroupVersionResource(gvr string) (*GroupVersionResource, *GroupResource) {
	parts := strings.SplitN(gvr, ".", 3)

	var gvrptr *GroupVersionResource

	if len(parts) == 3 {
		gvrptr = &GroupVersionResource{Group: parts[2], Version: parts[1], Resource: parts[0]}
	}

	return gvrptr, ParseGroupResource(gvr)
}

func ParseGroupVersionKind(gvk string) (*GroupVersionKind, *GroupKind) {
	parts := strings.SplitN(gvk, ".", 3)

	var gvkptr *GroupVersionKind

	if len(parts) == 3 {
		gvkptr = &GroupVersionKind{Group: parts[2], Version: parts[1], Kind: parts[0]}
	}

	return gvkptr, ParseGroupKind(gvk)
}

func ParseGroupResource(gr string) *GroupResource {
	if index := strings.Index(gr, "."); index != -1 {
		return &GroupResource{Group: gr[index+1:], Resource: gr[0:index]}
	}

	return &GroupResource{Resource: gr}
}

func ParseGroupKind(gk string) *GroupKind {
	if index := strings.Index(gk, "."); index != -1 {
		return &GroupKind{Group: gk[index+1:], Kind: gk[0:index]}
	}

	return &GroupKind{Kind: gk}
}

func ParseGroupVersion(gv string) (*GroupVersion, error) {
	if len(gv) == 0 || gv == "/" {
		return nil, nil
	}

	switch strings.Count(gv, "/") {
	case 0:
		return &GroupVersion{Version: gv}, nil
	case 1:
		index := strings.Index(gv, "/")
		return &GroupVersion{Group: gv[0:index], Version: gv[index+1:]}, nil
	default:
		return nil, fmt.Errorf("unrecognized GroupVersion string")
	}
}

type GroupResource struct {
	Group    string
	Resource string
}

func (gr *GroupResource) WithVersion(version string) *GroupVersionResource {
	return &GroupVersionResource{Group: gr.Group, Version: version, Resource: gr.Resource}
}

func (gr *GroupResource) Empty() bool {
	return gr.Group == "" && gr.Resource == ""
}

func (gr *GroupResource) String() string {
	if gr.Group == "" {
		return gr.Resource
	}

	return gr.Resource + "." + gr.Group
}

type GroupVersionResource struct {
	Group    string
	Version  string
	Resource string
}

func (gvr *GroupVersionResource) GroupResource() *GroupResource {
	return &GroupResource{Group: gvr.Group, Resource: gvr.Resource}
}

func (gvr *GroupVersionResource) GroupVersion() *GroupVersion {
	return &GroupVersion{Group: gvr.Group, Version: gvr.Version}
}

func (gvr *GroupVersionResource) Empty() bool {
	return gvr.Group == "" && gvr.Version == "" && gvr.Resource == ""
}

func (gvr *GroupVersionResource) String() string {
	return gvr.Group + "/" + gvr.Version + ", Resource:" + gvr.Resource
}

type GroupKind struct {
	Group string
	Kind  string
}

func (gk *GroupKind) Empty() bool {
	return gk.Group == "" && gk.Kind == ""
}

func (gk *GroupKind) WithVersion(version string) *GroupVersionKind {
	return &GroupVersionKind{Group: gk.Group, Version: version, Kind: gk.Kind}
}

func (gk *GroupKind) String() string {
	return gk.Kind + "." + gk.Group
}

type GroupVersionKind struct {
	Group   string
	Version string
	Kind    string
}

func (gvk *GroupVersionKind) Empty() bool {
	return gvk.Group == "" && gvk.Version == "" && gvk.Kind == ""
}

func (gvk *GroupVersionKind) GroupKind() *GroupKind {
	return &GroupKind{Group: gvk.Group, Kind: gvk.Kind}
}

func (gvk *GroupVersionKind) GroupVersion() *GroupVersion {
	return &GroupVersion{Group: gvk.Group, Version: gvk.Version}
}

func (gvk *GroupVersionKind) String() string {
	return gvk.Group + "/" + gvk.Version + ", Kind:" + gvk.Kind
}

type GroupVersion struct {
	Group   string
	Version string
}

func (gv *GroupVersion) Empty() bool {
	return gv.Group == "" && gv.Version == ""
}

func (gv *GroupVersion) String() string {
	return gv.Group + "/" + gv.Version
}

func (gv *GroupVersion) Identifier() string {
	return gv.String()
}

func (gv *GroupVersion) WithKind(kind string) *GroupVersionKind {
	return &GroupVersionKind{Group: gv.Group, Version: gv.Version, Kind: kind}
}

func (gv *GroupVersion) WithResource(resource string) *GroupVersionResource {
	return &GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: resource}
}

func (gv *GroupVersion) KindForGroupVersionKind(gvks []*GroupVersionKind) (*GroupVersionKind, bool) {
	for _, gvk := range gvks {
		if gvk.Group == gv.Group && gvk.Version == gv.Version {
			return gvk, true
		}
	}

	for _, gvk := range gvks {
		if gvk.Group == gv.Group {
			return gv.WithKind(gvk.Kind), true
		}
	}

	return nil, false
}

type GroupVersions []GroupVersion

func (gvs GroupVersions) String() string {
	s := make([]string, 0, len(gvs))

	for _, gv := range gvs {
		s = append(s, gv.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(s, ", "))
}

func (gvs GroupVersions) Identifier() string {
	return gvs.String()
}

func (gvs GroupVersions) KindForGroupVersionKind(gvks []*GroupVersionKind) (*GroupVersionKind, bool) {
	targets := make([]*GroupVersionKind, 0)

	for _, gv := range gvs {
		target, ok := gv.KindForGroupVersionKind(gvks)
		if !ok {
			continue
		}

		targets = append(targets, target)
	}

	if len(targets) == 1 {
		return targets[0], true
	}

	if len(targets) > 1 {
		return BestMatch(gvks, targets), true
	}

	return nil, false
}

func BestMatch(gvks, targets []*GroupVersionKind) *GroupVersionKind {
	for _, gvk := range gvks {
		for _, target := range targets {
			if gvk == target {
				return target
			}
		}
	}

	return targets[0]
}

func (gvk *GroupVersionKind) ToAPIVersionAndKind() (string, string) {
	return gvk.GroupVersion().String(), gvk.Kind
}

func FromAPIVersionAndKind(apiversion, kind string) *GroupVersionKind {
	gv, err := ParseGroupVersion(apiversion)
	if err != nil {
		return &GroupVersionKind{Kind: kind}
	}

	return gv.WithKind(kind)
}
