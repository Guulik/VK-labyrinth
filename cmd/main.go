package main

import (
	"VK-contora/internal/input"
	"VK-contora/internal/solver"
	"fmt"
	"os"
)

func main() {
	labyrinth, start, end, err := input.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	path, err := solver.ShortestPath(labyrinth, start, end)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, p := range path {
		fmt.Printf("%d %d\n", p.X, p.Y)
	}
	fmt.Println(".")
}
