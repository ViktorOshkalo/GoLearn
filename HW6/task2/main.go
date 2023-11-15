package main

import (
	"fmt"
	"main/passenger"
	"main/route"
	"main/vehicles"
)

func main() {
	fmt.Println("GO!")

	var r = route.Route{From: "Kyiv"}
	r.AddNextDestination(route.Destination{Name: "Warsaw", Vehicle: vehicles.Car{PassengerTransferable: vehicles.PassengerTransferable{MaxCapacity: 4}}})
	r.AddNextDestination(route.Destination{Name: "Berlin", Vehicle: vehicles.Car{PassengerTransferable: vehicles.PassengerTransferable{MaxCapacity: 5}}})
	r.AddNextDestination(route.Destination{Name: "London", Vehicle: vehicles.Airplane{PassengerTransferable: vehicles.PassengerTransferable{MaxCapacity: 180}}})
	r.AddNextDestination(route.Destination{Name: "New York", Vehicle: vehicles.Car{PassengerTransferable: vehicles.PassengerTransferable{MaxCapacity: 320}}})
	r.AddNextDestination(route.Destination{Name: "Los Angeles", IsFinal: true, Vehicle: vehicles.Train{PassengerTransferable: vehicles.PassengerTransferable{MaxCapacity: 1200}}})

	passengers := []passenger.Passenger{
		{Name: "Nick"},
		{Name: "Mike"},
		{Name: "Alex"},
		{Name: "John"},
		{Name: "Rick"},
		{Name: "Duke"},
		{Name: "Grok"},
		{Name: "Bros"},
		{Name: "Brol"},
		{Name: "Bill"},
	}

	r.RunRoute(passengers)
}
