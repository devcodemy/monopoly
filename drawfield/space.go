package drawfield

import "fmt"

func Draw() {
	fmt.Print("|")
	for i := 0; i < 6; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("^")
		}
		fmt.Print("\n|")
	}
}
