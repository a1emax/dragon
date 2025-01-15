package label

import (
	"image/color"

	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"dragon/pkg/window/common"
)

type Label[R scene.Region] interface {
	common.Element[R]
}

type Config struct {
	StateFunc func(state State) State
}

type State struct {
	IsInactive bool
	scene.Outline
	Text         string
	TextFontFace font.Face
	TextColor    color.Color
}

type labelImpl[R scene.Region] struct {
	scene.BaseElement[*ebiten.Image, R]
	textPart

	region    R
	state     State
	stateFunc func(state State) State
}

func New[R scene.Region](region R, config Config) Label[R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	return &labelImpl[R]{
		region:    region,
		stateFunc: config.StateFunc,
	}
}

func (l *labelImpl[R]) Region() R {
	return l.region
}

func (l *labelImpl[R]) IsActive() bool {
	return !l.state.IsInactive
}

func (l *labelImpl[R]) Outline() scene.Outline {
	return l.state.Outline
}

func (l *labelImpl[R]) Refresh() {
	l.state = l.stateFunc(l.state)
}

func (l *labelImpl[R]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	l.setupText()
	l.drawText(screen)
}

func (l *labelImpl[R]) Dispose() {
	l.disposeText()
}
