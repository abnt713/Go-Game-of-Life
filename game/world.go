package game

import (
	"fmt"
	"sync"
)

// World
type world struct {
	cellGridFac cellGridFactory
	width       int
	height      int
	canvas      worldCanvas

	cells       [][]cell
	currStep    chan struct{}
	mapUpdateWg sync.WaitGroup
	stepCount   int
}

func newWorld(cellGridFac cellGridFactory, canvas worldCanvas, width int, height int) *world {
	return &world{
		cellGridFac: cellGridFac,
		width:       width,
		height:      height,
		canvas:      canvas,
		stepCount:   0,
	}
}

func (w *world) build() {
	w.cells = w.cellGridFac.createCellGrid(w.width, w.height)
}

func (w *world) getCellAt(x, y int) (cell, error) {
	coordErr := w.checkCoordinates(x, y)
	if coordErr != nil {
		return nil, coordErr
	}

	return w.cells[y][x], nil
}

func (w *world) setCellStatusAt(x, y int, alive bool) error {
	coordErr := w.checkCoordinates(x, y)
	if coordErr != nil {
		return coordErr
	}

	w.mapUpdateWg.Add(1)
	go func() {
		<-w.currStep
		tgt := w.cells[y][x]
		if alive {
			tgt.born()
		} else {
			tgt.die()
		}
		w.mapUpdateWg.Done()
	}()

	return nil
}

func (w *world) checkCoordinates(x, y int) error {
	xLimit := w.width - 1
	if x < 0 || x > xLimit {
		return newOutOfBoundsError(x, xLimit)
	}

	yLimit := w.height - 1
	if y < 0 || y > yLimit {
		return newOutOfBoundsError(y, yLimit)
	}

	return nil
}

func (w *world) step() {
	w.stepCount++
	if w.stepCount == 1 {
		w.canvas.draw(w.stepCount, w.cells)
		return
	}

	w.currStep = make(chan struct{})
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			cell := w.cells[y][x]
			cell.step()
		}
	}

	close(w.currStep)
	w.mapUpdateWg.Wait()

	w.canvas.draw(w.stepCount, w.cells)
}

// Services
type cellGridFactory interface {
	createCellGrid(width, height int) [][]cell
}

type worldCanvas interface {
	draw(step int, cells [][]cell)
}

// Errors
type outOfBoundsError struct {
	coord int
	limit int
}

func newOutOfBoundsError(coord int, limit int) *outOfBoundsError {
	return &outOfBoundsError{
		coord: coord,
		limit: limit,
	}
}

func (err *outOfBoundsError) Error() string {
	return fmt.Sprintf("Invalid coordinate: %v (limit is %v and not negative)", err.coord, err.limit)
}
