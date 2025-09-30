package main

import (
	"fmt"
	"os"

	"github.com/bitfield/booty"
)

func main() {
	crew, err := booty.AskInt(os.Stdin, os.Stdout, "How many crew?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// AskInt enforces >= 1, so this check is redundant, but keep it as an extra guard.
	if crew <= 0 {
		fmt.Fprintln(os.Stderr, "Number of crew must be greater than zero.")
		os.Exit(1)
	}

	total, err := booty.AskInt(os.Stdin, os.Stdout, "How many pieces of eight?")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// AskInt enforces >= 1, so negative/zero totals are already rejected.

	share, remainder := booty.CalculateShares(total, crew)

	fmt.Printf("Each crew member gets %d piece(s) of eight.\n", share)
	fmt.Printf("The captain gets %d piece(s) of eight.\n", share*2)
	if remainder != 0 {
		fmt.Printf("Remainder: %d piece(s) of eight left over.\n", remainder)
	}
}
