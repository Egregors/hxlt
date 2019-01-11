package diff

import "testing"

func Test_diff(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0:45", args{0, 45}, 45},
		{"0:180", args{0, 180}, 180},
		{"0:190", args{0, 190}, 170},
		{"120:280", args{120, 280}, 160},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diff(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("diff() = %v, want %v", got, tt.want)
			}
		})
	}
}
