package game

import (
	"time"
)

// Engine exposes the game
type Engine struct {
	running bool
}

// NewEngine creates a new engine struct
func NewEngine() *Engine {
	return &Engine{
		running: false,
	}
}

// Start starts the game
func (eng *Engine) Start(width, height int, pace time.Duration) {
	eng.running = true

	cellFac := newGameCellFactory()
	cellGridFac := newSmallExploder(cellFac)
	canvas := newTerminalCanvas()
	gameWorld := newWorld(cellGridFac, canvas, width, height)

	cellFac.setWorld(gameWorld)
	gameWorld.build()

	for eng.running {
		gameWorld.step()
		time.Sleep(pace)
	}
}

// Stop stops the game
func (eng *Engine) Stop() {
	eng.running = false
}
