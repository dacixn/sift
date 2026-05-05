package main

import (
	"github.com/alecthomas/kong"
	"github.com/dacixn/sift/internal/config"
)

var cli struct {
	Path   string `type:"existingdir" arg:""`
	Dry    bool
	Config string `type:"existingdir"`
}

func main() {
	var cfg config.Config
	cfg.Init()

	ctx := kong.Parse(&cli)
	switch ctx.Command() {
	case "rm <path>":
	case "ls":
	default:
		panic(ctx.Command())
	}
}
