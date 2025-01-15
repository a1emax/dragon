//go:build android

package button

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input/element/touchscreentouch"
)

func (b *buttonImpl[R]) initController(config Config) {
	b.controller = touchscreentouch.NewController(touchscreentouch.ControllerConfig[basic.None]{
		Touchscreen: config.Input.Touchscreen(),
		Clock:       config.Clock,

		HitTest: func(position basic.Vec2) bool {
			return b.region.Rect().Contains(position)
		},

		OnHover: func(event touchscreentouch.HoverEvent[basic.None]) {
			b.isPressed = true

			if config.OnPress != nil {
				config.OnPress(PressEvent{
					Duration: event.Duration,
				})
			}
		},
		OnEnd: func(event touchscreentouch.EndEvent[basic.None]) {
			b.isPressed = false

			if config.OnClick != nil {
				config.OnClick(ClickEvent{})
			}
		},
		OnGone: func(event touchscreentouch.GoneEvent) {
			b.isPressed = false
		},
	})
}
