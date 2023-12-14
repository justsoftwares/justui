package explorer

import (
	"gioui.org/app"
	"gioui.org/layout"
	"github.com/justsoftwares/justui"
)

func (e *Explorer) ShowWindow() {
	u := justui.NewUI(app.NewWindow(), e.Theme, func(gtx layout.Context) {
		layout.Flex{Axis: layout.Vertical}.Layout(
			gtx,
			layout.Rigid(e.Widget()),
		)
	})
	u.AddFrameEventHandlers(e.SelectClickableClickedEvent, e.DirectoryClickableClickedEvent)
	u.Run()
}
