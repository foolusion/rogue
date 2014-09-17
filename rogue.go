/*
   Copyright 2014 Andrew O'Neill

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
	x, y   int
	ch     rune
	fg, bg termbox.Attribute
}

var player = &entity{
	x:  0,
	y:  0,
	ch: '@',
	fg: termbox.ColorDefault,
	bg: termbox.ColorDefault,
}

func main() {
	f, err := os.Create("log.txt")
	defer f.Close()
	if err != nil {
		log.Fatal()
	}
	log.SetOutput(f)

	if err := termbox.Init(); err != nil {
		termbox.Close()
		log.Fatal(err)
	}
	termbox.SetInputMode(termbox.InputAlt)

	draw()

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
			termbox.Close()
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
		draw()
	}
}

func draw() {
	if err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault); err != nil {
		termbox.Close()
		log.Fatal(err)
	}
	termbox.SetCell(player.x, player.y, player.ch, player.fg, player.bg)
	if err := termbox.Flush(); err != nil {
		termbox.Close()
		log.Fatal(err)
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
