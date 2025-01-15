package intern

import (
	"context"
	"fmt"

	"github.com/a1emax/youngine/fault"
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"dragon/pkg/global/tools"
	"dragon/pkg/global/vars"
	"dragon/pkg/kernel"
)

func init() {
	mobile.SetGame(kernel.EbitenGame())
}

func SetFilesDir(value string) error {
	vars.Extern.FilesDir = value

	return nil
}

func Activate() error {
	return fault.Recover(kernel.Activate)
}

func Suspend() error {
	return fault.Recover(func() {
		kernel.IfRunning(func() {
			err := tools.StoreSyncer.Save(context.Background())
			if err != nil {
				tools.Logger.Error(fmt.Sprintf("%+v", err))
			} else {
				tools.Logger.Debug("store file is updated on suspending")
			}
		})
	})
}

func Resume() error {
	return nil
}

// TODO: Call kernel.Close on exit.
