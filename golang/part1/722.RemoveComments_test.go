package part1

import (
	"reflect"
	"testing"
)

func Test_removeComments(t *testing.T) {
	type args struct {
		source []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "item1",
			args: args{
				source: []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"},
			},
			want: []string{"int main()", "{ ", "  ", "int a, b, c;", "a = b + c;", "}"},
		},
		{
			name: "item2",
			args: args{
				source: []string{"a/*comment", "line", "more_comment*/b"},
			},
			want: []string{"ab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeComments(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
