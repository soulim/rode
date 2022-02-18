package main

import (
	"os"

	"github.com/soulim/rode/internal/rode"
)

func main() {
	if err := rode.Run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
