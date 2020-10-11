package math32

import (
	"reflect"
	"testing"
)

func TestPoint3FromArray(t *testing.T) {
	type args struct {
		array [3]float32
	}
	tests := []struct {
		name string
		args args
		want *Point3
	}{
		{"Point3FromArray runs correctly", args{[3]float32{0.1, 0.1, 0.1}}, &Point3{0.1, 0.1, 0.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Point3FromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3FromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_AsArray(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	tests := []struct {
		name   string
		fields fields
		want   [3]float32
	}{
		{"AsArray runs correctly", fields{0.1, 0.1, 0.1}, [3]float32{0.1, 0.1, 0.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.AsArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.AsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Neg(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	tests := []struct {
		name   string
		fields fields
		want   *Point3
	}{
		{"Neg runs correctly", fields{0.1, -0.3, 0.4}, &Point3{-0.1, 0.3, -0.4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Add(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		vb *Point3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{"Add 01", fields{0.1, 0.2, 0.2}, args{&Point3{-0.1, 0.1, -0.1}}, &Point3{0.0, 0.3, 0.1}},
		{"Add 02", fields{-0.3, 0.3, 3.3}, args{&Point3{0.3, 4.0, 33.0}}, &Point3{0.0, 4.3, 36.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Add(tt.args.vb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_Substract(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		vb *Point3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{"Substract 01", fields{0.1, 0.2, 0.2}, args{&Point3{-0.1, 0.1, -0.1}}, &Vector3{0.2, 0.1, 0.3}},
		{"Substract 02", fields{-0.3, 0.3, 3.3}, args{&Point3{0.3, 4.0, 33.0}}, &Vector3{-0.6, -3.7, -29.7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Substract(tt.args.vb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.Substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_MultScalar(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{"MultScalar 01", fields{0.1, 0.2, 0.6}, args{2.0}, &Point3{0.2, 0.4, 1.2}},
		{"MultScalar 02", fields{-0.34, 0.33, 3.3}, args{-3.0}, &Point3{1.02, -0.99, -9.9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.MultScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_DivideScalar(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{"DivideScalar 01", fields{2.0, 4.0, 10.0}, args{2.0}, &Point3{1.0, 2.0, 5.0}},
		{"DivideScalar 02", fields{-15.0, -3.0, 9.0}, args{-3.0}, &Point3{5.0, 1.0, -3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.DivideScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.DivideScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint3_SetFromArray(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		array [3]float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{"SetFromArray runs correctly", fields{}, args{[3]float32{0.1, 0.2, 0.3}}, &Point3{0.1, 0.2, 0.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := &Point3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := point.SetFromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Point3.SetFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
