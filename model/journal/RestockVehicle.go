package journal
import "time"

type RestockVehicle struct {
	Cost      int64  `json:"Cost"`
	Count     int64  `json:"Count"`
	Loadout   string `json:"Loadout"`
	Type      string `json:"Type"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e RestockVehicle) GetEvent() string {
       return e.Event
}

func (e RestockVehicle) GetTimestamp() time.Time {
       return e.Timestamp
}
