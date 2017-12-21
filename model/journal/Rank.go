package journal
import "time"

type Rank struct {
	Cqc        int64  `json:"CQC"`
	Combat     int64  `json:"Combat"`
	Empire     int64  `json:"Empire"`
	Explore    int64  `json:"Explore"`
	Federation int64  `json:"Federation"`
	Trade      int64  `json:"Trade"`
	Event      string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Rank) GetEvent() string {
       return e.Event
}

func (e Rank) GetTimestamp() time.Time {
       return e.Timestamp
}
