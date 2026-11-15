package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "zero size", size: 0},
		{name: "negative size", size: -1},
		{name: "single element", size: 1},
		{name: "multiple elements", size: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			expectedLength := tt.size
			if expectedLength < 0 {
				expectedLength = 0
			}

			if len(data) != expectedLength {
				t.Fatalf(
					"expected length %d, got %d",
					expectedLength,
					len(data),
				)
			}

			for i, value := range data {
				if value <= 0 {
					t.Errorf(
						"element %d is not positive: %d",
						i,
						value,
					)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "maximum at beginning",
			data: []int{99, 2, 3, 4},
			want: 99,
		},
		{
			name: "maximum in middle",
			data: []int{1, 2, 99, 4},
			want: 99,
		},
		{
			name: "maximum at end",
			data: []int{1, 2, 3, 99},
			want: 99,
		},
		{
			name: "equal elements",
			data: []int{7, 7, 7, 7},
			want: 7,
		},
		{
			name: "negative elements",
			data: []int{-10, -3, -20},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf(
					"maximum(%v) = %d, want %d",
					tt.data,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "less than eight elements",
			data: []int{3, 9, 1, 5},
			want: 9,
		},
		{
			name: "exactly eight elements",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			name: "size not divisible by eight",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99},
			want: 99,
		},
		{
			name: "maximum at beginning",
			data: []int{99, 2, 3, 4, 5, 6, 7, 8},
			want: 99,
		},
		{
			name: "maximum in middle",
			data: []int{1, 2, 3, 4, 99, 6, 7, 8},
			want: 99,
		},
		{
			name: "equal elements",
			data: []int{7, 7, 7, 7, 7, 7, 7, 7},
			want: 7,
		},
		{
			name: "negative elements",
			data: []int{-10, -3, -20, -4, -5, -6, -7, -8},
			want: -3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf(
					"maxChunks(%v) = %d, want %d",
					tt.data,
					got,
					tt.want,
				)
			}
		})
	}
}
