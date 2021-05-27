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
	columnList                []Column
	floorRequestButtonList    []FloorRequestButton
}

func trueFalseinator(i int) bool {
	if i == 0 {
		return true
	} else {

		return false
	}
}
func (e *Battery) Init(_id int, _amountOfColumns int, _amountOfFloors int, _amountOfBasements int, _amountOfElevatorPerColumn int) {
	e.ID = _id
	e.amountOfColumns = _amountOfColumns
	e.amountOfFloors = _amountOfFloors
	e.amountOfBasements = _amountOfBasements
	e.amountOfElevatorPerColumn = _amountOfElevatorPerColumn
	//Column creator
	for i := 0; i < _amountOfColumns; i++ {
		column := new(Column)
		column.Init(i, _amountOfElevatorPerColumn, trueFalseinator(i))
		e.columnList = append(e.columnList, *column)
		//Assigning floors to each elevator column
		if i == 0 {
			for x := 0; x < 6; x++ {
				e.columnList[0].servedFloors = append(e.columnList[0].servedFloors, x-6)
			}
			e.columnList[0].servedFloors = append(e.columnList[0].servedFloors, 1)
		} //Sorry, I
		if i == 1 {
			for x := 0; x < 20; x++ {
				e.columnList[1].servedFloors = append(e.columnList[1].servedFloors, x+1)
			}
		} //Didn't make
		if i == 2 {
			for x := 20; x < 40; x++ {
				e.columnList[2].servedFloors = append(e.columnList[2].servedFloors, x+1)
			}
			e.columnList[2].servedFloors = append(e.columnList[2].servedFloors, 1)
		} //This simpler
		if i == 3 {
			for x := 40; x < 60; x++ {
				e.columnList[3].servedFloors = append(e.columnList[3].servedFloors, x+1)
			}
			e.columnList[3].servedFloors = append(e.columnList[3].servedFloors, 1)
		}
	}
	//Lobby only up and down buttons
	for i := 0; i < 1; i++ {
		upFloorRequestButtonCreator := new(FloorRequestButton)
		upFloorRequestButtonCreator.Init(i, i+1, "up")
		e.floorRequestButtonList = append(e.floorRequestButtonList, *upFloorRequestButtonCreator)

	}
	for i := 0; i < 1; i++ {
		downFloorRequestButtonCreator := new(FloorRequestButton)
		downFloorRequestButtonCreator.Init(i, i+1, "down")
		e.floorRequestButtonList = append(e.floorRequestButtonList, *downFloorRequestButtonCreator)
	}

}

type Column struct {
	ID                int
	status            string
	amountOfElevators int
	servedFloors      []int
	isBasement        bool
	elevatorList      []Elevator
	callButtonList    []CallButton
}

func (e *Column) Init(_id int, _amountOfElevators int, _isBasement bool) {
	e.ID = _id
	e.status = "built"
	e.isBasement = _isBasement
	//Elevator creator
	for i := 0; i < _amountOfElevators; i++ {
		elevator := new(Elevator)
		elevator.Init(i, "idle", 1)
		e.elevatorList = append(e.elevatorList, *elevator)
	}
	//Creates the CallButton for the basement floors
	if _isBasement == true {
		for i := 0; i < 6; i++ {
			upCallButtonCreator := new(CallButton)
			upCallButtonCreator.Init(i, -6-i, "up")
			e.callButtonList = append(e.callButtonList, *upCallButtonCreator)
		}
	}
	//Creates the CallButton for everything else
	if _isBasement == false {
		for i := 2; i <= 60; i++ {
			downCallButtonCreator := new(CallButton)
			downCallButtonCreator.Init(i, i, "down")
			e.callButtonList = append(e.callButtonList, *downCallButtonCreator)
		}
	}
}

type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	direction             string
	floorRequestList      []int
	completedRequestsList []int
	door                  []Door //try making this a for loop
}

func (e *Elevator) Init(_id int, _status string, _currentFloor int) {
	e.ID = _id
	e.status = _status
	e.currentFloor = _currentFloor
}

type CallButton struct {
	ID        int
	floor     int
	direction string
}

func (e *CallButton) Init(_id int, _floor int, _direction string) {
	e.ID = _id
	e.floor = _floor
	e.direction = _direction
}

type FloorRequestButton struct {
	ID        int
	floor     int
	direction string
}

func (e *FloorRequestButton) Init(_id int, _floor int, _direction string) {
	e.ID = _id
	e.floor = _floor
	e.direction = _direction
}

type Door struct {
	ID int
}

//Build Door func, add it into the elevator Init.
//Then on to the fun methods and selectors tomorrow >:)
func main() {
	// elevator1 := new(Elevator)
	// elevator1.Init(1, "idle", 1)
	// column1 := new(Column)
	// column1.Init(1, 5, false)
	//fmt.Print(column1.ID, column1.amountOfElevators)
	//fmt.Print(elevator1.ID, elevator1.currentFloor)
	battery1 := new(Battery)
	battery1.Init(1, 4, 60, 6, 5)
	fmt.Print(battery1.columnList[1].callButtonList)
}
