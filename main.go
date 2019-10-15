package main

import (
	"github.com/gen2brain/beeep"
)

func main() {
	err := beeep.Notify("Title", "Message body", "drinklogo.png")
	if err != nil {
		panic(err)
	}
}
