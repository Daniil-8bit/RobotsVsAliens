package main

import "fmt"

type Warriorable interface {
	showInfo()
	makeSound()
	causeDamage(*Warriorable)
	move(int, int) (int, int)
}

type Location struct {
	x int
	y int
}

type Robot struct {
	name    string
	health  int
	damage  int
	sound   string
	defense int
	Location
}

func (r *Robot) showInfo() {

	fmt.Printf("Robot: %s\nHealth: %d\n", r.name, r.health)
}

func (r *Robot) makeSound() {

	fmt.Println("bz-bz-bz")
}

func (r *Robot) causeDamage(a *Alien) {

	a.health = a.health - r.damage
}

func (r *Robot) move(x int, y int) {
	r.x = x
	r.y = y
}

type Alien struct {
	name    string
	health  int
	damage  int
	sound   string
	defense int
	Location
}
