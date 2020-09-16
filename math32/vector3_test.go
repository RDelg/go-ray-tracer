package math32

import (
	"math"
	"reflect"
	"testing"
)

func TestVector3FromArray(t *testing.T) {
	type args struct {
		array [3]float32
	}
	tests := []struct {
		name string
		args args
		want *Vector3
	}{
		{"Vector3FromArray runs correctly", args{[3]float32{0.1, 0.1, 0.1}}, &Vector3{0.1, 0.1, 0.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vector3FromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3FromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_AsArray(t *testing.T) {
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
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.AsArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.AsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Neg(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector3
	}{
		{"Neg runs correctly", fields{0.1, -0.3, 0.4}, &Vector3{-0.1, 0.3, -0.4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Add(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		vb *Vector3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{"Add 01", fields{0.1, 0.2, 0.2}, args{&Vector3{-0.1, 0.1, -0.1}}, &Vector3{0.0, 0.3, 0.1}},
		{"Add 02", fields{-0.3, 0.3, 3.3}, args{&Vector3{0.3, 4.0, 33.0}}, &Vector3{0.0, 4.3, 36.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Add(tt.args.vb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Substract(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		vb *Vector3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{"Substract 01", fields{0.1, 0.2, 0.2}, args{&Vector3{-0.1, 0.1, -0.1}}, &Vector3{0.2, 0.1, 0.3}},
		{"Substract 02", fields{-0.3, 0.3, 3.3}, args{&Vector3{0.3, 4.0, 33.0}}, &Vector3{-0.6, -3.7, -29.7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Substract(tt.args.vb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_MultScalar(t *testing.T) {
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
		want   *Vector3
	}{
		{"MultScalar 01", fields{0.1, 0.2, 0.6}, args{2.0}, &Vector3{0.2, 0.4, 1.2}},
		{"MultScalar 02", fields{-0.34, 0.33, 3.3}, args{-3.0}, &Vector3{1.02, -0.99, -9.9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.MultScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_DivideScalar(t *testing.T) {
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
		want   *Vector3
	}{
		{"DivideScalar 01", fields{2.0, 4.0, 10.0}, args{2.0}, &Vector3{1.0, 2.0, 5.0}},
		{"DivideScalar 02", fields{-15.0, -3.0, 9.0}, args{-3.0}, &Vector3{5.0, 1.0, -3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.DivideScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.DivideScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Magnitude(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{"Magnitude 01", fields{0.0, 1.0, 0.0}, 1.0},
		{"Magnitude 02", fields{0.0, 0.0, 1.0}, 1.0},
		{"Magnitude 03", fields{1.0, 2.0, 3.0}, float32(math.Sqrt(14.0))},
		{"Magnitude 03", fields{-1.0, -2.0, -3.0}, float32(math.Sqrt(14.0))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Magnitude(); got != tt.want {
				t.Errorf("Vector3.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Normalize(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	sqrt14 := float32(math.Sqrt(float64(14.0)))
	tests := []struct {
		name   string
		fields fields
		want   *Vector3
	}{
		{"Normalize 01", fields{4.0, 0.0, 0.0}, &Vector3{1.0, 0.0, 0.0}},
		{"Normalize 02", fields{1.0, 2.0, 3.0}, &Vector3{1.0 / sqrt14, 2.0 / sqrt14, 3.0 / sqrt14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Dot(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		vB *Vector3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{"Dot 01", fields{1.0, 2.0, 3.0}, args{&Vector3{2.0, 3.0, 4.0}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Dot(tt.args.vB); got != tt.want {
				t.Errorf("Vector3.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_Cross(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
	}
	type args struct {
		vB *Vector3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{"Cross 01", fields{1.0, 2.0, 3.0}, args{&Vector3{2.0, 3.0, 4.0}}, &Vector3{-1.0, 2.0, -1.0}},
		{"Cross 01", fields{2.0, 3.0, 4.0}, args{&Vector3{1.0, 2.0, 3.0}}, &Vector3{1.0, -2.0, 1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vA.Cross(tt.args.vB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.Cross() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector3_SetFromArray(t *testing.T) {
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
		want   *Vector3
	}{
		{"SetFromArray runs correctly", fields{}, args{[3]float32{0.1, 0.2, 0.3}}, &Vector3{0.1, 0.2, 0.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector3{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := vector.SetFromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector3.SetFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
