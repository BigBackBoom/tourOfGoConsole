package main

import (
	"tourOfGoConsole/tofgo"
)

type injector struct {
}

type Injector interface {
	CreateTourOfGo() (tofgo.TourOfGo, error)
}

func NewInjector() Injector {
	return &injector{}
}

func (i *injector) CreateTourOfGo() (tofgo.TourOfGo, error) {
	return tofgo.CreateTourOfGo()
}
