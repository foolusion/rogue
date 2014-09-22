package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

func main() {
	engine := newEngine()
	defer engine.shutdown()

	for !engine.done {
		engine.update()
		engine.render()
		if err := termbox.Flush(); err != nil {
			log.Fatal(err)
		}
	}
}
