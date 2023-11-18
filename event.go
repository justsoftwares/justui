package justui

import (
	"gioui.org/io/event"
	"gioui.org/layout"
)

type EventHandler struct {
	Event   func() bool
	Handler func(u *UI, gtx layout.Context, e event.Event)
}
