package vehicles

import (
	"fmt"
	"main/passenger"
	"main/route/interfaces"
)

type PassengerTransferable struct {
	MaxCapacity int
	Passengers  []passenger.Passenger
}

func (pt PassengerTransferable) GetMaxCapacity() int {
	return pt.MaxCapacity
}

func (pt PassengerTransferable) OnBoarding(passengers []passenger.Passenger) error {
	if len(passengers) > pt.MaxCapacity {
		panic(fmt.Sprintf("there is not enough capacity, max capacity: %d, passangers count: %d, over: %d", pt.MaxCapacity, len(passengers), len(passengers)-pt.MaxCapacity))
	}
	pt.Passengers = passengers
	fmt.Printf("Passengers are boarded, count: %d\n", len(pt.Passengers))
	var names []string
	for _, p := range pt.Passengers {
		names = append(names, p.Name)
	}
	fmt.Println("Passenger's names: ", names)
	return nil
}

func (pt PassengerTransferable) OffBoarding() {
	fmt.Println("Passengers are offboarding.")
	pt.Passengers = nil
}

var _ interfaces.IPassengerTransferable = PassengerTransferable{}
