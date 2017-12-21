package journal
import "time"

type MiningRefined struct {
	Type          string `json:"Type"`
	TypeLocalised string `json:"Type_Localised"`
	Event         string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MiningRefined) GetEvent() string {
       return e.Event
}

func (e MiningRefined) GetTimestamp() time.Time {
       return e.Timestamp
}
