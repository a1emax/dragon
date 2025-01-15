package directionpad

import (
	"math"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/scene"
	"github.com/hajimehoshi/ebiten/v2"

	"dragon/pkg/domain/space"
	"dragon/pkg/global/assets"
	"dragon/pkg/global/vars"
	"dragon/pkg/window/common"
)

type DirectionPad[R scene.Region] interface {
	common.Element[R]
}

type directionPadImpl[R scene.Region] struct {
	common.BaseElement[R]

	region     R
	controller input.Controller[basic.None]
}

func New[R scene.Region](region R) DirectionPad[R] {
	d := &directionPadImpl[R]{}

	d.region = region
	d.initController()

	return d
}

func (d *directionPadImpl[R]) Region() R {
	return d.region
}

func (d *directionPadImpl[R]) Outline() scene.Outline {
	var outline scene.Outline

	outline.SetWidth(400.0)
	outline.SetHeight(400.0)

	return outline
}

func (d *directionPadImpl[R]) Actuate() {
	d.controller.Actuate(basic.None{})
}

func (d *directionPadImpl[R]) Inhibit() {
	d.controller.Inhibit()
}

func (d *directionPadImpl[R]) Draw(screen *ebiten.Image) {
	dst := screen.SubImage(d.region.Rect().Image()).(*ebiten.Image)
	dst.Fill(assets.Colors.GamePlayDirectionPad)
}

func (d *directionPadImpl[R]) handlePress(position basic.Vec2) {
	if vars.GamePlay.IgnoreDirectionPad {
		return
	}

	r := d.region.Rect()
	c := r.Min.Add(r.Size.DivAll(2.0))
	p := position
	a := math.Abs(math.Atan2(c.Y()-p.Y(), p.X()-c.X()))

	var direction space.Direction
	switch {
	case a < 0.25*math.Pi:
		direction = space.East
	case a > 0.75*math.Pi:
		direction = space.West
	case p.Y() < c.Y():
		direction = space.North
	default:
		direction = space.South
	}

	dragon := vars.GamePlay.Session.Dragon()
	dragon.SetDirection(direction)
	dragon.Move()
}

func (d *directionPadImpl[R]) handleRelease() {
	if vars.GamePlay.IgnoreDirectionPad || vars.GamePlay.Session == nil {
		return
	}

	vars.GamePlay.Session.Dragon().Stop()
}
