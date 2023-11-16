package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/justsoftwares/justui"
	"log"
	"os"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Just Converter"),
			app.Size(unit.Dp(400), unit.Dp(400)),
		)
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	u := justui.NewUI(window)
	var ()
	u.Layout = justui.NewLayout(layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceStart,
	})
	c := justui.NewLayout(layout.Flex{
		Axis:    layout.Horizontal,
		Spacing: layout.SpaceStart,
	})
	u.Layout.Add(layout.Rigid(c.Layout))
	return u.Run()
}
