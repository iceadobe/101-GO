package mypackage

import "fmt"

type AnimalType int

const (
	MAMAL AnimalType = iota
	AMPHIBIAN
	REPTILE
)

func (tp AnimalType) ToString() string {
	names := []string{"MAMAL", "AMPHIBIAN", "REPTILE"}
	return names[tp-1]
}

type Animal interface {
	Move() string
}

type Dog struct {
	Animal_type AnimalType
	Name        string
}

func (d Dog) Move() string {
	return "WALK"
}

type Frog struct {
	Animal_type AnimalType
	Name        string
}

func (f Frog) Move() string {
	return "JUMP"
}

type Snake struct {
	Animal_type AnimalType
	Name        string
}

func (s Snake) Move() string {
	return "CRAWL"
}

func TakeMyPetForAWalk(pet Animal) {
	fmt.Println("What kind of animal are you ?")
	switch pet.(type) {
	case Dog:
		fmt.Println("A Doggy named", pet.(Dog).Name)
	case Snake:
		fmt.Println("A Snake named", pet.(Snake).Name)
		// will compile but give error at runtime
		//fmt.Println("A Snake named", pet.(Dog).name)
	case Frog:
		fmt.Println("A Frog named", pet.(Frog).Name)
	}
	fmt.Println("and I can", pet.Move())
}
