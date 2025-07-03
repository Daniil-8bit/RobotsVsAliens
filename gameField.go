package main

import "fmt"

type GameField struct {
	width  int
	length int
}

var robotsPositions []Robot
var aliansPositions []Alien

func showRobotsPositions(r []Robot) {

	for k, v := range r {
		fmt.Printf("Robot%d: %s\nHealth: %d\nPosition on the field: (%d;%d)\n\n", k+1, v.name, v.health, v.x, v.y)
	}
}

func showAliensPositions(a []Alien) {

	for k, v := range a {
		fmt.Printf("Alien%d: %s\nHealth: %d\nPosition on the field: (%d;%d)\n\n", k+1, v.name, v.health, v.x, v.y)
	}
}
