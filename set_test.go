package goset

import (
	"testing"
)

func TestSet_AddAndContains(t *testing.T) {
	tests := []struct {
		name    string
		element interface{}
		check   interface{}
		want    bool
	}{
		{name: "Add and check string", element: "apple", want: true},
		{name: "Add and check int", element: 42, want: true},
		{name: "Add and check struct", element: struct{ name string }{"test"}, want: true},
		{name: "Add and check int with no hit", element: 42, check: 10, want: false},
		{name: "Add and check string with no hit", element: "apple", check: "banana", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Add(tt.element)
			var got bool
			if tt.check != nil {
				got = s.Contains(tt.check)
			} else {
				got = s.Contains(tt.element)
			}
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_RemoveAndContains(t *testing.T) {
	tests := []struct {
		name    string
		element interface{}
		want    bool
	}{
		{name: "Remove and check string", element: "apple", want: false},
		{name: "Remove and check int", element: 42, want: false},
		{name: "Remove and check struct", element: struct{ name string }{"test"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.Add(tt.element)
			s.Remove(tt.element)
			if got := s.Contains(tt.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Size(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		want     int
	}{
		{name: "Empty set", elements: []interface{}{}, want: 0},
		{name: "Single element", elements: []interface{}{"apple"}, want: 1},
		{name: "Multiple elements", elements: []interface{}{"apple", 42, struct{ name string }{"test"}}, want: 3},
		{name: "Duplicate elements", elements: []interface{}{"apple", "apple", 42, 42}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			for _, element := range tt.elements {
				s.Add(element)
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_AddAllAndContains(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		check    interface{}
		want     bool
	}{
		{name: "AddAll and check string", elements: []interface{}{"apple", "banana"}, check: "banana", want: true},
		{name: "AddAll and check int", elements: []interface{}{1, 2, 3}, check: 2, want: true},
		{name: "AddAll and check struct", elements: []interface{}{struct{ name string }{"test1"}, struct{ name string }{"test2"}}, check: struct{ name string }{"test2"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.AddAll(tt.elements)
			if got := s.Contains(tt.check); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_AddAllAndSize(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		want     int
	}{
		{name: "AddAll with empty slice", elements: []interface{}{}, want: 0},
		{name: "AddAll with single element", elements: []interface{}{"apple"}, want: 1},
		{name: "AddAll with multiple elements", elements: []interface{}{"apple", 42, struct{ name string }{"test"}}, want: 3},
		{name: "AddAll with duplicate elements", elements: []interface{}{"apple", "apple", 42, 42}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New()
			s.AddAll(tt.elements)
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
