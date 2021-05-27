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

	for i := 0; i < _amountOfColumns; i++ {
		column := new(Column)
		column.Init(i, _amountOfElevatorPerColumn, i, trueFalseinator(i))
		e.columnList = append(e.columnList, *column)
		if i == 0 {
			for x := 0; x < 6; x++ {
				//columnList[0].servedFloors.Add(x - 6) adds basements
			}
			//columnList[0].servedFloors.Add(1) adding lobby
		}
		if i == 1 {
			for x := 0; x < 20; x++ {
				//columnList[1].servedFloors.Add(x + 1)
			}
		}
		if i == 2 {
			for x := 20; x < 40; x++ {
				//columnList[2].servedFloors.Add(x + 1)
			}
			//columnList[2].servedFloors.Add(1) adding lobby
		}
		if i == 3 {
			for x := 40; x < 60; x++ {
				//columnList[3].servedFloors.Add(x + 1)
			}
			//columnList[3].servedFloors.Add(1) adding lobby
		}
	}
	//Lobby only up and down buttons
	for i := 0; i < 1; i++ {
		//upFloorRequestButtonCreator = new FloorRequestButton(i, i + 1, "up");
		//floorRequestButtonsList.Add(upFloorRequestButtonCreator);
	}
	for i := 0; i < 1; i++ {
		//downFloorRequestButtonCreator = new FloorRequestButton(i, i + 1, "down");
		//floorRequestButtonsList.Add(downFloorRequestButtonCreator);
	}

}

type Column struct {
	ID                int
	status            string
	amountOfElevators int
	servedFloors      int
	isBasement        bool
	elevatorList      []Elevator
	callButtonList    []CallButton
}

func (e *Column) Init(_id int, _amountOfElevators int, _servedFloors int, _isBasement bool) {
	e.ID = _id
	e.status = "built"
	e.servedFloors = _servedFloors
	e.isBasement = _isBasement
	for i := 0; i < _amountOfElevators; i++ {
		//elevator := new(Elevator)
		//elevator.Init(i, "")
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
	//elevator1.Init(1, "idle", 1, "")
	column1 := new(Column)
	column1.Init(1, 5, 0, false)
	fmt.Print(column1.ID, column1.amountOfElevators)
	fmt.Print(elevator1.ID, elevator1.currentFloor)
	battery1 := new(Battery)
	battery1.Init(1, 4, 60, 6, 5)
	fmt.Print(battery1.columnList)
}
