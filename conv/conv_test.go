package conv

import (
	"math"
	"testing"
)

func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 32, want: 0},
	}

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvintoCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: -273.15, want: 0},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFarenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 273.2},
		{input: 32, want: 89.6},
	}

	for _, tc := range tests {
		got := CelsiusToFarenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFarenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 373.15, want: 212},
		{input: 0, want: -459.67},
	}

	for _, tc := range tests {
		got := KelvinToFarenheit(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
		{input: 212, want: 373.15},
	}

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if math.Abs(tc.want-got) > 0.01 {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
