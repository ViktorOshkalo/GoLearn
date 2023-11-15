package interfaces

import "main/passenger"

type IVehicle interface {
	Move(destinationName string)
	Stop()
	Accelerate()
}

type IPassengerTransferable interface {
	GetMaxCapacity() int
	OnBoarding(pessanger []passenger.Passenger) error
	OffBoarding()
}

type IPassengerVehicle interface {
	IVehicle
	IPassengerTransferable
	PrintInfo()
}
