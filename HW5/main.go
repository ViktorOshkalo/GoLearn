package main

import (
	"fmt"
	"main/matrix"
)

const (
	playgroundSize int = 3
	diagonalsCount int = 2
)

var players = []Player{
	initPlayer("John", 'X'),
	initPlayer("Mark", 'O'),
	//initPlayer("Richard", 'Y'),
}

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

func isAllLinesPartiallyTakenByAtLeastTwoPlayers() bool {
	// check columns
	for i := 0; i < playgroundSize; i++ {
		playersTookALineCount := 0
		for _, player := range players {
			column := player.playgroundMatrix.GetColumn(i)
			if matrix.IsLinePartiallyTaken(column) {
				playersTookALineCount++
			}
		}
		if playersTookALineCount < 2 {
			return false
		}
	}

	// check rows
	for i := 0; i < playgroundSize; i++ {
		playersTookALineCount := 0
		for _, player := range players {
			row := player.playgroundMatrix.GetRow(i)
			if matrix.IsLinePartiallyTaken(row) {
				playersTookALineCount++
			}
		}
		if playersTookALineCount < 2 {
			return false
		}
	}

	// check diagonal left ro right
	playersTookALineCount := 0
	for _, player := range players {
		line := player.playgroundMatrix.GetDiagonalLeftToRight()
		if matrix.IsLinePartiallyTaken(line) {
			playersTookALineCount++
		}
	}
	if playersTookALineCount < 2 {
		return false
	}

	// check diagonal right to left
	playersTookALineCount = 0
	for _, player := range players {
		line := player.playgroundMatrix.GetDiagonalLeftToRight()
		if matrix.IsLinePartiallyTaken(line) {
			playersTookALineCount++
		}
	}
	if playersTookALineCount < 2 {
		return false
	}

	return true
}

func isDraw() (isDraw bool) {
	if isNotEnoughMovesYet() {
		return false
	}

	return isAllLinesPartiallyTakenByAtLeastTwoPlayers()
}

func isAnyLineFullyTakenByPlayer() bool {
	// check rows
	for i := 0; i < playgroundSize; i++ {
		for _, player := range players {
			row := player.playgroundMatrix.GetRow(i)
			if matrix.IsLineFullyTaken(row) {
				return true
			}
		}
	}

	// check columns
	for i := 0; i < playgroundSize; i++ {
		for _, player := range players {
			column := player.playgroundMatrix.GetColumn(i)
			if matrix.IsLineFullyTaken(column) {
				return true
			}
		}
	}

	// check diagonal left to right
	playersTookALineCount := 0
	for _, player := range players {
		diagonal := player.playgroundMatrix.GetDiagonalLeftToRight()
		if matrix.IsLineFullyTaken(diagonal) {
			return true
		}
	}
	if playersTookALineCount < 2 {
		return false
	}

	// check diagonal right to left
	for _, player := range players {
		diagonal := player.playgroundMatrix.GetDiagonalLeftToRight()
		if matrix.IsLineFullyTaken(diagonal) {
			return true
		}
	}

	return false
}

func isPlayerWiner(player Player) (win bool) {
	if isNotEnoughMovesYet() {
		return false
	}

	return isAnyLineFullyTakenByPlayer()
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
	for cell, player := range moves {
		if cell.X == i && cell.Y == j {
			return player.Symbol
		}
	}
	return ' '
}

func printPlayground() {
	var separatorLine string
	for i := 1; i <= playgroundSize; i++ {
		separatorLine += "--------"
	}

	for j := 1; j <= playgroundSize; j++ {
		var line string
		for i := 1; i <= playgroundSize; i++ {
			line += "   "
			line += string(getCell(i, j))
			if i != playgroundSize {
				line += "\t|"
			}
		}

		fmt.Println(line)

		if j != playgroundSize {
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
