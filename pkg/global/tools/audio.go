package tools

import (
	"github.com/a1emax/youngine/audio"
	"github.com/a1emax/youngine/audio/driver/ebitenaudio"
	"github.com/a1emax/youngine/x/scope"
)

var AudioFactory ebitenaudio.Factory

func initAudio(lc scope.Lifecycle) {
	AudioFactory = ebitenaudio.NewFactory(audio.SampleRate44100Hz)
}
