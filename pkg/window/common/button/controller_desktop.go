//go:build windows

package button

import (
	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/input"
	"github.com/a1emax/youngine/input/element/mousebutton"
	"github.com/a1emax/youngine/input/element/mousecursor"
)

func (b *buttonImpl[R]) initController(config Config) {
	b.controller = mousecursor.NewController(mousecursor.ControllerConfig[basic.None]{
		Cursor: config.Input.Mouse().Cursor(),

		HitTest: func(position basic.Vec2) bool {
			return b.region.Rect().Contains(position)
		},

		Slave: mousebutton.NewController(mousebutton.ControllerConfig[mousecursor.Background[basic.None]]{
			Button: config.Input.Mouse().Button(input.MouseButtonCodeLeft),
			Clock:  config.Clock,

			OnPress: func(event mousebutton.PressEvent[mousecursor.Background[basic.None]]) {
				b.isPressed = true

				if config.OnPress != nil {
					config.OnPress(PressEvent{
						Duration: event.Duration,
					})
				}
			},
			OnUp: func(event mousebutton.UpEvent[mousecursor.Background[basic.None]]) {
				b.isPressed = false

				if config.OnClick != nil {
					config.OnClick(ClickEvent{})
				}
			},
			OnGone: func(event mousebutton.GoneEvent) {
				b.isPressed = false
			},
		}),
	})
}
