package vehicles

import (
	"fmt"
	interfaces "main/route/interfaces"
)

type Train struct {
}

func (car Train) Move() {
	fmt.Println("Train is going")
}

func (car Train) Stop() {
	fmt.Println("Train stopped")
}

func (car Train) Accelerate() {
	fmt.Println("Train is speeding up!!!")
}

func (car Train) OnBoarding() {
	fmt.Println("Pessangeers are boarding into Train")
}

func (car Train) OffBoarding() {
	fmt.Println("Pessangeers are offboarding from Train")
}

func (car Train) PrintInfo() {
	fmt.Println("Vehicle type: Train")
}

var _ interfaces.PassengerVehicle = Train{}
