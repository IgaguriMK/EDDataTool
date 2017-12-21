package journal
import "time"

type FSDJump struct {
	BoostUsed    int64  `json:"BoostUsed"`
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
	FuelLevel                 float64   `json:"FuelLevel"`
	FuelUsed                  float64   `json:"FuelUsed"`
	JumpDist                  float64   `json:"JumpDist"`
	Population                int64     `json:"Population"`
	StarPos                   []float64 `json:"StarPos"`
	StarSystem                string    `json:"StarSystem"`
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
func (e FSDJump) GetEvent() string {
       return e.Event
}

func (e FSDJump) GetTimestamp() time.Time {
       return e.Timestamp
}
