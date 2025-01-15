package res

import (
	"embed"
)

//go:embed colors/*
//go:embed font-faces/*
//go:embed fonts/*
//go:embed texts/*
var FS embed.FS
