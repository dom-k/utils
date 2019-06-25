package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// Make use of command line arguments via flag package.
	deadline := flag.Int("minutes", 0, "Countdown time in minutes")
	// Parse flags after defining them.
	flag.Parse()
	// Copy deadline value.
	counter := *deadline

	// Make deadline flag required.
	if *deadline == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Starting counter for %02d Minutes . . .\n", counter)

	// Print remaining time every minute until deadline is reached.
	for range time.Tick(1 * time.Minute) {
		if counter <= 0 {
			fmt.Println("Beep. Time's up!")
			break
		}
		counter--
		fmt.Printf("Minutes left: %02d\n", counter)
	}
}
