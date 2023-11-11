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

func isDraw() (isDraw bool) {
	if isNotEnoughMovesYet() {
		return false
	}

	var rowsSelector = matrix.GetRow
	var columnsSelector = matrix.GetColumn
	var diagonalsSelector = matrix.GetDiagonal

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

func isNotEnoughMovesYet() bool {
	var repeatsBeforeFirstCheck = playgroundSize - 1
	var movesBeforeFirstCheck = repeatsBeforeFirstCheck * len(players)
	var firstMoveToCheck = movesBeforeFirstCheck + 1
	return len(moves) < firstMoveToCheck
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

func getCellSymbolFromPlayers(i int, j int) (symbol rune) {
	var result []rune = make([]rune, 0)
	for _, player := range players {
		if player.playgroundMatrix[i][j] == 1 {
			result = append(result, player.Symbol)
		}
	}
	if len(result) > 1 {
		panic("Collision in player matricies") // should be prevented by 'moves' map
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
			line += string(getCellSymbolFromPlayers(i, j))
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
		panic("Playground size must be grather than 1")
	}

	if len(players) < 1 {
		panic("Players count must be greather than 1")
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
