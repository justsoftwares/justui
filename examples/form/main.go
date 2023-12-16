package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"github.com/justsoftwares/justui/widgets/form"
	"log"
)

func main() {
	w := app.NewWindow(
		app.Title("MyUI"),
		app.Size(unit.Dp(400), unit.Dp(400)),
		app.MinSize(unit.Dp(400), unit.Dp(400)),
	)
	t := material.NewTheme()
	nameEditor := &widget.Editor{}
	surnameEditor := &widget.Editor{}
	submitClickable := &widget.Clickable{}
	f := form.NewForm(t, "Registration", material.Button(t, submitClickable, "Submit").Layout)
	f.AddField("Name", material.Editor(t, nameEditor, "John").Layout)
	f.AddField("Surname", material.Editor(t, surnameEditor, "Smith").Layout)
	u := justui.NewUI(w, t, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(f.Layout),
		)
	})
	u.AddFrameEventHandlers(justui.EventHandler{
		Event: submitClickable.Clicked,
		Handler: func(gtx layout.Context, e event.Event) {
			log.Println("Name:", nameEditor.Text(), "Surname:", surnameEditor.Text())
		},
	})

	u.Run(true)
}
