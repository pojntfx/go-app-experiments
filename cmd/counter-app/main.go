package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/pojntfx/go-app-experiments/pkg/components"
)

func main() {
	app.Route("/", &components.App{Count: 0})

	app.Run()
}
