package main

import "log"

type Tanker interface {
	Attack()
	Defense()
}

type WizardAdapter struct {
	*Wizard
}

func (w *WizardAdapter) Attack() {
	w.Wizard.Attack()
}

func (w *WizardAdapter) Defense() {
	log.Println("WizardAdapter:Defense")
}

type Wizard struct {
	name string
}

func (w *Wizard) Attack() {
	log.Println("Wizard:Attack")
}

func (w *Wizard) Escape() {
	log.Println("Wizard:Escape")
}
