package main

import "testing"

func TestView_isDriver(t *testing.T) {
	type args struct {
		templateName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "driver",
			args: args{templateName: "index.html"},
			want: true,
		},
		{
			name: "helper",
			args: args{templateName: "partial/header.html"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := View{}
			if got := v.isDriver(tt.args.templateName); got != tt.want {
				t.Errorf("View.isDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}
