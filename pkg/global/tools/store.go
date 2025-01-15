package tools

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/a1emax/youngine/store"
	"github.com/a1emax/youngine/store/host/file"
	"github.com/a1emax/youngine/x/scope"

	"dragon/pkg/global/vars"
)

const storeFileName = "dragon.yml"

type StoreData struct {
	MuteAudio bool `yaml:"mute_audio"`
}

var StoreBuffer store.Buffer[StoreData]

var StoreSyncer store.Syncer[StoreData]

func initStore(lc scope.Lifecycle) {
	locker := store.NewLocker[StoreData]()

	filePath := path.Join(vars.Extern.FilesDir, storeFileName)
	accessor := file.NewAccessor[StoreData](filePath)
	Logger.Debug("store file: " + filePath)

	syncer := store.NewSyncer(locker, accessor)

	err := syncer.Load(context.Background())
	if err != nil {
		Logger.Error(fmt.Sprintf("%+v", err))
	}

	buffer := store.NewBuffer(locker)

	buffer.Pull()

	stop := make(chan struct{})
	done := make(chan struct{})
	go func() {
		defer close(done)

		t := time.NewTicker(10 * time.Second)
		defer t.Stop()

		select {
		case <-t.C:
			err := syncer.Save(context.Background())
			if err != nil {
				Logger.Error(fmt.Sprintf("%+v", err))
			} else {
				Logger.Debug("store file is updated in background")
			}
		case <-stop:
			return
		}
	}()

	lc.Defer(func() {
		close(stop)
		<-done

		err := syncer.Save(context.Background())
		if err != nil {
			Logger.Error(fmt.Sprintf("%+v", err))
		} else {
			Logger.Debug("store file is updated on exit")
		}
	})

	StoreBuffer = buffer
	StoreSyncer = syncer
}
