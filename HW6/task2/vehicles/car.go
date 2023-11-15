package vehicles

import (
	"fmt"
	interfaces "main/route/interfaces"
)

type Car struct {
	PassengerTransferable
}

func (car Car) Move(to string) {
	fmt.Printf("Car is going to: %s\n", to)
}

func (car Car) Stop() {
	fmt.Println("Car stopped")
}

func (car Car) Accelerate() {
	fmt.Println("Car is speeding up!!!")
}

func (car Car) PrintInfo() {
	fmt.Println("Vehicle type: Car")
	fmt.Printf("Max passengers: %d\n", car.MaxCapacity)
}

var _ interfaces.IPassengerVehicle = Car{}
