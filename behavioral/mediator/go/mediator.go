package main

type Mediator interface {
	canArrive(Trian) bool
	notifyAboutDeparture()
}
