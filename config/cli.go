package config

import "github.com/alecthomas/kong"

type CLI struct {
	ConfigFile string `type:"existingfile" default:"vanity.json" help:"Path to the configuration file"`
	OutputDir  string `type:"existingdir" default:"." help:"Path to the output directory"`
}

func ParseCLIOpts() *CLI {
	cli := CLI{}
	kong.Parse(&cli)
	return &cli
}
