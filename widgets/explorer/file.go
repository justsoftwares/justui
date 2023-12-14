package explorer

import "gioui.org/widget"

type FileElement struct {
	Path           string
	IsSelectedBool *widget.Bool
	//openClickable             *widget.Clickable
	//OpenClickableClickedEvent justui.EventHandler // here should to be a any method for handle this events,
	//												   but now I can`t do it...
	//explorer                  *Explorer
}

func NewFileElement(path string) *FileElement {
	f := &FileElement{
		Path:           path,
		IsSelectedBool: &widget.Bool{},
		//openClickable:  &widget.Clickable{},
		//explorer:       explorer,
	}
	/*f.OpenClickableClickedEvent = justui.EventHandler{
		Event: f.openClickable.Clicked,
		Handler: func(_ *justui.UI, _ layout.Context, _ event.Event) {
			f.explorer.directoryEditor.SetText(f.Path)
			f.explorer.Refresh()
		},
	}*/
	return f
}

func (f *FileElement) String() string {
	return f.Path
}

func (f *FileElement) Widget() {

}
