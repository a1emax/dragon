package world

import (
	"image/color"

	"github.com/a1emax/youngine/x/colors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"dragon/pkg/domain/space"
)

var images struct {
	floor        *ebiten.Image
	magma        *ebiten.Image
	wall         *ebiten.Image
	smallWall    *ebiten.Image
	door         *ebiten.Image
	treasure     *ebiten.Image
	knight       *ebiten.Image
	loadedKnight *ebiten.Image
	dragonHead   *ebiten.Image
	dragonBody   *ebiten.Image
	dragonTail   *ebiten.Image
}

func init() {
	images.floor = newFilledImage(colors.RGBA{0xCC, 0xA7, 0x66, 0xFF})
	images.magma = newFilledImage(colors.RGBA{0xFA, 0x50, 0x20, 0xFF})
	images.wall = newSquareImage(1, colors.RGBA{0x6E, 0x5A, 0x37, 0xFF})
	images.smallWall = newSquareImage(1, colors.RGBA{0x94, 0x79, 0x4A, 0xFF})
	images.door = newSquareImage(1, colors.RGBA{0x40, 0x34, 0x20, 0xFF})
	images.treasure = newSquareImage(1, colors.RGBA{0xFC, 0xC4, 0x28, 0xFF})
	images.knight = newCircleImage(0.8, colors.RGBA{0x60, 0x60, 0x60, 0xFF})
	images.loadedKnight = newCircleImage(0.8, colors.RGBA{0xFC, 0xC4, 0x28, 0xFF})
	images.dragonHead = newCircleImage(0.8, colors.RGBA{0x3D, 0x57, 0x0F, 0xFF})
	images.dragonBody = newCircleImage(0.8, colors.RGBA{0x55, 0x7A, 0x15, 0xFF})
	images.dragonTail = newCircleImage(0.8, colors.RGBA{0x69, 0x94, 0x1F, 0xFF})
}

func newFilledImage(clr color.Color) *ebiten.Image {
	img := ebiten.NewImage(space.WorldPerTile, space.WorldPerTile)
	img.Fill(clr)

	return img
}

func newSquareImage(offset float32, clr color.Color) *ebiten.Image {
	img := ebiten.NewImage(space.WorldPerTile, space.WorldPerTile)
	xy := offset
	wh := space.WorldPerTile - offset*2
	vector.DrawFilledRect(img, xy, xy, wh, wh, clr, false)

	return img
}

func newCircleImage(scale float32, clr color.Color) *ebiten.Image {
	img := ebiten.NewImage(space.WorldPerTile, space.WorldPerTile)
	cxy := float32(space.WorldPerTile) / 2.0
	r := space.WorldPerTile * scale / 2.0
	vector.DrawFilledCircle(img, cxy, cxy, r, clr, true)

	return img
}
