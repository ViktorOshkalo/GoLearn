package interfaces

type Vehicle interface {
	Move()
	Stop()
	Accelerate()
	PrintInfo()
}

type PassengerVehicle interface {
	Vehicle
	OnBoarding()
	OffBoarding()
}
