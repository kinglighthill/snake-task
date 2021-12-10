package main

type Snake struct {
	body      []Food
	length    int
	direction Direction
}

const (
	empty rune = '0'
	head rune = 'o'
	body rune = 'x'
)

func newSnake() Snake {
	return Snake{
		body: []Food{
			newFood(2, 3),
			newFood(2, 2),
			newFood(2, 1),
		},
		length:    1,
		direction: Right,
	}
}

func (snake Snake) Draw(board *Board, left, top int) {
	for i, boardRow := range board.cells {
		for j, boardColumn := range boardRow {
			if boardColumn.position == head || boardColumn.position == body {
				board.cells[i][j].position = empty
			}
		}
	}

	for i, j := range snake.body {
		if i == 0 {
			board.cells[j.X()][j.Y()].position = head
		} else {
			board.cells[j.X()][j.Y()].position = body
		}
	}

	for i, boardRow := range board.cells {
		for j, boardColumn := range boardRow {
			if boardColumn.position == head {
				board.printOn(left+j+1, top+i+1, head)
			} else if boardColumn.position == body {
				board.printOn(left+j+1, top+i+1, body)
			}
		}
	}
}

func (snake *Snake) Head() Food {
	return snake.body[0]
}

func (snake *Snake) Tail() Food {
	return snake.body[len(snake.body)-1]
}

func (snake *Snake) Move() {
	head := snake.Head()

	switch snake.direction {
	case Left:
		head.updateY(-1)
	case Up:
		head.updateX(-1)
	case Right:
		head.updateY(1)
	case Down:
		head.updateX(1)
	}

	body := make([]Food, 1)
	body[0] = head
	body = append(body, snake.body[:len(snake.body)-1]...)

	if snake.length > len(snake.body) - 2 {
		tail := snake.Tail()
		tail.updateY(-1)
		body = append(body, tail)
	}

	snake.body = body
}

func (snake *Snake) ReDirect(new Direction) {
	if snake.direction == new {
		return
	}

	directionOpposites := map[Direction]Direction{
		Left:  Right,
		Up:    Down,
		Right: Left,
		Down:  Up,
	}

	if directionOpp, ok := directionOpposites[new]; !ok || snake.direction == directionOpp {
		return
	}

	snake.direction = new
}
