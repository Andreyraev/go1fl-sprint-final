package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "отрицательный размер", size: -5, wantLen: 0},
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "один элемент", size: 1, wantLen: 1},
		{name: "много элементов", size: 1000, wantLen: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if data == nil {
				t.Fatal("ожидался непустой (не nil) слайс")
			}
			if len(data) != tt.wantLen {
				t.Fatalf("длина слайса = %d, ожидалось %d", len(data), tt.wantLen)
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("элемент %d = %d, ожидалось положительное число", i, value)
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
		{name: "nil слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 1, 2, 3}, want: 9},
		{name: "максимум в середине", data: []int{1, 9, 2, 3}, want: 9},
		{name: "максимум в конце", data: []int{1, 2, 3, 9}, want: 9},
		{name: "одинаковые элементы", data: []int{5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum() = %d, ожидалось %d", got, tt.want)
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
		{name: "nil слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "меньше элементов, чем частей", data: []int{3, 7, 1}, want: 7},
		{name: "ровно CHUNKS элементов", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "максимум в остатке последней части", data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, want: 99},
		{name: "максимум в первой части", data: []int{99, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, want: 99},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Fatalf("maxChunks() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	data := generateRandomElements(100_003)
	if got, want := maxChunks(data), maximum(data); got != want {
		t.Fatalf("maxChunks() = %d, maximum() = %d", got, want)
	}
}
