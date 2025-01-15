package world

import (
	"github.com/a1emax/youngine/x/colors"
	"github.com/hajimehoshi/ebiten/v2"

	"dragon/pkg/domain"
	"dragon/pkg/domain/space"
	"dragon/pkg/global/vars"
)

func (w *worldImpl[R]) Draw(screen *ebiten.Image) {
	dst := screen.SubImage(w.region.Rect().Image()).(*ebiten.Image)
	dst.Fill(colors.RGBA{0x00, 0x00, 0x00, 0xFF})

	session := vars.GamePlay.Session
	offset, from, to := space.ClipViewVec(session.TileMap().Size(), session.Dragon().Head(), w.region.Rect().Size)
	for y := from.Y(); y <= to.Y(); y++ {
		for x := from.X(); x <= to.X(); x++ {
			p := space.TileToWorldVec(space.TileVec{x, y}.Sub(from)).Add(offset).Add(w.region.Rect().Min)
			tile := session.TileMap().Tile(x, y)

			w.drawGround(dst, p, tile)
			w.drawObject(dst, p, tile)
			w.drawKnight(dst, p, tile)
			w.drawDragon(dst, p, tile)
		}
	}
}

func (w *worldImpl[R]) drawGround(dst *ebiten.Image, p space.WorldVec, tile *domain.Tile) {
	switch tile.GroundTag() {
	case domain.Floor:
		w.drawImage(dst, p, images.floor)
	case domain.Magma:
		w.drawImage(dst, p, images.magma)
	}
}

func (w *worldImpl[R]) drawObject(dst *ebiten.Image, p space.WorldVec, tile *domain.Tile) {
	switch tile.ObjectTag() {
	case domain.Wall:
		w.drawImage(dst, p, images.wall)
	case domain.SmallWall:
		w.drawImage(dst, p, images.smallWall)
	case domain.Door:
		w.drawImage(dst, p, images.door)
	case domain.Treasure:
		w.drawImage(dst, p, images.treasure)
	}
}

func (w *worldImpl[R]) drawKnight(dst *ebiten.Image, p space.WorldVec, tile *domain.Tile) {
	if tile.KnightID() == 0 {
		return
	}

	if vars.GamePlay.Session.KnightSet().Get(tile.KnightID()).IsLoaded() {
		w.drawImage(dst, p, images.loadedKnight)
	} else {
		w.drawImage(dst, p, images.knight)
	}
}

func (w *worldImpl[R]) drawDragon(dst *ebiten.Image, p space.WorldVec, tile *domain.Tile) {
	switch tile.DragonPartTag() {
	case domain.DragonHead:
		w.drawImage(dst, p, images.dragonHead)
	case domain.DragonBody:
		w.drawImage(dst, p, images.dragonBody)
	case domain.DragonTail:
		w.drawImage(dst, p, images.dragonTail)
	}
}

func (w *worldImpl[R]) drawImage(dst *ebiten.Image, p space.WorldVec, img *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X(), p.Y())

	dst.DrawImage(img, op)
}
