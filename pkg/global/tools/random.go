package tools

import (
	"math/rand"
	"time"

	"github.com/a1emax/youngine/x/scope"
)

var Random *rand.Rand

func initRandom(lc scope.Lifecycle) {
	Random = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
}
