package journal
import "time"

type Interdicted struct {
	Faction     string `json:"Faction"`
	Interdictor string `json:"Interdictor"`
	IsPlayer    bool   `json:"IsPlayer"`
	Submitted   bool   `json:"Submitted"`
	Event       string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Interdicted) GetEvent() string {
       return e.Event
}

func (e Interdicted) GetTimestamp() time.Time {
       return e.Timestamp
}
