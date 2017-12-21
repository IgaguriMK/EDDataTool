package journal
import "time"

type CommunityGoal struct {
	CurrentGoals []struct {
		Bonus                int64  `json:"Bonus"`
		Cgid                 int64  `json:"CGID"`
		CurrentTotal         int64  `json:"CurrentTotal"`
		Expiry               string `json:"Expiry"`
		IsComplete           bool   `json:"IsComplete"`
		MarketName           string `json:"MarketName"`
		NumContributors      int64  `json:"NumContributors"`
		PlayerContribution   int64  `json:"PlayerContribution"`
		PlayerInTopRank      bool   `json:"PlayerInTopRank"`
		PlayerPercentileBand int64  `json:"PlayerPercentileBand"`
		SystemName           string `json:"SystemName"`
		TierReached          string `json:"TierReached"`
		Title                string `json:"Title"`
		TopRankSize          int64  `json:"TopRankSize"`
	} `json:"CurrentGoals"`
	Event     string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e CommunityGoal) GetEvent() string {
       return e.Event
}

func (e CommunityGoal) GetTimestamp() time.Time {
       return e.Timestamp
}
