package justui

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	"log"
	"os"
)

type UI struct {
	Window             *app.Window
	Theme              *material.Theme
	Layout             func(u *UI, gtx layout.Context)
	FrameEventHandlers []EventHandler
	HandleOtherEvents  func(u *UI, e event.Event, ops *op.Ops)
}

func NewUI(window *app.Window, theme *material.Theme, layoutFunc func(u *UI, gtx layout.Context)) *UI {
	return &UI{
		Window: window,
		Theme:  theme,
		Layout: layoutFunc,
	}
}

func (u *UI) Run() {
	go func() {
		if err := u.loop(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func (u *UI) AddFrameEventHandlers(handler ...EventHandler) {
	u.FrameEventHandlers = append(u.FrameEventHandlers, handler...)
}

func (u *UI) HandleFrameEvents(gtx layout.Context, e event.Event) {
	for _, eh := range u.FrameEventHandlers {
		if eh.Event() {
			eh.Handler(u, gtx, e)
		}
	}
}

func (u *UI) loop() error {
	var ops op.Ops

	for e := range u.Window.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			u.HandleFrameEvents(gtx, e)
			u.Layout(u, gtx)
			e.Frame(gtx.Ops)
		}
		if u.HandleOtherEvents != nil {
			u.HandleOtherEvents(u, e, &ops)
		}
	}

	return nil
}
