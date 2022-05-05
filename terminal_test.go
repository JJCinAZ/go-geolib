package geolib

import (
	"math"
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
		wantΦ2 float64
		wantΛ2 float64
	}{
		{"North", args{32.251006, -110.738187, 2, 0}, 32.268995, -110.738187},
		{"NorthEast30", args{32.251006, -110.738187, 1, 30}, 32.258795, -110.732862},
		{"NorthEast40", args{32.251006, -110.738187, 6, 40}, 32.292340, -110.697144},
		{"East", args{32.251006, -110.738187, 5, 90}, 32.250995, -110.685004},
		{"West", args{32.251006, -110.738187, 5, 270}, 32.250995, -110.791356},
		{"South", args{32.251006, -110.738187, 2, 180}, 32.233017, -110.738180},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotΦ2, gotΛ2 := Terminal(tt.args.φ1, tt.args.λ1, tt.args.distance, tt.args.bearing)
			if testround(gotΦ2) != testround(tt.wantΦ2) {
				t.Errorf("Terminal() gotΦ2 = %v, want %v", testround(gotΦ2), testround(tt.wantΦ2))
			}
			if testround(gotΛ2) != testround(tt.wantΛ2) {
				t.Errorf("Terminal() gotΛ2 = %v, want %v", testround(gotΛ2), testround(tt.wantΛ2))
			}
		})
	}
}

func testround(x float64) float64 {
	return math.Round(x*1000000) / 1000000
}
