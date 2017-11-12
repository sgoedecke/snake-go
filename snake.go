package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	// initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256) // set 256-color mode

	// initialize an event queue and poll eternally, sending events to a channel
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	snake := Snake{2, []SnakeNode{}, nil}
  snake.Body = append(snake.Body, SnakeNode{5,5})
	world := World{80, 40, &snake, nil}
  world.PlaceFood()
	snake.World = &world

	go world.Loop()

	// listen for key presses
	for {
		event := <-eventQueue
		if event.Type == termbox.EventKey {
			switch {
			case event.Ch == 'w' || event.Key == termbox.KeyArrowUp:
				snake.GoUp()
			case event.Ch == 's' || event.Key == termbox.KeyArrowDown:
				snake.GoDown()
			case event.Ch == 'a' || event.Key == termbox.KeyArrowLeft:
				snake.GoLeft()
			case event.Ch == 'd' || event.Key == termbox.KeyArrowRight:
				snake.GoRight()
			case event.Ch == 'q' || event.Key == termbox.KeyEsc || event.Key == termbox.KeyCtrlC || event.Key == termbox.KeyCtrlD:
				return
			}
		}
	}
}
