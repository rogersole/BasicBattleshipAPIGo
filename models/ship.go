package models

import (
	"fmt"
	"log"
)

type Ship struct {
	Id int
	InitialPosition [2]int
	Positions map[string]bool
}

func NewShip(id int, x int, y int) *Ship {
	ship := Ship{}
	ship.Id = id
	ship.InitialPosition = [2]int{x, y}
	ship.Positions = make(map[string]bool)
	ship.Positions[fmt.Sprintf("%v,%v", x, y)] = false
	ship.Positions[fmt.Sprintf("%v,%v", x, y-1)] = false
	ship.Positions[fmt.Sprintf("%v,%v", x, y-2)] = false
	return &ship
}

func (s *Ship) IsBoundingBox(x int, y int) bool {
	if _, ok := s.Positions[fmt.Sprintf("%v,%v", x, y)]; ok {
		return true
	}
	return false
}

func (s *Ship) Attack(x int, y int) string {
	log.Printf("Attacking %v,%v", x, y)
	s.Positions[fmt.Sprintf("%v,%v", x, y)] = true
	allSunk := true

	for _, v := range s.Positions {
		allSunk = allSunk && v
	}

	if allSunk {
		return "sunk"
	}
	return "hit"
}

func (s *Ship) PositionTouched(x int, y int) bool {
	return s.Positions[fmt.Sprintf("%v,%v", x, y)]
}
