package space

import (
	"reflect"
	"testing"
)

func TestNewPoint3D(t *testing.T) {
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name string
		args args
		want Point3D
	}{
		{"Creates Point3D correctly", args{0.1, 0.1, 0.1}, Point3D{Tuple4DFloat{[4]float32{0.1, 0.1, 0.1, 1.0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint3D(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint3D() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVector3D(t *testing.T) {
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name string
		args args
		want Vector3D
	}{
		{"Creates Vector3D correctly", args{0.1, 0.1, 0.1}, Vector3D{Tuple4DFloat{[4]float32{0.1, 0.1, 0.1, 0.0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector3D(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector3D() = %v, want %v", got, tt.want)
			}
		})
	}
}
