package main

import (
	"flag"
	"fmt"
	"github.com/gen2brain/beeep"
	"os"
	"time"
)

func main() {
	// Make use of command line arguments via flag package.
	deadline := flag.Int("minutes", 0, "Countdown time in minutes")
	// Parse flags after defining them.
	flag.Parse()
	// Copy deadline value in seconds to minimize delay in for loop.
	counter := *deadline * 60

	// Make deadline flag required.
	if *deadline == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Starting counter for %d Minutes . . .\n", *deadline)

	// Print remaining time every minute until deadline is reached.
	for range time.Tick(1 * time.Second) {
		if counter <= 0 {
			fmt.Println("Beep. Time's up!")
			err := beeep.Notify("timer.go", "Beep. Time's up!", "")
			if err != nil {
				panic(err)
			}
			break
		}
		counter--
		// Don't just modify deadline. Why? See:
		// https://reactjs.org/tutorial/tutorial.html#why-immutability-is-important
		if counter%60 == 0 {
			fmt.Printf("Minutes left: %d\n", counter/60)
		}
	}
}
