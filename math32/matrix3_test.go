package math32

import (
	"reflect"
	"testing"
)

func TestNewMatrix3Ones(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix3
	}{
		{
			"Return a new Matrix3 full of 1s",
			&Matrix3{[3][3]float32{
				{1.0, 1.0, 1.0},
				{1.0, 1.0, 1.0},
				{1.0, 1.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix3Ones(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix3Ones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix3Zeros(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix3
	}{
		{
			"Return a new Matrix3 full of 0s",
			&Matrix3{[3][3]float32{
				{0.0, 0.0, 0.0},
				{0.0, 0.0, 0.0},
				{0.0, 0.0, 0.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix3Zeros(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix3Zeros() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix3Identity(t *testing.T) {
	tests := []struct {
		name string
		want *Matrix3
	}{
		{
			"Return the identity matrix",
			&Matrix3{[3][3]float32{
				{1.0, 0.0, 0.0},
				{0.0, 1.0, 0.0},
				{0.0, 0.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix3Identity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix3Identity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_Add(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		other *Matrix3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix3
	}{
		{
			"Adds runs correctly 01",
			fields{},
			args{
				&Matrix3{[3][3]float32{
					{1.0, 0.0, 0.0},
					{0.0, 1.0, 0.0},
					{0.0, 0.0, 1.0},
				}}},
			&Matrix3{[3][3]float32{
				{1.0, 0.0, 0.0},
				{0.0, 1.0, 0.0},
				{0.0, 0.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_Mult(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		other *Matrix3
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix3
	}{
		{
			"Mult runs correctly 01",
			fields{NewMatrix3Ones().Values},
			args{
				&Matrix3{[3][3]float32{
					{1.0, 0.0, 0.0},
					{0.0, 1.0, 0.0},
					{0.0, 0.0, 1.0},
				}}},
			&Matrix3{[3][3]float32{
				{1.0, 1.0, 1.0},
				{1.0, 1.0, 1.0},
				{1.0, 1.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.Mult(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_Transpose(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   *Matrix3
	}{
		{
			"Transpose of identity returns identity",
			fields{NewMatrix3Identity().Values},
			&Matrix3{[3][3]float32{
				{1.0, 0.0, 0.0},
				{0.0, 1.0, 0.0},
				{0.0, 0.0, 1.0},
			}},
		},
		{
			"Transpose runs correctly",
			fields{
				[3][3]float32{
					{1.0, -1.0, 2.0},
					{1.0, 1.0, 6.0},
					{3.0, 4.0, 1.0},
				},
			},
			&Matrix3{[3][3]float32{
				{1.0, 1.0, 3.0},
				{-1.0, 1.0, 4.0},
				{2.0, 6.0, 1.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_Determinant(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	tests := []struct {
		name   string
		fields fields
		want   float32
	}{
		{
			"Determinannt of identity returns 1.0",
			fields{NewMatrix3Identity().Values},
			1.0,
		},
		{
			"Determinannt of runs correctly",
			fields{
				[3][3]float32{
					{1.0, -1.0, 2.0},
					{1.0, 1.0, 6.0},
					{3.0, 4.0, 1.0},
				},
			},
			-38.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.Determinant(); got != tt.want {
				t.Errorf("Matrix3.Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_Equal(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		other *Matrix3
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
				[3][3]float32{
					{1.0, -1.0, 2.0},
					{1.0, 1.0, 6.0},
					{3.0, 4.0, 1.0},
				},
			},
			args{&Matrix3{[3][3]float32{
				{1.0, -1.0, 2.0},
				{1.0, 1.0, 6.0},
				{3.0, 4.0, 1.0},
			}}},
			true,
		},
		{
			"Equal returns False",
			fields{
				[3][3]float32{
					{1.0, -1.0, 2.0},
					{1.0, 1.0, 9.0},
					{3.0, 4.0, 1.0},
				},
			},
			args{&Matrix3{[3][3]float32{
				{1.0, 1.0, 3.0},
				{-1.0, 1.0, 4.0},
				{2.0, 6.0, 1.0},
			}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.Equal(tt.args.other); got != tt.want {
				t.Errorf("Matrix3.Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_AddVector3(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		vector     *Vector3
		columnWise bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix3
	}{
		{
			"AddVector3 runs correctly with column wise true",
			fields{},
			args{&Vector3{1.0, 2.0, 3.0}, true},
			&Matrix3{[3][3]float32{
				{1.0, 2.0, 3.0},
				{1.0, 2.0, 3.0},
				{1.0, 2.0, 3.0},
			}},
		},
		{
			"AddVector3 runs correctly with column wise false",
			fields{},
			args{&Vector3{1.0, 2.0, 3.0}, false},
			&Matrix3{[3][3]float32{
				{1.0, 1.0, 1.0},
				{2.0, 2.0, 2.0},
				{3.0, 3.0, 3.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.AddVector3(tt.args.vector, tt.args.columnWise); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.AddVector3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_MultVector3(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		vector     *Vector3
		columnWise bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Vector3
	}{
		{
			"MultVector3 runs correctly with column wise true",
			fields{NewMatrix3Ones().Values},
			args{&Vector3{1.0, 2.0, 3.0}, true},
			&Vector3{6.0, 6.0, 6.0},
		},
		{
			"MultVector3 runs correctly with column wise false",
			fields{NewMatrix3Ones().Values},
			args{&Vector3{1.0, 2.0, 3.0}, false},
			&Vector3{6.0, 6.0, 6.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.MultVector3(tt.args.vector, tt.args.columnWise); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.MultVector3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_MultPoint3(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		vector     *Point3
		columnWise bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Point3
	}{
		{
			"MultPoint3 runs correctly with column wise true",
			fields{NewMatrix3Ones().Values},
			args{&Point3{1.0, 2.0, 3.0}, true},
			&Point3{6.0, 6.0, 6.0},
		},
		{
			"MultPoint3 runs correctly with column wise false",
			fields{NewMatrix3Ones().Values},
			args{&Point3{1.0, 2.0, 3.0}, false},
			&Point3{6.0, 6.0, 6.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.MultPoint3(tt.args.vector, tt.args.columnWise); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.MultPoint3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_AddScalar(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix3
	}{
		{
			"AddScalar runs correctly",
			fields{NewMatrix3Ones().Values},
			args{2.0},
			&Matrix3{[3][3]float32{
				{3.0, 3.0, 3.0},
				{3.0, 3.0, 3.0},
				{3.0, 3.0, 3.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.AddScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.AddScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_MultScalar(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix3
	}{
		{
			"MultScalar runs correctly",
			fields{NewMatrix3Ones().Values},
			args{2.0},
			&Matrix3{[3][3]float32{
				{2.0, 2.0, 2.0},
				{2.0, 2.0, 2.0},
				{2.0, 2.0, 2.0},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.MultScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_DivScalar(t *testing.T) {
	type fields struct {
		Values [3][3]float32
	}
	type args struct {
		scalar float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix3
	}{
		{
			"DivScalar runs correctly",
			fields{NewMatrix3Ones().Values},
			args{2.0},
			&Matrix3{[3][3]float32{
				{0.5, 0.5, 0.5},
				{0.5, 0.5, 0.5},
				{0.5, 0.5, 0.5},
			}},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix3{
				Values: tt.fields.Values,
			}
			if got := m.DivScalar(tt.args.scalar); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix3.DivScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}
