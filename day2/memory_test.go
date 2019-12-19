package day2

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		program string
		tokens  []int
	}{
		{"short program", "1,9,10,70,", []int{1, 9, 10, 70}},
		{"non-integer panic", "1,9,f,79,", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tokens == nil {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf(`"%s" should have paniced`, tt.name)
					}
				}()
			}
			tokens := Parse(tt.program)
			if !reflect.DeepEqual(tokens, tt.tokens) {
				t.Errorf("%s - got: %v, want: %v", tt.name, tokens, tt.tokens)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name   string
		tokens []int
		noun   int
		verb   int
		result int
	}{
		{"add", []int{1, 0, 0, 0, 99}, 4, 0, 100},
		{"multiply", []int{2, 0, 0, 0, 99, 0}, 0, 4, 198},
		{"longer", []int{1, 0, 0, 5, 2, 0, 0, 0, 99}, 4, 3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Run(tt.noun, tt.verb, tt.tokens)
			if result != tt.result {
				t.Errorf("%v - got: %d, want: %d", tt.tokens, result, tt.result)
			}
		})
	}
}

func TestInitialize(t *testing.T) {
	tests := []struct {
		name   string
		tokens []int
		result []int
	}{
		{"short", []int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{"medium", []int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{"longer", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Initialize(tt.tokens)
			if !reflect.DeepEqual(result, tt.result) {
				t.Errorf("%v - got: %d, want: %d", tt.tokens, result, tt.result)
			}
		})
	}
}
