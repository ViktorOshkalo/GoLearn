package vehicles

import (
	"fmt"
	interfaces "main/route/interfaces"
)

type Car struct {
}

func (car Car) Move() {
	fmt.Println("Car is riding")
}

func (car Car) Stop() {
	fmt.Println("Car stopped")
}

func (car Car) Accelerate() {
	fmt.Println("Car is speeding up!!!")
}

func (car Car) OnBoarding() {
	fmt.Println("Pessangeers are boarding into Car")
}

func (car Car) OffBoarding() {
	fmt.Println("Pessangeers are offboarding from Car")
}

func (car Car) PrintInfo() {
	fmt.Println("Vehicle type: Car")
}

var _ interfaces.PassengerVehicle = Car{}
