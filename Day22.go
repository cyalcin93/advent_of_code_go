package advent_of_code_go

import (
	"fmt"
	"strings"
)

func Day22(player1Deck []int, player2Deck []int) int {
	rounds := []Round{ {
		Number:      1,
		Player1Deck: player1Deck,
		Player2Deck: player2Deck,
	} }

	for len(getLastRound(rounds).Player1Deck) != 0 && len(getLastRound(rounds).Player2Deck) != 0 {
		newRound := getLastRound(rounds)
		newRound.Number += 1
		if newRound.Player1Deck[0] > newRound.Player2Deck[0] {
			var poppedPlayer1Card int
			var poppedPlayer2Card int
			poppedPlayer2Card, newRound.Player2Deck = popTopCards(newRound.Player2Deck)
			poppedPlayer1Card, newRound.Player1Deck = popTopCards(newRound.Player1Deck)
			newRound.Player1Deck = append(newRound.Player1Deck, poppedPlayer1Card, poppedPlayer2Card)
		} else {
			var poppedPlayer1Card int
			var poppedPlayer2Card int
			poppedPlayer2Card, newRound.Player2Deck = popTopCards(newRound.Player2Deck)
			poppedPlayer1Card, newRound.Player1Deck = popTopCards(newRound.Player1Deck)
			newRound.Player2Deck = append(newRound.Player2Deck, poppedPlayer2Card, poppedPlayer1Card)
		}

		rounds = append(rounds, newRound)
	}

	for i, r := range rounds {
		if i == len(rounds) - 1 {
			postGameResults(r)
		} else {
			printRound(r)
		}
	}

	return calculatePoints(rounds[len(rounds)-1])
}

func postGameResults(r Round) {
	fmt.Printf("== Post-game results ==\n")
	fmt.Printf("Player 1's deck: %s\n", arrayToString(r.Player1Deck, ", "))
	fmt.Printf("Player 2's deck: %s\n", arrayToString(r.Player2Deck, ", "))
}

func popTopCards(deck []int) (int, []int) {
	return deck[0], deck[1:]
}

func getLastRound(rounds []Round) Round {
	return rounds[len(rounds)-1]
}

func printRound(round Round) {
	fmt.Printf("-- Round %d --\n", round.Number)
	fmt.Printf("Player 1's deck: %s\n", arrayToString(round.Player1Deck, ", "))
	fmt.Printf("Player 2's deck: %s\n", arrayToString(round.Player2Deck, ", "))
	fmt.Printf("Player 1 plays: %d\n", round.Player1Deck[0])
	fmt.Printf("Player 2 plays: %d\n", round.Player2Deck[0])
	if round.Player1Deck[0] > round.Player2Deck[0] {
		println("Player 1 wins the round!")
	} else {
		println("Player 2 wins the round!")
	}
}

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func calculatePoints(lastRound Round) int {
	winningDeck := winningDeck(lastRound)
	deckIndex := 0
	points := 0
	for i := len(winningDeck); i > 0; i-- {
		points += i * winningDeck[deckIndex]
		deckIndex++
	}

	return points
}

func winningDeck(round Round) []int {
	if len(round.Player2Deck) > 0 {
		return round.Player2Deck
	} else {
		return round.Player1Deck
	}
}

type Round struct {
	Number			int
	Player1Deck		[]int
	Player2Deck		[]int
}