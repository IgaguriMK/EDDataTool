package journal
import "time"

type DatalinkScan struct {
	Message          string `json:"Message"`
	MessageLocalised string `json:"Message_Localised"`
	Event            string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e DatalinkScan) GetEvent() string {
       return e.Event
}

func (e DatalinkScan) GetTimestamp() time.Time {
       return e.Timestamp
}
