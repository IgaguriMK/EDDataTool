package journal
import "time"

type ShipyardTransfer struct {
	Distance      float64 `json:"Distance"`
	ShipID        int64   `json:"ShipID"`
	ShipType      string  `json:"ShipType"`
	System        string  `json:"System"`
	TransferPrice int64   `json:"TransferPrice"`
	TransferTime  int64   `json:"TransferTime"`
	Event         string  `json:"event"`
	Timestamp time.Time  `json:"timestamp"`
}
func (e ShipyardTransfer) GetEvent() string {
       return e.Event
}

func (e ShipyardTransfer) GetTimestamp() time.Time {
       return e.Timestamp
}
