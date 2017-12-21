package journal
import "time"

type Location struct {
	Body         string `json:"Body"`
	BodyType     string `json:"BodyType"`
	Docked       bool   `json:"Docked"`
	FactionState string `json:"FactionState"`
	Factions     []struct {
		Allegiance    string  `json:"Allegiance"`
		FactionState  string  `json:"FactionState"`
		Government    string  `json:"Government"`
		Influence     float64 `json:"Influence"`
		Name          string  `json:"Name"`
		PendingStates []struct {
			State string `json:"State"`
			Trend int64  `json:"Trend"`
		} `json:"PendingStates"`
		RecoveringStates []struct {
			State string `json:"State"`
			Trend int64  `json:"Trend"`
		} `json:"RecoveringStates"`
	} `json:"Factions"`
	Latitude                  float64   `json:"Latitude"`
	Longitude                 float64   `json:"Longitude"`
	Population                int64     `json:"Population"`
	StarPos                   []float64 `json:"StarPos"`
	StarSystem                string    `json:"StarSystem"`
	StationName               string    `json:"StationName"`
	StationType               string    `json:"StationType"`
	SystemAllegiance          string    `json:"SystemAllegiance"`
	SystemEconomy             string    `json:"SystemEconomy"`
	SystemEconomyLocalised    string    `json:"SystemEconomy_Localised"`
	SystemFaction             string    `json:"SystemFaction"`
	SystemGovernment          string    `json:"SystemGovernment"`
	SystemGovernmentLocalised string    `json:"SystemGovernment_Localised"`
	SystemSecurity            string    `json:"SystemSecurity"`
	SystemSecurityLocalised   string    `json:"SystemSecurity_Localised"`
	Event                     string    `json:"event"`
	Timestamp time.Time    `json:"timestamp"`
}
func (e Location) GetEvent() string {
       return e.Event
}

func (e Location) GetTimestamp() time.Time {
       return e.Timestamp
}
