package main

import "fmt"

//Animaler any animal can implement this interface
type Animaler interface {
	SetName(v string)
	Name() string
	GetLegs() int
	SetLegs(v int)
	Speak() string
}

type Dog struct {
	name  string
	legs  int
	speak string
}

func (d *Dog) SetName(v string) {
	d.name = v
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) GetLegs() int {
	return d.legs
}

func (d *Dog) SetLegs(v int) {
	d.legs = v
	return
}

func (d Dog) Speak() string {
	return d.speak
}

type Fish struct {
	name  string
	legs  int
	speak string
}

func (d *Fish) SetName(v string) {
	d.name = v
}

func (d Fish) Name() string {
	return d.name
}

func (d Fish) GetLegs() int {
	return d.legs
}

func (d *Fish) SetLegs(v int) {
	d.legs = v
	return
}

func (d Fish) Speak() string {
	return d.speak
}

func NewDog(name string, legs int, speak string) Animaler {
	return &Dog{name: name, legs: legs, speak: speak}
}

func NewFish(name string, legs int, speak string) Animaler {
	return &Dog{name: name, legs: legs, speak: speak}
}

//AnimalFactory make a new animal in species
func AnimalFactory(species string) Animaler {
	switch species {
	case "dog":
		return NewDog("Super dog", 4, "bow wow")
	case "dolphin":
		return NewFish("Dolphin", 0, "bleats")
	}
	panic(fmt.Sprintf("not support species %v", species))
	return nil
}
