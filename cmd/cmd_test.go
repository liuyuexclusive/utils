package cmd

import "testing"

func TestRun(t *testing.T) {
	type args struct {
		name string
		arg  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{args: args{name: "pwd"}, want: "/Users/wen/golang/utils/cmd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.name, tt.args.arg...); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
