package explorer

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/justsoftwares/justui"
	"log"
	"os"
)

type Explorer struct {
	Theme                               *material.Theme
	Files                               []*FileElement
	UI                                  *justui.UI
	filesInDir                          []*FileElement
	listWidget                          *widget.List
	directoryEditor                     *widget.Editor
	directoryClickable, selectClickable *widget.Clickable
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
	e.directoryEditor.SetText("C:\\")
	e.createUI()
	e.UI.AddFrameEventHandlers(justui.EventHandler{
		Event: e.selectClickable.Clicked,
		Handler: func(gtx layout.Context, evt event.Event) {
			e.Files = e.selectedFiles(e.Files)
			selectClickableClicked(gtx, evt)
			e.UI.Window.Perform(system.ActionClose)
		},
	}, justui.EventHandler{
		Event:   e.directoryClickable.Clicked,
		Handler: e.directoryClickableClicked,
	})
	return e
}

func (e *Explorer) GetSelectedFiles() []string {
	var res []string
	for _, fileElement := range e.Files {
		res = append(res, fileElement.FullPath())
	}
	return res
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
	e.Files = e.selectedFiles(e.Files)
	e.filesInDir = files
}

func (e *Explorer) Run() {
	e.UI.Run(false)
}

func (e *Explorer) createUI() {
	e.UI = justui.NewUI(app.NewWindow(), e.Theme, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(
			gtx,
			layout.Rigid(e.widget()),
		)
	})
}

func (e *Explorer) widget() layout.Widget {
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

func (e *Explorer) fileElement(f *FileElement, gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(
		gtx,
		layout.Rigid(material.CheckBox(e.Theme, f.IsSelectedBool, f.Name).Layout),
		//layout.Rigid(material.Button(e.Theme, f.openClickable, "Open").Layout),
	)
}

func (e *Explorer) selectedFiles(currentSelected []*FileElement) []*FileElement {
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

func (e *Explorer) getFilesInDirectory(path string) ([]*FileElement, error) {
	var files []*FileElement

	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			log.Println(err)
		}
	}(dir)

	fileInfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		file := NewFileElement(dir.Name(), fileInfo.Name())
		for _, el := range e.Files {
			if el.Name == fileInfo.Name() {
				file = el
				break
			}
		}
		files = append(files, file)
	}

	return files, nil
}

func (e *Explorer) directoryClickableClicked(_ layout.Context, _ event.Event) {
	e.Refresh()
}
