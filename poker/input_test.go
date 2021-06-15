package poker

import (
	"testing"
)

func TestInput_Compare(t *testing.T) {
	in := InitMethod()

	type args struct {
		p1 [][]int
		p2 [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "同花顺与顺子",
			args: args{
				p1: [][]int{{1, 1}, {2, 1}, {3, 1}},
				p2: [][]int{{1, 2}, {2, 2}, {3, 3}},
			},
			want: 1,
		},
		{
			name: "两组同花顺A/2/3与A/K/Q",
			args: args{
				p1: [][]int{{1, 1}, {2, 1}, {3, 1}},
				p2: [][]int{{1, 2}, {13, 2}, {12, 2}},
			},
			want: -1,
		},
		{
			name: "两组同花顺K/Q/J与A/K/Q",
			args: args{
				p1: [][]int{{13, 1}, {12, 1}, {11, 1}},
				p2: [][]int{{1, 2}, {13, 2}, {12, 2}},
			},
			want: -1,
		},
		{
			name: "同花与顺子",
			args: args{
				p1: [][]int{{7, 1}, {2, 1}, {3, 1}},
				p2: [][]int{{1, 2}, {2, 2}, {3, 3}},
			},
			want: 1,
		},
		{
			name: "豹子",
			args: args{
				p1: [][]int{{7, 1}, {7, 2}, {7, 4}},
				p2: [][]int{{10, 2}, {10, 1}, {10, 3}},
			},
			want: -1,
		},
		{
			name: "豹子与对子",
			args: args{
				p1: [][]int{{7, 1}, {7, 2}, {7, 4}},
				p2: [][]int{{10, 2}, {10, 1}, {12, 1}},
			},
			want: 1,
		},
		{
			name: "对子",
			args: args{
				p1: [][]int{{7, 1}, {7, 2}, {10, 4}},
				p2: [][]int{{7, 3}, {7, 4}, {12, 1}},
			},
			want: -1,
		},
		{
			name: "单张",
			args: args{
				p1: [][]int{{7, 1}, {2, 2}, {10, 4}},
				p2: [][]int{{7, 3}, {2, 4}, {1, 1}},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := in.Compare(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_search(t *testing.T) {

	in := InitMethod()
	type args struct {
		p [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "同花顺RANK1",
			args: args{
				p: [][]int{{3, 1}, {2, 1}, {4, 1}},
			},
			want: RANK1,
		},
		{
			name: "一对RANK5",
			args: args{
				p: [][]int{{3, 1}, {3, 2}, {4, 1}},
			},
			want: RANK5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := in.search(tt.args.p); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compare(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "大于",
			args: args{3, 2},
			want: 1,
		},
		{
			name: "等于",
			args: args{3, 3},
			want: 0,
		},
		{
			name: "小于",
			args: args{5, 12},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
