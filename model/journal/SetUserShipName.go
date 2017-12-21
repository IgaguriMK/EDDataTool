package journal
import "time"

type SetUserShipName struct {
	Ship         string `json:"Ship"`
	ShipID       int64  `json:"ShipID"`
	UserShipID   string `json:"UserShipId"`
	UserShipName string `json:"UserShipName"`
	Event        string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e SetUserShipName) GetEvent() string {
       return e.Event
}

func (e SetUserShipName) GetTimestamp() time.Time {
       return e.Timestamp
}
