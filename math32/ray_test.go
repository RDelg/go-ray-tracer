package math32

import (
	"reflect"
	"testing"
)

func TestRay_Intersect(t *testing.T) {
	type fields struct {
		Origin    *Point3
		Direction *Vector3
	}
	type args struct {
		sphere *Sphere
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
		want1  []float32
	}{
		{
			"Intersect in 2 points",
			fields{
				&Point3{0, 0, -5},
				&Vector3{0, 0, 1},
			},
			args{
				NewSphere(&Point3{0, 0, -0}, 1.0),
			},
			2,
			[]float32{4.0, 6.0},
		},
		{
			"Intersect with tangent",
			fields{
				&Point3{0, 1, -5},
				&Vector3{0, 0, 1},
			},
			args{
				NewSphere(&Point3{0, 0, -0}, 1.0),
			},
			2,
			[]float32{5., 5.},
		},
		{
			"Not intersect",
			fields{
				&Point3{0, 2, -5},
				&Vector3{0, 0, 1},
			},
			args{
				NewSphere(&Point3{0, 0, -0}, 1.0),
			},
			0,
			nil,
		},
		{
			"From the center returns one negative time",
			fields{
				&Point3{0, 0, 0},
				&Vector3{0, 0, 1},
			},
			args{
				NewSphere(&Point3{0, 0, -0}, 1.0),
			},
			2,
			[]float32{-1., 1.},
		},
		{
			"Sphere behind the ray returns both times as negative",
			fields{
				&Point3{0, 0, 5},
				&Vector3{0, 0, 1},
			},
			args{
				NewSphere(&Point3{0, 0, -0}, 1.0),
			},
			2,
			[]float32{-6., -4.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				Origin:    tt.fields.Origin,
				Direction: tt.fields.Direction,
			}
			got, got1 := r.Intersect(tt.args.sphere)
			if got != tt.want {
				t.Errorf("Ray.Intersect() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Ray.Intersect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
