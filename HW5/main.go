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
	playgroundMatrix matrix.Matrix // represents playground matrix with only palyer moves
}

func initPlayer(name string, symbol rune) Player {
	return Player{
		Name:             name,
		Symbol:           symbol,
		playgroundMatrix: matrix.InitMatrix(playgroundSize, playgroundSize),
	}
}

func isAllLinesPartiallyTakenByAtLeastTwoPlayers(iterationsOnMatrix int, lineSelector func(i int, matrix matrix.Matrix) []int) bool {
	for i := 0; i < iterationsOnMatrix; i++ {
		playersTookALineCount := 0
		for _, player := range players {
			if matrix.IsLineTaken(lineSelector(i, player.playgroundMatrix)) {
				playersTookALineCount++
			}
		}
		if playersTookALineCount < 2 {
			return false
		}
	}
	return true
}

func isDraw() (isDraw bool) {
	if isNotEnoughMoves() {
		return false
	}

	var rowsSelector = matrix.GetRow
	var columnsSelector = matrix.GetColumn
	var diagonalsSelector = matrix.GetDiagonal

	return isAllLinesPartiallyTakenByAtLeastTwoPlayers(playgroundSize, rowsSelector) &&
		isAllLinesPartiallyTakenByAtLeastTwoPlayers(playgroundSize, columnsSelector) &&
		isAllLinesPartiallyTakenByAtLeastTwoPlayers(2, diagonalsSelector)
}

func isAnyLineFullyTakenByOnePlayer(player Player, iterationsOnMatrix int, lineSelector func(i int, matrix matrix.Matrix) []int) bool {
	for i := 0; i < iterationsOnMatrix; i++ {
		sumLine := matrix.GetSum(lineSelector(i, player.playgroundMatrix))
		if sumLine == playgroundSize {
			return true
		}
	}
	return false
}

func isNotEnoughMoves() bool {
	return len(moves) < (playgroundSize-1)*(len(players))+1
}

func isPlayerWiner(player Player) (win bool) {
	if isNotEnoughMoves() {
		return false
	}

	var rowsSelector = matrix.GetRow
	var columnsSelector = matrix.GetColumn
	var diagonalsSelector = matrix.GetDiagonal

	return isAnyLineFullyTakenByOnePlayer(player, playgroundSize, rowsSelector) ||
		isAnyLineFullyTakenByOnePlayer(player, playgroundSize, columnsSelector) ||
		isAnyLineFullyTakenByOnePlayer(player, 2, diagonalsSelector)
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

		player.playgroundMatrix[cell.Y-1][cell.X-1] = 1
		break
	}
}

func getCellSymbol(i int, j int) (symbol rune) {
	var result []rune = make([]rune, 0)
	for _, player := range players {
		if player.playgroundMatrix[i][j] == 1 { // should be garanteed it's single by 'moves' map
			result = append(result, player.Symbol)
		}
	}
	if len(result) > 1 {
		panic("Collision in player matricies")
	}
	if len(result) == 1 {
		return result[0]
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
			line += string(getCellSymbol(i, j))
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

func main() {
	fmt.Println("Game is starting...")
	printPlayers()

	fmt.Println("\nInfo:")
	fmt.Println("x cordinate - from left to right")
	fmt.Println("y cordinate - from top to bottom")
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
