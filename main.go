package main

import (
	"log"
	"tourOfGoConsole/tofgo"
)

func main() {

	t := tofgo.TourOfGo{}
	err := t.InitTofGo()

	if err != nil {
		log.Panicln(err)
	}
	defer t.Close()

	t.CrtLayout()

	if err := t.SetKeyBnd(); err != nil {
		log.Panicln(err)
	}

	if err := t.Run(); err != nil {
		log.Panicln(err)
	}
}
