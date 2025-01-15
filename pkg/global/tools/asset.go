package tools

import (
	"github.com/a1emax/youngine/asset"
	"github.com/a1emax/youngine/asset/format/image"
	"github.com/a1emax/youngine/asset/format/kage"
	"github.com/a1emax/youngine/asset/format/mp3"
	"github.com/a1emax/youngine/asset/format/rgba"
	"github.com/a1emax/youngine/asset/format/sfnt"
	"github.com/a1emax/youngine/asset/format/text"
	"github.com/a1emax/youngine/asset/format/wav"
	"github.com/a1emax/youngine/asset/host/filesystem"
	"github.com/a1emax/youngine/x/scope"

	"dragon/res"
)

var AssetMapper asset.Mapper

var AssetLoader asset.Loader

func initAsset(lc scope.Lifecycle) {
	mapper := asset.NewMapper()
	loader := asset.NewLoader(mapper)

	fetcher := filesystem.NewFetcher(res.FS)

	asset.Map[image.Asset](mapper, image.NewProvider(fetcher))
	asset.Map[kage.Asset](mapper, kage.NewProvider(fetcher))
	asset.Map[mp3.Asset](mapper, mp3.NewProvider(fetcher, 0))
	asset.Map[rgba.Asset](mapper, rgba.NewProvider(fetcher))
	asset.Map[sfnt.Asset](mapper, sfnt.NewProvider(fetcher))
	asset.Map[sfnt.FaceAsset](mapper, sfnt.NewFaceProvider(fetcher, loader))
	asset.Map[text.Asset](mapper, text.NewProvider(fetcher))
	asset.Map[wav.Asset](mapper, wav.NewProvider(fetcher, 0))

	AssetMapper = mapper
	AssetLoader = loader
}
