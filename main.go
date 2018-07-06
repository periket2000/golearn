package main

import "fmt"

var deckSize int

func main() {
	deckSize = 20
	fmt.Println("Hey mate!")
	fmt.Println(deckSize)

	f := file{
		filetName: "shit",
		fileType:  "pdf",
	}

	// Using it in a canonical way (see file.go)
	pf := &f
	fmt.Println(f)
	pf.updateName("Marco")
	fmt.Println(f)

	//Using it with Go infering the pointer for us (see file.go)
	f.updateName("whatever")
	fmt.Println(f)
}
