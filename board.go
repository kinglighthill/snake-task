package main

import (
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	backgroundColor = termbox.ColorDefault

	defaultBoardHeight = 15
	defaultBoardWeight = 75
)

type Board struct {
	cells [][]Cell
	width int
	height int
	left int
	top int
	right int
	bottom int
}

func newBoard(width, height int) Board {
	cells := make([][]Cell, height)

	for i := range cells {
		cells[i] = make([]Cell, width)
	}

	return Board{
		cells: cells,
		width: width,
		height: height,
	}
}

func (board Board) Draw(left, top, bottom, right int) {
	verticalRule(left - 1, top, board.height + 1, '│')
	verticalRule(right + 1, top, board.height + 1, '│')

	termbox.SetCell(left-1, top, '┌', defaultColor, backgroundColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, backgroundColor)
	termbox.SetCell(left+board.width, top, '┐', defaultColor, backgroundColor)
	termbox.SetCell(left+board.width, bottom, '┘', defaultColor, backgroundColor)

	horizontalRule(left, top, board.width, termbox.Cell{Ch: '─'})
	horizontalRule(left, bottom, board.width, termbox.Cell{Ch: '─'})
}

func (board *Board) printOn(x, y int, char rune) {
	if (x < board.left || x > board.right || y < board.top || y > board.bottom) {
		return
	}

	termbox.SetCell(x, y, char, defaultColor, defaultColor)
}

func horizontalRule(x, y, w int, cell termbox.Cell) {
	for dx := 0; dx < w; dx++ {
		termbox.SetCell(x+dx, y, '─', defaultColor, backgroundColor)
	}
}

func verticalRule(x, y, height int, char rune) {
	for dy := 0; dy < height; dy++ {
		termbox.SetCell(x, y+dy, '│', defaultColor, backgroundColor)
	}
}