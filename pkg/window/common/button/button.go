package button

import (
	"image/color"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/clock"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"dragon/pkg/window/common"
)

type Button[R scene.Region] interface {
	common.Element[R]
}

type Config struct {
	StateFunc func(state State) State
	Clock     clock.Clock
	Input     input.Input
	OnPress   func(event PressEvent)
	OnClick   func(event ClickEvent)
}

type State struct {
	IsInactive bool
	scene.Outline
	CornerRadius     basic.Opt[basic.Float]
	PrimaryColor     color.Color
	PressedColor     color.Color
	Text             string
	TextFontFace     font.Face
	TextPrimaryColor color.Color
	TextPressedColor color.Color
}

type buttonImpl[R scene.Region] struct {
	scene.BaseElement[*ebiten.Image, R]
	shapePart
	textPart

	region     R
	state      State
	stateFunc  func(state State) State
	controller input.Controller[basic.None]

	isPressed bool
}

func New[R scene.Region](region R, config Config) Button[R] {
	if config.StateFunc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Input == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}
	if config.Clock == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b := &buttonImpl[R]{}

	b.region = region
	b.stateFunc = config.StateFunc
	b.initController(config)

	return b
}

func (b *buttonImpl[R]) Region() R {
	return b.region
}

func (b *buttonImpl[R]) IsActive() bool {
	return !b.state.IsInactive
}

func (b *buttonImpl[R]) Outline() scene.Outline {
	return b.state.Outline
}

func (b *buttonImpl[R]) Refresh() {
	b.state = b.stateFunc(b.state)
}

func (b *buttonImpl[R]) Actuate() {
	b.controller.Actuate(basic.None{})
}

func (b *buttonImpl[R]) Inhibit() {
	b.controller.Inhibit()
}

func (b *buttonImpl[R]) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	b.setupShape()
	b.drawShape(screen)

	b.setupText()
	b.drawText(screen)
}

func (b *buttonImpl[R]) Dispose() {
	b.disposeShape()
	b.disposeText()
}
