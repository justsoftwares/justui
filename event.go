package justui

import "gioui.org/layout"

type EventHandler struct {
	Event   func() bool
	Handler func(gtx layout.Context)
}
