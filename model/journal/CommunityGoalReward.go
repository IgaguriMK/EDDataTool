package journal
import "time"

type CommunityGoalReward struct {
	Name      string `json:"Name"`
	Reward    int64  `json:"Reward"`
	System    string `json:"System"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e CommunityGoalReward) GetEvent() string {
       return e.Event
}

func (e CommunityGoalReward) GetTimestamp() time.Time {
       return e.Timestamp
}
