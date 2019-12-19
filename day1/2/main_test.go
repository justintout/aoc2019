package main

import (
	"fmt"
	"testing"
)

func TestModuleFuel(t *testing.T) {
	tests := []struct {
		mass     int
		required int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("mass %d", tt.mass), func(t *testing.T) {
			required := moduleFuel(tt.mass)
			if required != tt.required {
				t.Errorf("got %d, want %d", required, tt.required)
			}
		})
	}
}

func TestMassFuel(t *testing.T) {
	tests := []struct {
		mass     int
		required int
	}{
		{14, 2},
		{2, 0},
		{1969, 654},
		{654, 216},
		{216, 70},
		{100756, 33583},
		{0, 0},
		{43, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("mass %d", tt.mass), func(t *testing.T) {
			required := massFuel(tt.mass)
			if required != tt.required {
				t.Errorf("got: %d, want: %d", required, tt.required)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := []struct {
		in    int
		floor int
		out   int
	}{
		{14, 1, 14},
		{-2, 0, 0},
		{1235412351, 9999999999999, 9999999999999},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d floor %d", tt.in, tt.floor), func(t *testing.T) {
			val := floor(tt.in, tt.floor)
			if val != tt.out {
				t.Errorf("got: %d, want: %d", val, tt.out)
			}
		})
	}
}
