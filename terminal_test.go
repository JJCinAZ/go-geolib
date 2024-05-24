package geolib

import (
	"fmt"
	"testing"
)

func TestTerminal(t *testing.T) {
	type args struct {
		φ1       float64
		λ1       float64
		distance float64
		bearing  float64
	}
	tests := []struct {
		name   string
		args   args
		wantφ2 string
		wantΛ2 string
	}{
		{"North", args{32.251006, -110.738187, 2, 0}, "32.269", "-110.738"},
		{"NorthEast30", args{32.251006, -110.738187, 1, 30}, "32.259", "-110.733"},
		{"NorthEast40", args{32.251006, -110.738187, 6, 40}, "32.292", "-110.697"},
		{"East", args{32.251006, -110.738187, 5, 90}, "32.251", "-110.685"},
		{"West", args{32.251006, -110.738187, 5, 270}, "32.251", "-110.791"},
		{"South", args{32.251006, -110.738187, 2, 180}, "32.233", "-110.738"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotφ2, gotΛ2 := Terminal(tt.args.φ1, tt.args.λ1, tt.args.distance, tt.args.bearing)
			sgotφ2 := fmt.Sprintf("%.3f", gotφ2)
			sgotΛ2 := fmt.Sprintf("%.3f", gotΛ2)
			if sgotφ2 != tt.wantφ2 {
				t.Errorf("Terminal() gotΦ2 = %s, want %s", sgotφ2, tt.wantφ2)
			}
			if sgotΛ2 != tt.wantΛ2 {
				t.Errorf("Terminal() gotΛ2 = %s, want %s", sgotΛ2, tt.wantΛ2)
			}
		})
	}
}
