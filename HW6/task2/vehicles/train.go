package vehicles

import (
	"fmt"
	interfaces "main/route/interfaces"
)

type Train struct {
	PassengerTransferable
}

func (train Train) Move(to string) {
	fmt.Printf("Train is going to: %s\n", to)
}

func (train Train) Stop() {
	fmt.Println("Train stopped")
}

func (train Train) Accelerate() {
	fmt.Println("Train is speeding up!!!")
}

func (train Train) PrintInfo() {
	fmt.Println("Vehicle type: Train")
	fmt.Printf("Max passengers: %d\n", train.MaxCapacity)
}

var _ interfaces.IPassengerVehicle = Train{}
