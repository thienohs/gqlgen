package api

import (
	"github.com/thienohs/gqlgen/codegen/config"
	"github.com/thienohs/gqlgen/plugin"
)

type Option func(cfg *config.Config, plugins *[]plugin.Plugin)

func NoPlugins() Option {
	return func(cfg *config.Config, plugins *[]plugin.Plugin) {
		*plugins = nil
	}
}

func AddPlugin(p plugin.Plugin) Option {
	return func(cfg *config.Config, plugins *[]plugin.Plugin) {
		*plugins = append(*plugins, p)
	}
}
