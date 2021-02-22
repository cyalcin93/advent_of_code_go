package advent_of_code_go

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestDay22(t *testing.T) {
	Equal(t, Day22([]int {9, 2, 6, 3, 1}, []int {5, 8, 4, 7, 10}), 306)
}

func TestDay22Part2(t *testing.T) {
	Equal(t, Day22([]int {9, 2, 6, 3, 1}, []int {5, 8, 4, 7, 10}), 306)
}

func TestDay22MyInput(t *testing.T) {
	Equal(t, Day22(
		[]int {39, 15, 13, 23, 12, 49, 36, 44, 8, 21, 28, 37, 40, 42, 6, 47, 2, 38, 18,	31,	20,	10,	16,	43,	5},
		[]int {29, 26, 19, 35, 34, 4, 41, 11, 3, 50, 33, 22, 48, 7, 17, 32, 27, 45, 46, 9, 25, 30, 1, 24, 14}),
		35202)
}
