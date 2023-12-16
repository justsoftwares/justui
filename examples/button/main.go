package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"log"
)

func main() {
	w := app.NewWindow(
		app.Title("MyUI"),
		app.Size(unit.Dp(400), unit.Dp(400)),
		app.MinSize(unit.Dp(400), unit.Dp(400)),
	)
	t := material.NewTheme()
	clickMeBtn := &widget.Clickable{}
	u := justui.NewUI(w, t, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.H1(t, "Hello, World!").Layout),
			layout.Rigid(material.Button(t, clickMeBtn, "Click me!").Layout),
		)
	})

	u.AddFrameEventHandlers(justui.EventHandler{
		Event: clickMeBtn.Clicked,
		Handler: func(gtx layout.Context, e event.Event) {
			log.Println("I was clicked!")
		},
	})

	u.Run(true)
}
