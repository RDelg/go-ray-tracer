package math32

import (
	"reflect"
	"testing"
)

func TestColorFromName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{"darkmagenta returns {0.545, 0.000, 0.545}", args{"darkmagenta"}, &Color{0.545, 0.000, 0.545}},
		{"not existing color name returns nil", args{"nullcolor"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorFromName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_SetRGB(t *testing.T) {
	type fields struct {
		R float32
		G float32
		B float32
	}
	type args struct {
		r float32
		g float32
		b float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{"SetRGB runs correctly 01", fields{}, args{0.1, 0.2, 0.3}, &Color{0.1, 0.2, 0.3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := c.SetRGB(tt.args.r, tt.args.g, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.SetRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_SetByName(t *testing.T) {
	type fields struct {
		R float32
		G float32
		B float32
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
		want1  bool
	}{
		{"set darkslateblue returns {0.282, 0.239, 0.545}", fields{}, args{"darkslateblue"}, &Color{0.282, 0.239, 0.545}, true},
		{"set nullcolor don't change color and return false as ok", fields{}, args{"nullcolor"}, &Color{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			got, got1 := c.SetByName(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.SetByName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Color.SetByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestColor_Add(t *testing.T) {
	type fields struct {
		R float32
		G float32
		B float32
	}
	type args struct {
		other *Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{"Add adds correctly 01", fields{0.1, 0.2, 0.3}, args{&Color{0.1, 0.2, 0.3}}, &Color{0.2, 0.4, 0.6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := c.Add(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_AddScalar(t *testing.T) {
	type fields struct {
		R float32
		G float32
		B float32
	}
	type args struct {
		s float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{"AddScalar adds correctly 01", fields{0.1, 0.2, 0.3}, args{0.1}, &Color{0.2, 0.3, 0.4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := c.AddScalar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.AddScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Mult(t *testing.T) {
	type fields struct {
		R float32
		G float32
		B float32
	}
	type args struct {
		other *Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{"Mult multiplies correctly 01", fields{0.25, 0.8, 0.3}, args{&Color{0.4, 0.16, 0.16}}, &Color{0.1, 0.128, 0.048}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := c.Mult(tt.args.other); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_MultScalar(t *testing.T) {
	type fields struct {
		R float32
		G float32
		B float32
	}
	type args struct {
		s float32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{"Mult multiplies correctly 01", fields{0.25, 0.46, 0.38}, args{0.4}, &Color{0.1, 0.184, 0.152}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				R: tt.fields.R,
				G: tt.fields.G,
				B: tt.fields.B,
			}
			if got := c.MultScalar(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Color.MultScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}
