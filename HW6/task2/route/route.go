package route

import (
	"fmt"
	"main/passenger"
	"main/route/interfaces"
)

type Destination struct {
	Name    string
	Vehicle interfaces.IPassengerVehicle
	IsFinal bool
}

type Route struct {
	From         string
	Destinations []Destination
}

func (r *Route) AddNextDestination(dest Destination) {
	r.Destinations = append(r.Destinations, dest)
}

func (route Route) RunRoute(passengers []passenger.Passenger) {
	fmt.Println("Route is starting. Enjoy your run!")

	fmt.Printf("Starting point: %s\n", route.From)
	from := route.From
	for _, destination := range route.Destinations {
		to := destination.Name
		fmt.Printf("\nNext destination : %s\n", to)
		destination.Vehicle.PrintInfo()
		batchStart := 0
		for {
			batchEnd := batchStart + destination.Vehicle.GetMaxCapacity()
			isLastBatch := batchEnd >= len(passengers)
			if isLastBatch {
				batchEnd = len(passengers)
			}

			destination.Vehicle.OnBoarding(passengers[batchStart:batchEnd])
			destination.Vehicle.Accelerate()
			destination.Vehicle.Move(to)
			destination.Vehicle.Stop()
			destination.Vehicle.OffBoarding()

			if !isLastBatch {
				// move vehicle back to pick up next passangers batch
				destination.Vehicle.Accelerate()
				destination.Vehicle.Move(from)
				destination.Vehicle.Stop()
				batchStart = batchEnd
				continue
			}

			break
		}

		if destination.IsFinal {
			fmt.Println("DONE! Final destination is reached.")
			break
		}

		from = to // remember prev destination name
	}

	fmt.Println("\nRoute finished!")
}
