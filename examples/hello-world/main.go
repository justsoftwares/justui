package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
)

func main() {
	w := app.NewWindow(
		app.Title("MyUI"),
		app.Size(unit.Dp(400), unit.Dp(400)),
		app.MinSize(unit.Dp(400), unit.Dp(400)),
	)
	t := material.NewTheme()
	u := justui.NewUI(w, t, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.H1(t, "Hello, World!").Layout),
		)
	})
	u.Run(true)
}
