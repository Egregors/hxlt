package point

import (
	"testing"
)

func Test_getQuadtant(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0;0", args{makePoint(0, 0)}, 0},
		{"5;0", args{makePoint(5, 0)}, 0},
		{"1;5", args{makePoint(1, 5)}, 1},
		{"-3;10", args{makePoint(-3, 10)}, 2},
		{"-2;-5", args{makePoint(-2, -5)}, 3},
		{"4;-1", args{makePoint(4, -1)}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getQuadtant(tt.args.p); got != tt.want {
				t.Errorf("getQuadtant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateDistance(t *testing.T) {
	type args struct {
		p1 *Point
		p2 *Point
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDistance(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("calculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
