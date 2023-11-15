package route

import (
	"fmt"
	interfaces "main/route/interfaces"
)

type Route struct {
	Vehicles []interfaces.PassengerVehicle
}

func (route *Route) AddVehicle(vehicle interfaces.PassengerVehicle) {
	route.Vehicles = append(route.Vehicles, vehicle)
}

func (route Route) PrintVehicles() {
	fmt.Println("\nVehicles on the route: ")
	for _, v := range route.Vehicles {
		v.PrintInfo()
	}
}

func (route Route) RunRoute() {
	fmt.Println("\nRoute is running...")
	for _, v := range route.Vehicles {
		fmt.Println("\nNext vehicle is running: ")
		v.PrintInfo()
		v.OnBoarding()
		v.Accelerate()
		v.Move()
		v.Stop()
		v.OffBoarding()
	}
	fmt.Println("\nRoute finished!")
}
