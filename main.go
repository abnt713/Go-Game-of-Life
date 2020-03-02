package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/abnt713/gameoflife/game"
)

func main() {
	eng := game.NewEngine()
	go eng.Start(40, 40, 500*time.Millisecond)
	waitForInterrupt()
	eng.Stop()
}

func waitForInterrupt() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	<-signalChannel
}
