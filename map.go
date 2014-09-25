package main

import "github.com/nsf/termbox-go"

type tile struct {
	canWalk bool
}

const (
	roomMaxSize = 12
	roomMinSize = 6
)

func newTile() *tile {
	return &tile{canWalk: false}
}

type rMap struct {
	e             *engine
	width, height int
	tiles         []tile
}

func newRMap(e *engine, width, height int) *rMap {
	m := &rMap{
		e:      e,
		width:  width,
		height: height,
		tiles:  make([]tile, width*height),
	}
	x1, y1 := e.rng.Intn(width), e.rng.Intn(height)
	x2, y2 := x1+roomMinSize+e.rng.Intn(roomMaxSize-roomMinSize), y1+roomMinSize+e.rng.Intn(roomMaxSize-roomMinSize)
	m.createRoom(true, x1, y1, x2, y2)

	return m
}

func (m *rMap) createRoom(first bool, x1, y1, x2, y2 int) {
	m.dig(x1, y1, x2, y2)
	if first {
		m.e.player.x = (x1 + x2) / 2
		m.e.player.y = (y1 + y2) / 2
	} else {
		if m.e.rng.Intn(3) == 0 {
			m.e.actors = append(
				m.e.actors,
				newActor((x1+x2)/2,
					(y1+y2)/2,
					'@',
					termbox.ColorDefault),
			)
		}
	}
}

func (m *rMap) dig(x1, y1, x2, y2 int) {
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	for tx := x1; tx <= x2; tx++ {
		for ty := y1; ty <= y2; ty++ {
			m.tiles[tx+ty*m.width].canWalk = true
		}
	}
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
