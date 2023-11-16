package justui

import "gioui.org/layout"

type Layout struct {
	RootLayout    layout.Flex
	Children      []layout.FlexChild
	EventHandlers []EventHandler
}

func NewLayout(rootLayout layout.Flex) *Layout {
	return &Layout{RootLayout: rootLayout}
}

func (l *Layout) Add(c layout.FlexChild, handlers ...EventHandler) {
	l.Children = append(l.Children, c)
	l.EventHandlers = append(l.EventHandlers, handlers...)
}

func (l *Layout) Layout(gtx layout.Context) layout.Dimensions {
	return l.RootLayout.Layout(gtx, l.Children...)
}
