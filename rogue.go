/* Copyright 2014 Andrew O'Neill

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"log"
	"os"

	"github.com/nsf/termbox-go"
)

type entity struct {
	x, y     int
	graphics GraphicsComponent
}

var player = &entity{
	x: 0,
	y: 0,
	graphics: TermboxGraphicsComponent{
		ch: '@',
		fg: termbox.ColorDefault,
		bg: termbox.ColorDefault,
	},
}

type GraphicsComponent interface {
	Draw(e *entity)
}

type TermboxGraphicsComponent struct {
	ch     rune
	fg, bg termbox.Attribute
}

func (t TermboxGraphicsComponent) Draw(e *entity) {
	termbox.SetCell(e.x, e.y, t.ch, t.fg, t.bg)
}

// Graphics is the interface that will draw things to the screen.
// TODO(foolusion@gmail.com): change this to be a graphics manager
// and give the player graphics components.
type Graphics interface {
	Init() error
	Draw([]GraphicsComponent) error
	Shutdown()
}

// TermboxGraphics implements the graphics interface for termbox.
type TermboxGraphics struct {
	gc []*entity
}

// Start implements the graphics interface.
func (t TermboxGraphics) Init() error {
	if err := termbox.Init(); err != nil {
		termbox.Close()
		return err
	}
	return nil
}

// Draw implements the graphics interface
func (t TermboxGraphics) Draw() error {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		return err
	}
	for _, v := range t.gc {
		v.graphics.Draw(v)
	}
	if err := termbox.Flush(); err != nil {
		return err
	}
	return nil
}

func (t TermboxGraphics) Shutdown() {
	termbox.Close()
}

func main() {
	f, err := os.Create("log.txt")
	defer f.Close()
	if err != nil {
		log.Fatal()
	}
	log.SetOutput(f)

	graphics := TermboxGraphics{
		gc: []*entity{
			player,
		},
	}
	if err := graphics.Init(); err != nil {
		log.Fatal(err)
	}
	defer graphics.Shutdown()

	//TODO(foolusion@gmail.com): Move this into an input manager interface
	termbox.SetInputMode(termbox.InputAlt)

	if err := graphics.Draw(); err != nil {
		log.Fatal(err)
	}

	ch := make(chan termbox.Event)
	go func() {
		for {
			ch <- termbox.PollEvent()
		}
	}()
	for {
		e := <-ch
		log.Printf("%+v", e)
		switch {
		case e.Ch == 'q', e.Ch == 'Q':
			return
		case e.Key == termbox.KeyEsc:
			termbox.Close()
		case e.Key == termbox.KeyArrowUp:
			doUpAction(player)
		case e.Key == termbox.KeyArrowDown:
			doDownAction(player)
		case e.Key == termbox.KeyArrowLeft:
			doLeftAction(player)
		case e.Key == termbox.KeyArrowRight:
			doRightAction(player)
		}
		graphics.Draw()
	}
}

func doUpAction(e *entity) {
	if player.y > 0 {
		player.y--
	}
}

func doDownAction(e *entity) {
	_, h := termbox.Size()
	if player.y < h-1 {
		player.y++
	}
}

func doLeftAction(e *entity) {
	if player.x > 0 {
		player.x--
	}
}
func doRightAction(e *entity) {
	w, _ := termbox.Size()
	if player.x < w-1 {
		player.x++
	}
}
