package main

import (
	"fmt"
)

type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

type animal struct {
	name string
	species string
	age int 
	sound string
}

func NewAnimal(name, species string, age int, sound string) Animal {
	return &animal{name: name, species: species, age: age, sound: sound}
}

func (a *animal) MakeSound() string {
	return a.sound
}

func (a *animal) GetName() string {
	return a.name
}

func (a *animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}


func ZooShow(animals []Animal) {
	for _, element := range animals {
		fmt.Println(element.GetInfo())
		fmt.Println(element.MakeSound())
	}
}

type ZooKeeper struct {}

func (z *ZooKeeper) Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!\n", animal.GetName(), animal.MakeSound())
}