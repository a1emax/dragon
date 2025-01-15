package space

import (
	"math"
)

func ClipViewDim(area, focus TileDim, viewport WorldDim) (offset WorldDim, from, to TileDim) {
	if wArea := TileToWorldDim(area); wArea < viewport {
		return (viewport - wArea) / 2.0, 0, area - 1
	}

	wFocus := TileToWorldCtrDim(focus)

	wFrom := wFocus - viewport/2.0
	from = WorldToTileDim(wFrom)
	if from < 0 {
		wFrom = 0.0
		from = 0
	}

	offset = TileToWorldDim(from) - wFrom

	to = WorldToTileDim(wFrom + viewport - 1.0)
	if to >= area {
		tViewport := TileDim(math.Ceil(viewport / WorldPerTile))

		offset = viewport - TileToWorldDim(tViewport)
		from = area - tViewport
		to = area - 1
	}

	return offset, from, to
}

func ClipViewVec(area, focus TileVec, viewport WorldVec) (offset WorldVec, from, to TileVec) {
	offset[0], from[0], to[0] = ClipViewDim(area[0], focus[0], viewport[0])
	offset[1], from[1], to[1] = ClipViewDim(area[1], focus[1], viewport[1])

	return offset, from, to
}
