package math32

import (
	"reflect"
	"testing"
)

func TestNewMatrix4Ones(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix4
	}{
		{
			"Return a new Matrix4 full of 1s",
			&Matrix4{[4][4]float32{
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{1.0, 1.0, 1.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix4Ones(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix4Ones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix4Zeros(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix4
	}{
		{
			"Return a new Matrix4 full of 0s",
			&Matrix4{[4][4]float32{
				[4]float32{0.0, 0.0, 0.0, 0.0},
				[4]float32{0.0, 0.0, 0.0, 0.0},
				[4]float32{0.0, 0.0, 0.0, 0.0},
				[4]float32{0.0, 0.0, 0.0, 0.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix4Zeros(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix4Zeros() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix4Identity(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix4
	}{
		{
			"Return the identity matrix",
			&Matrix4{[4][4]float32{
				[4]float32{1.0},
				[4]float32{0.0, 1.0},
				[4]float32{0.0, 0.0, 1.0},
				[4]float32{0.0, 0.0, 0.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix4Identity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix4Identity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Add(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		other *Matrix4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix4
	}{
		{
			"Adds runs correctly 01",
			fields{},
			args{
				&Matrix4{[4][4]float32{
					[4]float32{1.0},
					[4]float32{0.0, 1.0},
					[4]float32{0.0, 0.0, 1.0},
					[4]float32{0.0, 0.0, 0.0, 1.0},
				}}},
			&Matrix4{[4][4]float32{
				[4]float32{1.0},
				[4]float32{0.0, 1.0},
				[4]float32{0.0, 0.0, 1.0},
				[4]float32{0.0, 0.0, 0.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Mult(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		other *Matrix4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix4
	}{
		{
			"Mult runs correctly 01",
			fields{NewMatrix4Ones().Values},
			args{
				&Matrix4{[4][4]float32{
					[4]float32{1.0},
					[4]float32{0.0, 1.0},
					[4]float32{0.0, 0.0, 1.0},
					[4]float32{0.0, 0.0, 0.0, 1.0},
				}}},
			&Matrix4{[4][4]float32{
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{1.0, 1.0, 1.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.Mult(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Transpose(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   *Matrix4
	}{
		{
			"Transpose of identity returns identity",
			fields{NewMatrix4Identity().Values},
			&Matrix4{[4][4]float32{
				[4]float32{1.0},
				[4]float32{0.0, 1.0},
				[4]float32{0.0, 0.0, 1.0},
				[4]float32{0.0, 0.0, 0.0, 1.0},
			}},
		},
		{
			"Transpose runs correctly",
			fields{
				[4][4]float32{
					[4]float32{1.0, -1.0, 2.0, 3.0},
					[4]float32{1.0, 1.0, 6.0, 0.0},
					[4]float32{3.0, 4.0, 1.0, 1.0},
					[4]float32{3.0, 4.0, 1.0, 1.0},
				},
			},
			&Matrix4{[4][4]float32{
				[4]float32{1.0, 1.0, 3.0, 3.0},
				[4]float32{-1.0, 1.0, 4.0, 4.0},
				[4]float32{2.0, 6.0, 1.0, 1.0},
				[4]float32{3.0, 0.0, 1.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Determinant(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			"Determinannt of identity returns 1.0",
			fields{NewMatrix4Identity().Values},
			1.0,
		},
		{
			"Determinannt of runs correctly",
			fields{
				[4][4]float32{
					[4]float32{1.0, -1.0, 2.0, 1.0},
					[4]float32{1.0, 1.0, 6.0, 1.0},
					[4]float32{3.0, 4.0, 1.0, 1.0},
					[4]float32{4.0, 2.0, 1.0, 6.0},
				},
			},
			-152.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.Determinant(); got != tt.want {
				t.Errorf("Matrix4.Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Equal(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		other *Matrix4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"Equal returns true",
			fields{
				[4][4]float32{
					[4]float32{1.0, -1.0, 2.0, 1.0},
					[4]float32{1.0, 1.0, 6.0, 1.0},
					[4]float32{3.0, 4.0, 1.0, 1.0},
					[4]float32{6.0, 3.0, 2.0, 1.0},
				},
			},
			args{&Matrix4{[4][4]float32{
				[4]float32{1.0, -1.0, 2.0, 1.0},
				[4]float32{1.0, 1.0, 6.0, 1.0},
				[4]float32{3.0, 4.0, 1.0, 1.0},
				[4]float32{6.0, 3.0, 2.0, 1.0},
			}}},
			true,
		},
		{
			"Equal returns False",
			fields{
				[4][4]float32{
					[4]float32{1.0, -1.0, 2.0, 1.0},
					[4]float32{1.0, 1.0, 6.0, 1.0},
					[4]float32{3.0, 4.0, 1.0, 1.0},
					[4]float32{6.0, 3.0, 2.0, 1.0},
				},
			},
			args{&Matrix4{[4][4]float32{
				[4]float32{1.0, -1.0, 2.0, 1.0},
				[4]float32{1.0, 1.0, 6.0, 1.0},
				[4]float32{3.0, 4.0, 3.0, 1.0},
				[4]float32{6.0, 3.0, 2.0, 1.0},
			}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.Equal(tt.args.other); got != tt.want {
				t.Errorf("Matrix4.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_AddVector4(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		vector     *Vector4
		columnWise bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix4
	}{
		{
			"AddVector4 runs correctly with column wise false",
			fields{},
			args{&Vector4{1.0, 2.0, 3.0, 4.0}, false},
			&Matrix4{[4][4]float32{
				[4]float32{1.0, 2.0, 3.0, 4.0},
				[4]float32{1.0, 2.0, 3.0, 4.0},
				[4]float32{1.0, 2.0, 3.0, 4.0},
				[4]float32{1.0, 2.0, 3.0, 4.0},
			}},
		},
		{
			"AddVector4 runs correctly with column wise true",
			fields{},
			args{&Vector4{1.0, 2.0, 3.0, 4.0}, true},
			&Matrix4{[4][4]float32{
				[4]float32{1.0, 1.0, 1.0, 1.0},
				[4]float32{2.0, 2.0, 2.0, 2.0},
				[4]float32{3.0, 3.0, 3.0, 3.0},
				[4]float32{4.0, 4.0, 4.0, 4.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.AddVector4(tt.args.vector, tt.args.columnWise); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.AddVector4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_MultVector4(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		vector     *Vector4
		columnWise bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector4
	}{
		{
			"MultVector4 runs correctly with column wise false",
			fields{NewMatrix4Ones().Values},
			args{&Vector4{1.0, 2.0, 3.0, 4.0}, false},
			&Vector4{10.0, 10.0, 10.0, 10.0},
		},
		{
			"MultVector4 runs correctly with column wise true",
			fields{NewMatrix4Ones().Values},
			args{&Vector4{1.0, 2.0, 3.0, 4.0}, true},
			&Vector4{10.0, 10.0, 10.0, 10.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.MultVector4(tt.args.vector, tt.args.columnWise); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.MultVector4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_AddScalar(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix4
	}{
		{
			"AddScalar runs correctly",
			fields{NewMatrix4Ones().Values},
			args{2.0},
			&Matrix4{[4][4]float32{
				[4]float32{3.0, 3.0, 3.0, 3.0},
				[4]float32{3.0, 3.0, 3.0, 3.0},
				[4]float32{3.0, 3.0, 3.0, 3.0},
				[4]float32{3.0, 3.0, 3.0, 3.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.AddScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.AddScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_MultScalar(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix4
	}{
		{
			"MultScalar runs correctly",
			fields{NewMatrix4Ones().Values},
			args{2.0},
			&Matrix4{[4][4]float32{
				[4]float32{2.0, 2.0, 2.0, 2.0},
				[4]float32{2.0, 2.0, 2.0, 2.0},
				[4]float32{2.0, 2.0, 2.0, 2.0},
				[4]float32{2.0, 2.0, 2.0, 2.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.MultScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_DivScalar(t *testing.T) {
	type fields struct {
		Values [4][4]float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix4
	}{
		{
			"DivScalar runs correctly",
			fields{NewMatrix4Ones().Values},
			args{2.0},
			&Matrix4{[4][4]float32{
				[4]float32{0.5, 0.5, 0.5, 0.5},
				[4]float32{0.5, 0.5, 0.5, 0.5},
				[4]float32{0.5, 0.5, 0.5, 0.5},
				[4]float32{0.5, 0.5, 0.5, 0.5},
			}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix4{
				Values: tt.fields.Values,
			}
			if got := m.DivScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix4.DivScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}
