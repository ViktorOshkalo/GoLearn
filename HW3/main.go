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
	Decision1    Decision
	Decision2    Decision
	Condition    string
	PlayerChoise int
	LastScene    bool
	Player       *Player
}

func CreateScene(sceneId int, player *Player) (scene Scene, success bool) {
	switch sceneId {
	case 1:
		return Scene{
			Id:     sceneId,
			Player: player,
			Name:   "Reach the islands",
			Location: Location{
				Name:        "River side",
				Description: "Strong water running. Water temparture is low. Sniper is watching for you."},
			Condition: "You have to reach other side of the river.\n" +
				"There are different islands in the middle of the river on your way.\n" +
				"A sniper sitting aside and will try to kill you when you crossing a river. You found a boat and hydrosuit inside.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Use a boat.",
				Result:          "You reached the nearest island but sniper shoot your leg.",
				HitPointsImpact: -20,
				NextSceneId:     11,
			},
			Decision2: Decision{
				Description:     "Swimm in hydrosuit.",
				Result:          "Water was runninng fast and you got to the other island far away from you destination. Sniper haven't noticed you. But you got hypothermia.",
				HitPointsImpact: -10,
				NextSceneId:     12,
			},
		}, true
	case 11:
		return Scene{
			Id:     sceneId,
			Player: player,
			Name:   "Wild nature",
			Location: Location{
				Name:        "Island with forest",
				Description: "Forest full of snakes. Night is comming.",
			},
			Condition: "You reached island with forest on it, but you have been whounded by a sniper.\n" +
				"You have to cross forest to reach other side of the island.\n" +
				"Forest is full of snakes. They will be slipping at night.\n" +
				"On the other hand you are bleeding right now, you need to stop bood loosing.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "No stop. Go at night.",
				Result:          "You reached other side of the island but you loose a lot of blood.",
				HitPointsImpact: -30,
				NextSceneId:     111,
			},
			Decision2: Decision{
				Description:     "Take a rest, stop bleeding. Go at morning.",
				Result:          "You was hit by snake and you are poisoned now. Under poison you messed up with direction and got to a sligtly different point fo the island.",
				HitPointsImpact: -50,
				NextSceneId:     112,
			},
		}, true
	case 111:
		return Scene{
			Id:     sceneId,
			Player: player,
			Name:   "Aligator fight",
			Location: Location{
				Name:        "Other side of island",
				Description: "Morning. River full of aligators. Sun is up and going to fire you. River is narrow here and you can fastly reach other side.",
			},
			Condition: "You reached other side of the forest island, loose a loot of blood and still bleeding.\n" +
				"There are aligators swimming in the river, but river is narrow here.\n" +
				"On the other hand you found a trees and can make a raft.\n" +
				"Sun is up and starts activly fire you.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Sun is super hot. Try fastly swimm a river full of aligators.",
				Result:          "You reached other side of the river but aligator felt your blood smell and bit you leg!.",
				HitPointsImpact: -60,
			},
			Decision2: Decision{
				Description:     "Make a raft under hot sun and cross the river.",
				Result:          "You made a raft and crossed the river, but sun was realy hot and you got huge skin burns.",
				HitPointsImpact: -30,
			},
		}, true
	case 112:
		return Scene{
			Id:     sceneId,
			Player: player,
			Name:   "Lost on island",
			Location: Location{
				Name:        "Unknown place",
				Description: "Morning. Rain and wind are coming. River is wide here.",
			},
			Condition: "You reached a beach. You see a land far away but you don't now if it is correct direction.\n" +
				"You have a trees to build a raft.\n" +
				"Or you can stay on island and try to find out right direction.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description: "Build a raft and cross a river.",
				Result: "It was took time to build a raft and you started feel much better after snake bite. You reached other side of the river fast as wind was in the same direction.\n" +
					"But direction was wrong! You reached other island!",
				HitPointsImpact: +10,
				NextSceneId:     12,
			},
			Decision2: Decision{
				Description:     "Stay on island and find correct direction.",
				Result:          "You made a raft and crossed the river, but sun was realy hot and you got huge skin burns.",
				HitPointsImpact: -30,
			},
		}, true
	case 12:
		return Scene{
			Id:     sceneId,
			Player: player,
			Name:   "Don't explode yourself",
			Location: Location{
				Name:        "Mine fields",
				Description: "Mine fields. Night is comming. Cold wind.",
			},
			Condition: "You reached other island far away from finish point.\n" +
				"To get to the other side of island you have to go through a mine fields installed here during WW2.\n" +
				"Night is coming and you are freezing and there are no trees to make a fire.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "No stop. Go at night through a mine feilds before you get frozen.",
				Result:          "You havn't noticed a mine at a night and got exploded. So sad :(",
				HitPointsImpact: -100,
			},
			Decision2: Decision{
				Description:     "Try to still warm as much as possible. Go at morning.",
				Result:          "You was moving actively at a night in order to keep warm, then you successfully crossed mine fields and reached the other side of the island. But you are very exaust now.",
				HitPointsImpact: -25,
				NextSceneId:     122,
			},
		}, true

	default:
		return Scene{}, false
	}
}

func (scene *Scene) RunScene() {
	fmt.Printf("\nScene: %s.\n", scene.Name)
	fmt.Println(scene.Condition)

	for {
		fmt.Printf("1. %s\n", scene.Decision1.Description)
		fmt.Printf("2. %s\n", scene.Decision2.Description)

		var choise int
		fmt.Printf("Make a choise, %s: ", scene.Player.Name)
		fmt.Scan(&choise)
		fmt.Println()

		scene.PlayerChoise = choise

		decision, success := scene.GetDecision()
		if !success {
			fmt.Println("Wrong decision. Please enter correct value.")
			continue
		}

		fmt.Printf("Result: %s\n", decision.Result)

		scene.Player.HitPoints = scene.Player.HitPoints + decision.HitPointsImpact

		fmt.Printf("HP affected: %d\n", decision.HitPointsImpact)
		fmt.Printf("HP current: %d\n", scene.Player.HitPoints)
		break
	}
}

func (scene Scene) GetDecision() (decision Decision, success bool) {

	fmt.Println(scene.PlayerChoise)

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
		return -1, false
	}
}

func main() {
	var player Player = Player{Name: "Viktor", HitPoints: 100}
	//fmt.Print("Enter your name: ")
	//fmt.Scan(&player.Name)
	fmt.Printf("Hello %s. Game is starting.\n", player.Name)
	fmt.Printf("Player HP: %d\n", player.HitPoints)

	sceneId := 1
	for {
		scene, success := CreateScene(sceneId, &player)
		if !success {
			fmt.Printf("Something whent wrong: unable to get scene by id %d.\n", sceneId)
			break
		}

		scene.RunScene()

		if scene.Player.HitPoints <= 0 {
			fmt.Println("GAME OVER. You are ded!")
			break
		}

		if scene.LastScene && scene.Player.HitPoints > 0 {
			fmt.Println("Congratulations! You survived!")
			break
		}

		fmt.Println("\nContinue? (y/n)")
		var temp string
		fmt.Scan(&temp)
		if temp == "n" {
			fmt.Println("Exit game.")
			break
		}

		nextSceneId, success := scene.GetNextSceneId()
		if !success {
			fmt.Printf("Something whent wrong: unable to get next scene, scene id %d.\n", scene.Id)
			break
		}
		sceneId = nextSceneId
	}
}
