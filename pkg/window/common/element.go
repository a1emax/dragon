package common

import (
	"github.com/a1emax/youngine/scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type Element[R scene.Region] interface {
	scene.Element[*ebiten.Image, R]
}

type BaseElement[R scene.Region] struct {
	scene.BaseElement[*ebiten.Image, R]
}
