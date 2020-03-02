package game

type cellFactory interface {
	createCell(x, y int, alive bool) cell
}

type gameCellFactory struct {
	w *world
}

func newGameCellFactory() *gameCellFactory {
	return &gameCellFactory{}
}

func (gameCellFac *gameCellFactory) setWorld(w *world) {
	gameCellFac.w = w
}

func (gameCellFac *gameCellFactory) createCell(x, y int, alive bool) cell {
	return newWorldCell(gameCellFac.w, x, y, alive)
}

type smallExploder struct {
	cellFac  cellFactory
	exploder [][]int8
}

func newSmallExploder(cellFac cellFactory) *smallExploder {

	exploder := [][]int8{
		[]int8{0, 1, 0},
		[]int8{1, 1, 1},
		[]int8{1, 0, 1},
		[]int8{0, 1, 0},
	}

	return &smallExploder{
		cellFac:  cellFac,
		exploder: exploder,
	}
}

func (builder *smallExploder) createCellGrid(width, height int) [][]cell {
	// halfwayX := width / 2
	// halfwayY := height / 2

	// exploderX := halfwayX - 1
	// exploderY := halfwayY - 1

	exploderX := 5
	exploderY := 5

	cells := [][]cell{}

	for y := 0; y < height; y++ {
		xCells := []cell{}
		for x := 0; x < width; x++ {
			alive := builder.shouldBeAlive(x, y, exploderX, exploderY)
			cell := builder.cellFac.createCell(x, y, alive)
			xCells = append(xCells, cell)
		}
		cells = append(cells, xCells)
	}

	return cells
}

func (builder *smallExploder) shouldBeAlive(x, y, exploderX, exploderY int) bool {
	if x < exploderX || x > exploderX+2 {
		return false
	}

	if y < exploderY || y > exploderY+3 {
		return false
	}

	adjustedY := y - exploderY
	adjustedX := x - exploderX

	if builder.exploder[adjustedY][adjustedX] == 1 {
		return true
	}

	return false
}
