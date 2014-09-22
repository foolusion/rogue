package main

import "github.com/nsf/termbox-go"

type actor struct {
	x, y int
	ch   rune
	col  termbox.Attribute
}

func newActor(x, y int, ch rune, col termbox.Attribute) *actor {
	return &actor{x: x, y: y, ch: ch, col: col}
}

func (a *actor) render() {
	termbox.SetCell(a.x, a.y, a.ch, a.col, termbox.ColorDefault)
}
