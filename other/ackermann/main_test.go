package ackermann

import "testing"

func Test_ack(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 0", args{0, 0}, 1},
		{"0 1", args{0, 1}, 2},
		{"0 2", args{0, 2}, 3},
		{"0 3", args{0, 3}, 4},
		{"0 4", args{0, 4}, 5},
		{"1 0", args{1, 0}, 2},
		{"1 1", args{1, 1}, 3},
		{"1 2", args{1, 2}, 4},
		{"1 3", args{1, 3}, 5},
		{"1 4", args{1, 4}, 6},
		{"2 0", args{2, 0}, 3},
		{"2 1", args{2, 1}, 5},
		{"2 2", args{2, 2}, 7},
		{"2 3", args{2, 3}, 9},
		{"2 4", args{2, 4}, 11},
		{"3 0", args{3, 0}, 5},
		{"3 1", args{3, 1}, 13},
		{"3 2", args{3, 2}, 29},
		{"3 3", args{3, 3}, 61},
		{"3 4", args{3, 4}, 125},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ack(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("ack() = %v, want %v", got, tt.want)
			}
		})
	}
}
