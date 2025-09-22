package validation

import (
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		st   interface{}
		want bool
	}{
		{
			st: &struct {
				Name        string `validate:"name"`
				Description string `validate:"description"`
			}{
				Name:        "my-name",
				Description: "test1",
			},
			want: true,
		},
		{
			st: &struct {
				File string `validate:"file"`
			}{
				File: "non-existing file",
			},
			want: false,
		},
		{
			st: &struct {
				Dir string `validate:"dir"`
			}{
				Dir: "non-existing dir",
			},
			want: false,
		},
		{
			st: &struct {
				Name string `validate:"name"`
			}{
				Name: "-invalid.name",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		v := NewValidator(tt.st)
		errlist := v.Validate()

		if res := (len(errlist) == 0); res != tt.want {
			t.Errorf("validation.go source file has bug: want:%v got:%v\n", tt.want, res)
		}
	}
}
