package button

import (
	"github.com/a1emax/youngine/scene"

	"dragon/pkg/global/assets"
	"dragon/pkg/global/tools"
	"dragon/pkg/global/vars"
	"dragon/pkg/window/common"
	"dragon/pkg/window/common/button"
)

type Button[R scene.Region] interface {
	common.Element[R]
}

func New[R scene.Region](region R) Button[R] {
	return button.New(region, button.Config{
		StateFunc: func(state button.State) button.State {
			state = button.State{}

			state.SetWidth(420.0)
			state.SetHeight(120.0)
			state.PrimaryColor = assets.Colors.GameOverButtonPrimary
			state.PressedColor = assets.Colors.GameOverButtonPressed
			state.Text = assets.Texts.GameOverButton
			state.TextFontFace = assets.FontFaces.GameOverButtonText
			state.TextPrimaryColor = assets.Colors.GameOverButtonText
			state.TextPressedColor = assets.Colors.GameOverButtonText

			return state
		},

		Clock: tools.Clock,
		Input: tools.Input,

		OnClick: func(event button.ClickEvent) {
			vars.Window.Page = vars.WindowPageMainMenu
		},
	})
}
