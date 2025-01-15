package world

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/scene"

	"dragon/pkg/domain"
	"dragon/pkg/global/tools"
	"dragon/pkg/global/vars"
	"dragon/pkg/window/common"
)

type World[R scene.Region] interface {
	common.Element[R]
}

type worldImpl[R scene.Region] struct {
	common.BaseElement[R]

	region     R
	controller input.Controller[basic.None]

	isPrepared bool
}

func New[R scene.Region](region R) World[R] {
	w := &worldImpl[R]{}

	w.region = region
	w.initController()

	return w
}

func (w *worldImpl[R]) Region() R {
	return w.region
}

func (w *worldImpl[R]) Prepare() {
	if w.isPrepared {
		return
	}

	w.isPrepared = true
	vars.GamePlay.Session = domain.NewSession(domain.SessionConfig{
		Random:        tools.Random,
		Clock:         tools.Clock,
		OnKnightEnter: w.onKnightEnter,
		OnKnightLoad:  w.onKnightLoad,
		OnKnightKill:  w.onKnightKill,
		OnKnightExit:  w.onKnightExit,
		OnDragonChop:  w.onDragonChop,
		OnGameOver:    w.onGameOver,
	})
}

func (w *worldImpl[R]) Exclude() {
	w.cleanup()
}

func (w *worldImpl[R]) Actuate() {
	w.controller.Actuate(basic.None{})
}

func (w *worldImpl[R]) Inhibit() {
	w.controller.Inhibit()
}

func (w *worldImpl[R]) Update() {
	vars.GamePlay.Session.Update()
}

func (w *worldImpl[R]) Dispose() {
	w.cleanup()
}

func (w *worldImpl[R]) cleanup() {
	vars.GamePlay.Session = nil
	w.isPrepared = false
}
