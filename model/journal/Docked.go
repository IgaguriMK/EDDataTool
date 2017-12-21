package journal
import "time"

type Docked struct {
	DistFromStarLS             float64  `json:"DistFromStarLS"`
	FactionState               string   `json:"FactionState"`
	StarSystem                 string   `json:"StarSystem"`
	StationAllegiance          string   `json:"StationAllegiance"`
	StationEconomy             string   `json:"StationEconomy"`
	StationEconomyLocalised    string   `json:"StationEconomy_Localised"`
	StationFaction             string   `json:"StationFaction"`
	StationGovernment          string   `json:"StationGovernment"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised"`
	StationName                string   `json:"StationName"`
	StationServices            []string `json:"StationServices"`
	StationState               string   `json:"StationState"`
	StationType                string   `json:"StationType"`
	Event                      string   `json:"event"`
	Timestamp time.Time   `json:"timestamp"`
}
func (e Docked) GetEvent() string {
       return e.Event
}

func (e Docked) GetTimestamp() time.Time {
       return e.Timestamp
}
