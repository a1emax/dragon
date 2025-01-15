package assets

import (
	"context"

	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/fault"
	"github.com/a1emax/youngine/x/scope"

	"dragon/pkg/global/tools"
)

func Init(lc scope.Lifecycle) {
	if lc == nil {
		panic(fault.Trace(fault.ErrNilPointer))
	}

	initColors(lc)
	initFontFaces(lc)
	initTexts(lc)
}

func load[T any](lc scope.Lifecycle, uri string) T {
	typedAsset, unload := asset.MustLoad[T](context.Background(), tools.AssetLoader, uri)
	lc.Defer(unload)

	return typedAsset
}
