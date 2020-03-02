package game

type cell interface {
	isAlive() bool
	step()
	die()
	born()
}

type worldCell struct {
	w     *world
	x     int
	y     int
	alive bool
}

func newWorldCell(w *world, x, y int, alive bool) *worldCell {
	return &worldCell{
		w:     w,
		x:     x,
		y:     y,
		alive: alive,
	}
}

func (wCell *worldCell) step() {
	neighbors := wCell.countNeighbors()
	if wCell.isAlive() {
		wCell.behaveAsAlive(neighbors)
	} else {
		wCell.behaveAsEmpty(neighbors)
	}
}

func (wCell *worldCell) countNeighbors() int {
	neighbors := 0

	for yMod := -1; yMod < 2; yMod++ {
		for xMod := -1; xMod < 2; xMod++ {
			xCoord := wCell.x + xMod
			yCoord := wCell.y + yMod

			if yMod == 0 && xMod == 0 {
				continue
			}

			currCell, err := wCell.w.getCellAt(xCoord, yCoord)
			if err != nil {
				continue
			}

			if currCell.isAlive() {
				neighbors++
			}
		}
	}

	return neighbors
}

func (wCell *worldCell) isAlive() bool {
	return wCell.alive
}

func (wCell *worldCell) behaveAsAlive(neighbors int) {
	if neighbors < 2 || neighbors > 3 {
		wCell.w.setCellStatusAt(wCell.x, wCell.y, false)
	}
}

func (wCell *worldCell) behaveAsEmpty(neighbors int) {
	if neighbors == 3 {
		wCell.w.setCellStatusAt(wCell.x, wCell.y, true)
	}

}

func (wCell *worldCell) die() {
	wCell.alive = false
}

func (wCell *worldCell) born() {
	wCell.alive = true
}
