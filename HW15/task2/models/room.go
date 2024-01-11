package models

import (
	"fmt"
	"main/observer"
)

type Room struct {
	Name string
	observer.ObserverManager
}

func NewRoom(name string) Room {
	return Room{Name: name, ObserverManager: observer.ObserverManager{}}
}

func (r *Room) AddPlayer(p Player) {
	r.ObserverManager.Register(p)
}

func (r *Room) RemovePlayer(p Player) {
	r.ObserverManager.Unregister(p)
}

func (r *Room) Move(p Player, move string) {
	r.ObserverManager.NotifyOthers(p, fmt.Sprintf("Room %s. Player %s. Move: %s\n", r.Name, p.Name, move))
}
