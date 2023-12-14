package form

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Field struct {
	Theme      *material.Theme
	Title      string
	Widget     layout.Widget
	rootLayout *layout.Flex
}

func NewField(theme *material.Theme, title string, widget layout.Widget, rootLayout *layout.Flex) *Field {
	if rootLayout == nil {
		rootLayout = &layout.Flex{
			Axis: layout.Horizontal,
		}
	}
	return &Field{
		Theme:      theme,
		Title:      title,
		Widget:     widget,
		rootLayout: rootLayout,
	}
}

func (f *Field) Layout(gtx layout.Context) layout.Dimensions {
	return f.rootLayout.Layout(gtx, layout.Rigid(material.Body1(f.Theme, f.Title+": ").Layout), layout.Rigid(f.Widget))
}
