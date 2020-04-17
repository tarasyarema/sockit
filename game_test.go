package main

import (
	"testing"
)

func TestCheckBoard(t *testing.T) {
	boards := []struct {
		want         bool
		target, size int
		board        []int
	}{
		{want: true, target: 0, size: 3, board: []int{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{want: false, target: 1, size: 3, board: []int{0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{want: true, target: 1, size: 3, board: []int{1, 1, 1, 0, 0, 0, 0, 0, 0}},
		{want: true, target: 2, size: 3, board: []int{2, 1, 0, 2, 0, 0, 2, 0, 0}},
		{want: true, target: 2, size: 3, board: []int{2, 0, 0, 0, 2, 0, 0, 0, 2}},
		{want: true, target: 1, size: 3, board: []int{0, 0, 1, 0, 1, 0, 1, 0, 0}},
		{want: true, target: 1, size: 4, board: []int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{want: true, target: 1, size: 4, board: []int{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0}},
		{want: true, target: 4, size: 4, board: []int{4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 4, 0, 0, 0, 0, 4}},
		{want: true, target: 3, size: 4, board: []int{4, 0, 0, 3, 0, 4, 3, 0, 0, 3, 4, 0, 3, 0, 0, 4}},
	}

	for _, b := range boards {
		got := checkBoard(b.target, b.size, b.board)
		if b.want != got {
			t.Errorf("error: got %v with %v", got, b)
		}
	}
}

func BenchmarkCheckBoard3(b *testing.B) {
	target, size := 1, 3
	board := []int{
		0, 0, 1, 
		0, 1, 0, 
		1, 0, 0,
	}

	for n := 0; n < b.N; n++ {
		checkBoard(target, size, board)
	}
}

func BenchmarkCheckBoard4(b *testing.B) {
	target, size := 3, 4
	board := []int{
		4, 0, 0, 3, 
		0, 4, 3, 0,
		0, 3, 4, 0, 
		3, 0, 0, 4,
	}

	for n := 0; n < b.N; n++ {
		checkBoard(target, size, board)
	}
}

func BenchmarkCheckBoard5(b *testing.B) {
	target, size := 1, 5
	board := []int{
		1, 0, 0, 3, 0,
		4, 1, 0, 0, 3,
		4, 0, 1, 0, 0,
		4, 0, 3, 1, 0,
		4, 0, 3, 0, 1,
	}

	for n := 0; n < b.N; n++ {
		checkBoard(target, size, board)
	}
}
