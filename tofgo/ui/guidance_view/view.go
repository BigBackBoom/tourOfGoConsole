package guidance_view

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"tourOfGoConsole/tofgo/ui"
)

type guidanceView struct {
	goCUI *gocui.Gui
}

func CreateGuidanceView(g *gocui.Gui) ui.View {
	return &guidanceView{g}
}

func Layout(g *gocui.Gui) error {

	_, maxY := g.Size()
	if v, err := g.SetView("side", -1, -1, 30, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Item 1")
		fmt.Fprintln(v, "Item 2")
		fmt.Fprintln(v, "Item 3")
		fmt.Fprint(v, "\rWill be")
		fmt.Fprint(v, "deleted\rItem 4\nItem 5")
	}
	return nil
}

func (gv *guidanceView) InitKeyBindings() error {
	if err := gv.goCUI.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		return err
	}
	return nil
}
