package main
CounterTerrroristDressType
type CounterTerroristDress struct {
	color string
}

func (t *CounterTerroristDress) getColor() string {
	return t.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{
		color: "green",
	}
}
