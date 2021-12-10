package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)

func main() {
	width, height := readBoardDimen()
	game := NewGame(width, height)

	if err := game.Start(); err != nil {
		log.Fatal(err)
	}
}

func readBoardDimen() (width, height int) {
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the width of the board (preferably 75): ")
	input.Scan()
	widthInput := input.Text()

	width, err := strconv.Atoi(widthInput)

	if err != nil {
		width = defaultBoardWeight
	}

	fmt.Println("Enter the height of the board (preferably 10): ")
	input.Scan()
	heightInput := input.Text()

	height, err = strconv.Atoi(heightInput)

	if err != nil {
		height = defaultBoardHeight
	}

	return
}