package journal
import "time"

type ReceiveText struct {
	Channel          string `json:"Channel"`
	From             string `json:"From"`
	FromLocalised    string `json:"From_Localised"`
	Message          string `json:"Message"`
	MessageLocalised string `json:"Message_Localised"`
	Event            string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ReceiveText) GetEvent() string {
       return e.Event
}

func (e ReceiveText) GetTimestamp() time.Time {
       return e.Timestamp
}
