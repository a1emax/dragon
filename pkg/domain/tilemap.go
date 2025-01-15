package domain

import (
	"dragon/pkg/domain/space"
)

type TileMap struct {
	tiles  []Tile
	width  int
	height int
}

func newTileMap(width, height int) *TileMap {
	return &TileMap{
		tiles:  make([]Tile, width*height),
		width:  width,
		height: height,
	}
}

func (t *TileMap) Width() int {
	return t.width
}

func (t *TileMap) Height() int {
	return t.height
}

func (t *TileMap) Size() space.TileVec {
	return space.TileVec{t.width, t.height}
}

func (t *TileMap) Tile(x, y int) *Tile {
	if x < 0 || x >= t.width || y < 0 || y >= t.height {
		return nil
	}

	return &t.tiles[y*t.width+x]
}

func (t *TileMap) TileAt(p space.TileVec) *Tile {
	return t.Tile(p.X(), p.Y())
}
