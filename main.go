package main

func main() {

	robotsAmount := make([]Robot, 0)
	aliensAmount := make([]Alien, 0)

	robotsAmount = []Robot{
		Robot{name: "Small Robot", health: 30, Location: Location{x: 7, y: 1}},
		Robot{name: "Middle Robot", health: 50, Location: Location{x: 5, y: 3}},
		Robot{name: "Big Robot", health: 90, Location: Location{x: 6, y: 2}},
	}

	aliensAmount = []Alien{
		Alien{name: "Small Aien", health: 45, Location: Location{x: 1, y: 5}},
		Alien{name: "Middle Aien", health: 20, Location: Location{x: 3, y: 1}},
		Alien{name: "Big Aien", health: 75, Location: Location{x: 2, y: 3}},
	}

	showRobotsPositions(robotsAmount)
	showAliensPositions(aliensAmount)
}
