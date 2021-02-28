package advent_of_code_go

import (
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestDay22Part1(t *testing.T) {
	Equal(t, 306, Day22Part1(
		[]Round{{Number: 0, Player1Deck: []int {9, 2, 6, 3, 1}, Player2Deck: []int {5, 8, 4, 7, 10}}}))
}

func TestDay22Part1MyInput(t *testing.T) {
	Equal(t, 35202, Day22Part1(
		[]Round{{
			Number: 1,
			Player1Deck: []int {39, 15, 13, 23, 12, 49, 36, 44, 8, 21, 28, 37, 40, 42, 6, 47, 2, 38, 18,	31,	20,	10,	16,	43,	5},
			Player2Deck: []int {29, 26, 19, 35, 34, 4, 41, 11, 3, 50, 33, 22, 48, 7, 17, 32, 27, 45, 46, 9, 25, 30, 1, 24, 14}}}))
}

func TestDay22Part2InfiniteRule(t *testing.T) {
	Equal(t, 105, Day22Part2(Game{
			Number: 1,
			Rounds: []Round{{Number: 1, Player1Deck: []int {43, 19}, Player2Deck: []int {2, 29, 14}}},
		},
	))
}

func TestDay22Part2(t *testing.T) {
	Equal(t, 147, Day22Part2(Game{
			Number: 1,
			Rounds: []Round{{Number: 1, Player1Deck: []int{9, 2, 6, 3, 1}, Player2Deck: []int{5, 8, 4, 7, 10}}},
		},
	))
}


func TestDay22Part2MyInput(t *testing.T) {
	Equal(t, 35202, Day22Part2(Game{
		Number: 1,
		Rounds: []Round{
			{
				Number: 1,
				Player1Deck: []int {39, 15, 13, 23, 12, 49, 36, 44, 8, 21, 28, 37, 40, 42, 6, 47, 2, 38, 18,	31,	20,	10,	16,	43,	5},
				Player2Deck: []int {29, 26, 19, 35, 34, 4, 41, 11, 3, 50, 33, 22, 48, 7, 17, 32, 27, 45, 46, 9, 25, 30, 1, 24, 14}},
			},
		},
	))
}
