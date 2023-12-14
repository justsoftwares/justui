package explorer

import (
	"gioui.org/io/event"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"log"
	"os"
)

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

func (e *FileElement) String() string {
	return e.Path
}

type Explorer struct {
	Theme                                                       *material.Theme
	Files                                                       []*FileElement
	filesInDir                                                  []*FileElement
	listWidget                                                  *widget.List
	directoryEditor                                             *widget.Editor
	directoryClickable, selectClickable                         *widget.Clickable
	DirectoryClickableClickedEvent, SelectClickableClickedEvent justui.EventHandler
}

func NewExplorer(theme *material.Theme, selectClickableClicked justui.Handler) *Explorer {
	e := &Explorer{
		Theme: theme,
		listWidget: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
		directoryEditor:    &widget.Editor{},
		directoryClickable: &widget.Clickable{},
		selectClickable:    &widget.Clickable{},
	}
	e.SelectClickableClickedEvent = justui.EventHandler{
		Event: e.selectClickable.Clicked,
		Handler: func(u *justui.UI, gtx layout.Context, evt event.Event) {
			e.Files = e.SelectedFiles(e.Files)
			selectClickableClicked(u, gtx, evt)
		},
	}
	e.DirectoryClickableClickedEvent = justui.EventHandler{
		Event:   e.directoryClickable.Clicked,
		Handler: e.directoryClickableClicked,
	}
	e.directoryEditor.SetText("C:\\")
	return e
}

func (e *Explorer) SelectFiles() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		buttonsWidget := func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Axis: layout.Horizontal,
			}.Layout(gtx, layout.Flexed(0.5, material.Button(e.Theme, e.directoryClickable, "Show").Layout),
				layout.Flexed(0.5, material.Button(e.Theme, e.selectClickable, "Select").Layout))
		}
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(material.Editor(e.Theme, e.directoryEditor, "C:\\").Layout),
			layout.Rigid(buttonsWidget),
			layout.Rigid(e.fileList()),
		)
	}
}

func (e *Explorer) fileList() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return material.List(e.Theme, e.listWidget).Layout(
			gtx,
			len(e.filesInDir),
			func(gtx layout.Context, index int) layout.Dimensions {
				f := e.filesInDir[index]

				return e.fileElement(f, gtx)
			},
		)
	}
}

func (e *Explorer) directoryClickableClicked(_ *justui.UI, _ layout.Context, _ event.Event) {
	e.Refresh()
}

func (e *Explorer) fileElement(f *FileElement, gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Rigid(material.CheckBox(e.Theme, f.IsSelectedBool, f.Path).Layout),
		//layout.Rigid(material.Button(e.Theme, f.openClickable, "Open").Layout),
	)
}

func (e *Explorer) getFilesInDirectory(path string) ([]*FileElement, error) {
	var files []*FileElement

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		file := NewFileElement(fileInfo.Name())
		for _, el := range e.Files {
			if el.Path == fileInfo.Name() {
				file = el
				break
			}
		}
		files = append(files, file)
	}

	return files, nil
}

func (e *Explorer) Refresh() {
	currentPath := e.directoryEditor.Text()

	if currentPath == "" {
		return
	}

	files, err := e.getFilesInDirectory(currentPath)
	if err != nil {
		log.Println("Error getting files:", err)
		return
	}
	e.Files = e.SelectedFiles(e.Files)
	e.filesInDir = files
}

func (e *Explorer) SelectedFiles(currentSelected []*FileElement) []*FileElement {
	for _, fileInDir := range e.filesInDir {
		exist := false
		for csI, csFile := range currentSelected {
			if csFile == fileInDir {
				if fileInDir.IsSelectedBool.Value {
					exist = true
				} else {
					currentSelected = append(currentSelected[:csI], currentSelected[csI+1:]...)
				}
				break
			}
		}
		if !exist && fileInDir.IsSelectedBool.Value {
			currentSelected = append(currentSelected, fileInDir)
		}
	}
	return currentSelected
}
