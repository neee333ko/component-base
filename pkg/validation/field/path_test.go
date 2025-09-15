package field

import (
	"strconv"
	"testing"
)

func TestPath(t *testing.T) {
	type Struct struct {
		names      []string
		wantString string
	}

	tests := []Struct{
		{names: []string{"struct1", "field1", "struct2", "field2"}, wantString: "struct1.field1.struct2.field2"},
		{names: []string{"struct1", "field1", "1"}, wantString: "struct1.field1[1]"},
	}

	path1 := NewPath(tests[0].names[0], tests[0].names[1:]...)
	result1 := path1.String()

	if result1 != tests[0].wantString {
		t.Errorf("String(error)string has error:\twant:%v\tgot:%v\n", tests[0].wantString, result1)
	}

	path2 := NewPath(tests[1].names[0], tests[1].names[1])
	i, _ := strconv.Atoi(tests[1].names[2])
	path2 = path2.Index(i)
	result2 := path2.String()

	if result2 != tests[1].wantString {
		t.Errorf("String(error)string has error:\twant:%v\tgot:%v\n", tests[1].wantString, result2)
	}
}
