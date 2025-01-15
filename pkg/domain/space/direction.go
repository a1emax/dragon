package space

import (
	"fmt"
)

type Direction int

const (
	_ Direction = iota
	North
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return fmt.Sprintf("<invalid direction %d>", d)
	}
}

func (d Direction) TurnCW() Direction {
	if d < North || d > West {
		return 0 // Panic?
	}

	d++
	if d > West {
		d = North
	}

	return d
}

func (d Direction) TurnCCW() Direction {
	if d < North || d > West {
		return 0 // Panic?
	}

	d--
	if d < North {
		d = West
	}

	return d
}

func (d Direction) Invert() Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	default:
		return 0 // Panic?
	}
}

func AddDirection(t TileVec, d Direction) TileVec {
	switch d {
	case North:
		return t.Add(TileVec{0, -1})
	case East:
		return t.Add(TileVec{1, 0})
	case South:
		return t.Add(TileVec{0, 1})
	case West:
		return t.Add(TileVec{-1, 0})
	default:
		return t // Panic?
	}
}

func GetDirection(from, to TileVec) (Direction, bool) {
	v := to.Sub(from)
	switch {
	case v.X() == 0 && v.Y() < 0:
		return North, true
	case v.X() > 0 && v.Y() == 0:
		return East, true
	case v.X() == 0 && v.Y() > 0:
		return South, true
	case v.X() < 0 && v.Y() == 0:
		return West, true
	default:
		return 0, false
	}
}
