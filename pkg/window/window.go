package window

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/pageset"

	"dragon/pkg/global/vars"
	"dragon/pkg/window/common"
	"dragon/pkg/window/debuginfo"
	"dragon/pkg/window/gameover"
	"dragon/pkg/window/gameplay"
	"dragon/pkg/window/mainmenu"
)

type Window[R scene.Region] interface {
	common.Element[R]
}

func New[R scene.Region](region R) Window[R] {
	return flexbox.New(region, flexbox.Config{
		StateFunc: func(state flexbox.State) flexbox.State {
			state = flexbox.State{}

			state.Direction = flexbox.DirectionColumn
			state.JustifyContent = flexbox.JustifyCenter
			state.AlignItems = flexbox.AlignCenter

			return state
		},
	},
		debuginfo.New(flexbox.NewRegion(flexbox.RegionConfig{
			StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
				state = flexbox.RegionState{}

				return state
			},
		})),

		pageset.New(flexbox.NewRegion(flexbox.RegionConfig{
			StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
				state = flexbox.RegionState{}

				return state
			},
		}), pageset.Config{
			StateFunc: func(state pageset.State) pageset.State {
				state = pageset.State{}

				state.Page = vars.Window.Page

				return state
			},
		},
			mainmenu.New(pageset.NewRegion()),
			gameplay.New(pageset.NewRegion()),
			gameover.New(pageset.NewRegion()),
		),
	)
}
