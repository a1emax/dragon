package colorarea

import (
	"image/color"

	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
	"github.com/hajimehoshi/ebiten/v2"

	"dragon/pkg/window/common"
)

type ColorArea[R scene.Region] interface {
	common.Element[R]
}

type Config struct {
	StateFunc func(state State) State
}

type State struct {
	IsInactive bool
	scene.Outline
	Color color.Color
}

type colorAreaImpl[R scene.Region] struct {
	scene.BaseElement[*ebiten.Image, R]
	fillPart

	region    R
	state     State
	stateFunc func(state State) State
}

func New[R scene.Region](region R, config Config) ColorArea[R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &colorAreaImpl[R]{
		region:    region,
		stateFunc: config.StateFunc,
	}
}

func (c *colorAreaImpl[R]) Region() R {
	return c.region
}

func (c *colorAreaImpl[R]) IsActive() bool {
	return !c.state.IsInactive
}

func (c *colorAreaImpl[R]) Outline() scene.Outline {
	return c.state.Outline
}

func (c *colorAreaImpl[R]) Refresh() {
	c.state = c.stateFunc(c.state)
}

func (c *colorAreaImpl[R]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	c.setupFill()
	c.drawFill(screen)
}

func (c *colorAreaImpl[R]) Dispose() {
	c.disposeFill()
}
