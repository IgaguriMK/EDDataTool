package journal
import "time"

type CommunityGoalJoin struct {
	Name      string `json:"Name"`
	System    string `json:"System"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e CommunityGoalJoin) GetEvent() string {
       return e.Event
}

func (e CommunityGoalJoin) GetTimestamp() time.Time {
       return e.Timestamp
}
