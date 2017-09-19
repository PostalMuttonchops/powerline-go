package main

import (
	"fmt"
	"os"
	"path"
)

func segmentVirtualEnv(p *powerline) {
	var env string
	var root string
	if env == "" {
		env, _ = os.LookupEnv("VIRTUAL_ENV")
	}
	if env == "" {
		env, _ = os.LookupEnv("CONDA_ENV_PATH")
	}
	if env == "" {
		env, _ = os.LookupEnv("CONDA_DEFAULT_ENV")
	}
	root, _ = os.LookupEnv("CONDA_DEFAULT_ENV")
	if env == "" || root == "root" {
		return
	} else {
		envName := path.Base(env)
		p.appendSegment("venv", segment{
			content:    fmt.Sprintf(" %s ", envName),
			foreground: p.theme.VirtualEnvFg,
			background: p.theme.VirtualEnvBg,
		})
	}
}
