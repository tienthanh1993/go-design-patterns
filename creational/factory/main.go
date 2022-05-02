package main

import "log"

func main() {
	dog := AnimalFactory("dog")
	dolphin := AnimalFactory("dolphin")
	log.Printf("This is %v, it has %v legs and it bark sound '%v'", dog.Name(), dog.GetLegs(), dog.Speak())
	log.Printf("This is %v, it has %v leg and it speak sound '%v'", dolphin.Name(), dolphin.GetLegs(), dolphin.Speak())
}
