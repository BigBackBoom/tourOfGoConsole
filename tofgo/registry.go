package tofgo

import (
	"github.com/jroimartin/gocui"
	"tourOfGoConsole/tofgo/ui/handler"
)

type interactor struct {
	gui *gocui.Gui
}

type Iteractor interface {
}

func NewInteractor(gui *gocui.Gui) Iteractor {
	return &interactor{gui}
}
func (i *interactor) NewAppHandler() handler.AppHandler {
	return i.NewUserHandler()
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler()
}
