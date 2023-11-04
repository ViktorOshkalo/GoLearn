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
	Player       Player
	FinishScene  bool
}

func GetScene(sceneId int) (scene Scene, success bool) {

	fmt.Printf("Next scene: %d.\n", sceneId)

	switch sceneId {
	case 1:
		return Scene{
			Id:   sceneId,
			Name: "Reach the islands",
			Location: Location{
				Name:        "River side",
				Description: "Strong water running. Water temparture is low. Sniper is watching for you."},
			Condition: "You are trying to reach other side of the river.\n" +
				"There are different islands in the middle of the river on your way.\n" +
				"A sniper sitting aside and trying to kill you. You found a boat and hydrosuit inside.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "Use a boat.",
				Result:          "You reached the nearest island but sniper shoot your leg.",
				HitPointsImpact: -30,
				NextSceneId:     2,
			},
			Decision2: Decision{
				Description:     "Swimm in hydrosuit.",
				Result:          "Water was runninng fast and you got to the other island far away from you destination. Sniper haven't noticed you. But you got hypothermia.",
				HitPointsImpact: -15,
				NextSceneId:     3,
			},
		}, true
	case 2:
		return Scene{
			Id:   sceneId,
			Name: "Wild nature",
			Location: Location{
				Name:        "Island 1",
				Description: "Forest full of snakes. Night is comming.",
			},
			Condition: "You reached nearest island but you have whounded by sniper.\n" +
				"You need to cross forest to reach other side of the island.\n" +
				"Better to go at night as a snakes are slipping, but on the other side -  you are bleeding right now.\n" +
				"Other option is to stay on the beach, get rest and stop bleading. Then go at the morning.\n" +
				"What decision will u make?\n",
			Decision1: Decision{
				Description:     "No stop. Go at night.",
				Result:          "You reached other side of the island but you loosed a lot of blood.",
				HitPointsImpact: -50,
				NextSceneId:     4,
			},
			Decision2: Decision{
				Description:     "Take a rest. Go at morning.",
				Result:          "You reached other side of the island but snake bit you. You are poisoned.",
				HitPointsImpact: -20,
				NextSceneId:     5,
			},
		}, true
	default:
		return Scene{}, false
	}
}

func (scene *Scene) RunScene() {
	fmt.Printf("Scene: %s.\n", scene.Name)
	fmt.Println(scene.Condition)

	if scene.FinishScene {
		return
	}

	fmt.Printf("1. %s\n", scene.Decision1.Description)
	fmt.Printf("2. %s\n", scene.Decision2.Description)

	for {
		var choise int
		fmt.Printf("Make a choise, %s: ", scene.Player.Name)
		fmt.Scan(&choise)
		fmt.Println()

		var decision Decision
		if choise == 1 {
			decision = scene.Decision1
		} else if choise == 2 {
			decision = scene.Decision2
		} else {
			fmt.Println("Bad number. Enter correct value.")
			continue
		}

		fmt.Printf("Decision result: %s\n", decision.Result)
		scene.PlayerChoise = choise
		scene.Player.HitPoints = scene.Player.HitPoints + decision.HitPointsImpact
		break
	}
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

func main() {
	var player Player = Player{Name: "Viktor"}
	//fmt.Print("Enter your name: ")
	//fmt.Scan(&player.Name)
	fmt.Printf("Hello %s. Game is starting.\n", player.Name)

	sceneId := 1
	for {
		scene, success := GetScene(sceneId)
		if !success {
			fmt.Printf("Something whent wrong: unable to get scene by id %d.\n", sceneId)
			break
		}

		scene.RunScene()
		if scene.FinishScene {
			fmt.Println("GAME END.")
			break
		}

		decision, success := scene.GetDecision()
		if !success {
			fmt.Printf("Something whent wrong: unable to get next scene, scene id %d.\n", scene.Id)
			break
		}
		sceneId = decision.NextSceneId
	}
}
