package button

import (
	"image/color"
)

func (b *buttonImpl[R]) color(primary, pressed color.Color) color.Color {
	if b.isPressed {
		return pressed
	}

	return primary
}
