package main

import "fmt"

type Zoo struct {
	Name string
}

type Cage struct {
	Zoo    *Zoo
	Number int
	Type   string
}

type Animal struct {
	Cage  *Cage
	Name  string
	Color string
}

func (animal *Animal) FreeAnimal() {
	animal.Cage = nil // maybe bad idea :)
}

func (animal *Animal) LockAnimal(cage *Cage) {
	animal.Cage = cage
}

func (animal Animal) GetInfo() string {

	animalInfo := fmt.Sprintf("Animal name: %s, color: %s", animal.Name, animal.Color)

	var cageInfo string
	if animal.Cage != nil {
		cageInfo = fmt.Sprintf("Cage number: %d, type: %s", animal.Cage.Number, animal.Cage.Type)
	} else {
		cageInfo = "Cage is unknown, animal is FREE!"
	}

	return fmt.Sprintf("%s. %s", animalInfo, cageInfo)
}

// animals
type Lion struct {
	Animal
	Pride string
}

type Mouse struct {
	Animal
	TailLength float32
}

type Octopus struct {
	Animal
	LegsCount int
}

type Shark struct {
	Animal
	TeethCount int
}

type Owl struct {
	Animal
	ViewDistance float32
}
