package math32

import (
	"math"
	"reflect"
	"testing"
)

func TestVector4FromArray(t *testing.T) {
	type args struct {
		array [4]float32
	}
	tests := []struct {
		name string
		args args
		want *Vector4
	}{
		{"Vector4FromArray runs correctly", args{[4]float32{0.1, 0.1, 0.1, 0.1}}, &Vector4{0.1, 0.1, 0.1, 0.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vector4FromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4FromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_AsArray(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	tests := []struct {
		name   string
		fields fields
		want   [4]float32
	}{
		{"AsArray runs correctly", fields{0.1, 0.1, 0.1, 0.1}, [4]float32{0.1, 0.1, 0.1, 0.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.AsArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.AsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_Neg(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	tests := []struct {
		name   string
		fields fields
		want   *Vector4
	}{
		{"Neg runs correctly", fields{0.1, -0.3, 0.4, 2.1}, &Vector4{-0.1, 0.3, -0.4, -2.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_Add(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		vb *Vector4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector4
	}{
		{"Add 01", fields{0.1, 0.2, 0.2, 0.0}, args{&Vector4{-0.1, 0.1, -0.1, 0.1}}, &Vector4{0.0, 0.3, 0.1, 0.1}},
		{"Add 02", fields{-0.3, 0.3, 3.3, 2.2}, args{&Vector4{0.3, 4.0, 33.0, 1.0}}, &Vector4{0.0, 4.3, 36.3, 3.2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.Add(tt.args.vb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_Substract(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		vb *Vector4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector4
	}{
		{"Substract 01", fields{0.1, 0.2, 0.2, 0.1}, args{&Vector4{-0.1, 0.1, -0.1, 1.0}}, &Vector4{0.2, 0.1, 0.3, -0.9}},
		{"Substract 02", fields{-0.3, 0.3, 3.3, 2.3}, args{&Vector4{0.3, 4.0, 33.0, 3.3}}, &Vector4{-0.6, -3.7, -29.7, -1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.Substract(tt.args.vb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.Substract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_MultScalar(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector4
	}{
		{"MultScalar 01", fields{0.1, 0.2, 0.6, 1.0}, args{2.0}, &Vector4{0.2, 0.4, 1.2, 2.0}},
		{"MultScalar 02", fields{-0.34, 0.33, 3.3, 1.0}, args{-3.0}, &Vector4{1.02, -0.99, -9.9, -3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.MultScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_DivideScalar(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector4
	}{
		{"DivideScalar 01", fields{2.0, 4.0, 10.0, 5.0}, args{2.0}, &Vector4{1.0, 2.0, 5.0, 2.5}},
		{"DivideScalar 02", fields{-15.0, -3.0, 9.0, 15.0}, args{-3.0}, &Vector4{5.0, 1.0, -3.0, -5.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.DivideScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.DivideScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_Magnitude(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{"Magnitude 01", fields{0.0, 1.0, 0.0, 0.0}, 1.0},
		{"Magnitude 02", fields{0.0, 0.0, 1.0, 0.0}, 1.0},
		{"Magnitude 03", fields{1.0, 2.0, 3.0, 0.0}, float32(math.Sqrt(14.0))},
		{"Magnitude 03", fields{-1.0, -2.0, -3.0, 0.0}, float32(math.Sqrt(14.0))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.Magnitude(); got != tt.want {
				t.Errorf("Vector4.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_Normalize(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	sqrt14 := float32(math.Sqrt(float64(14.0)))
	tests := []struct {
		name   string
		fields fields
		want   *Vector4
	}{
		{"Normalize 01", fields{4.0, 0.0, 0.0, 0.0}, &Vector4{1.0, 0.0, 0.0, 0.0}},
		{"Normalize 02", fields{1.0, 2.0, 3.0, 0.0}, &Vector4{1.0 / sqrt14, 2.0 / sqrt14, 3.0 / sqrt14, 0.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_Dot(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		vB *Vector4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float32
	}{
		{"Dot 01", fields{1.0, 2.0, 3.0, 4.0}, args{&Vector4{2.0, 3.0, 4.0, 4.0}}, 36.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vA := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vA.Dot(tt.args.vB); got != tt.want {
				t.Errorf("Vector4.Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector4_SetFromArray(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		array [4]float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector4
	}{
		{"SetFromArray runs correctly", fields{}, args{[4]float32{0.1, 0.2, 0.3, 0.4}}, &Vector4{0.1, 0.2, 0.3, 0.4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vector := &Vector4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := vector.SetFromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector4.SetFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
