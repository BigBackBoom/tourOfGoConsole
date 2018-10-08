package tofgo

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
)

type tourOfGo struct {
	gui *gocui.Gui
}

type TourOfGo interface {
	InitTofGo() error
	Run() error
	Close()
}

func CreateTourOfGo() (TourOfGo, error) {
	gui, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		return nil, err
	}
	return &tourOfGo{gui}, nil
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
	if err := t.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
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
