package field

import (
	"bytes"
	"fmt"
	"strconv"
)

type Path struct {
	name   string
	index  string
	parent *Path
}

func NewPath(pname string, elems ...string) *Path {
	p := &Path{
		name:   pname,
		parent: nil,
	}

	for _, elem := range elems {
		p = &Path{name: elem, parent: p}
	}

	return p
}

func (p *Path) Root() *Path {
	for ; p.parent != nil; p = p.parent {
	}

	return p
}

func (p *Path) Child(pname string, elems ...string) *Path {
	c := NewPath(pname, elems...)
	c.Root().parent = p

	return c
}

func (p *Path) Index(i int) *Path {
	c := &Path{index: strconv.Itoa(i), parent: p}
	c.parent = p

	return c
}

func (p *Path) Key(key string) *Path {
	c := &Path{index: key, parent: p}
	c.parent = p

	return c
}

func (p *Path) String() string {
	pathList := make([]*Path, 0)

	for ; p != nil; p = p.parent {
		pathList = append(pathList, p)
	}

	len := len(pathList)

	buffer := bytes.NewBuffer(nil)

	for i := range pathList {
		path := pathList[len-1-i]

		if i >= 1 && path.name != "" {
			buffer.WriteString(".")
		}

		if path.name != "" {
			buffer.WriteString(path.name)
		}

		if path.index != "" {
			buffer.WriteString(fmt.Sprintf("[%s]", path.index))
		}
	}

	return buffer.String()
}
