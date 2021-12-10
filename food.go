package main

const (
	food rune = '*'
)


type Food struct {
	coordinate [2]int32
}

func newFood(x, y int32) Food {
	food := Food{}
	food.coordinate[0] = x
	food.coordinate[1] = y

	return food
}

func (food *Food) X() int32 {
	return food.coordinate[0]
}

func (food *Food) Y() int32 {
	return food.coordinate[1]
}

func (food *Food) updateX(x int32) {
	food.coordinate[0] += x
}

func (food *Food) updateY(y int32) {
	food.coordinate[1] += y
}

func (f Food) Draw(board *Board, left, top int) {
	board.cells[f.X()][f.Y()].position = food

	for i, boardRow := range board.cells {
		for j, boardColumn := range boardRow {
			if boardColumn.position == food {
				board.printOn(left + j + 1, top + i + 1, food)
			}
		}
	}
}