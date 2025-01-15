package kernel

import (
	"sync"
	"sync/atomic"

	"github.com/a1emax/youngine/basic"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/scene"
	"github.com/a1emax/youngine/x/scope"
	"github.com/hajimehoshi/ebiten/v2"

	"dragon/pkg/global/assets"
	"dragon/pkg/global/tools"
	"dragon/pkg/global/vars"
	"dragon/pkg/window"
)

var k struct {
	status   atomic.Int64
	closeMu  sync.RWMutex
	initErr  error
	drawErr  error
	teardown scope.TeardownFunc
	window   window.Window[scene.RegionFunc]
}

const (
	statusInitial int64 = iota
	statusActivated
	statusRunning
	statusClosed
)

func Activate() {
	if !k.status.CompareAndSwap(statusInitial, statusActivated) {
		panic(fault.Trace(fault.ErrInvalidUse))
	}
}

func IfRunning(f func()) {
	if f == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	k.closeMu.RLock()
	defer k.closeMu.RUnlock()

	if k.status.Load() != statusRunning {
		return
	}

	f()
}

type ebitenGame struct{}

func EbitenGame() ebiten.Game {
	return ebitenGame{}
}

func (ebitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	vars.Ebiten.OutsideWidth = outsideWidth
	vars.Ebiten.OutsideHeight = outsideHeight

	x, y := 720.0, 1280.0
	if r := float64(outsideHeight) / float64(outsideWidth); r > y/x {
		y = x * r
	}

	vars.Ebiten.ScreenWidth = int(x)
	vars.Ebiten.ScreenHeight = int(y)

	return vars.Ebiten.ScreenWidth, vars.Ebiten.ScreenHeight
}

func (ebitenGame) Update() error {
	if k.initErr != nil {
		return k.initErr
	}
	if k.drawErr != nil {
		return k.drawErr
	}

	k.closeMu.RLock()
	defer k.closeMu.RUnlock()

	switch k.status.Load() {

	case statusActivated:
		k.initErr = fault.Recover(func() {
			k.teardown = scope.MustSetup(func(lc scope.Lifecycle) {
				tools.Init(lc)
				assets.Init(lc)

				k.window = window.New(scene.RegionFunc(func() basic.Rect {
					return basic.Rect{
						Size: basic.Vec2{
							basic.Float(vars.Ebiten.ScreenWidth),
							basic.Float(vars.Ebiten.ScreenHeight),
						},
					}
				}))
				lc.Defer(func() {
					scene.Dispose(k.window)
				})
			})
		})
		if k.initErr != nil {
			return k.initErr
		}

		k.status.Store(statusRunning)

		fallthrough

	case statusRunning:
		if vars.Kernel.IsTerminated {
			return ebiten.Termination
		}

		return fault.Recover(func() {
			tools.Clock.Update()
			tools.Input.Update()

			scene.Update(k.window)
		})

	default:
		return nil
	}
}

func (ebitenGame) Draw(screen *ebiten.Image) {
	if screen == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	IfRunning(func() {
		k.drawErr = fault.Recover(func() {
			scene.Draw(k.window, screen)
		})
	})
}

func Close() error {
	k.closeMu.Lock()
	defer k.closeMu.Unlock()

	wasNotRunning := k.status.Load() != statusRunning
	k.status.Store(statusClosed)

	if wasNotRunning {
		return nil
	}

	return fault.Recover(func() {
		k.teardown()
		k.teardown = nil

		k.window = nil
	})
}
