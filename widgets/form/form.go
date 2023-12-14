package form

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Form struct {
	Theme        *material.Theme
	Title        string
	rootLayout   *layout.Flex
	SubmitButton layout.Widget
	Fields       []*Field
}

func NewForm(theme *material.Theme, title string, rootLayout *layout.Flex, submitButton layout.Widget, fields ...*Field) *Form {
	if rootLayout == nil {
		rootLayout = &layout.Flex{
			Axis: layout.Vertical,
		}
	}
	return &Form{
		Theme:        theme,
		Title:        title,
		rootLayout:   rootLayout,
		SubmitButton: submitButton,
		Fields:       fields,
	}
}

func (f *Form) AddField(title string, widget layout.Widget) {
	f.Fields = append(f.Fields, NewField(f.Theme, title, widget, nil))
}

func (f *Form) Layout(gtx layout.Context) layout.Dimensions {
	el := layout.Flex{
		Axis: layout.Vertical,
	}
	var children []layout.FlexChild
	children = append(children, layout.Rigid(material.Caption(f.Theme, f.Title).Layout))
	for _, field := range f.Fields {
		children = append(children, layout.Rigid(field.Layout))
	}
	children = append(children, layout.Rigid(f.SubmitButton))
	return el.Layout(gtx, children...)
}
