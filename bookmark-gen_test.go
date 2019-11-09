package main

import "testing"

func Test_formatLine(t *testing.T) {
	type args struct {
		input string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Basic test case", args{"link", 1},
			"<DT><A HREF=\"link\">Link #1</DT>"},
		{"Empty string test case", args{"", 2},
			"<DT><A HREF=\"\">Link #2</DT>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatLine(tt.args.input, tt.args.count); got != tt.want {
				t.Errorf("formatLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
