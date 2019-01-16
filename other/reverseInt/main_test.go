package reverseint

import "testing"

func Test_r(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"13", args{13}, 31},
		{"-123", args{-123}, -321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := r(tt.args.n); got != tt.want {
				t.Errorf("r() = %v, want %v", got, tt.want)
			}
		})
	}
}
