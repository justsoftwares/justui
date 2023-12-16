package explorer

import (
	"fmt"
	"gioui.org/widget"
	"os"
)

type FileElement struct {
	Root           string
	Name           string
	IsSelectedBool *widget.Bool
	//openClickable             *widget.Clickable
	//OpenClickableClickedEvent justui.EventHandler // here should to be a any method for handle this events,
	//												   but now I can`t do it...
	//explorer                  *Explorer
}

func NewFileElement(root, name string) *FileElement {
	f := &FileElement{
		Root:           root,
		Name:           name,
		IsSelectedBool: &widget.Bool{},
		//openClickable:  &widget.Clickable{},
		//explorer:       explorer,
	}
	/*f.OpenClickableClickedEvent = justui.EventHandler{
		Event: f.openClickable.Clicked,
		Handler: func(_ *justui.UI, _ layout.Context, _ event.Event) {
			f.explorer.directoryEditor.SetText(f.Name)
			f.explorer.Refresh()
		},
	}*/
	return f
}

func (f *FileElement) String() string {
	return f.FullPath()
}

func (f *FileElement) FullPath() string {
	return fmt.Sprintf("%s%c%s", f.Root, os.PathSeparator, f.Name)
}

//func (f *FileElement) layout() {
//}
