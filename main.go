package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"tourOfGoConsole/tofgo"
)

func main() {
	gocui.NewGui(gocui.OutputNormal)

	r := tofgo.NewInjector()

	g, err := r.CreateTourOfGo()

	if err != nil {
		log.Panicln(err)
	}

	if err := g.InitTofGo(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	if err := g.Run(); err != nil {
		log.Panicln(err)
	}
}
