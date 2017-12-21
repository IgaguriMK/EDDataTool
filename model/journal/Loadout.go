package journal

import "time"

type Loadout struct {
	Modules []struct {
		AmmoInClip        int64   `json:"AmmoInClip"`
		AmmoInHopper      int64   `json:"AmmoInHopper"`
		EngineerBlueprint string  `json:"EngineerBlueprint"`
		EngineerLevel     int64   `json:"EngineerLevel"`
		Health            float64 `json:"Health"`
		Item              string  `json:"Item"`
		On                bool    `json:"On"`
		Priority          int64   `json:"Priority"`
		Slot              string  `json:"Slot"`
		Value             int64   `json:"Value"`
	} `json:"Modules"`
	Ship      string    `json:"Ship"`
	ShipID    int64     `json:"ShipID"`
	ShipIdent string    `json:"ShipIdent"`
	ShipName  string    `json:"ShipName"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Loadout) GetEvent() string {
	return e.Event
}

func (e Loadout) GetTimestamp() time.Time {
	return e.Timestamp
}
