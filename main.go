package main

import (
	"fmt"
	//"container/list"
)

type Battery struct {
	ID                        int
	status                    string
	bestColumn                Column
	bestElevator              Elevator
	bestScore                 int
	referenceGap              int
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnsList               []Column
	floorRequestButtonsList   []FloorRequestButton
}

func (e Battery) findBestColumn(_requestedFloor int) Column {
	bestColumn := e.columnsList[0]
	return bestColumn
}

func (e Battery) assignElevator(_requestedFloor int, _direction string) Elevator {
	bestElevator := e.columnsList[0].elevatorsList[1]
	fmt.Print(bestElevator)
	return bestElevator

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
	e.status = "charged"
	e.amountOfColumns = _amountOfColumns
	e.amountOfFloors = _amountOfFloors
	e.amountOfBasements = _amountOfBasements
	e.amountOfElevatorPerColumn = _amountOfElevatorPerColumn
	//Column creator
	for i := 0; i < _amountOfColumns; i++ {
		column := new(Column)
		column.Init(i, _amountOfElevatorPerColumn, trueFalseinator(i))
		e.columnsList = append(e.columnsList, *column)
		//Assigning floors to each elevator column
		if i == 0 {
			for x := 0; x < 6; x++ {
				e.columnsList[0].servedFloors = append(e.columnsList[0].servedFloors, x-6)
			}
			e.columnsList[0].servedFloors = append(e.columnsList[0].servedFloors, 1)
		} //Sorry, I
		if i == 1 {
			for x := 0; x < 20; x++ {
				e.columnsList[1].servedFloors = append(e.columnsList[1].servedFloors, x+1)
			}
		} //Didn't make
		if i == 2 {
			for x := 20; x < 40; x++ {
				e.columnsList[2].servedFloors = append(e.columnsList[2].servedFloors, x+1)
			}
			e.columnsList[2].servedFloors = append(e.columnsList[2].servedFloors, 1)
		} //This simpler
		if i == 3 {
			for x := 40; x < 60; x++ {
				e.columnsList[3].servedFloors = append(e.columnsList[3].servedFloors, x+1)
			}
			e.columnsList[3].servedFloors = append(e.columnsList[3].servedFloors, 1)
		}
	}
	//Lobby only up and down buttons
	for i := 0; i < 1; i++ {
		upFloorRequestButtonCreator := new(FloorRequestButton)
		upFloorRequestButtonCreator.Init(i, i+1, "up")
		e.floorRequestButtonsList = append(e.floorRequestButtonsList, *upFloorRequestButtonCreator)

	}
	for i := 0; i < 1; i++ {
		downFloorRequestButtonCreator := new(FloorRequestButton)
		downFloorRequestButtonCreator.Init(i, i+1, "down")
		e.floorRequestButtonsList = append(e.floorRequestButtonsList, *downFloorRequestButtonCreator)
	}

}

type Column struct {
	ID                int
	status            string
	amountOfElevators int
	servedFloors      []int
	isBasement        bool
	bestElevator1     Elevator
	bestScore1        int
	referenceGap1     int
	elevatorsList     []Elevator
	callButtonsList   []CallButton
}

//func (e Battery) findBestColumn(_requestedFloor int) Column
func (e Column) requestElevator(_requestedFloor int, _direction string) Elevator {
	bestElevator1 := e.elevatorsList[0]
	return bestElevator1
}
func (e *Column) Init(_id int, _amountOfElevators int, _isBasement bool) {
	e.ID = _id
	e.status = "built"
	e.isBasement = _isBasement
	//Elevator creator
	for i := 0; i < _amountOfElevators; i++ {
		elevator := new(Elevator)
		elevator.Init(i, "idle", 1)
		e.elevatorsList = append(e.elevatorsList, *elevator)
	}
	//Creates the CallButton for the basement floors
	if _isBasement == true {
		for i := 0; i < 6; i++ {
			upCallButtonCreator := new(CallButton)
			upCallButtonCreator.Init(i, -6-i, "up")
			e.callButtonsList = append(e.callButtonsList, *upCallButtonCreator)
		}
	}
	//Creates the CallButton for everything else
	if _isBasement == false {
		for i := 2; i <= 60; i++ {
			downCallButtonCreator := new(CallButton)
			downCallButtonCreator.Init(i, i, "down")
			e.callButtonsList = append(e.callButtonsList, *downCallButtonCreator)
		}
	}
}

type Elevator struct {
	ID                    int
	status                string
	currentFloor          int
	direction             string
	door                  []Door
	floorRequestsList     []int
	completedRequestsList []int
}

func (e *Elevator) Init(_id int, _status string, _currentFloor int) {
	e.ID = _id
	e.status = _status
	e.currentFloor = _currentFloor
	newDoorCreator := new(Door)
	newDoorCreator.Init(_id)
	e.door = append(e.door, *newDoorCreator)

}

type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func (e *CallButton) Init(_id int, _floor int, _direction string) {
	e.ID = _id
	e.status = "off"
	e.floor = _floor
	e.direction = _direction
}

type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func (e *FloorRequestButton) Init(_id int, _floor int, _direction string) {
	e.ID = _id
	e.status = "off"
	e.floor = _floor
	e.direction = _direction
}

type Door struct {
	ID     int
	status string
}

func (e *Door) Init(_id int) {
	e.ID = _id
	e.status = "closed"
}

//Add door functionality, door.status = "open" etc
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
	fmt.Print(len(battery1.columnsList))
	fmt.Print("\ncolumnList ^")
	fmt.Print(len(battery1.floorRequestButtonsList))
	fmt.Print("\nfloorRequestButtonsList ^")
	fmt.Print(len(battery1.columnsList[0].elevatorsList))
	fmt.Print("\nelevatorsList ^")
	fmt.Print(len(battery1.columnsList[0].callButtonsList))
	fmt.Print("\ncallButtonsList ^")
	fmt.Print(len(battery1.columnsList[1].servedFloors))
	fmt.Print("\nservedFloors ^")
	fmt.Print((battery1.columnsList[1].elevatorsList[0].door))
	fmt.Print("\nelevatorsList[0].door ^")
	fmt.Print((battery1.columnsList[1].isBasement))
	fmt.Print("\nisBasement ^")
	//battery1.assignElevator(3, "up")
	fmt.Print("\nwhy does the import fmt get deleted if you save without having an fmt called. That's stupic.")
}
