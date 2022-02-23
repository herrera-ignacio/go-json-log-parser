package parser

import (
	"reflect"
	"testing"
)

func Test_parseJson(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  interface{}
	}{
		{"empty string", "", nil},
		{
			"json string",
			`{"name": "Nacho"}`,
			map[string]interface{}{
				"name": "Nacho",
			},
		},
		{
			"nested json",
			`{"name": "Nacho", "pet": { "name": "Pipo" }}`,
			map[string]interface{}{
				"name": "Nacho",
				"pet": map[string]interface{}{
					"name": "Pipo",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseJson(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getJsonRegex(t *testing.T) {
	type args struct {
		depth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"depth 0",
			args{depth: 0},
			"{(?:[^{}]|)*}",
		},
		{
			"depth 1",
			args{depth: 1},
			"{(?:[^{}]|(?:{(?:[^{}]|)*}))*}",
		},
		{
			"depth 2",
			args{depth: 2},
			"{(?:[^{}]|(?:{(?:[^{}]|(?:{(?:[^{}]|)*}))*}))*}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getJsonRegex(tt.args.depth); got != tt.want {
				t.Errorf("getJsonRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllJSONStrings(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{"empty string", "", nil},
		{
			"text with single json at the end",
			`this is just a test{"name": "Nacho"}`,
			[]string{`{"name": "Nacho"}`},
		},
		{
			"text with single json at the middle",
			`this is just{"name": "Nacho"}a test`,
			[]string{`{"name": "Nacho"}`},
		},
		{
			"text with line breaks and single json at the middle",
			`this is\n\n just{"name": "Nacho"}\na test`,
			[]string{`{"name": "Nacho"}`},
		},
		{
			"text with multiple json",
			`this is a test\n {"name": "Nacho"}{"age": 23}\nLet's see what happens { "name": "Roman" }`,
			[]string{`{"name": "Nacho"}`, `{"age": 23}`, `{ "name": "Roman" }`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllJSONStrings(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllJSONStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
