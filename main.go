package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// init game

	// load resources

	// game loop

	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	for {
		//update screen
		printScreen()

		//process input

		//process movement

		//process collisions

		//check game over
		break
	}
}

func loadMaze() error {

	f, err := os.Open("maze01.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

var maze []string

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}
