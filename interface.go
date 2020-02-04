package main

import (
	MyTypes "./mypackage"
	"fmt"
)

func main() {
	// Creating three Animals
	d := MyTypes.Dog{MyTypes.MAMAL, "Tommy"}
	s := MyTypes.Snake{MyTypes.REPTILE, "Shanky"}
	f := MyTypes.Frog{MyTypes.AMPHIBIAN, "Gavin"}

	MyTypes.TakeMyPetForAWalk(d)
	fmt.Println()
	MyTypes.TakeMyPetForAWalk(s)
	fmt.Println()
	MyTypes.TakeMyPetForAWalk(f)
}
