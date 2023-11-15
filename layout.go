package justui

import "gioui.org/layout"

type LayoutChild interface {
	Layout(gtx layout.Context) layout.Dimensions
}

type Layout struct {
	RootLayout     layout.Flex
	layoutChildren []layout.FlexChild
	eventHandlers  []EventHandler
}

func (l *Layout) Add(child LayoutChild, eventHandlers ...EventHandler) {
	l.layoutChildren = append(l.layoutChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return child.Layout(gtx)
	}))
	l.eventHandlers = append(l.eventHandlers, eventHandlers...)
}

func (l *Layout) GetChildren() []layout.FlexChild {
	return l.layoutChildren
}

func (l *Layout) GetEventHandlers() []EventHandler {
	return l.eventHandlers
}
