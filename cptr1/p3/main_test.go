package perfect

import "testing"

func Test_isPerfect(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"6", args{6}, true},
		{"11", args{11}, false},
		{"28", args{28}, true},
		{"496", args{496}, true},
		{"8128", args{8128}, true},
		{"123", args{123}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfect(tt.args.n); got != tt.want {
				t.Errorf("isPerfect() = %v, want %v", got, tt.want)
			}
		})
	}
}
