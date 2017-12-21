package journal
import "time"

type LaunchSRV struct {
	Loadout          string `json:"Loadout"`
	PlayerControlled bool   `json:"PlayerControlled"`
	Event            string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e LaunchSRV) GetEvent() string {
       return e.Event
}

func (e LaunchSRV) GetTimestamp() time.Time {
       return e.Timestamp
}
