package tofgo

import (
	"github.com/jroimartin/gocui"
	"tourOfGoConsole/tofgo/ui"
	"tourOfGoConsole/tofgo/ui/guidance_view"
)

type injector struct {
}

type Injector interface {
	CreateTourOfGo() (TourOfGo, error)
}

var gui, guiErr = gocui.NewGui(gocui.Output256)

func NewInjector() Injector {
	return &injector{}
}

func (i *injector) CreateGoCUI() (*gocui.Gui, error) {
	if guiErr != nil {
		return nil, guiErr
	}
	return gui, nil
}

func (i *injector) CreateTourOfGo() (TourOfGo, error) {
	g, err := i.CreateGoCUI()
	if err != nil {
		return nil, err
	}

	gv, err := i.CreateGuidanceView()
	if err != nil {
		return nil, err
	}

	return CreateTourOfGo(g, gv), nil
}

func (i *injector) CreateGuidanceView() (ui.View, error) {
	g, err := i.CreateGoCUI()
	if err != nil {
		return nil, err
	}
	return guidance_view.CreateGuidanceView(g), nil
}
