package math32

import (
	"reflect"
	"testing"
)

func TestSphere_Intersect(t *testing.T) {
	type fields struct {
		Center *Point3
		Radius float32
	}
	type args struct {
		r *Ray
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
				&Point3{0, 0, -0},
				1.0,
			},
			args{
				NewRay(
					&Point3{0, 0, -5},
					&Vector3{0, 0, 1},
				),
			},
			2,
			[]float32{4.0, 6.0},
		},
		{
			"Intersect with tangent",
			fields{
				&Point3{0, 0, -0},
				1.0,
			},
			args{
				NewRay(
					&Point3{0, 1, -5},
					&Vector3{0, 0, 1}),
			},
			2,
			[]float32{5., 5.},
		},
		{
			"Not intersect",
			fields{
				&Point3{0, 0, -0}, 1.0,
			},
			args{
				NewRay(
					&Point3{0, 2, -5},
					&Vector3{0, 0, 1},
				),
			},
			0,
			nil,
		},
		{
			"From the center returns one negative time",
			fields{
				&Point3{0, 0, -0}, 1.0,
			},
			args{
				NewRay(
					&Point3{0, 0, 0},
					&Vector3{0, 0, 1},
				),
			},
			2,
			[]float32{-1., 1.},
		},
		{
			"Sphere behind the ray returns both times as negative",
			fields{
				&Point3{0, 0, -0}, 1.0,
			},
			args{
				NewRay(
					&Point3{0, 0, 5},
					&Vector3{0, 0, 1},
				),
			},
			2,
			[]float32{-6., -4.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Sphere{
				Center: tt.fields.Center,
				Radius: tt.fields.Radius,
			}
			got, got1 := s.Intersect(tt.args.r)
			if got != tt.want {
				t.Errorf("Sphere.Intersect() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Sphere.Intersect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewSphere(t *testing.T) {
	type args struct {
		center *Point3
		radius float32
	}
	tests := []struct {
		name string
		args args
		want *Sphere
	}{
		{
			"Runs corretly",
			args{&Point3{0., 0., 0.}, 1.},
			&Sphere{&Point3{0., 0., 0.}, 1.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSphere(tt.args.center, tt.args.radius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSphere() = %v, want %v", got, tt.want)
			}
		})
	}
}
