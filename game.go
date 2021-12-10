package main

import (
	"math/rand"
	"errors"
	"fmt"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Game struct {
	board Board
	snake Snake
	food Food
	score int16
	paused bool
	round int16
}

func NewGame(width, height int) Game {
	game := Game{
		board: newBoard(width, height),
		snake: newSnake(),
	}

	game.food = game.createFood()

	return game
}

func (game Game) Start() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	defer termbox.Close()

	channel := make(chan KeyPress)
	go onKeyPressListener(channel)

	if err := game.Render(); err != nil {
		return err
	}

Main: 
	for {

		select {
		case key := <-channel:
			switch key.action {
			case Move:
				if !game.paused {
					game.snake.ReDirect(key.direction)
				}
			case Pause:
				game.paused = !game.paused
			case Exit:
				break Main
			}
		default:

		}

		if err := game.MoveSnake(); err != nil {
			return err
		}

		if err := game.Render(); err != nil {
			return err
		}

		time.Sleep(500 * time.Millisecond)
	}	
	return nil
}

func (game Game) Render() error {
	err := termbox.Clear(defaultColor, defaultColor)

	if err != nil {
		return err
	}

	var (
		width, height = termbox.Size()
		midY = height / 2
		midX = width / 2
		left = midX - (game.board.width / 2)
		right = midX + (game.board.width / 2)
		top = midY - (game.board.height / 2)
		bottom = midY + (game.board.height / 2) + 1
	)

	game.board.left = left
	game.board.top = top
	game.board.right = right
	game.board.bottom = bottom

	printState(left, top - 7, fmt.Sprintf("Board Width: %d", game.board.width))
	printState(left, top - 6, fmt.Sprintf("Board Height: %d", game.board.height))
	printState(left, top - 5, fmt.Sprintf("Current Round: %d", game.round))
	printState(left, top - 4, fmt.Sprintf("Score: %d", game.score))
	printState(left, top - 3, fmt.Sprintf("Snake Length: %d", game.snake.length))
	printState(left, top - 2, fmt.Sprintf("Snake Head: %d", game.snake.Head().coordinate))
	printState(left, top - 1, "Press SPACE to pause, press again to resume then ESC to quit")

	game.board.Draw(left, top, bottom, right)
	game.snake.Draw(&game.board, left, top)
	game.food.Draw(&game.board, left, top)

	return termbox.Flush()
}

func (game *Game) MoveSnake() error {
	if game.paused {
		return nil
	}

	game.snake.Move()
	game.IncrementRound()

	if game.snakeAteFood() {
		game.IncrementScore()
		game.snake.length++
		game.food = game.createFood()
	}

	if game.snakeHitSelf() {
		return errors.New("snake has hit itself")
	}

	if game.snakeHitWall() {
		return errors.New("snake has hit wall")
	}

	return nil
}

func (game *Game) snakeAteFood() bool {
	head := game.snake.Head()
	food := game.food
	
	return head.X() == food.X() && head.Y() == food.Y()
}

func (game *Game) snakeHitWall() bool {
	head := game.snake.Head()
	board := game.board

	return head.X() < 0 || head.Y() < 0 || head.X() > int32(board.height - 1) || head.Y() > int32(board.width - 1)
}

func (game *Game) snakeHitSelf() bool {
	snake := game.snake
	head := snake.Head()

	for _, body := range snake.body[1:] {
		if head.X() == body.X() && head.Y() == body.Y() {
			return true
		}
	}

	return false
}

func (game *Game) createFood() Food {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	boardRow := random.Intn(game.board.height - 1)
	boardColumn := random.Intn(game.board.width - 1)

	if game.board.cells[boardRow][boardColumn].position == body {
		return game.createFood()
	}

	return newFood(int32(boardRow), int32(boardColumn))
}

func (game *Game) IncrementRound() {
	game.round += 1
}
func (game *Game) IncrementScore() {
	game.score += 1
}

func printState(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, termbox.ColorCyan, defaultColor)
		x += runewidth.RuneWidth(c)
	}
}