package journal
import "time"

type ModuleSwap struct {
	FromItem          string `json:"FromItem"`
	FromItemLocalised string `json:"FromItem_Localised"`
	FromSlot          string `json:"FromSlot"`
	Ship              string `json:"Ship"`
	ShipID            int64  `json:"ShipID"`
	ToItem            string `json:"ToItem"`
	ToItemLocalised   string `json:"ToItem_Localised"`
	ToSlot            string `json:"ToSlot"`
	Event             string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ModuleSwap) GetEvent() string {
       return e.Event
}

func (e ModuleSwap) GetTimestamp() time.Time {
       return e.Timestamp
}
