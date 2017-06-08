package models

import (
	"fmt"
	"errors"
	"strconv"
)

const BOARD_SIDE = 10

type Board struct {
	Ships []*Ship
}

func NewBoard(shipPositions [][]int) (*Board, error) {

	board := Board{}

	for idx, shipPosition := range shipPositions {
		if !inBoundaries(shipPosition[0]) {
			return nil, errors.New("X positions out of boundaries")
		}
		if !inBoundaries(shipPosition[1]) {
			return nil, errors.New("Y positions out of boundaries")
		}
		ship := NewShip(idx + 1, shipPosition[0], shipPosition[1])
		board.Ships = append(board.Ships, ship)
	}

	return &board, nil
}

func (b *Board) Attack(x int, y int) (string, error) {

	if !inBoundaries(x) {
		return "", errors.New("X Position out of boundaries")
	}

	if !inBoundaries(y) {
		return "", errors.New("Y Position out of boundaries")
	}

	result := "miss"
	for _, ship := range b.Ships {
		if ship.IsBoundingBox(x, y) {
			result = ship.Attack(x, y)
		}
	}

	return result, nil
}

func inBoundaries(x int) bool {
	return x >= 0 && x < BOARD_SIDE
}

func (b *Board) ShowBattlefield() {

	// initialize battlefield
	battlefield := make([]string, BOARD_SIDE*BOARD_SIDE)
	for idx := range battlefield {
		battlefield[idx] = "0"
	}

	// update battlefield
	for _, ship := range b.Ships {
		sx := ship.InitialPosition[0]
		sy := ship.InitialPosition[1]

		gy := BOARD_SIDE - ship.InitialPosition[1] + 1

		var touched bool
		shipId := strconv.Itoa(ship.Id)
		idx0 := gy * BOARD_SIDE + sx
		if touched = ship.PositionTouched(sx, sy - 2); touched {
			battlefield[idx0] = "X"
		} else {
			battlefield[idx0] = shipId
		}

		idx1 := (gy - 1) * BOARD_SIDE + sx
		if touched = ship.PositionTouched(sx, sy - 1); touched {
			battlefield[idx1] = "X"
		} else {
			battlefield[idx1] = shipId
		}

		idx2 := (gy - 2) * BOARD_SIDE + sx
		if touched = ship.PositionTouched(sx, sy); touched {
			battlefield[idx2] = "X"
		} else {
			battlefield[idx2] = shipId
		}

	}

	// print battlefield
	fmt.Print("\n==== BATTLEFIELD STATUS =============================================")
	xrange := make([]int, BOARD_SIDE * BOARD_SIDE)
	for idx, _ := range xrange {
		if idx % BOARD_SIDE == 0 {
			fmt.Println("")
		}
		x := idx % BOARD_SIDE
		y := BOARD_SIDE - (idx/BOARD_SIDE) - 1
		fmt.Printf("%v%v [%v] ", x, y, battlefield[idx])
	}
	fmt.Println("\n=====================================================================\n")
}
