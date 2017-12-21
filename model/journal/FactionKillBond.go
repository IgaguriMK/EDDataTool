package journal
import "time"

type FactionKillBond struct {
	AwardingFaction          string `json:"AwardingFaction"`
	AwardingFactionLocalised string `json:"AwardingFaction_Localised"`
	Reward                   int64  `json:"Reward"`
	VictimFaction            string `json:"VictimFaction"`
	Event                    string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e FactionKillBond) GetEvent() string {
       return e.Event
}

func (e FactionKillBond) GetTimestamp() time.Time {
       return e.Timestamp
}
