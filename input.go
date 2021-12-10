package main

import (
	"github.com/nsf/termbox-go"
)

type KeyAction int 
type Direction int

const (
	Move KeyAction = iota + 1
	Pause 
	Exit
)

const (
	Left Direction = iota + 1
	Up
	Right
	Down
)

type KeyPress struct {
	action KeyAction
	direction Direction
}

func onKeyPressListener(channel chan KeyPress) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			switch event.Key {
			case termbox.KeyArrowLeft:
				channel <- KeyPress{action: Move, direction: Left}
			case termbox.KeyArrowUp:
				channel <- KeyPress{action: Move, direction: Up}
			case termbox.KeyArrowRight:
				channel <- KeyPress{action: Move, direction: Right}
			case termbox.KeyArrowDown:
				channel <- KeyPress{action: Move, direction: Down}
			case termbox.KeySpace:
				channel <- KeyPress{action: Pause}
			case termbox.KeyEsc:
				channel <- KeyPress{action: Exit}
			default:
			}
		case termbox.EventError:
			panic(event.Err)
		}
	}
}