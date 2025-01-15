package tools

import (
	"github.com/a1emax/youngine/clock/driver/ebitenclock"
	"github.com/a1emax/youngine/x/scope"
)

var Clock ebitenclock.Clock

func initClock(lc scope.Lifecycle) {
	Clock = ebitenclock.New()
}
