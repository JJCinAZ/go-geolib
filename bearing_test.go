package geolib

import (
	"fmt"
	"testing"
)

func TestBearing(t *testing.T) {
	type args struct {
		flag int
		φ1   float64
		λ1   float64
		φ2   float64
		λ2   float64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test1", args{0, 32.222417, -110.852072, 32.226925, -110.861259}, "300.12", false},
		{"Test2", args{1, 32.222417, -110.852072, 32.226925, -110.861259}, "120.12", false},
		{"Test3", args{0, 50.116667, 8.683333, 52.516667, 13.383333}, "48.93", false},
		{"Test4", args{1, 52.516667, 13.383333, 50.116667, 8.683333}, "52.60", false},
		{"Straight North", args{0, 32.222417, -110.852072, 32.226925, -110.852072}, "0.00", false},
		{"Return South", args{1, 32.222417, -110.852072, 32.226925, -110.852072}, "180.00", false},
		{"Straight East", args{0, 32.222417, -110.861259, 32.222417, -110.852072}, "90.00", false},
		{"Return West", args{1, 32.222417, -110.861259, 32.222417, -110.852072}, "270.00", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bearing(tt.args.flag, tt.args.φ1, tt.args.λ1, tt.args.φ2, tt.args.λ2)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bearing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sgot := fmt.Sprintf("%.2f", got)
			if sgot != tt.want {
				t.Errorf("Bearing() got = %v, want %v", sgot, tt.want)
			}
		})
	}
}
