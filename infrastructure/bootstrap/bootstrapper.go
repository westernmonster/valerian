package bootstrap

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*gin.Engine
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Engine:       gin.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (p *Bootstrapper) SetupErrorHandlers() {
}

const (
	StaticAssets = "./public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

func (p *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(p)
	}
}

func (p *Bootstrapper) Bootstrap() *Bootstrapper {
	p.SetupErrorHandlers()

	p.Use(gin.Logger())
	p.Use(gin.Recovery())
	return p
}
