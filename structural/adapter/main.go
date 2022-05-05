package main

func main() {
	var fighter Tanker
	//fighter = &Wizard{} // invalid because Wizard is not compatible with Tanker
	fighter = &WizardAdapter{
		&Wizard{},
	}
	fighter.Attack()
	fighter.Defense()
}
