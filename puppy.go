package puppy

import "fmt"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! woof! woof! Woof!"
}

func BigBark() string {
	return WhenGrowsup(Bark())
}

func BigBarks() string {
	return WhenGrowsup(Barks())
}
func Fromv1() {
	fmt.Println("V!.0.0")
}
