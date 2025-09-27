package stringutil

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		base    []string
		exclude []string
		want    []string
	}{
		{
			base:    []string{"1", "2", "3"},
			exclude: []string{"3", "4", "5"},
			want:    []string{"1", "2"},
		},
		{
			base:    []string{"1", "2"},
			exclude: []string{"1", "2", "3"},
			want:    []string{},
		},
	}

	for _, tt := range tests {
		res := Diff(tt.base, tt.exclude)

		if !reflect.DeepEqual(res, tt.want) {
			t.Errorf("Diff has an error: got:%v want:%v\n", res, tt.want)
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		ss   []string
		want []string
	}{
		{
			ss:   []string{"1", "2", "1"},
			want: []string{"1", "2"},
		},
		{
			ss:   []string{},
			want: []string{},
		},
	}

	for _, tt := range tests {
		res := Unique(tt.ss)

		if !reflect.DeepEqual(res, tt.want) {
			t.Errorf("Unique has an error: got:%v want:%v\n", res, tt.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{
			str:  "我爱你",
			want: "你爱我",
		},
		{
			str:  "I love you",
			want: "uoy evol I",
		},
	}

	for _, tt := range tests {
		res := Reverse(tt.str)

		if res != tt.want {
			t.Errorf("Reverse has an error: got:%v want%v\n", res, tt.want)
		}
	}
}
