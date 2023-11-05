package main

import "fmt"

type Player struct {
	Name      string
	HitPoints int
}

type Location struct {
	Name        string
	Description string
}

type Decision struct {
	Description     string
	Result          string
	HitPointsImpact int
	NextSceneId     int
}

type Scene struct {
	Id           int
	Name         string
	Location     Location
	Condition    string
	Decision1    Decision
	Decision2    Decision
	PlayerChoise int
	LastScene    bool
}

func (scene Scene) GetDecision() (decision Decision, success bool) {
	switch scene.PlayerChoise {
	case 1:
		return scene.Decision1, true
	case 2:
		return scene.Decision2, true
	default:
		return Decision{}, false
	}
}

func (scene Scene) GetNextSceneId() (nextSceneId int, success bool) {
	decision, success := scene.GetDecision()
	if success {
		return decision.NextSceneId, true
	} else {
		return 0, false
	}
}

func (scene *Scene) RunScene(player *Player) {
	fmt.Printf("\nPlayer HP: %d\n", player.HitPoints)

	fmt.Printf("\nScene: %s\n", scene.Name)
	fmt.Printf("Location: %s\n", scene.Location.Name)
	fmt.Printf("Conditions: %s\n", scene.Location.Description)

	fmt.Printf("\n%s\n", scene.Condition)

	if scene.LastScene {
		return
	}

	for {
		fmt.Printf("1. %s\n", scene.Decision1.Description)
		fmt.Printf("2. %s\n", scene.Decision2.Description)

		var choise int
		fmt.Printf("\nMake a choise, %s: ", player.Name)
		fmt.Scan(&choise)
		fmt.Println()

		scene.PlayerChoise = choise

		decision, success := scene.GetDecision()
		if !success {
			fmt.Println("Wrong decision. Please enter correct value.")
			continue
		}

		player.HitPoints = player.HitPoints + decision.HitPointsImpact

		fmt.Printf("Decision result: %s\n", decision.Result)
		fmt.Printf("HP affected: %d\n", decision.HitPointsImpact)
		fmt.Printf("HP current: %d\n", player.HitPoints)
		break
	}
}

func GetSceneById(sceneId int) (scene Scene, success bool) {
	switch sceneId {
	case 1:
		return Scene{
			Id:   sceneId,
			Name: "Reach the islands",
			Location: Location{
				Name:        "River side",
				Description: "Strong water running. Water temparture is low. Sniper is watching for you"},
			Condition: "You want to get to the other side of the river.\n" +
				"There are too islands in the middle of the river.\n" +
				"First is near to the final point. There are forest full of snakes.\n" +
				"Second is far from the final point. There are minefields on it.\n" +
				"A sniper sitting aside and will try to kill you when you are crossing a river. You found a boat and hydrasuit.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Use a boat",
				Result:          "Sniper shoot your leg!",
				HitPointsImpact: -25,
				NextSceneId:     11,
			},
			Decision2: Decision{
				Description:     "Swimm in hydrasuit",
				Result:          "Sniper haven't noticed you but you spent more energy during a swimm",
				HitPointsImpact: -10,
				NextSceneId:     12,
			},
		}, true
	case 11:
		return Scene{
			Id:   sceneId,
			Name: "Wild nature",
			Location: Location{
				Name:        "Island with forest",
				Description: "Forest full of snakes. Night is comming",
			},
			Condition: "You reached first island with forest, but you have been whounded by a sniper.\n" +
				"You have to cross a forest to reach other side of the island.\n" +
				"Forest is full of snakes. Night is coming and snakes are slipping at night.\n" +
				"On the other hand you are bleeding right now, probably you need take a rest and stop bleeding.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "No stop. Go night",
				Result:          "You loose a lot of blood",
				HitPointsImpact: -30,
				NextSceneId:     111,
			},
			Decision2: Decision{
				Description:     "Take a rest, stop bleeding. Go morning",
				Result:          "Snake hit you and you are poisoned. You messed up with direction",
				HitPointsImpact: -50,
				NextSceneId:     112,
			},
		}, true
	case 111:
		return Scene{
			Id:   sceneId,
			Name: "Aligator's trap",
			Location: Location{
				Name:        "Other side of first island",
				Description: "Morning. River full of aligators. Sun is up and going to be super hot. River is narrow here",
			},
			Condition: "You reached other side of the forest island, but loose a lot of blood.\n" +
				"There are aligators swimming in the river, but river is narrow here and you can swimm it fast.\n" +
				"Also you found a trees and can build a raft, but sun is rising and promise to be very hot.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Fastly swimm a river",
				Result:          "Aligator bit your leg!",
				HitPointsImpact: -60,
				NextSceneId:     2,
			},
			Decision2: Decision{
				Description:     "Build a raft",
				Result:          "Sun was realy hot and you got huge skin burns",
				HitPointsImpact: -30,
				NextSceneId:     2,
			},
		}, true
	case 112:
		return Scene{
			Id:   sceneId,
			Name: "Lost on island",
			Location: Location{
				Name:        "Unknown place",
				Description: "Morning. Rain and wind are coming. River is wide here",
			},
			Condition: "You reached a beach. You see a land far away but you don't now if it is correct direction.\n" +
				"You have a trees to build a raft.\n" +
				"Or you can stay on island and make sure you are on right direction.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Build a raft and cross a river",
				Result:          "You recover a bit after be poisoned",
				HitPointsImpact: +10,
				NextSceneId:     12,
			},
			Decision2: Decision{
				Description:     "Stay on island and find correct direction",
				Result:          "Snake bite you again!",
				HitPointsImpact: -50,
				NextSceneId:     111,
			},
		}, true
	case 12:
		return Scene{
			Id:   sceneId,
			Name: "Don't explode yourself",
			Location: Location{
				Name:        "Mine fields",
				Description: "Mine fields. Night is comming. Cold wind",
			},
			Condition: "You reached other island far away from finish point.\n" +
				"To get to the other side of island you have to go through a mine fields installed here during WW2.\n" +
				"Night is coming and you are freezing and there are no trees to make a fire.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "No stop. Go at night through a mine feilds before you get frozen",
				Result:          "You havn't noticed a mine at a night and got exploded. So sad :(",
				HitPointsImpact: -200,
				NextSceneId:     121,
			},
			Decision2: Decision{
				Description:     "Try to keep warm as much as possible. Go at morning",
				Result:          "Spent a lot of energy to keep warm",
				HitPointsImpact: -25,
				NextSceneId:     121,
			},
		}, true
	case 121:
		return Scene{
			Id:   sceneId,
			Name: "Long water run",
			Location: Location{
				Name:        "Far away beach",
				Description: "The perfect wheter",
			},
			Condition: "You reached the other side of the island.\n" +
				"The beach is located far away from a finish point.\n" +
				"You found several trees with fruits but you don't know if they eatable.\n" +
				"But the only way to get to finish point is swimming and you need some power for it.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Eat fruits and go",
				Result:          "You got additional power, but you spent some energy during a long swimm",
				HitPointsImpact: -25,
				NextSceneId:     2,
			},
			Decision2: Decision{
				Description:     "Do not eat fruits. Don't take that risk",
				Result:          "You spent more energy during a long swimm",
				HitPointsImpact: -50,
				NextSceneId:     2,
			},
		}, true
	case 2:
		return Scene{
			Id:        sceneId,
			Name:      "Final",
			LastScene: true,
			Location: Location{
				Name:        "Beer paradise",
				Description: "Perfect conditions, free cold bear on the beach",
			},
			Condition: "That was a long trip but you did it!.\n" +
				"Take a rest, heal you wounds, open a beer and be ready for a next home work.\n",
		}, true
	default:
		return Scene{}, false
	}
}

func main() {
	var playerName string
	var attempts int

	fmt.Print("Enter your name: ")
	fmt.Scan(&playerName)

	attempts = 3

	fmt.Printf("Hello %s. Game is starting.\n", playerName)
	fmt.Println("Goal: You are on the one side of the river and you want to reach other side in order to get to a some beautiful place called 'Beer paradise'")
	fmt.Printf("You will have %d attempts\n", attempts)

	for i := 1; i <= attempts; i++ {
		if res := askForContinue(); !res {
			return
		}

		fmt.Printf("Attempt %d\n", i)
		var player Player = Player{HitPoints: 100, Name: playerName}
		sceneId := 1
		for {
			scene, success := GetSceneById(sceneId)
			if !success {
				fmt.Printf("Something whent wrong: unable to get scene by id %d.\n", sceneId)
				break
			}

			scene.RunScene(&player)

			if player.HitPoints <= 0 {
				fmt.Println("GAME OVER. You are DEAD!")
				break
			}

			if scene.LastScene && player.HitPoints > 0 {
				fmt.Println("Congratulations! You survived!")
				return
			}

			if res := askForContinue(); !res {
				return
			}

			nextSceneId, success := scene.GetNextSceneId()
			if !success {
				fmt.Printf("Something whent wrong: unable to get next scene, scene id %d.\n", scene.Id)
				break
			}
			sceneId = nextSceneId
		}
	}
}

func askForContinue() bool {
	fmt.Println("\nContinue? (y/n)")
	var temp string
	fmt.Scan(&temp)
	if temp == "n" {
		fmt.Println("Exit game.")
		return false
	}
	return true
}
