package main

import "log"

func main() {
	var (
		builder = NewCharacterBuilder()
		hero1   = builder.SetName("Tanker").SetClass(ClassBarbarian).SetStrength(10).SetAgility(10).Build()
		hero2   = builder.SetName("Assassin").SetClass(ClassAssassin).SetStrength(10).SetAgility(10).Build()
	)
	log.Println(hero1.Attack(hero2.Name()))
	log.Println(hero2.Attack(hero1.Name()))
}
