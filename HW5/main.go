package main

import (
	"fmt"
)

var playgroundSize = 4
var moves map[Cell]Player = make(map[Cell]Player)

type Cell struct {
	X int
	Y int
}

type Matrix [][]int

type Player struct {
	Name             string
	Symbol           rune
	playgroundMatrix Matrix
}

func initPlayer(name string, symbol rune) Player {
	return Player{
		Name:             name,
		Symbol:           symbol,
		playgroundMatrix: initMatrix(playgroundSize, playgroundSize),
	}
}

func initMatrix(rows int, columns int) Matrix {
	var matrix Matrix = make([][]int, rows, rows)
	for i := range matrix {
		matrix[i] = make([]int, columns, columns)
	}
	return matrix
}

func (matrix Matrix) getRow(i int) []int {
	return matrix[i]
}

func (matrix Matrix) getColumn(j int) []int {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][j]
	}
	return output
}

func (matrix Matrix) getDiagonal1() []int {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][i]
	}
	return output
}

func (matrix Matrix) getDiagonal2() []int {
	var output []int = make([]int, len(matrix))
	for i := range matrix {
		output[i] = matrix[i][len(matrix)-1-i]
	}
	return output
}

func getSum(arr []int) int {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return sum
}

func isLineTaken(line []int) bool {
	return getSum(line) > 0
}

func checkDraw(players []Player) (isDraw bool) {
	// check rows and colums if at least 2 players are there
	for i := 0; i < playgroundSize; i++ {
		playersOnRow := 0
		playersOnColumn := 0
		for _, player := range players {
			if isLineTaken(player.playgroundMatrix.getRow(i)) {
				playersOnRow++
			}
			if isLineTaken(player.playgroundMatrix.getColumn(i)) {
				playersOnColumn++
			}
		}

		if playersOnRow < 2 || playersOnColumn < 2 {
			return false
		}
	}

	// check diagonals if at least 2 players are there
	for i := 0; i < 2; i++ {
		playersOnDiagonal1 := 0
		playersOnDiagonal2 := 0
		for _, player := range players {
			if isLineTaken(player.playgroundMatrix.getDiagonal1()) {
				playersOnDiagonal1++
			}
			if isLineTaken(player.playgroundMatrix.getDiagonal2()) {
				playersOnDiagonal2++
			}
		}

		if playersOnDiagonal1 < 2 || playersOnDiagonal2 < 2 {
			return false
		}
	}

	return true
}

func checkPlaygroundMatrixForWin(player Player) (win bool) {
	for i := 0; i < playgroundSize; i++ {
		sumRow := getSum(player.playgroundMatrix.getRow(i))
		if sumRow == playgroundSize {
			return true
		}
		sumColumn := getSum(player.playgroundMatrix.getColumn(i))
		if sumColumn == playgroundSize {
			return true
		}
	}

	sumDiagonal1 := getSum(player.playgroundMatrix.getDiagonal1())
	if sumDiagonal1 == playgroundSize {
		return true
	}

	sumDiagonal2 := getSum(player.playgroundMatrix.getDiagonal2())
	if sumDiagonal2 == playgroundSize {
		return true
	}

	return false
}

func getSymbolByPlayerMatricies(players []Player, i int, j int) (symbol rune) {
	for _, player := range players {
		if player.playgroundMatrix[i][j] == 1 {
			return player.Symbol
		}
	}
	return ' '
}

func printPlayground(players []Player) {
	var separatorLine string
	for i := 0; i < playgroundSize; i++ {
		separatorLine += "--------"
	}

	for i := 0; i < playgroundSize; i++ {
		var line string
		for j := 0; j < playgroundSize; j++ {
			line += "   "
			line += string(getSymbolByPlayerMatricies(players, i, j))
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

func printPlayers(players []Player) {
	fmt.Println("\nPlayers:")
	for _, player := range players {
		fmt.Printf("Player: %s, Symbol: %c\n", player.Name, player.Symbol)
	}
}

func validateInput(x int, y int) (err error) {
	if x < 1 || x > playgroundSize {
		return fmt.Errorf("x value must be from 1 to %d.", playgroundSize)
	}
	if y < 1 || y > playgroundSize {
		return fmt.Errorf("y value must be from 1 to %d.", playgroundSize)
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

func main() {
	players := []Player{
		initPlayer("John", 'X'),
		initPlayer("Mark", 'O'),
		initPlayer("Richard", 'Y'),
	}

	fmt.Println("Game is starting...")
	printPlayers(players)
	fmt.Println("\nInfo:")
	fmt.Println("x cordinate - from left to right")
	fmt.Println("y cordinate - from top to bottom")
	fmt.Printf("playground size: %d x %d\n", playgroundSize, playgroundSize)
	fmt.Println("\nGood luck!")

	printPlayground(players)
	var nextPlayerNumber = 0
	for {
		player := players[nextPlayerNumber]
		askForMove(&player)
		printPlayground(players)
		if win := checkPlaygroundMatrixForWin(player); win {
			fmt.Printf("Congrats %s. You are winner!", player.Name)
			break
		}

		if draw := checkDraw(players); draw {
			fmt.Println("It's DRAW!")
			break
		}

		nextPlayerNumber++
		if nextPlayerNumber == len(players) {
			nextPlayerNumber = 0
		}
	}
}
