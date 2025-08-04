package services

import (
	"fmt"
	"math/rand"
	"parking/constants"
	"parking/models"
	"time"
)

type SpotManager struct {
	Spots    []*models.Spot
	Tofill   map[int64]int64
	Unfilled []int64
}

func (sm *SpotManager) Init() {
	sm.Spots = []*models.Spot{}
	for floor := 1; floor <= constants.FLOORS; floor++ {
		for slot := 1; slot <= constants.SLOTS; slot++ {
			sm.Spots = append(sm.Spots, &models.Spot{
				Floor:  floor,
				Slot:   slot,
				Status: false,
			})
		}
	}
	sm.Tofill, sm.Unfilled = populate(constants.FLOORS, constants.SLOTS)
}

func populate(m, n int) (map[int64]int64, []int64) {
	tofill := make(map[int64]int64)
	unfilled := make([]int64, 0, m*n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			id := int64(i*100 + j)
			tofill[int64(id)] = 0
			unfilled = append(unfilled, int64(id))
		}
	}
	return tofill, unfilled
}

func (sm *SpotManager) Park(CarType string, CarRegd int64) {
	cars := models.Cars{
		Type:  CarType,
		Regd:  CarRegd,
		Entry: time.Now(),
	}
	spotId := sm.pickRandom()
	slotId, floor := floorSpot(int(spotId))
	fmt.Printf("Car with %d reg id is alloted %d floor at %d slot\n", CarRegd, floor, slotId)
	// Find the spot and update it
	for _, s := range sm.Spots {
		if s.Floor == floor && s.Slot == slotId {
			s.Car = &cars
			s.Status = true
			break
		}
	}
	sm.Tofill[CarRegd] = spotId
}

func (sm *SpotManager) pickRandom() int64 {
	n := len(sm.Unfilled)
	rand.Seed(time.Now().UnixNano()) // Always seed once, usually at start

	m := rand.Intn(n)
	val := sm.Unfilled[m]
	sm.Unfilled = append(sm.Unfilled[:m], sm.Unfilled[m+1:]...)
	return val
}

func floorSpot(slot int) (floor int, spot int) {
	spot = slot % 100
	floor = slot - spot
	return floor / 100, spot
}

func (sm *SpotManager) Unpark(CarRegd int64) {
	fsp := sm.Tofill[CarRegd]
	if fsp == 0 {
		fmt.Println("Car isn't park")
		return
	}
	floor, slot := floorSpot(int(fsp))
	index := 0
	for i := range sm.Spots {
		if sm.Spots[i].Floor == floor && sm.Spots[i].Slot == slot {
			index = i
			break
		}
	}
	sm.Spots = append(sm.Spots[index:], sm.Spots[:index+1]...)
	sm.Unfilled = append(sm.Unfilled, fsp)
	sm.Tofill[CarRegd] = 0
	fmt.Printf("Car with regd %d has been unparked\n", CarRegd)
}

// statusall
func (sm *SpotManager) Statusall() {
	n := len(sm.Unfilled)
	
	fmt.Println("These many slots are empty", n)
}

// checkmycar
func (sm *SpotManager) Checkmycar(CarRegd int64) {
	fsp := sm.Tofill[CarRegd]
	if fsp == 0 {
		fmt.Println("Car isn't park")
		return
	}
	floor, slot := floorSpot(int(fsp))

	fmt.Printf("Car with %d reg id is present on %d floor at %d slot\n", CarRegd, floor, slot)
}
