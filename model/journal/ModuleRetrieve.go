package journal
import "time"

type ModuleRetrieve struct {
	Cost                   int64  `json:"Cost"`
	EngineerModifications  string `json:"EngineerModifications"`
	RetrievedItem          string `json:"RetrievedItem"`
	RetrievedItemLocalised string `json:"RetrievedItem_Localised"`
	Ship                   string `json:"Ship"`
	ShipID                 int64  `json:"ShipID"`
	Slot                   string `json:"Slot"`
	SwapOutItem            string `json:"SwapOutItem"`
	SwapOutItemLocalised   string `json:"SwapOutItem_Localised"`
	Event                  string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ModuleRetrieve) GetEvent() string {
       return e.Event
}

func (e ModuleRetrieve) GetTimestamp() time.Time {
       return e.Timestamp
}
