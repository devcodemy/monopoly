package main

import (
	"fmt"
	"mono/drawfield"
)
import "mono/linked"

func do() {
	drawfield.Draw()
}

func main() {
	f1 := linked.Field{
		Cell: " --- GO GO GO ---",
		Rent: 0,
	}
	f2 := linked.Field{
		Cell: "PHILIPS",
		Rent: -100,
	}
	f3 := linked.Field{
		Cell: "SIEMENS",
		Rent: -200,
	}
	f4 := linked.Field{
		Cell: "BOSH",
		Rent: -290,
	}

	bg := linked.Node{}
	bg.Insert(&f1)
	last, _ := bg.Insert(&f2)
	bg.Insert(&f3)
	bg.Insert(&f4)
	//last := bg.Next.Next.Next
	last.Next = &bg
	fmt.Println("---")
	bg.Traverse()

	//st := linked.Node{Value: &f2, Previous: nil, Next: nil}
	//st.Print()
	//st = *st.Insert(&f1)
	//st.Traverse()
	//st.Previous.Print()
	//st.Print()
	//st.Next.Print()
	//st.Next.Print()
	//st.Next.Next.Print()
	//st.Next.Next.Next.Print()
	//f1.Action()
	//f2.Action()
}
