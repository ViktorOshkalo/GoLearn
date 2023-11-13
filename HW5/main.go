package main

import (
	"fmt"
	"main/matrix"
)

var players = []Player{
	initPlayer("John", 'X'),
	initPlayer("Mark", 'O'),
	//initPlayer("Richard", 'Y'),
}

var playgroundSize = 3
var moves map[Cell]Player = make(map[Cell]Player)

type Cell struct {
	X int
	Y int
}

type Player struct {
	Name             string
	Symbol           rune
	playgroundMatrix matrix.Matrix // represents playground matrix which contains only palyer moves
}

func initPlayer(name string, symbol rune) Player {
	return Player{
		Name:             name,
		Symbol:           symbol,
		playgroundMatrix: matrix.InitMatrix(playgroundSize, playgroundSize),
	}
}

func isNotEnoughMovesYet() bool {
	var repeatsBeforeFirstCheck = playgroundSize - 1
	var movesBeforeFirstCheck = repeatsBeforeFirstCheck * len(players)
	var firstMoveToCheck = movesBeforeFirstCheck + 1
	return len(moves) < firstMoveToCheck
}

func isAllLinesPartiallyTakenAtLeastByTwoPlayers(iterationsOnMatrix int, lineSelector func(i int, matrix matrix.Matrix) matrix.Line) bool {
	for i := 0; i < iterationsOnMatrix; i++ {
		playersTookALineCount := 0
		for _, player := range players {
			line := lineSelector(i, player.playgroundMatrix)
			if matrix.IsLinePartiallyTaken(line) {
				playersTookALineCount++
			}
		}
		if playersTookALineCount < 2 {
			return false
		}
	}
	return true
}

func moreThenOnePlayer(players []Player,
	predicate func(player Player) bool,
) bool {
	count := 0
	for _, player := range players {
		if predicate(player) {
			count++
		}
	}
	return count > 1
}

func forAllLines(players []Player,
	iterationsOnMatrix int,
	lineSelector func(i int, matrix matrix.Matrix) matrix.Line,
	linesPredicate func(line matrix.Line) bool,
	playersPredicate func(players []Player, predicate func(player Player) bool) bool,
) bool {
	for i := 0; i < iterationsOnMatrix; i++ {

		linesSelectorPredicate := func(player Player) bool {
			line := lineSelector(i, player.playgroundMatrix)
			return linesPredicate(line)
		}

		if !playersPredicate(players, linesSelectorPredicate) {
			return false
		}
	}
	return true
}

func isDraw() (isDraw bool) {
	if isNotEnoughMovesYet() {
		return false
	}

	var rowsSelector = matrix.GetRow
	var columnsSelector = matrix.GetColumn
	var diagonalsSelector = matrix.GetDiagonal

	test := forAllLines(players, playgroundSize, rowsSelector, matrix.IsLinePartiallyTaken, moreThenOnePlayer)
	fmt.Printf("Test: %t", test)

	return isAllLinesPartiallyTakenAtLeastByTwoPlayers(playgroundSize, rowsSelector) &&
		isAllLinesPartiallyTakenAtLeastByTwoPlayers(playgroundSize, columnsSelector) &&
		isAllLinesPartiallyTakenAtLeastByTwoPlayers(2, diagonalsSelector)
}

func isAnyLineFullyTakenByPlayer(player Player, iterationsOnMatrix int, lineSelector func(i int, matrix matrix.Matrix) matrix.Line) bool {
	for i := 0; i < iterationsOnMatrix; i++ {
		line := lineSelector(i, player.playgroundMatrix)
		var isFullyTaken = matrix.IsLineFullyTaken(line)
		if isFullyTaken {
			return true
		}
	}
	return false
}

func isPlayerWiner(player Player) (win bool) {
	if isNotEnoughMovesYet() {
		return false
	}

	var rowsSelector = matrix.GetRow
	var columnsSelector = matrix.GetColumn
	var diagonalsSelector = matrix.GetDiagonal

	return isAnyLineFullyTakenByPlayer(player, playgroundSize, rowsSelector) ||
		isAnyLineFullyTakenByPlayer(player, playgroundSize, columnsSelector) ||
		isAnyLineFullyTakenByPlayer(player, 2, diagonalsSelector)
}

func validateInput(x int, y int) (err error) {
	if x < 1 || x > playgroundSize {
		return fmt.Errorf("x value must be from 1 to %d", playgroundSize)
	}
	if y < 1 || y > playgroundSize {
		return fmt.Errorf("y value must be from 1 to %d", playgroundSize)
	}
	return nil
}

func askForMove(player *Player) {
	var inputX, inputY int
	for {
		fmt.Printf("%s, please enter your move (%c). Format: x y\n", player.Name, player.Symbol)
		fmt.Scanln(&inputX, &inputY)

		if err := validateInput(inputX, inputY); err != nil {
			fmt.Println(err)
			continue
		}

		cell := Cell{X: inputX, Y: inputY}
		if _, exists := moves[cell]; !exists {
			moves[cell] = *player
		} else {
			fmt.Println("Cell is already taken. Choose another one.")
			continue
		}

		player.playgroundMatrix[cell.Y-1][cell.X-1] = 1 // transform from x,y to m,n style
		break
	}
}

func getCell(i int, j int) (symbol rune) {
	m := i + 1
	n := j + 1
	for cell, player := range moves {
		if cell.X == m && cell.Y == n {
			return player.Symbol
		}
	}
	return ' '
}

func printPlayground() {
	var separatorLine string
	for i := 0; i < playgroundSize; i++ {
		separatorLine += "--------"
	}

	for i := 0; i < playgroundSize; i++ {
		var line string
		for j := 0; j < playgroundSize; j++ {
			line += "   "
			line += string(getCell(i, j))
			if j != playgroundSize-1 {
				line += "\t|"
			}
		}
		fmt.Println(line)

		if i != playgroundSize-1 {
			fmt.Println(separatorLine)
		}
	}
}

func printPlayers() {
	fmt.Println("\nPlayers:")
	for _, player := range players {
		fmt.Printf("Player: %s, Symbol: %c\n", player.Name, player.Symbol)
	}
}

func checkSettings() {
	if playgroundSize < 1 {
		panic("Playground size must be grather than 0")
	}

	if len(players) < 1 {
		panic("Player list must be non empty")
	}
}

func main() {
	checkSettings()

	fmt.Println("Game is starting...")
	printPlayers()

	fmt.Println("\nInfo:")
	fmt.Printf("x cordinate - from left to right, from 1 to %d\n", playgroundSize)
	fmt.Printf("y cordinate - from top to bottom, from 1 to %d\n", playgroundSize)
	fmt.Printf("playground size: %d x %d\n", playgroundSize, playgroundSize)
	fmt.Println("\nGood luck!")

	printPlayground()

	var nextPlayerNumber = 0
	for {
		player := players[nextPlayerNumber]
		askForMove(&player)
		printPlayground()

		if win := isPlayerWiner(player); win {
			fmt.Printf("Congrats %s. You are winner!", player.Name)
			break
		}

		if draw := isDraw(); draw {
			fmt.Println("It's DRAW!")
			break
		}

		nextPlayerNumber++
		if nextPlayerNumber == len(players) {
			nextPlayerNumber = 0
		}
	}
}
