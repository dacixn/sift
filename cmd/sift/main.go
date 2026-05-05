package main

import (
	"github.com/alecthomas/kong"
	"github.com/dacixn/sift/internal/config"
	"github.com/dacixn/sift/internal/filesort"
)

var cli struct {
	Path   string `type:"existingdir" arg:""`
	Dry    bool
	Config string `type:"existingdir"`
}

func main() {
	var cfg config.Config
	cfg.Init()
	kong.Parse(&cli)

	filesort.SortFiles(cfg.Groups, cli.Path)
}
