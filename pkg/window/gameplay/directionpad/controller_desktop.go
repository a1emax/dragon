//go:build windows

package directionpad

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/input/element/mousebutton"
	"github.com/a1emax/youngine/input/element/mousecursor"

	"dragon/pkg/global/tools"
)

func (d *directionPadImpl[R]) initController() {
	d.controller = input.MultiController[basic.None]{
		mousecursor.NewController(mousecursor.ControllerConfig[basic.None]{
			Cursor: tools.Input.Mouse().Cursor(),

			HitTest: func(position basic.Vec2) bool {
				return d.region.Rect().Contains(position)
			},

			Slave: mousebutton.NewController(mousebutton.ControllerConfig[mousecursor.Background[basic.None]]{
				Button: tools.Input.Mouse().Button(input.MouseButtonCodeLeft),
				Clock:  tools.Clock,

				OnPress: func(event mousebutton.PressEvent[mousecursor.Background[basic.None]]) {
					d.handlePress(event.Background.Position)
				},
				OnUp: func(event mousebutton.UpEvent[mousecursor.Background[basic.None]]) {
					d.handleRelease()
				},
				OnGone: func(event mousebutton.GoneEvent) {
					d.handleRelease()
				},
			}),
		}),
	}
}
