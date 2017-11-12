package main

import (
	"github.com/nsf/termbox-go"
)

func (w *World) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  // draw background
  for y := 0; y < w.Height+1; y++ {
    for x := 0; x < w.Width+1; x++ {
      termbox.SetCell(x, y, 32, termbox.ColorDefault, termbox.ColorWhite)
    }
  }

	w.Snake.Draw()
	w.Food.Draw()

	_ = termbox.Flush()
}

func (s *Snake) Draw() {
	color := termbox.ColorRed
	for _, n := range s.Body {
		termbox.SetCell(n.X, n.Y, 32, termbox.ColorDefault, color)
	}
}

func (f *Food) Draw() {
	color := termbox.ColorCyan
	termbox.SetCell(f.X, f.Y, 32, termbox.ColorDefault, color)
}
