package colorarea

import (
	"image/color"

	"github.com/a1emax/youngine/basic"
	"github.com/hajimehoshi/ebiten/v2"
)

type fillPart struct {
	fillKey basic.Opt[fillKey]
	fill    *ebiten.Image
}

type fillKey struct {
	color color.Color
}

func (c *colorAreaImpl[R]) setupFill() {
	key := fillKey{
		color: c.state.Color,
	}

	if c.fillKey.IsSet() && c.fillKey.Get() == key {
		return
	}

	c.disposeFill()

	if key.color == nil {
		return
	}

	img := ebiten.NewImage(1, 1)
	img.Fill(key.color)

	c.fillKey = basic.SetOpt(key)
	c.fill = img
}

func (c *colorAreaImpl[R]) drawFill(screen *ebiten.Image) {
	if !c.fillKey.IsSet() {
		return
	}

	r := c.region.Rect()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(r.Width(), r.Height())
	op.GeoM.Translate(r.Left(), r.Top())

	screen.DrawImage(c.fill, op)
}

func (c *colorAreaImpl[R]) disposeFill() {
	if !c.fillKey.IsSet() {
		return
	}

	c.fill.Deallocate()

	c.fillPart = fillPart{}
}
