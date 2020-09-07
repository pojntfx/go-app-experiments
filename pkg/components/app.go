package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type App struct {
	app.Compo
	Count int
}

func (c *App) Render() app.UI {
	return app.Main().Body(
		app.Div().Class("container").Body(
			app.H1().Body(
				app.Text("Counter"),
			),
			app.Div().Body(app.Text(fmt.Sprintf("Current count: %v", c.Count))),
			app.Div().Class("btn-group").Body(
				app.Button().Class("btn").Body(app.Text("Decrease")).OnClick(func(ctx app.Context, e app.Event) { c.Count--; c.Update() }),
				app.Button().Class("btn").Body(app.Text("Increase")).OnClick(func(ctx app.Context, e app.Event) { c.Count++; c.Update() }),
			),
		),
	)
}
