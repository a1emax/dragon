//go:build android

package directionpad

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/input/element/touchscreentouch"

	"dragon/pkg/global/tools"
)

func (d *directionPadImpl[R]) initController() {
	d.controller = input.MultiController[basic.None]{
		touchscreentouch.NewController(touchscreentouch.ControllerConfig[basic.None]{
			Touchscreen: tools.Input.Touchscreen(),
			Clock:       tools.Clock,

			HitTest: func(position basic.Vec2) bool {
				return d.region.Rect().Contains(position)
			},

			OnHover: func(event touchscreentouch.HoverEvent[basic.None]) {
				d.handlePress(event.Position)
			},
			OnEnd: func(event touchscreentouch.EndEvent[basic.None]) {
				d.handleRelease()
			},
			OnGone: func(event touchscreentouch.GoneEvent) {
				d.handleRelease()
			},
		}),
	}
}
