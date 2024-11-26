package _matrix

import (
	"reflect"
	"testing"
)

func TestMultiply(t *testing.T) {
	type args struct {
		a [][]int
		b [][]int
	}
	type testCase struct {
		name string
		args args
		want [][]int
	}
	tests := []testCase{
		{
			name: "1",
			args: args{
				a: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
				b: [][]int{
					{10, 11},
					{20, 21},
					{30, 31},
				},
			},
			want: [][]int{
				{140, 146},
				{320, 335},
			},
		},
		{
			name: "2",
			args: args{
				a: [][]int{
					{5, 6, 7},
				},
				b: [][]int{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			want: [][]int{
				{5, 6, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
