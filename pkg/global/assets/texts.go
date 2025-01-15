package assets

import (
	"github.com/a1emax/youngine/asset/format/text"
	"github.com/a1emax/youngine/x/scope"
)

var Texts struct {
	GameOverButton text.Asset
	MainMenuButton text.Asset
}

func initTexts(lc scope.Lifecycle) {
	Texts.GameOverButton = load[text.Asset](lc, "texts/game-over-button.txt")
	Texts.MainMenuButton = load[text.Asset](lc, "texts/main-menu-button.txt")
}
