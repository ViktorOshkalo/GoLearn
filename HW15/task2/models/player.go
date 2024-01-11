package models

import (
	"fmt"
	"main/observer"
)

type Player struct {
	Name string
}

func (pl Player) OnMessage(message string) {
	fmt.Printf("Player %s recived message: \n%s\n", pl.Name, message)
}

var _ observer.Observer = Player{}
