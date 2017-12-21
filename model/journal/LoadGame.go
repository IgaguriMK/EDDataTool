package journal

import "time"

type LoadGame struct {
	Commander    string    `json:"Commander"`
	Credits      int64     `json:"Credits"`
	FuelCapacity float64   `json:"FuelCapacity"`
	FuelLevel    float64   `json:"FuelLevel"`
	GameMode     string    `json:"GameMode"`
	Group        string    `json:"Group"`
	Loan         int64     `json:"Loan"`
	Ship         string    `json:"Ship"`
	ShipID       int64     `json:"ShipID"`
	ShipIdent    string    `json:"ShipIdent"`
	ShipName     string    `json:"ShipName"`
	StartLanded  bool      `json:"StartLanded"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e LoadGame) GetEvent() string {
	return e.Event
}

func (e LoadGame) GetTimestamp() time.Time {
	return e.Timestamp
}
