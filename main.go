package main

import (
	"fmt"
	"math"
)

type Battery struct {
	ID                        int
	status                    string
	bestColumn                Column
	bestElevator              Elevator
	bestScore                 int
	referenceGap              float64
	amountOfColumns           int
	amountOfFloors            int
	amountOfBasements         int
	amountOfElevatorPerColumn int
	columnsList               []Column
	floorRequestButtonsList   []FloorRequestButton
}

//Used for the bool isBasement so the first Column is yes and the rest no
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

//Very specific to the Scenarios
func (this Battery) findBestColumn(_requestedFloor float64) Column {
	bestColumn := this.columnsList[0]
	if _requestedFloor < 1 {
		bestColumn = this.columnsList[0]
	}
	if _requestedFloor > 1 && _requestedFloor <= 20 {
		bestColumn = this.columnsList[1]
	}
	if _requestedFloor >= 21 && _requestedFloor <= 40 {
		bestColumn = this.columnsList[2]
	}
	if _requestedFloor >= 41 {
		bestColumn = this.columnsList[3]
	}
	fmt.Print("bestColumn is Array #")
	fmt.Print(bestColumn.ID)
	fmt.Print("\n")
	return bestColumn
}

func (this Battery) assignElevator(_requestedFloor float64, _direction string) Elevator {
	bestColumn := this.columnsList[0]
	bestElevator := this.columnsList[0].elevatorsList[1]
	bestScore := 5
	referenceGap := 10000
	//Did it this way because go doesn't have a .Contains operator. Ended up making a func later, but this worked so I kept it.
	if _requestedFloor < 1 {
		bestColumn = this.columnsList[0]
	}
	if _requestedFloor > 1 && _requestedFloor <= 20 {
		bestColumn = this.columnsList[1]
	}
	if _requestedFloor >= 21 && _requestedFloor <= 40 {
		bestColumn = this.columnsList[2]
	}
	if _requestedFloor >= 41 {
		bestColumn = this.columnsList[3]
	}
	if _requestedFloor > 1 {
		//Dont ask about what the referenceGap is doing, it works and thats all you need to know.
		for i := 0; i < len(bestColumn.elevatorsList); i++ {
			if bestColumn.elevatorsList[i].currentFloor == 1 && bestColumn.elevatorsList[i].status == "stopped" {
				bestElevator = bestColumn.elevatorsList[i]
				bestScore = 1
				referenceGap = int(math.Abs(float64(bestColumn.elevatorsList[i].currentFloor - 1)))
			}
			if bestColumn.elevatorsList[i].currentFloor == 1 && bestColumn.elevatorsList[i].status == "idle" {
				bestElevator = bestColumn.elevatorsList[i]
				bestScore = 2
				referenceGap = int(math.Abs(float64(bestColumn.elevatorsList[i].currentFloor - 1)))
			}
			if contains(bestColumn.elevatorsList[i].floorRequestsList, 1) == true && int(math.Abs(float64(bestColumn.elevatorsList[i].currentFloor-1))) < referenceGap && (bestColumn.elevatorsList[i].direction == "down") {
				bestElevator = bestColumn.elevatorsList[i]
				bestScore = 3
				referenceGap = int(math.Abs(float64(bestColumn.elevatorsList[i].currentFloor - 1)))
			}
			if bestColumn.elevatorsList[i].status == "idle" {
				bestElevator = bestColumn.elevatorsList[i]
				bestScore = 4
				referenceGap = int(math.Abs(float64(bestColumn.elevatorsList[i].currentFloor - 1)))
			}
		}
	}
	//Adds floor to floorRequestList
	bestElevator.floorRequestsList = append(bestElevator.floorRequestsList, _requestedFloor)
	//Moves elevator to requestedFloor
	for bestElevator.currentFloor > 1 {
		bestElevator.currentFloor--
		fmt.Print("Elevator is on floor ")
		fmt.Print(bestElevator.currentFloor)
		fmt.Print(".\n")
	}
	for bestElevator.currentFloor > _requestedFloor {
		bestElevator.currentFloor--
		bestElevator.status = "moving"
		fmt.Print("Elevator is on floor ")
		fmt.Print(bestElevator.currentFloor)
		fmt.Print(".\n")
	}
	for bestElevator.currentFloor < _requestedFloor {
		bestElevator.currentFloor++
		bestElevator.status = "moving"
		fmt.Print("Elevator is on floor ")
		fmt.Print(bestElevator.currentFloor)
		fmt.Print(".\n")
	}
	for bestElevator.currentFloor == _requestedFloor {
		bestElevator.completedRequestsList = append(bestElevator.completedRequestsList, _requestedFloor)
		bestElevator.status = "idle"
		fmt.Print("*DING* Elevator has arrived at floor ")
		fmt.Print(_requestedFloor)
		fmt.Print(".\n")
		break
	}
	//Scenario testing purposes
	fmt.Print("ReferenceGap = ")
	fmt.Print(referenceGap)
	fmt.Print("\n")
	fmt.Print("bestScore = ")
	fmt.Print(bestScore)
	fmt.Print("\n")
	fmt.Print("bestColumn is Array #")
	fmt.Print(bestColumn.ID)
	fmt.Print("\n")
	fmt.Print("bestElevator Array ID  = ")
	fmt.Print(bestElevator.ID)
	fmt.Print("\n")
	return bestElevator

}

//Because go does not have a .Contains method I needed to make one. e is the variable you want to check the slice for.
func contains(floorRequestsList []float64, e float64) bool {
	for _, a := range floorRequestsList {
		if a == e {
			return true
		}
	}
	return false
}

type Column struct {
	ID                int
	status            string
	amountOfElevators int
	servedFloors      []int
	isBasement        bool
	bestElevator1     Elevator
	bestScore1        int
	referenceGap1     float64
	elevatorsList     []Elevator
	callButtonsList   []CallButton
}

//The requestedFloor parameter is the floor the user is currently on, since they only have the option to head to the lobby.
func (this Column) requestElevator(_requestedFloor float64, _direction string) Elevator {
	bestElevator1 := this.elevatorsList[0]
	bestScore1 := 5
	referenceGap1 := 1000000
	for i := 0; i < len(this.elevatorsList); i++ {
		//For basement floors
		if _requestedFloor > -7 && _requestedFloor < 1 {
			if this.elevatorsList[i].currentFloor == _requestedFloor && this.elevatorsList[i].direction == "up" {
				bestElevator1 = this.elevatorsList[i]
				bestScore1 = 1
				referenceGap1 = int(math.Abs(float64(this.elevatorsList[i].currentFloor - _requestedFloor)))
			}
			if contains(this.elevatorsList[i].floorRequestsList, 1) == true && this.elevatorsList[i].currentFloor < _requestedFloor && (this.elevatorsList[i].direction == "up") {
				bestElevator1 = this.elevatorsList[i]
				bestScore1 = 2
				referenceGap1 = int(math.Abs(float64(this.elevatorsList[i].currentFloor - _requestedFloor)))
			}
			if this.elevatorsList[i].status == "idle" {
				bestElevator1 = this.elevatorsList[i]
				bestScore1 = 3
				referenceGap1 = int(math.Abs(float64(this.elevatorsList[i].currentFloor - _requestedFloor)))
			}
		}
		//Same thing but for floors above the lobby, just changed up for down essentially
		if _requestedFloor > 1 && _requestedFloor < 61 {
			if this.elevatorsList[i].currentFloor == _requestedFloor && this.elevatorsList[i].direction == "down" {
				bestElevator1 = this.elevatorsList[i]
				bestScore1 = 1
				referenceGap1 = int(math.Abs(float64(this.elevatorsList[i].currentFloor - _requestedFloor)))
			}
			if contains(this.elevatorsList[i].floorRequestsList, 1) == true && this.elevatorsList[i].currentFloor > _requestedFloor && int(math.Abs(float64(this.elevatorsList[i].currentFloor-_requestedFloor))) < referenceGap1 && (this.elevatorsList[i].direction == "down") {
				bestElevator1 = this.elevatorsList[i]
				bestScore1 = 2
				referenceGap1 = int(math.Abs(float64(this.elevatorsList[i].currentFloor - _requestedFloor)))
			}
			if this.elevatorsList[i].status == "idle" {
				bestElevator1 = this.elevatorsList[i]
				bestScore1 = 3
				referenceGap1 = int(math.Abs(float64(this.elevatorsList[i].currentFloor - _requestedFloor)))
			}
		}
	}
	//Moves elevator to lobby and prints a beautiful message each time.
	for bestElevator1.currentFloor > 1 {
		for bestElevator1.currentFloor == _requestedFloor {
			fmt.Print("*DING* Elevator doors are open, please enter.\n")
			break
		}
		bestElevator1.currentFloor--
		bestElevator1.status = "moving"
		fmt.Print("Elevator is on floor ")
		fmt.Print(bestElevator1.currentFloor)
		fmt.Print(".\n")
	}
	for bestElevator1.currentFloor < -1 {
		bestElevator1.currentFloor++
		bestElevator1.status = "moving"
		fmt.Print("Elevator is on floor ")
		fmt.Print(bestElevator1.currentFloor)
		fmt.Print(".\n")
	}
	for bestElevator1.currentFloor == -1 {
		bestElevator1.currentFloor++
	}
	for bestElevator1.currentFloor == 0 {
		bestElevator1.currentFloor++
	}
	for bestElevator1.currentFloor == 1 {
		bestElevator1.status = "idle"
		fmt.Print("*DING* Elevator has arrived at Lobby.\n")
		break
	}
	//Scenario testing is done here so I don't need to write it in main
	fmt.Print("ReferenceGap = ")
	fmt.Print(referenceGap1)
	fmt.Print("\n")
	fmt.Print("bestScore = ")
	fmt.Print(bestScore1)
	fmt.Print("\n")
	fmt.Print("bestElevator Array ID = ")
	fmt.Print(bestElevator1.ID)
	fmt.Print("\n")
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
	currentFloor          float64
	direction             string
	door                  []Door
	floorRequestsList     []float64
	completedRequestsList []float64
}

func (e *Elevator) Init(_id int, _status string, _currentFloor float64) {
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
func main() {
	//Test Scenario 1
	battery1 := new(Battery)
	battery1.Init(1, 4, 60, 6, 5)
	battery1.findBestColumn(20)
	battery1.columnsList[1].elevatorsList[0].currentFloor = 20
	battery1.columnsList[1].elevatorsList[1].currentFloor = 3
	battery1.columnsList[1].elevatorsList[2].currentFloor = 13
	battery1.columnsList[1].elevatorsList[3].currentFloor = 15
	battery1.columnsList[1].elevatorsList[4].currentFloor = 6
	battery1.columnsList[1].elevatorsList[0].floorRequestsList = append(battery1.columnsList[1].elevatorsList[0].floorRequestsList, 5)
	battery1.columnsList[1].elevatorsList[1].floorRequestsList = append(battery1.columnsList[1].elevatorsList[1].floorRequestsList, 15)
	battery1.columnsList[1].elevatorsList[2].floorRequestsList = append(battery1.columnsList[1].elevatorsList[2].floorRequestsList, 1)
	battery1.columnsList[1].elevatorsList[3].floorRequestsList = append(battery1.columnsList[1].elevatorsList[3].floorRequestsList, 2)
	battery1.columnsList[1].elevatorsList[4].floorRequestsList = append(battery1.columnsList[1].elevatorsList[4].floorRequestsList, 1)
	battery1.columnsList[1].elevatorsList[0].status = "moving"
	battery1.columnsList[1].elevatorsList[1].status = "moving"
	battery1.columnsList[1].elevatorsList[2].status = "moving"
	battery1.columnsList[1].elevatorsList[3].status = "moving"
	battery1.columnsList[1].elevatorsList[4].status = "moving"
	battery1.columnsList[1].elevatorsList[0].direction = "down"
	battery1.columnsList[1].elevatorsList[1].direction = "up"
	battery1.columnsList[1].elevatorsList[2].direction = "down"
	battery1.columnsList[1].elevatorsList[3].direction = "down"
	battery1.columnsList[1].elevatorsList[4].direction = "down"
	battery1.assignElevator(20, "up")

	//Test Scenario 2
	// battery2 := new(Battery)
	// battery2.Init(1, 4, 60, 6, 5)
	// battery2.findBestColumn(36)
	// battery2.columnsList[2].elevatorsList[0].currentFloor = 1
	// battery2.columnsList[2].elevatorsList[1].currentFloor = 23
	// battery2.columnsList[2].elevatorsList[2].currentFloor = 33
	// battery2.columnsList[2].elevatorsList[3].currentFloor = 40
	// battery2.columnsList[2].elevatorsList[4].currentFloor = 39
	// battery2.columnsList[2].elevatorsList[0].floorRequestsList = append(battery2.columnsList[2].elevatorsList[0].floorRequestsList, 21)
	// battery2.columnsList[2].elevatorsList[1].floorRequestsList = append(battery2.columnsList[2].elevatorsList[1].floorRequestsList, 28)
	// battery2.columnsList[2].elevatorsList[2].floorRequestsList = append(battery2.columnsList[2].elevatorsList[2].floorRequestsList, 1)
	// battery2.columnsList[2].elevatorsList[3].floorRequestsList = append(battery2.columnsList[2].elevatorsList[3].floorRequestsList, 24)
	// battery2.columnsList[2].elevatorsList[4].floorRequestsList = append(battery2.columnsList[2].elevatorsList[4].floorRequestsList, 39)
	// battery2.columnsList[2].elevatorsList[0].status = "stopped"
	// battery2.columnsList[2].elevatorsList[1].status = "moving"
	// battery2.columnsList[2].elevatorsList[2].status = "moving"
	// battery2.columnsList[2].elevatorsList[3].status = "moving"
	// battery2.columnsList[2].elevatorsList[4].status = "moving"
	// battery2.columnsList[2].elevatorsList[0].direction = ""
	// battery2.columnsList[2].elevatorsList[1].direction = "up"
	// battery2.columnsList[2].elevatorsList[2].direction = "down"
	// battery2.columnsList[2].elevatorsList[3].direction = "down"
	// battery2.columnsList[2].elevatorsList[4].direction = "down"
	// battery2.assignElevator(36, "up")

	//Test Scenario 3
	// battery3 := new(Battery)
	// battery3.Init(2, 4, 60, 6, 5)
	// battery3.columnsList[3].elevatorsList[0].currentFloor = 58
	// battery3.columnsList[3].elevatorsList[1].currentFloor = 50
	// battery3.columnsList[3].elevatorsList[2].currentFloor = 46
	// battery3.columnsList[3].elevatorsList[3].currentFloor = 1
	// battery3.columnsList[3].elevatorsList[4].currentFloor = 60
	// battery3.columnsList[3].elevatorsList[0].floorRequestsList = append(battery3.columnsList[3].elevatorsList[0].floorRequestsList, 1)
	// battery3.columnsList[3].elevatorsList[1].floorRequestsList = append(battery3.columnsList[3].elevatorsList[1].floorRequestsList, 60)
	// battery3.columnsList[3].elevatorsList[2].floorRequestsList = append(battery3.columnsList[3].elevatorsList[2].floorRequestsList, 58)
	// battery3.columnsList[3].elevatorsList[3].floorRequestsList = append(battery3.columnsList[3].elevatorsList[3].floorRequestsList, 54)
	// battery3.columnsList[3].elevatorsList[4].floorRequestsList = append(battery3.columnsList[3].elevatorsList[4].floorRequestsList, 1)
	// battery3.columnsList[3].elevatorsList[0].status = "moving"
	// battery3.columnsList[3].elevatorsList[1].status = "moving"
	// battery3.columnsList[3].elevatorsList[2].status = "moving"
	// battery3.columnsList[3].elevatorsList[3].status = "moving"
	// battery3.columnsList[3].elevatorsList[4].status = "moving"
	// battery3.columnsList[3].elevatorsList[0].direction = "down"
	// battery3.columnsList[3].elevatorsList[1].direction = "up"
	// battery3.columnsList[3].elevatorsList[2].direction = "up"
	// battery3.columnsList[3].elevatorsList[3].direction = "up"
	// battery3.columnsList[3].elevatorsList[4].direction = "down"
	// battery3.columnsList[3].requestElevator(54, "down")

	//Test Scenario 4
	// battery4 := new(Battery)
	// battery4.Init(2, 4, 60, 6, 5)
	// battery4.columnsList[0].elevatorsList[0].currentFloor = -4
	// battery4.columnsList[0].elevatorsList[1].currentFloor = 1
	// battery4.columnsList[0].elevatorsList[2].currentFloor = -3
	// battery4.columnsList[0].elevatorsList[3].currentFloor = -6
	// battery4.columnsList[0].elevatorsList[4].currentFloor = -1
	// battery4.columnsList[0].elevatorsList[2].floorRequestsList = append(battery4.columnsList[0].elevatorsList[2].floorRequestsList, -5)
	// battery4.columnsList[0].elevatorsList[3].floorRequestsList = append(battery4.columnsList[0].elevatorsList[3].floorRequestsList, 1)
	// battery4.columnsList[0].elevatorsList[4].floorRequestsList = append(battery4.columnsList[0].elevatorsList[4].floorRequestsList, -6)
	// battery4.columnsList[0].elevatorsList[0].status = "idle"
	// battery4.columnsList[0].elevatorsList[1].status = "idle"
	// battery4.columnsList[0].elevatorsList[2].status = "moving"
	// battery4.columnsList[0].elevatorsList[3].status = "moving"
	// battery4.columnsList[0].elevatorsList[4].status = "moving"
	// battery4.columnsList[0].elevatorsList[0].direction = ""
	// battery4.columnsList[0].elevatorsList[1].direction = ""
	// battery4.columnsList[0].elevatorsList[2].direction = "down"
	// battery4.columnsList[0].elevatorsList[3].direction = "up"
	// battery4.columnsList[0].elevatorsList[4].direction = "down"
	// battery4.columnsList[0].requestElevator(-3, "up")

}
