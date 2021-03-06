package linter

import (
	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/lint/astcache"
	"github.com/golangci/golangci-lint/pkg/logutils"
	"github.com/golangci/golangci-lint/pkg/packages"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
)

type Context struct {
	PkgProgram           *packages.Program
	Cfg                  *config.Config
	Program              *loader.Program
	SSAProgram           *ssa.Program
	LoaderConfig         *loader.Config
	ASTCache             *astcache.Cache
	NotCompilingPackages []*loader.PackageInfo
	Log                  logutils.Log
}

func (c *Context) Settings() *config.LintersSettings {
	return &c.Cfg.LintersSettings
}
