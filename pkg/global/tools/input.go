package tools

import (
	"github.com/a1emax/youngine/input/driver/ebiteninput"
	"github.com/a1emax/youngine/x/scope"
)

var Input ebiteninput.Input

func initInput(lc scope.Lifecycle) {
	Input = ebiteninput.New(Clock)
}
