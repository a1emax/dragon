package button

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/x/roundrect"
	"github.com/hajimehoshi/ebiten/v2"
)

type shapePart struct {
	shapeKey basic.Opt[shapeKey]
	shape    *ebiten.Image
}

type shapeKey struct {
	size         basic.Vec2
	cornerRadius basic.Float
}

func (b *buttonImpl[R]) setupShape() {
	r := b.region.Rect()

	var cornerRadius basic.Float
	if b.state.CornerRadius.IsSet() {
		cornerRadius = b.state.CornerRadius.Get()
	} else {
		cornerRadius = r.Height() / 2
	}

	key := shapeKey{
		size:         r.Size,
		cornerRadius: cornerRadius,
	}

	if b.shapeKey.IsSet() && b.shapeKey.Get() == key {
		return
	}

	b.disposeShape()

	bmp := roundrect.Fill(key.size.X(), key.size.Y(), key.cornerRadius, key.cornerRadius)
	img := ebiten.NewImage(bmp.Width(), bmp.Height())
	img.WritePixels(bmp.Data())

	b.shapeKey = basic.SetOpt(key)
	b.shape = img
}

func (b *buttonImpl[R]) drawShape(screen *ebiten.Image) {
	if !b.shapeKey.IsSet() {
		return
	}

	clr := b.color(b.state.PrimaryColor, b.state.PressedColor)
	if clr == nil {
		return
	}

	r := b.region.Rect()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.Left(), r.Top())
	op.ColorScale.ScaleWithColor(clr)

	screen.DrawImage(b.shape, op)
}

func (b *buttonImpl[R]) disposeShape() {
	if !b.shapeKey.IsSet() {
		return
	}

	b.shape.Deallocate()

	b.shapePart = shapePart{}
}
