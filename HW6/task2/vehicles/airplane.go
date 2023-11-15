package vehicles

import (
	"fmt"
	interfaces "main/route/interfaces"
)

type Airplane struct {
}

func (car Airplane) Move() {
	fmt.Println("Airplane is flying")
}

func (car Airplane) Stop() {
	fmt.Println("Airplane landed")
}

func (car Airplane) Accelerate() {
	fmt.Println("Airplane is taking off!!!")
}

func (car Airplane) OnBoarding() {
	fmt.Println("Pessangeers are boarding into Airplane")
}

func (car Airplane) OffBoarding() {
	fmt.Println("Pessangeers are offboarding from Airplane")
}

func (car Airplane) PrintInfo() {
	fmt.Println("Vehicle type: Airplane")
}

var _ interfaces.PassengerVehicle = Airplane{}
