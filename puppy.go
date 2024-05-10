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
func Fromv3() {
	fmt.Println("V3.0.0")
}

// to Add the diff tags with different code
// 1) first push the code
// 2) then add the tag using "git tag v1.0.0"
// 3) then push the tags "git push origin --tags"
