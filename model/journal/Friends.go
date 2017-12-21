package journal
import "time"

type Friends struct {
	Name      string `json:"Name"`
	Status    string `json:"Status"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Friends) GetEvent() string {
       return e.Event
}

func (e Friends) GetTimestamp() time.Time {
       return e.Timestamp
}
