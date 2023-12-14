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
	Layout             func(gtx layout.Context)
	FrameEventHandlers []EventHandler
	HandleOtherEvents  func(e event.Event, ops *op.Ops)
}

func NewUI(window *app.Window, theme *material.Theme, layoutFunc func(gtx layout.Context)) *UI {
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

func (u *UI) RemoveFrameEventHandlers(handler ...EventHandler) {
	for _, eventHandler := range handler {
		for i2, frameEventHandler := range u.FrameEventHandlers {
			if &eventHandler == &frameEventHandler {
				u.FrameEventHandlers = append(u.FrameEventHandlers[:i2], u.FrameEventHandlers[i2+1:]...)
			}
		}
	}
}

func (u *UI) HandleFrameEvents(gtx layout.Context, e event.Event) {
	for _, eh := range u.FrameEventHandlers {
		if eh.Event() {
			eh.Handler(gtx, e)
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
			u.Layout(gtx)
			e.Frame(gtx.Ops)
		}
		if u.HandleOtherEvents != nil {
			u.HandleOtherEvents(e, &ops)
		}
	}

	return nil
}
