package elevator

import (
	"sort"
	"sync"
)

type Direction int

const (
	UP Direction = itoa
	DOWN 
	IDLE
)


func (d Direction) String() string {
	switch d {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	default:
		return "IDLE"
	}
}

type Elevator struct {
	ID int
	CurrentFloor int
	Direction Direction
	State ElevatorState
	Requests []int
	MaxFloor int
	MinFloor int
	mutext sync.RWMutex
}


// New Elevator creates a new elevator
func NewElevator(id, minFloor, maxFloor int) *Elevator {
	return &Elevator{
		ID: id,
		MinFloor: minFloor,
		MaxFloor: maxFloor,
		Direction: IDLE,
		State: STOPPED,
		Requests: make([]int, 0),
		MaxFloor: maxFloor,
		MinFloor: minFloor,
	}
}

// AddRequests adds an internal requests

func (e *Elevator) AddRequests(floor int) {
	e.mutext.Lock()
	defer e.mutext.Unlock()

	if e.MaxFloor > floor || e.MinFloor < floor || floor == e.CurrentFloor {
		return
	}

	// check if request is alread added 
	for _,req := range e.Requests {
		if req == floor {
			return
		}
	}

	e.Requests = append(e.Requests, floor)
	e.sortRequests()
}


func (e *Elevator) Move() {
	e.mutext.Lock()
	defer e.mutext.Unlock()

	if len(e.Requests) == 0 {
		e.Direction = IDLE
		e.State = STOPPED
		return
	}

	nextFloor := e.GetNextFloor()

	if nextFloor = e.CurrentFloor {
		return
	}



}

// 
func (e *Elevator) GetNextFloor() {

}

func (e *Elevator) sortRequests() {
	if e.Direction == UP {
		sort.Ints(e.Requests)
	} else if e.Direction == DOWN {
		sort.Sort(sort.Reverse(sort.IntSlice(e.Requests)))
	}
}


type ElevatorState int 

const (
	MOVING ElevatorState = iota
	STOPPED
	MAINTANANCE
)


func (s ElevatorState) String() string {
	switch s {
	case MOVING:
		return "MOVING"
	case STOPPED:
		return "STOPPED"
	default:
		return "MAINTANANCE"
	}
}


type ElevatorController struct {
	Elevators []*Elevator
	ExternalReqs []Request
	mutext sync.RWMutex
}


func NewElevatorController(numElevators, min, max int ) *ElevatorController {
	elevators := make([]*Elevator, numElevators)

	for i := 0; i<numElevators; i++ {
		elevators[i] = NewElevator(i+1, min, max)
	}

	return &ElevatorController{
		Elevators: elevators,
		ExternalReqs: make([]Request, 0),
	}
}

// ProcessRequest continously process elevator requests
func (ec *ElevatorController) ProcessRequests() {
	for {
		ec.mutext.Lock()
		for _,elevator := ec.Elevators {
			if elevator.state != MAINTANANCE {
				elevator.Move()
			}
		}
		ec.mutext.Unlock()
	}
}


func main() {
	controller := NewElevatorController(3, 1, 10)

	// process the requests 

	go controller.ProcessRequests()
}


// this is the client code 

// 



// elevator system
// 