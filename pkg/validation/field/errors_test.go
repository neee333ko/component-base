package field

import (
	"testing"

	utilerror "github.com/neee333ko/errors"
)

func TestErrorPrint(t *testing.T) {
	type Struct struct {
		err  *Error
		want string
	}

	test := []Struct{
		{
			err:  NotFound(NewPath("Struct1.Field1"), 123),
			want: "Struct1.Field1: ErrorNotFound: 123",
		},
		{
			err:  Required(NewPath("Struct1.Field1[1]")),
			want: "Struct1.Field1[1]: ErrorRequired",
		},
		{
			err:  Duplicate(NewPath("Struct1.Field1[1].Field2"), 123),
			want: "Struct1.Field1[1].Field2: ErrorDuplicate: 123",
		},
		{
			err:  Invalid(NewPath("Struct1.Field2.Struct2.Field1"), "abc", "this is a invalid value"),
			want: "Struct1.Field2.Struct2.Field1: ErrorInvalid: abc: this is a invalid value",
		},
		{
			err:  NotSupport(NewPath("Struct1.Field1"), 123, []string{"a", "b", "c"}),
			want: `Struct1.Field1: ErrorNotSupport: 123: supported values: "a", "b", "c"`,
		},
		{
			err:  Forbidden(NewPath("Struct1.Field1"), "this field must be int under this condition"),
			want: "Struct1.Field1: ErrorForbidden: this field must be int under this condition",
		},
		{
			err:  TooLong(NewPath("Struct1.Field1"), 99999999, 123),
			want: "Struct1.Field1: ErrorTooLong: must have at most 123 bytes",
		},
		{
			err:  TooMany(NewPath("Struct1.Field1"), 1000, 100),
			want: "Struct1.Field1: ErrorTooMany: 1000: must have at most 100",
		},
		{
			err:  Internal(NewPath("Struct1.Field1"), utilerror.New("newerror")),
			want: "Struct1.Field1: ErrorInternal: newerror",
		},
	}

	for _, tt := range test {
		if tt.err.Error() != tt.want {
			t.Errorf("Error() string has error: want: %v got: %v\n", tt.want, tt.err.Error())
		}
	}
}

func TestErrorListPrint(t *testing.T) {
	type Struct struct {
		errlist ErrorList
		want    string
	}

	tests := []Struct{
		{
			errlist: ErrorList{
				Invalid(NewPath("Struct1.Field1"), 123, ""),
				Invalid(NewPath("Struct1.Field1[0]"), 123, "invalid"),
			},
			want: "[1]Struct1.Field1: ErrorInvalid: 123; [2]Struct1.Field1[0]: ErrorInvalid: 123: invalid",
		},
		{
			errlist: ErrorList{
				Invalid(NewPath("Struct1.Field1"), 123, ""),
				Invalid(NewPath("Struct1.Field1[0]"), 123, "invalid"),
				Invalid(NewPath("Struct2.Field1"), "abc", ""),
			},
			want: "[1]Struct1.Field1: ErrorInvalid: 123; [2]Struct1.Field1[0]: ErrorInvalid: 123: invalid; [3]Struct2.Field1: ErrorInvalid: abc",
		},
	}

	for _, tt := range tests {
		if result := tt.errlist.ToAggregate().Error(); result != tt.want {
			t.Errorf("Error() string has error: want:%v got:%v\n", tt.want, result)
		}
	}
}
