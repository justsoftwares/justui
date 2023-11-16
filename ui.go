package justui

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
)

type UI struct {
	Window *app.Window
	Theme  *material.Theme
	Layout *Layout
}

func NewUI(window *app.Window) *UI {
	return &UI{
		Window: window,
		Theme:  material.NewTheme(),
	}
}

func (u *UI) Run() error {
	return u.loop()
}

func (u *UI) layout(gtx layout.Context) {
	u.Layout.RootLayout.Layout(gtx, u.Layout.Children...)
}

func (u *UI) handlers(gtx layout.Context) {
	for _, eh := range u.Layout.EventHandlers {
		if eh.Event() {
			eh.Handler(gtx)
		}
	}
}

func (u *UI) loop() error {
	var ops op.Ops

	for e := range u.Window.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			u.layout(gtx)
			u.handlers(gtx)
			e.Frame(gtx.Ops)

		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
