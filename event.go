package justui

import (
	"gioui.org/io/event"
	"gioui.org/layout"
)

type Handler func(gtx layout.Context, e event.Event)

type EventHandler struct {
	Event   func(gtx layout.Context) bool
	Handler Handler
}
