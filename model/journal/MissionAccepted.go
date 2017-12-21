package journal
import "time"

type MissionAccepted struct {
	Commodity          string `json:"Commodity"`
	CommodityLocalised string `json:"Commodity_Localised"`
	Count              int64  `json:"Count"`
	DestinationStation string `json:"DestinationStation"`
	DestinationSystem  string `json:"DestinationSystem"`
	Expiry             string `json:"Expiry"`
	Faction            string `json:"Faction"`
	Influence          string `json:"Influence"`
	KillCount          int64  `json:"KillCount"`
	LocalisedName      string `json:"LocalisedName"`
	MissionID          int64  `json:"MissionID"`
	Name               string `json:"Name"`
	PassengerCount     int64  `json:"PassengerCount"`
	PassengerType      string `json:"PassengerType"`
	PassengerVIPs      bool   `json:"PassengerVIPs"`
	PassengerWanted    bool   `json:"PassengerWanted"`
	Reputation         string `json:"Reputation"`
	Reward             int64  `json:"Reward"`
	TargetFaction      string `json:"TargetFaction"`
	Event              string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e MissionAccepted) GetEvent() string {
       return e.Event
}

func (e MissionAccepted) GetTimestamp() time.Time {
       return e.Timestamp
}
