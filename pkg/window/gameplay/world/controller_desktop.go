//go:build windows

package world

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/input/element/keyboardkey"

	"dragon/pkg/domain/space"
	"dragon/pkg/global/tools"
	"dragon/pkg/global/vars"
)

func (w *worldImpl[R]) initController() {
	w.controller = input.MultiController[basic.None]{
		keyboardkey.NewController(keyboardkey.ControllerConfig[basic.None]{
			Key:   tools.Input.Keyboard().Key(input.KeyboardKeyCodeArrowUp),
			Clock: tools.Clock,

			OnDown: func(event keyboardkey.DownEvent[basic.None]) {
				vars.GamePlay.Session.Dragon().SetDirection(space.North)
			},
		}),

		keyboardkey.NewController(keyboardkey.ControllerConfig[basic.None]{
			Key:   tools.Input.Keyboard().Key(input.KeyboardKeyCodeArrowDown),
			Clock: tools.Clock,

			OnDown: func(event keyboardkey.DownEvent[basic.None]) {
				vars.GamePlay.Session.Dragon().SetDirection(space.South)
			},
		}),

		keyboardkey.NewController(keyboardkey.ControllerConfig[basic.None]{
			Key:   tools.Input.Keyboard().Key(input.KeyboardKeyCodeArrowRight),
			Clock: tools.Clock,

			OnDown: func(event keyboardkey.DownEvent[basic.None]) {
				vars.GamePlay.Session.Dragon().SetDirection(space.East)
			},
		}),

		keyboardkey.NewController(keyboardkey.ControllerConfig[basic.None]{
			Key:   tools.Input.Keyboard().Key(input.KeyboardKeyCodeArrowLeft),
			Clock: tools.Clock,

			OnDown: func(event keyboardkey.DownEvent[basic.None]) {
				vars.GamePlay.Session.Dragon().SetDirection(space.West)
			},
		}),

		keyboardkey.NewController(keyboardkey.ControllerConfig[basic.None]{
			Key:   tools.Input.Keyboard().Key(input.KeyboardKeyCodeControlLeft),
			Clock: tools.Clock,

			OnPress: func(event keyboardkey.PressEvent[basic.None]) {
				vars.GamePlay.IgnoreDirectionPad = true
				vars.GamePlay.Session.Dragon().Move()
			},
			OnUp: func(event keyboardkey.UpEvent[basic.None]) {
				vars.GamePlay.IgnoreDirectionPad = false
				vars.GamePlay.Session.Dragon().Stop()
			},
			OnGone: func(event keyboardkey.GoneEvent) {
				if vars.GamePlay.Session == nil {
					return
				}

				vars.GamePlay.IgnoreDirectionPad = false
				vars.GamePlay.Session.Dragon().Stop()
			},
		}),

		keyboardkey.NewController(keyboardkey.ControllerConfig[basic.None]{
			Key:   tools.Input.Keyboard().Key(input.KeyboardKeyCodeEscape),
			Clock: tools.Clock,

			OnUp: func(event keyboardkey.UpEvent[basic.None]) {
				vars.Window.Page = vars.WindowPageMainMenu
			},
		}),
	}
}
