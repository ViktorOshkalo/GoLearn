package main

import (
	"fmt"
	m "main/models"
)

func main() {
	fmt.Println("Yooo")

	player1 := m.Player{Name: "John"}
	player2 := m.Player{Name: "Mark"}

	room := m.NewRoom("FunnyChess")
	room.AddPlayer(player1)
	room.AddPlayer(player2)

	room.Move(player1, "E2E4")
	room.Move(player2, "A2A3")
	room.Move(player1, "D2D3")
	room.Move(player2, "H2H3")
}
