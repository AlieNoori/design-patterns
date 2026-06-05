package main

import "fmt"

const (
	// TerroristDressType terrorist dress type
	TerroristDressType = "tDress"

	// CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var dressFactorySingleInstance = &DressFactory{
	dressMap: make(map[string]Dress),
}

type DressFactory struct {
	dressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if dress, ok := d.dressMap[dressType]; ok {
		return dress, nil
	}

	if dressType == TerroristDressType {
		newTDress := newTerroristDress()
		d.dressMap[dressType] = newTDress
		return newTDress, nil
	}

	if dressType == TerroristDressType {
		newCTDress := newCounterTerroristDress()
		d.dressMap[dressType] = newCTDress
		return newCTDress, nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
