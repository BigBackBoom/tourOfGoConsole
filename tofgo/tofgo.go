package tofgo

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type TourOfGo struct {
	gui *gocui.Gui
}

func (t *TourOfGo) InitTofGo() error {

	var err error
	t.gui, err = gocui.NewGui(gocui.Output256)

	if err != nil {
		return err
	}
	return nil
}

func (t *TourOfGo) CrtLayout() {
	t.gui.SetManagerFunc(layout)
}

func (t *TourOfGo) SetKeyBnd() error {
	if err := t.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func (t *TourOfGo) Run() error {
	if err := t.gui.MainLoop(); err != nil {
		return err
	}
	return nil
}

func (t *TourOfGo) Close() {
	t.gui.Close()
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
