package advent_of_code_go

import (
	"fmt"
	"reflect"
	"strings"
)

func Day22Part1(rounds []Round) int {
	playAGamePart1(&rounds)
	postGameResults(rounds[len(rounds)-1])
	return calculatePoints(rounds[len(rounds)-1])
}

func Day22Part2(game Game) int {
	fmt.Printf("=== Game %d ===\n", game.Number)
	games := []Game{ game }
	playAGamePart2(&games)
	postGameResults(game.Rounds[len(game.Rounds)-1])
	return calculatePoints(game.Rounds[len(game.Rounds)-1])
}

func playAGamePart1(rounds *[]Round) {
	for len(getLastRound(*rounds).Player1Deck) != 0 && len(getLastRound(*rounds).Player2Deck) != 0 {
		newRound := getLastRound(*rounds)
		if len(*rounds) != 1 {
			newRound.Number =+ 1
		}
		*rounds = append(*rounds, newRound)
		printRound(newRound, 1)
		printWinner(newRound, 1)

		player1Card, player2Card := newRound.Player1Deck[0], newRound.Player2Deck[0]
		popTopCards(&newRound)

		if player1Card > player2Card {
			newRound.Player1Deck = append(newRound.Player1Deck, player1Card, player2Card)
		} else {
			newRound.Player2Deck = append(newRound.Player2Deck, player2Card, player1Card)
		}
	}
}

func playAGamePart2(games *[]Game) {
	currentGame := &(*games)[len(*games)-1]
	for len(getLastRound(currentGame.Rounds).Player1Deck) != 0 && len(getLastRound(currentGame.Rounds).Player2Deck) != 0 {
		currentGameRound := &currentGame.Rounds[len(currentGame.Rounds)-1]
		printRound(*currentGameRound, currentGame.Number)
		player1Card, player2Card := currentGameRound.Player1Deck[0], currentGameRound.Player2Deck[0]
		popTopCards(currentGameRound)

		if shouldPlayASubGame(*currentGameRound, player1Card, player2Card) {
			println("Playing a sub-games to determine the winner...")
			fmt.Printf("\n=== Game %d ===\n\n", currentGame.Number + 1)
			playASubGame(games, player1Card, player2Card)
			currentGameRound.ChildGame = &(*games)[len(*games)-1]
			switch currentGameRound.ChildGame.Winner {
			case 1:
				fmt.Printf("The winner of game %d is player 1!\n", currentGameRound.ChildGame.Number)
				currentGameRound.Player1Deck = append(currentGameRound.Player1Deck, player1Card, player2Card)
			case 2:
				fmt.Printf("The winner of game %d is player 2!\n", currentGameRound.ChildGame.Number)
				currentGameRound.Player2Deck = append(currentGameRound.Player2Deck, player2Card, player1Card)
			}
			fmt.Printf("...anyway, back to game %d.\n", currentGame.Number)
		} else if player1Card > player2Card  {
			currentGameRound.Player1Deck = append(currentGameRound.Player1Deck, player1Card, player2Card)
		} else {
			currentGameRound.Player2Deck = append(currentGameRound.Player2Deck, player2Card, player1Card)
		}
		printWinner(*currentGameRound, currentGame.Number)

		if roundHasAppearedBefore(currentGame.Rounds) {
			break
		}

		addNewRound(currentGame)
	}

	(*games)[len(*games)-1].Winner = getWinner(getLastRound((*games)[len(*games)-1].Rounds))
}

func addNewRound(game *Game) {
	newRound := getLastRound(game.Rounds)
	newRound.Number += 1
	newRound.ChildGame = nil
	game.Rounds = append(game.Rounds, newRound)
}

func playASubGame(games *[]Game, player1Card int, player2Card int) {
	latestRound := (*games)[len(*games)-1].Rounds[len((*games)[len(*games)-1].Rounds)-1]
	newGame := Game{Number: (*games)[len(*games)-1].Number + 1, Rounds: []Round{}}
	newGame.Rounds = append(newGame.Rounds, Round {
		Game: &newGame,
		Number: 1,
		Player1Deck: latestRound.Player1Deck[:player1Card],
		Player2Deck: latestRound.Player2Deck[:player2Card],
	})

	*games = append(*games, newGame)
	playAGamePart2(games)
}

func getWinner(r Round) int {
	if r.ChildGame != nil && r.ChildGame.Winner != 0 {
		return r.ChildGame.Winner
	}
	if len(r.Player2Deck) > 0 && len(r.Player1Deck) > 0 {
		return 1
	}
	if len(r.Player2Deck) > 0 {
		return 2
	} else {
		return 1
	}
}

func shouldPlayASubGame(r Round, player2Card, player1Card int) bool {
	if player1Card <= len(r.Player1Deck) && player2Card <= len(r.Player2Deck) {
		return true
	}

	return false
}

func roundHasAppearedBefore(rounds []Round) bool {
	previousRounds, lastRound := rounds[1:], rounds[0]
	for _, r := range previousRounds {
		if reflect.DeepEqual(lastRound.Player1Deck, r.Player1Deck) &&
			reflect.DeepEqual(lastRound.Player2Deck, r.Player2Deck) {
			return true
		}
	}

	return false
}

func postGameResults(r Round) {
	fmt.Printf("== Post-game results ==\n")
	fmt.Printf("Player 1's deck: %s\n", arrayToString(r.Player1Deck, ", "))
	fmt.Printf("Player 2's deck: %s\n", arrayToString(r.Player2Deck, ", "))
}

func popTopCards(r *Round) {
	(*r).Player1Deck = (*r).Player1Deck[1:]
	(*r).Player2Deck = (*r).Player2Deck[1:]
}

func getLastRound(rounds []Round) Round {
	return rounds[len(rounds)-1]
}

func printRound(round Round, game int) {
	fmt.Printf("-- Round %d (Game %d) --\n", round.Number, game)
	fmt.Printf("Player 1's deck: %s\n", arrayToString(round.Player1Deck, ", "))
	fmt.Printf("Player 2's deck: %s\n", arrayToString(round.Player2Deck, ", "))
	fmt.Printf("Player 1 plays: %d\n", round.Player1Deck[0])
	fmt.Printf("Player 2 plays: %d\n", round.Player2Deck[0])
}

func printWinner(round Round, game int) {
	winner := getWinner(round)
	fmt.Printf("Player %d wins round %d of game %d!\n\n", winner, round.Number, game)
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
	//if round.GameWinner == 1 {
	//	return round.Player1Deck
	//}
	//if round.GameWinner == 2 {
	//	return round.Player2Deck
	//}
	if len(round.Player2Deck) > 0 && len(round.Player1Deck) > 0 {
		return round.Player1Deck
	}
	if len(round.Player2Deck) > 0 {
		return round.Player2Deck
	} else {
		return round.Player1Deck
	}
}

type Round struct {
	Game			*Game
	ChildGame		*Game
	Number			int
	Player1Deck		[]int
	Player2Deck		[]int
}

type Game struct {
	Number	int
	Winner	int
	Rounds	[]Round
}