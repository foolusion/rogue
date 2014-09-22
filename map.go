package main

import "github.com/nsf/termbox-go"

type tile struct {
	canWalk bool
}

func newTile() *tile {
	return &tile{canWalk: true}
}

type rMap struct {
	width, height int
	tiles         []tile
}

func newRMap(width, height int) *rMap {
	m := &rMap{
		width:  width,
		height: height,
		tiles:  make([]tile, width*height),
	}
	for i := range m.tiles {
		m.tiles[i].canWalk = true
	}
	m.setWall(10, 15)
	m.setWall(20, 15)
	return m
}

func (m *rMap) isWall(x, y int) bool {
	return !m.tiles[x+y*m.width].canWalk
}

func (m *rMap) setWall(x, y int) {
	m.tiles[x+y*m.width].canWalk = false
}

func (m *rMap) render() {
	const (
		wallColor   termbox.Attribute = termbox.ColorBlue
		groundColor                   = termbox.ColorDefault
	)

	for x := 0; x < m.width; x++ {
		for y := 0; y < m.height; y++ {
			var color termbox.Attribute
			if m.isWall(x, y) {
				color = wallColor
			} else {
				color = groundColor
			}
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, color)
		}
	}
}
