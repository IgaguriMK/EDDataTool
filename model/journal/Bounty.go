package journal
import "time"

type Bounty struct {
	Rewards []struct {
		Faction string `json:"Faction"`
		Reward  int64  `json:"Reward"`
	} `json:"Rewards"`
	SharedWithOthers int64  `json:"SharedWithOthers"`
	Target           string `json:"Target"`
	TotalReward      int64  `json:"TotalReward"`
	VictimFaction    string `json:"VictimFaction"`
	Event            string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e Bounty) GetEvent() string {
       return e.Event
}

func (e Bounty) GetTimestamp() time.Time {
       return e.Timestamp
}
