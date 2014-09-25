package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

func main() {
	gEngine = newEngine()
	defer gEngine.shutdown()

	for !gEngine.done {
		gEngine.update()
		gEngine.render()
		if err := termbox.Flush(); err != nil {
			log.Fatal(err)
		}
	}
}
