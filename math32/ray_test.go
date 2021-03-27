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
			[]float32{4.0, 2.0},
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
