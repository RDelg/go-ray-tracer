package math32

import (
	"reflect"
	"testing"
)

func TestRay_Position(t *testing.T) {
	type fields struct {
		Origin    *Point3
		Direction *Vector3
	}
	type args struct {
		time float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"Runs corretly",
			fields{
				&Point3{1., 2., 3.},
				&Vector3{4., 5., 6.},
			},
			args{2.0},
			&Point3{9., 12., 15.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				Origin:    tt.fields.Origin,
				Direction: tt.fields.Direction,
			}
			if got := r.Position(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ray.Position() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRay(t *testing.T) {
	type args struct {
		origin    *Point3
		direction *Vector3
	}
	tests := []struct {
		name string
		args args
		want *Ray
	}{
		{
			"Runs corretly",
			args{
				&Point3{1., 2., 3.},
				&Vector3{4., 5., 6.},
			},
			&Ray{
				&Point3{1., 2., 3.},
				&Vector3{4., 5., 6.},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRay(tt.args.origin, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRay() = %v, want %v", got, tt.want)
			}
		})
	}
}
