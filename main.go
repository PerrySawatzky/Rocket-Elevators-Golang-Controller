package main

import (
	"fmt"
)

type Battery struct {
	id                        int
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
}
type Column struct {
	id                int
	amountOfElevators int
	servedFloors      int
	isBasement        bool
}
type Elevator struct {
	id                    int
	status                string
	currentFloor          int
	direction             string
	door                  int //figure this shit out later
	floorRequestList      int
	completedRequestsList int
}

func (e *Elevator) Init(id int, status string, currentFloor int, direction string, door int, floorRequestList int, completedRequestsList int) {
	e.id = id
	e.status = status
	e.currentFloor = currentFloor
	e.direction = direction
	e.door = door
	e.floorRequestList = floorRequestList
	e.completedRequestsList = completedRequestsList
}
func info(id int, status string, currentFloor int, direction string, door int, floorRequestList int, completedRequestsList int) *Elevator {
	e := new(Elevator)
	e.id = id
	e.status = status
	e.currentFloor = currentFloor
	e.direction = direction
	e.door = door
	e.floorRequestList = floorRequestList
	e.completedRequestsList = completedRequestsList
	return e
}
func main() {
	emp := new(Elevator)
	emp.Init(1, "idle", 1, "", 1, 1, 1)
	fmt.Print(emp.id, emp.currentFloor)
}

type CallButton struct {
	id        int
	floor     int
	direction string
}
type FloorRequestButton struct {
	id        int
	floor     int
	direction string
}
type Door struct {
	id int
}
