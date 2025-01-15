package tools

import (
	"github.com/a1emax/youngine/x/scope"
)

func Init(lc scope.Lifecycle) {
	initLogger(lc)
	initRandom(lc)
	initClock(lc)
	initInput(lc)
	initAudio(lc)
	initAsset(lc)
	initStore(lc)
}
