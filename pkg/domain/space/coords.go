package space

import (
	"math"

	"github.com/a1emax/youngine/basic"
)

const WorldPerTile = 32

type WorldDim = basic.Float

type TileDim = int

func WorldToTileDim(w WorldDim) TileDim {
	return TileDim(math.Floor(w / WorldPerTile))
}

func TileToWorldDim(t TileDim) WorldDim {
	return WorldDim(t * WorldPerTile)
}

func TileToWorldCtrDim(t TileDim) WorldDim {
	return TileToWorldDim(t) + WorldPerTile/2.0
}

type WorldVec = basic.Vec2

type TileVec = basic.Ivec2

func WorldToTileVec(w WorldVec) TileVec {
	return TileVec{WorldToTileDim(w[0]), WorldToTileDim(w[1])}
}

func TileToWorldVec(t TileVec) WorldVec {
	return WorldVec{TileToWorldDim(t[0]), TileToWorldDim(t[1])}
}

func TileToWorldCtrVec(t TileVec) WorldVec {
	return WorldVec{TileToWorldCtrDim(t[0]), TileToWorldCtrDim(t[1])}
}
