package models

type Spot struct {
	Floor  int
	Slot   int
	Status bool
	Car    *Cars
}
