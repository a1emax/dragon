package domain

import (
	"dragon/pkg/domain/space"
)

const (
	levelWidth  = 32
	levelHeight = 32
)

var levelASCIIMap = [levelHeight]string{
	"#                              #",
	"                                ",
	"    ***********~~***********    ",
	"    *                      *    ",
	"    ~          $$          ~    ",
	"    ~          $$          ~    ",
	"    *                      *    ",
	"    *+********+**+********+*    ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"        ~~~~~~~  ~~~~~~~        ",
	"        ~~~~~~~  ~~~~~~~        ",
	"        ~~~~~~~  ~~~~~~~        ",
	"       *~              ~*       ",
	"       *~    $$  $$    ~*       ",
	"       +      $$$$      +       ",
	"       *~    $$$$$$    ~*       ",
	"       *~~~~~~~~~~~~~~~~*       ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"~~~~~*+*~~~~~******~~~~~*+*~~~~~",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"#                              #",
}

func newLevel() (tileMap *TileMap, doors []space.TileVec, treasureCount int) {
	tileMap = newTileMap(levelWidth, levelHeight)

	for y := 0; y < levelHeight; y++ {
		for x := 0; x < levelWidth; x++ {
			tile := tileMap.Tile(x, y)

			tile.setGround(Floor)

			switch levelASCIIMap[y][x] {
			case '~':
				tile.setGround(Magma)
			case '*':
				tile.setObject(Wall)
			case '+':
				tile.setObject(SmallWall)
			case '#':
				tile.setObject(Door)
				doors = append(doors, space.TileVec{x, y})
			case '$':
				tile.setObject(Treasure)
				treasureCount++
			}
		}
	}

	return tileMap, doors, treasureCount
}
