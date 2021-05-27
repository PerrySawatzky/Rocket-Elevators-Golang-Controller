package main

import (
	"fmt"
	//"container/list"
)

type Battery struct {
	ID                        int
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
}
func (e *Battery) Init(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int){
	e.ID = _id
	e.amountOfColumns = _amountOfColumns
	e.amountOfFloors = _amountOfFloors
	e.amountOfBasements = _amountOfBasements
	e.amountOfElevatorPerColumn = _amountOfElevatorPerColumn
	for()
}
type Column struct {
	ID                int
	status            string
	amountOfElevators int
	servedFloors      int
	isBasement        bool
	elevatorList      int
	callButtonList    int
}

func (e *Column) Init(_id int, _amountOfElevators int, _servedFloors int, _isBasement bool, elevatorList int, callButtonList int) {
	e.ID = _id
	e.status = "built"
	e.servedFloors = _servedFloors
	e.isBasement = _isBasement
	e.elevatorList = elevatorList
	e.callButtonList = callButtonList
}
type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	direction             string
	door                  int //figure this shit out later
	floorRequestList      int
	completedRequestsList int
}

func (e *Elevator) Init(_id int, _status string, _currentFloor int, _direction string, _door int, _floorRequestList int, _completedRequestsList int) {
	e.ID = _id
	e.status = _status
	e.currentFloor = _currentFloor
	e.direction = _direction
	e.door = _door
	e.floorRequestList = _floorRequestList
	e.completedRequestsList = _completedRequestsList
}
type CallButton struct {
	ID        int
	floor     int
	direction string
}
type FloorRequestButton struct {
	ID        int
	floor     int
	direction string
}
type Door struct {
	ID int
}

func main() {
	elevator1 := new(Elevator)
	elevator1.Init(1, "idle", 1, "", 1, 1, 1)
	column1 := new(Column)
	column1.Init(1, 5, 0, false, 1, 1)
	fmt.Print(column1.ID, column1.amountOfElevators)
	fmt.Print(emp.ID, emp.currentFloor)
}