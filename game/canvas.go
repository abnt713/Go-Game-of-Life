package game

import (
	"fmt"
	"os"
	"os/exec"
)

type terminalCanvas struct{}

func newTerminalCanvas() *terminalCanvas {
	return &terminalCanvas{}
}

func (canvas *terminalCanvas) draw(step int, cells [][]cell) {
	canvas.clearTerminal()
	fmt.Printf("Step: %v\n\n", step)
	cellsHeight := len(cells)
	cellsWidth := len(cells[0])
	for y := 0; y < cellsHeight; y++ {
		for x := 0; x < cellsWidth; x++ {
			isAlive := cells[y][x].isAlive()
			if isAlive {
				fmt.Print("0")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (canvas *terminalCanvas) clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	cmd.Run()
}
