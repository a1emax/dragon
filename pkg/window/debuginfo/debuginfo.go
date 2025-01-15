package debuginfo

import (
	"fmt"
	"runtime"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/scene/element/flexbox"
	"github.com/a1emax/youngine/scene/element/overlay"
	"github.com/hajimehoshi/ebiten/v2"

	"dragon/pkg/global/assets"
	"dragon/pkg/global/vars"
	"dragon/pkg/window/common"
	"dragon/pkg/window/common/colorarea"
	"dragon/pkg/window/common/label"
)

type DebugInfo[R scene.Region] interface {
	common.Element[R]
}

func New[R scene.Region](region R) DebugInfo[R] {
	var maxMemSys uint64
	var maxMemPauseNs uint64

	return overlay.New(region, overlay.Config{
		StateFunc: func(state overlay.State) overlay.State {
			state = overlay.State{}

			state.SetHeight(24.0)

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

				state.Color = assets.Colors.DebugInfoBackground

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

				state.Direction = flexbox.DirectionRow
				state.JustifyContent = flexbox.JustifySpaceBetween
				state.AlignItems = flexbox.AlignCenter

				return state
			},
		},
			label.New(flexbox.NewRegion(flexbox.RegionConfig{
				StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
					state = flexbox.RegionState{}

					return state
				},
			}), label.Config{
				StateFunc: func(state label.State) label.State {
					state = label.State{}

					state.SetWidth(150.0)
					state.Text = fmt.Sprintf("%.2f / %.2f", ebiten.ActualFPS(), ebiten.ActualTPS())
					state.TextFontFace = assets.FontFaces.DebugInfoText
					state.TextColor = assets.Colors.DebugInfoText

					return state
				},
			}),

			label.New(flexbox.NewRegion(flexbox.RegionConfig{
				StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
					state = flexbox.RegionState{}

					return state
				},
			}), label.Config{
				StateFunc: func(state label.State) label.State {
					state = label.State{}

					state.SetWidth(150.0)
					state.Text = fmt.Sprintf("%dx%d / %dx%d",
						vars.Ebiten.ScreenWidth, vars.Ebiten.ScreenHeight,
						vars.Ebiten.OutsideWidth, vars.Ebiten.OutsideHeight,
					)
					state.TextFontFace = assets.FontFaces.DebugInfoText
					state.TextColor = assets.Colors.DebugInfoText

					return state
				},
			}),

			label.New(flexbox.NewRegion(flexbox.RegionConfig{
				StateFunc: func(state flexbox.RegionState) flexbox.RegionState {
					state = flexbox.RegionState{}

					return state
				},
			}), label.Config{
				StateFunc: func(state label.State) label.State {
					state = label.State{}

					var mem runtime.MemStats
					runtime.ReadMemStats(&mem)

					if mem.Sys > maxMemSys {
						maxMemSys = mem.Sys
					}

					memPauseNs := mem.PauseNs[(mem.NumGC+255)%256]
					if memPauseNs > maxMemPauseNs {
						maxMemPauseNs = memPauseNs
					}

					state.SetWidth(150.0)
					state.Text = fmt.Sprintf("%.2f MiB (%.2f ms)",
						basic.Float(maxMemSys)/(1024*1024),
						basic.Float(maxMemPauseNs)/1_000_000,
					)
					state.TextFontFace = assets.FontFaces.DebugInfoText
					state.TextColor = assets.Colors.DebugInfoText

					return state
				},
			}),
		),
	)
}
