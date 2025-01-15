package button

import (
	"github.com/a1emax/youngine/clock"
)

type PressEvent struct {
	Duration clock.Ticks
}

type ClickEvent struct {
}
