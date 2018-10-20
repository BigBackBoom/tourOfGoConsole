package tofgo

import (
	"github.com/jroimartin/gocui"
	"log"
	"tourOfGoConsole/tofgo/ui"
	"tourOfGoConsole/tofgo/ui/guidance_view"
)

type tourOfGo struct {
	gui          *gocui.Gui
	guidanceView ui.View
}

type TourOfGo interface {
	InitTofGo() error
	Run() error
	Close()
}

func CreateTourOfGo(g *gocui.Gui, gv ui.View) TourOfGo {
	return &tourOfGo{
		g,
		gv,
	}
}

func (t *tourOfGo) InitTofGo() error {

	t.crtLayout()

	if err := t.setKeyBnd(); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

func (t *tourOfGo) crtLayout() {
	t.gui.SetManagerFunc(layout)
}

func (t *tourOfGo) setKeyBnd() error {
	t.guidanceView.InitKeyBindings()
	return nil
}

func (t *tourOfGo) Run() error {
	if err := t.gui.MainLoop(); err != nil {
		return err
	}
	return nil
}

func (t *tourOfGo) Close() {
	t.gui.Close()
}

func layout(g *gocui.Gui) error {
	guidance_view.Layout(g)
	return nil
}
