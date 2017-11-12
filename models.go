package main

import (
	"time"
  "math/rand"
)

type Snake struct {
	Direction int // 0 up, 1 right, 2 down, 3 left
	Body      []SnakeNode
	World     *World
}

func (s *Snake) GoUp() {
	s.Direction = 0
}
func (s *Snake) GoRight() {
	s.Direction = 1
}
func (s *Snake) GoDown() {
	s.Direction = 2
}
func (s *Snake) GoLeft() {
	s.Direction = 3
}

type SnakeNode struct {
	X int
	Y int
}

func (s *Snake) Move() {
	head := s.Body[len(s.Body)-1]
	switch s.Direction {
	case 0:
		s.PlaceNode(head.X, head.Y-1)
		return
	case 1:
		s.PlaceNode(head.X+1, head.Y)
		return
	case 2:
		s.PlaceNode(head.X, head.Y+1)
		return
	case 3:
		s.PlaceNode(head.X-1, head.Y)
		return
	}
}

func (s *Snake) PlaceNode(x int, y int) {
	// check if the walls are hit
	if x < 0 || x > s.World.Width || y < 0 || y > s.World.Height {
		s.Die()
		return
	}

	// check if the snake has hit itself
	for _, n := range s.Body {
		if n.X == x && n.Y == y {
			s.Die()
			return
		}
	}

	// check if the snake has hit food
	food := s.World.Food
	if x == food.X && y == food.Y {
		s.World.PlaceFood()
	} else {
		// remove the last node
		s.Body = s.Body[1:]
	}

	// place the node
	sn := SnakeNode{x, y}
	s.Body = append(s.Body, sn)
}

func (s *Snake) Die() {
	s.Body = s.Body[len(s.Body)-1:]
}

type World struct {
	Width  int
	Height int
	Snake  *Snake
	Food   *Food
}

func (w *World) PlaceFood() {
  r := rand.New(rand.NewSource(time.Now().UnixNano())) // have to re-seed each time, apparently :(
  x := r.Intn(w.Width)
  y := r.Intn(w.Height)

	w.Food = &Food{x,y}
}

func (w *World) Loop() {
  for {
    w.Snake.Move()
    w.Draw()
    time.Sleep(100 * time.Millisecond)
  }
}

type Food struct {
	X int
	Y int
}
