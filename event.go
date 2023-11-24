package justui

import (
	"gioui.org/io/event"
	"gioui.org/layout"
)

type Handler func(u *UI, gtx layout.Context, e event.Event)

type EventHandler struct {
	Event   func() bool
	Handler Handler
}
