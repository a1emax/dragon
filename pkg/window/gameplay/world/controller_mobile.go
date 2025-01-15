//go:build android

package world

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
)

func (w *worldImpl[R]) initController() {
	w.controller = input.MultiController[basic.None]{}
}
