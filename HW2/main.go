package main

import "fmt"

func main() {
	//init
	zoo := Zoo{Name: "KyivZoo"}
	cage1 := Cage{Number: 1, Type: "Cage", Zoo: &zoo}
	cage2 := Cage{Number: 2, Type: "Aquarium", Zoo: &zoo}

	leo := Lion{
		Animal: Animal{&cage1, "Leo", "Gold"},
		Pride:  "TheLionsPride"}

	mikey := Mouse{
		Animal:     Animal{&cage1, "Mikey", "Grey"},
		TailLength: 5.56,
	}

	octo := Octopus{
		Animal:    Animal{&cage2, "Octo", "Blue"},
		LegsCount: 8,
	}

	shark := Shark{
		Animal:     Animal{&cage2, "BabyShark", "Blue"},
		TeethCount: 2048,
	}

	owl := Owl{
		Animal:       Animal{&cage1, "KnowOwl", "DarkGrey"},
		ViewDistance: 500,
	}

	owl2free := Owl{
		Animal:       Animal{Name: "FreeOwl", Color: "SuperDarkGrey"},
		ViewDistance: 900,
	}

	fmt.Println("Animals: ")
	fmt.Println(leo.GetInfo())
	fmt.Println(mikey.GetInfo())
	fmt.Println(octo.GetInfo())
	fmt.Println(shark.GetInfo())
	fmt.Println(owl.GetInfo())
	fmt.Println(owl2free.GetInfo())

	// make some animals free
	leo.FreeAnimal()
	shark.FreeAnimal()

	fmt.Println("\nAnimals after some free:")
	fmt.Println(leo.GetInfo())
	fmt.Println(mikey.GetInfo())
	fmt.Println(octo.GetInfo())
	fmt.Println(shark.GetInfo())
	fmt.Println(owl.GetInfo())
	fmt.Println(owl2free.GetInfo())

	cage3 := Cage{Number: 3, Type: "Cage", Zoo: &zoo}

	// lock animals
	leo.LockAnimal(&cage2)
	shark.LockAnimal(&cage2)
	owl2free.LockAnimal(&cage3)

	fmt.Println("\nAnimals after some locked back: ")
	fmt.Println(leo.GetInfo())
	fmt.Println(mikey.GetInfo())
	fmt.Println(octo.GetInfo())
	fmt.Println(shark.GetInfo())
	fmt.Println(owl.GetInfo())
	fmt.Println(owl2free.GetInfo())
}
