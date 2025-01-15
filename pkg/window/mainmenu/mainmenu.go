package mainmenu

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/overlay"

	"dragon/pkg/global/assets"
	"dragon/pkg/window/common"
	"dragon/pkg/window/common/colorarea"
	"dragon/pkg/window/mainmenu/button"
)

type MainMenu[R scene.Region] interface {
	common.Element[R]
}

func New[R scene.Region](region R) MainMenu[R] {
	return overlay.New(region, overlay.Config{
		StateFunc: func(state overlay.State) overlay.State {
			state = overlay.State{}

			return state
		},
	},
		colorarea.New(overlay.NewRegion(overlay.RegionConfig{
			StateFunc: func(state overlay.RegionState) overlay.RegionState {
				state = overlay.RegionState{}

				return state
			},
		}), colorarea.Config{
			StateFunc: func(state colorarea.State) colorarea.State {
				state = colorarea.State{}

				state.Color = assets.Colors.MainMenuBackground

				return state
			},
		}),

		button.New(overlay.NewRegion(overlay.RegionConfig{
			StateFunc: func(state overlay.RegionState) overlay.RegionState {
				state = overlay.RegionState{}

				return state
			},
		})),
	)
}
