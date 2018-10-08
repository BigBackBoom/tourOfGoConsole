package main

import (
	"tourOfGoConsole/tofgo"
)

type interactor struct {
}

type Iteractor interface {
	CreateTourOfGo() (tofgo.TourOfGo, error)
}

func NewInteractor() Iteractor {
	return &interactor{}
}

func (i *interactor) CreateTourOfGo() (tofgo.TourOfGo, error) {
	return tofgo.CreateTourOfGo()
}
