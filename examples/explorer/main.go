package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"github.com/justsoftwares/justui/widgets/explorer"
	"log"
)

func main() {
	w := app.NewWindow(
		app.Title("MyUI"),
		app.Size(unit.Dp(400), unit.Dp(400)),
		app.MinSize(unit.Dp(400), unit.Dp(400)),
	)
	t := material.NewTheme()
	var exp *explorer.Explorer
	exp = explorer.NewExplorer(t, func(_ layout.Context, _ event.Event) {
		log.Println(exp.Files)
	})
	showExpBtn := &widget.Clickable{}
	u := justui.NewUI(w, t, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.Button(t, showExpBtn, "Select Files").Layout),
		)
	})
	u.AddFrameEventHandlers(justui.EventHandler{
		Event: showExpBtn.Clicked,
		Handler: func(gtx layout.Context, e event.Event) {
			exp.Run()
		},
	})
	u.Run(true)
}
