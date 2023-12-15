package main

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"github.com/justsoftwares/justui/widgets/explorer"
	"log"
	"strconv"
	"strings"
)

func main() {
	w := app.NewWindow(
		app.Title("JustUI"),
		app.Size(unit.Dp(400), unit.Dp(400)),
		app.MinSize(unit.Dp(400), unit.Dp(400)),
	)
	t := material.NewTheme()
	counter := &widget.Float{}
	cb1 := &widget.Bool{}
	btn2 := &widget.Clickable{}
	edt1 := &widget.Editor{}
	var exp *explorer.Explorer
	exp = explorer.NewExplorer(t, func(_ layout.Context, _ event.Event) {
		log.Println(exp.Files)
	})
	showExpBtn := &widget.Clickable{}
	u := justui.NewUI(w, t, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.CheckBox(t, cb1, "track keys").Layout),
			layout.Rigid(material.Button(t, btn2, "+").Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				e := material.Editor(t, edt1, "")
				e.TextSize = 100
				return layout.Center.Layout(gtx, e.Layout)
			}),
			layout.Rigid(material.Slider(t, counter).Layout),
			layout.Rigid(material.Button(t, showExpBtn, "Select Files").Layout),
		)
	})
	u.AddFrameEventHandlers(justui.EventHandler{
		Event: cb1.Update,
		Handler: func(gtx layout.Context, e event.Event) {
			cb1.Value = !cb1.Value
		},
	}, justui.EventHandler{
		Event: func(gtx layout.Context) bool {
			return cb1.Value
		},
		Handler: func(gtx layout.Context, e event.Event) {
			key.InputOp{Tag: &cb1.Value}.Add(gtx.Ops)
			for _, evt := range gtx.Events(&cb1.Value) {
				if x, ok := evt.(key.Event); ok {
					log.Println(x.Name)
				}
			}
		},
	}, justui.EventHandler{
		Event: btn2.Clicked,
		Handler: func(gtx layout.Context, e event.Event) {
			val, _ := strconv.ParseFloat(strings.TrimSpace(edt1.Text()), 32)
			counter.Value = float32(val)
			counter.Value++
			edt1.SetText(strconv.FormatFloat(float64(counter.Value), 'f', -1, 32))
		},
	}, justui.EventHandler{
		Event: counter.Update,
		Handler: func(gtx layout.Context, e event.Event) {
			edt1.SetText(strconv.FormatFloat(float64(counter.Value), 'f', -1, 32))
		},
	}, justui.EventHandler{
		Event: showExpBtn.Clicked,
		Handler: func(gtx layout.Context, e event.Event) {
			exp.Run()
		},
	})

	u.Run(true)
}
