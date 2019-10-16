package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
)

var (
	//Default interval when the notifications are active
	ActiveFrom int
	ActiveTo   int

	//Default interval in hours
	Interval = 2
)

func main() {
	//Parse flags and set defaults in case there's no options provided when running the app
	flag.IntVar(&ActiveFrom, "ActiveFrom", 7, "Hour from when notifications become active")
	flag.IntVar(&ActiveTo, "ActiveTo", 17, "Hour when notifications become inactive")
	flag.Parse()

	if ActiveTo < ActiveFrom {
		fmt.Println("ActiveTo option cannot be bigger than ActiveFrom. Shutting down...")
		os.Exit(1)
	}

	for {
		//Delay by the interval
		time.Sleep(time.Hour * time.Duration(Interval))

		hour, _, _ := time.Now().Clock()

		//If current time is outside of active hours, then skip
		if hour < ActiveFrom || ActiveTo < hour {
			fmt.Println("Notifications are inactive...")
			continue
		}

		//Fire up the notification
		err := beeep.Notify("GoDrinker", "It's time to drink yo!", "drinklogo.png")
		fmt.Println("New notification! Next notification at:")
		if err != nil {
			panic(err)
		}
	}
}
