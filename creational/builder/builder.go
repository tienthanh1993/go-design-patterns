package main

import (
	"fmt"
	"strings"
)

const (
	ClassBarbarian = "barbarian"
	ClassAssassin  = "assassin"
)

var (
	BarbarianDamageFactor = [2]uint32{5, 2}
	AssassinDamageFactor  = [2]uint32{4, 1}
)

//Character what a character can do
type Character interface {
	Name() string
	Attack(target string) string
}

type hero struct {
	name     string
	class    string
	strength uint32
	agility  uint32
}

//Name return hero name
func (b *hero) Name() string {
	return b.name
}

//Attack returns damage to target base on hero class
func (b hero) Attack(target string) string {
	var ret strings.Builder
	switch b.class {
	case ClassBarbarian:
		ret.WriteString(fmt.Sprintf("%v attack %v with %v damage", b.name, target, b.strength*BarbarianDamageFactor[0]+b.agility*BarbarianDamageFactor[1]))
	case ClassAssassin:
		ret.WriteString(fmt.Sprintf("%v attack %v with %v damage", b.name, target, b.agility*AssassinDamageFactor[0]+b.strength*AssassinDamageFactor[1]))
	}
	return ret.String()
}

//CharacterBuilder build complex character with simple method
type CharacterBuilder interface {
	SetName(v string) CharacterBuilder
	SetClass(v string) CharacterBuilder
	SetStrength(v uint32) CharacterBuilder
	SetAgility(v uint32) CharacterBuilder
	Build() Character
}

type characterBuilder struct {
	name     string
	class    string
	strength uint32
	agility  uint32
}

func NewCharacterBuilder() *characterBuilder {
	return &characterBuilder{}
}

func (c *characterBuilder) SetStrength(v uint32) CharacterBuilder {
	c.strength = v
	return c
}

func (c *characterBuilder) SetAgility(v uint32) CharacterBuilder {
	c.agility = v
	return c
}

func (c *characterBuilder) SetName(v string) CharacterBuilder {
	c.name = v
	return c
}

func (c *characterBuilder) SetClass(v string) CharacterBuilder {
	c.class = v
	return c
}

func (c *characterBuilder) Build() Character {
	return &hero{
		name:     c.name,
		class:    c.class,
		strength: c.strength,
		agility:  c.agility,
	}
}
