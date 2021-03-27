package math32

import (
	"reflect"
	"testing"
)

func TestNewIntersections(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Intersections
	}{
		{
			"Runs corretly",
			args{10},
			&Intersections{0, make([]Intersection, 10)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntersections(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersections_Add(t *testing.T) {
	type fields struct {
		Count  uint32
		Memory []Intersection
	}
	type args struct {
		new *Intersection
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Intersections
	}{
		{
			"Runs corretly",
			fields{
				10,
				[]Intersection{},
			},
			args{&Intersection{}},
			&Intersections{11, []Intersection{{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Intersections{
				Count:  tt.fields.Count,
				Memory: tt.fields.Memory,
			}
			if got := i.Add(tt.args.new); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersections.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
