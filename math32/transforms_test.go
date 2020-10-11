package math32

import (
	"math"
	"reflect"
	"testing"
)

func TestVector3_Scale(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{
			"Scale runs correctly 01",
			fields{1., 1., 1.},
			args{2., 3., 4.},
			&Vector3{2., 3., 4.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vector.Scale(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_RotateX(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		r float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{
			"RotateX runs correctly 01",
			fields{1., 1., 1.},
			args{math.Pi / 4.},
			&Vector3{1., 0., 1.4142135},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vector.RotateX(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.RotateX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_RotateY(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		r float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{
			"RotateY runs correctly 01",
			fields{1., 1., 1.},
			args{math.Pi / 4.},
			&Vector3{1.4142135, 1., 0.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vector.RotateY(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.RotateY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_RotateZ(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		r float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{
			"RotateZ runs correctly 01",
			fields{0., 1., 0.},
			args{math.Pi / 4},
			&Vector3{-0.70710677, 0.70710677, 0.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vector.RotateZ(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.RotateZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Shear(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		xy float32
		xz float32
		yx float32
		yz float32
		zx float32
		zy float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{
			"RotateZ runs correctly 01",
			fields{2., 3., 4.},
			args{1., 0., 0., 0., 0., 0.},
			&Vector3{5., 3., 4.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vector.Shear(tt.args.xy, tt.args.xz, tt.args.yx, tt.args.yz, tt.args.zx, tt.args.zy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Shear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Scale(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"Scale runs correctly 01",
			fields{1., 1., 1.},
			args{2., 3., 4.},
			&Point3{2., 3., 4.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.Scale(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_RotateX(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		r float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"RotateX runs correctly 01",
			fields{1., 1., 1.},
			args{math.Pi / 4.},
			&Point3{1., 0., 1.4142135},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.RotateX(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.RotateX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_RotateY(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		r float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"RotateY runs correctly 01",
			fields{1., 1., 1.},
			args{math.Pi / 4.},
			&Point3{1.4142135, 1., 0.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.RotateY(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.RotateY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_RotateZ(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		r float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"RotateZ runs correctly 01",
			fields{0., 1., 0.},
			args{math.Pi / 4},
			&Point3{-0.70710677, 0.70710677, 0.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.RotateZ(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.RotateZ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Shear(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		xy float32
		xz float32
		yx float32
		yz float32
		zx float32
		zy float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"RotateZ runs correctly 01",
			fields{2., 3., 4.},
			args{1., 0., 0., 0., 0., 0.},
			&Point3{5., 3., 4.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.Shear(tt.args.xy, tt.args.xz, tt.args.yx, tt.args.yz, tt.args.zx, tt.args.zy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.Shear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Translate(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		x float32
		y float32
		z float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"Translate runs correctly 01",
			fields{0., 0., 0.},
			args{1., 2., 3.},
			&Point3{1., 2., 3.},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.Translate(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
