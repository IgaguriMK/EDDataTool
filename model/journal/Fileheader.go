package journal
import "time"

type Fileheader struct {
	Build       string `json:"build"`
	Event       string `json:"event"`
	Gameversion string `json:"gameversion"`
	Language    string `json:"language"`
	Part        int64  `json:"part"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Fileheader) GetEvent() string {
       return e.Event
}

func (e Fileheader) GetTimestamp() time.Time {
       return e.Timestamp
}
