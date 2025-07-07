package main

import (
	"fmt"
	"os"
	"path/filepath"
	twoplustwo "github.com/pepperonirollz/twoplustwo-go"
)

func main() {
	// Change to the root directory where HandRanks.dat should be created
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	// Go up two levels to the project root
	rootDir := filepath.Join(wd, "..", "..")
	if err := os.Chdir(rootDir); err != nil {
		panic(err)
	}
	
	fmt.Println("Starting HandRanks.dat generation...")
	twoplustwo.Generate()
	fmt.Println("HandRanks.dat generation complete!")
}