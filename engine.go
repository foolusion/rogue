package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

type engine struct {
	actors []*actor
	player *actor
	rmap   *rMap
	done   bool
}

func newEngine() *engine {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	termbox.SetInputMode(termbox.InputAlt)

	e := &engine{done: false}
	e.player = newActor(0, 0, '@', termbox.ColorDefault)
	e.actors = append(e.actors, e.player)
	e.actors = append(e.actors, newActor(10, 20, '@', termbox.ColorRed))
	w, h := termbox.Size()
	e.rmap = newRMap(w, h)
	return e
}

func (e *engine) update() {
	ev := termbox.PollEvent()
	switch {
	case ev.Ch == 'q', ev.Ch == 'Q':
		e.done = true
	case ev.Key == termbox.KeyArrowUp:
		if !e.rmap.isWall(e.player.x, e.player.y-1) {
			e.player.y--
		}
	case ev.Key == termbox.KeyArrowDown:
		if !e.rmap.isWall(e.player.x, e.player.y+1) {
			e.player.y++
		}
	case ev.Key == termbox.KeyArrowLeft:
		if !e.rmap.isWall(e.player.x-1, e.player.y) {
			e.player.x--
		}
	case ev.Key == termbox.KeyArrowRight:
		if !e.rmap.isWall(e.player.x+1, e.player.y) {
			e.player.x++
		}
	}
}

func (e *engine) render() {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		log.Fatal(err)
	}
	e.rmap.render()
	for _, a := range e.actors {
		termbox.SetCell(a.x, a.y, a.ch, a.col, termbox.ColorDefault)
	}
}

func (e *engine) shutdown() {
	termbox.Close()
}
