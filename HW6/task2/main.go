package main

import (
	"fmt"
	"main/route"
	"main/vehicles"
)

func main() {
	fmt.Println("GO!")

	var route = route.Route{}
	route.AddVehicle(vehicles.Car{})
	route.AddVehicle(vehicles.Train{})
	route.AddVehicle(vehicles.Airplane{})
	route.AddVehicle(vehicles.Train{})
	route.AddVehicle(vehicles.Car{})

	route.PrintVehicles()

	passengerName := "Mike"
	route.RunRoute(passengerName)
}
