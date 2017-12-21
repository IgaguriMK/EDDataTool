package journal
import "time"

type ShipyardSwap struct {
	ShipID       int64  `json:"ShipID"`
	ShipType     string `json:"ShipType"`
	StoreOldShip string `json:"StoreOldShip"`
	StoreShipID  int64  `json:"StoreShipID"`
	Event        string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ShipyardSwap) GetEvent() string {
       return e.Event
}

func (e ShipyardSwap) GetTimestamp() time.Time {
       return e.Timestamp
}
