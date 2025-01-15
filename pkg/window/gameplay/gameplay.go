package gameplay

import (
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/nothing"
	"github.com/a1emax/youngine/scene/element/overlay"
	"github.com/hajimehoshi/ebiten/v2"

	"dragon/pkg/global/assets"
	"dragon/pkg/window/common"
	"dragon/pkg/window/common/colorarea"
	"dragon/pkg/window/gameplay/directionpad"
	"dragon/pkg/window/gameplay/world"
)

type GamePlay[R scene.Region] interface {
	common.Element[R]
}

func New[R scene.Region](region R) GamePlay[R] {
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

				state.Color = assets.Colors.GamePlayBackground

				return state
			},
		}),

		flexbox.New(overlay.NewRegion(overlay.RegionConfig{
			StateFunc: func(state overlay.RegionState) overlay.RegionState {
				state = overlay.RegionState{}

				return state
			},
		}), flexbox.Config{
			StateFunc: func(state flexbox.State) flexbox.State {
				state = flexbox.State{}

				state.Direction = flexbox.DirectionColumn
				state.JustifyContent = flexbox.JustifySpaceBetween
				state.AlignItems = flexbox.AlignCenter

				return state
			},
		},
			nothing.New[*ebiten.Image](flexbox.NewRegion(flexbox.RegionConfig{
				StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
					state = flexbox.RegionState{}

					return state
				},
			}), nothing.Config{
				StateFunc: func(state nothing.State) nothing.State {
					state = nothing.State{}

					state.SetHeight(100.0)

					return state
				},
			}),

			world.New(flexbox.NewRegion(flexbox.RegionConfig{
				StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
					state = flexbox.RegionState{}

					return state
				},
			})),

			flexbox.New(flexbox.NewRegion(flexbox.RegionConfig{
				StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
					state = flexbox.RegionState{}

					return state
				},
			}), flexbox.Config{
				StateFunc: func(state flexbox.State) flexbox.State {
					state = flexbox.State{}

					state.SetHeight(500.0)
					state.Direction = flexbox.DirectionRow
					state.JustifyContent = flexbox.JustifyCenter
					state.AlignItems = flexbox.AlignCenter

					return state
				},
			},
				directionpad.New(flexbox.NewRegion(flexbox.RegionConfig{
					StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
						state = flexbox.RegionState{}

						return state
					},
				})),
			),
		),
	)
}
