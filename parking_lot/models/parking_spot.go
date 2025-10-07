package models

type SpotType string

const (
	Compact    SpotType = "Compact"
	Large      SpotType = "large"
	MotorCycle SpotType = "MotorCycle"
	Electric   SpotType = "Electric"
)

type ParkingSpot struct {
	ID          string
	SpotType    SpotType
	Vehicle     *Vehicle
	State       SpotState
	FloorNumber int
}

func (p *ParkingSpot) AssignVehicle(vehicle *Vehicle) {
	p.State.AssignVehicle(p, vehicle)
}

func (p *ParkingSpot) RemoveVehicle() {
	p.State.RemoveVehicle(p)
}

func (p *ParkingSpot) IsAvailable() bool {
	return p.State.IsAvailable()
}
