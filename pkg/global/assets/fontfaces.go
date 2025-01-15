package assets

import (
	"github.com/a1emax/youngine/asset/format/sfnt"
	"github.com/a1emax/youngine/x/scope"
)

var FontFaces struct {
	DebugInfoText      sfnt.FaceAsset
	GameOverButtonText sfnt.FaceAsset
	MainMenuButtonText sfnt.FaceAsset
}

func initFontFaces(lc scope.Lifecycle) {
	FontFaces.DebugInfoText = load[sfnt.FaceAsset](lc, "font-faces/debug-info-text.sff")
	FontFaces.GameOverButtonText = load[sfnt.FaceAsset](lc, "font-faces/game-over-button-text.sff")
	FontFaces.MainMenuButtonText = load[sfnt.FaceAsset](lc, "font-faces/main-menu-button-text.sff")
}
