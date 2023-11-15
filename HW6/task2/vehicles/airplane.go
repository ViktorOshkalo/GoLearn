package vehicles

import (
	"fmt"
	"main/route/interfaces"
)

type Airplane struct {
	PassengerTransferable
}

func (airplane Airplane) Move(to string) {
	fmt.Printf("Airplane is flying to: %s\n", to)
}

func (airplane Airplane) Stop() {
	fmt.Println("Airplane landed")
}

func (airplane Airplane) Accelerate() {
	fmt.Println("Airplane is taking off!!!")
}

func (airplane Airplane) PrintInfo() {
	fmt.Println("Vehicle type: Airplane")
	fmt.Printf("Max passengers: %d\n", airplane.MaxCapacity)
}

var _ interfaces.IPassengerVehicle = Airplane{}
