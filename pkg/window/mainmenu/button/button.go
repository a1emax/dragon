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
			state.PrimaryColor = assets.Colors.MainMenuButtonPrimary
			state.PressedColor = assets.Colors.MainMenuButtonPressed
			state.Text = assets.Texts.MainMenuButton
			state.TextFontFace = assets.FontFaces.MainMenuButtonText
			state.TextPrimaryColor = assets.Colors.MainMenuButtonText
			state.TextPressedColor = assets.Colors.MainMenuButtonText

			return state
		},

		Clock: tools.Clock,
		Input: tools.Input,

		OnClick: func(event button.ClickEvent) {
			vars.Window.Page = vars.WindowPageGamePlay
		},
	})
}
