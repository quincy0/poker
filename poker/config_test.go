package poker

import (
	"reflect"
	"testing"
)

func Test_classifyCards(t *testing.T) {
	type args struct {
		cards [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 []int
	}{
		{
			name: "classifyCards",
			args: args{cards: [][]int{{1,1},{2,2},{3,4}}},
			want: []int{1,2,3},
			want1: []int{1,2,4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := classifyCards(tt.args.cards)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("classifyCards() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("classifyCards() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getRank(t *testing.T) {
	type args struct {
		cards [][]int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "getRank",
			args: args{cards: [][]int{{1,1},{2,2},{3,4}}},
			want: RANK3,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getRank(tt.args.cards)
			if got != tt.want {
				t.Errorf("getRank() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRank() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isPair(t *testing.T) {
	type args struct {
		sortedScores []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isPair",
			args: args{sortedScores: []int{4,5,6}},
			want: false,
		},
		{
			name: "isPair",
			args: args{sortedScores: []int{4,4,6}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPair(tt.args.sortedScores); got != tt.want {
				t.Errorf("isPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSameScore(t *testing.T) {
	type args struct {
		sortedScores []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isSameScore",
			args: args{sortedScores: []int{4,4,4}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameScore(tt.args.sortedScores); got != tt.want {
				t.Errorf("isSameScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isStraight(t *testing.T) {
	type args struct {
		sortedScores []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isStraight-1",
			args: args{sortedScores: []int{1,2,3}},
			want: true,
		},
		{
			name: "isStraight-2",
			args: args{sortedScores: []int{1,12,13}},
			want: true,
		},
		{
			name: "isStraight-3",
			args: args{sortedScores: []int{10,12,13}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isStraight(tt.args.sortedScores); got != tt.want {
				t.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}
