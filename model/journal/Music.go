package journal
import "time"

type Music struct {
	MusicTrack string `json:"MusicTrack"`
	Event      string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Music) GetEvent() string {
       return e.Event
}

func (e Music) GetTimestamp() time.Time {
       return e.Timestamp
}
