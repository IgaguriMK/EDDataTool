package journal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type Event interface {
	GetEvent() string
	GetTimestamp() time.Time
}

const InvalidEvent = "!!INVALID!!"

func ParseEvent(jsonStr string) (Event, error) {
	jsonBytes := []byte(jsonStr)

	eventType := struct {
		Event string `json:"event"`
	}{
		Event: InvalidEvent,
	}

	err := json.Unmarshal(jsonBytes, &eventType)
	if err != nil {
		return nil, errors.Wrap(err, "Failed get event type")
	}

	event, err := parseWithType(jsonBytes, eventType.Event)
	if err != nil {
		return nil, errors.Wrap(err, "Failed parse event")
	}

	return event, nil
}

func parseWithType(bytes []byte, eventType string) (Event, error) {
	switch eventType {
	case "AfmuRepairs":
		var e AfmuRepairs
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ApproachSettlement":
		var e ApproachSettlement
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Bounty":
		var e Bounty
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "BuyAmmo":
		var e BuyAmmo
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "BuyDrones":
		var e BuyDrones
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "BuyExplorationData":
		var e BuyExplorationData
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "BuyTradeData":
		var e BuyTradeData
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Cargo":
		var e Cargo
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "CockpitBreached":
		var e CockpitBreached
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "CollectCargo":
		var e CollectCargo
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "CommitCrime":
		var e CommitCrime
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "CommunityGoal":
		var e CommunityGoal
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "CommunityGoalJoin":
		var e CommunityGoalJoin
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "CommunityGoalReward":
		var e CommunityGoalReward
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DatalinkScan":
		var e DatalinkScan
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DataScanned":
		var e DataScanned
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Died":
		var e Died
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Docked":
		var e Docked
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DockingCancelled":
		var e DockingCancelled
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DockingDenied":
		var e DockingDenied
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DockingGranted":
		var e DockingGranted
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DockingRequested":
		var e DockingRequested
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "DockSRV":
		var e DockSRV
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "EjectCargo":
		var e EjectCargo
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "EngineerApply":
		var e EngineerApply
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "EngineerContribution":
		var e EngineerContribution
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "EngineerCraft":
		var e EngineerCraft
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "EngineerProgress":
		var e EngineerProgress
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "EscapeInterdiction":
		var e EscapeInterdiction
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "FactionKillBond":
		var e FactionKillBond
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "FetchRemoteModule":
		var e FetchRemoteModule
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Fileheader":
		var e Fileheader
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Friends":
		var e Friends
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "FSDJump":
		var e FSDJump
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "FuelScoop":
		var e FuelScoop
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "HeatDamage":
		var e HeatDamage
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "HeatWarning":
		var e HeatWarning
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "HullDamage":
		var e HullDamage
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Interdicted":
		var e Interdicted
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "JetConeBoost":
		var e JetConeBoost
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "LaunchSRV":
		var e LaunchSRV
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Liftoff":
		var e Liftoff
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "LoadGame":
		var e LoadGame
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Loadout":
		var e Loadout
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Location":
		var e Location
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MarketBuy":
		var e MarketBuy
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MarketSell":
		var e MarketSell
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MaterialCollected":
		var e MaterialCollected
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MaterialDiscarded":
		var e MaterialDiscarded
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MaterialDiscovered":
		var e MaterialDiscovered
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Materials":
		var e Materials
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MiningRefined":
		var e MiningRefined
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MissionAbandoned":
		var e MissionAbandoned
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MissionAccepted":
		var e MissionAccepted
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MissionCompleted":
		var e MissionCompleted
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MissionFailed":
		var e MissionFailed
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "MissionRedirected":
		var e MissionRedirected
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ModuleBuy":
		var e ModuleBuy
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ModuleRetrieve":
		var e ModuleRetrieve
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ModuleSell":
		var e ModuleSell
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ModuleSellRemote":
		var e ModuleSellRemote
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ModuleStore":
		var e ModuleStore
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ModuleSwap":
		var e ModuleSwap
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Music":
		var e Music
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "NavBeaconScan":
		var e NavBeaconScan
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Passengers":
		var e Passengers
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "PayFines":
		var e PayFines
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Progress":
		var e Progress
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Promotion":
		var e Promotion
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Rank":
		var e Rank
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RebootRepair":
		var e RebootRepair
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ReceiveText":
		var e ReceiveText
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RedeemVoucher":
		var e RedeemVoucher
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RefuelAll":
		var e RefuelAll
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Repair":
		var e Repair
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RepairAll":
		var e RepairAll
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "RestockVehicle":
		var e RestockVehicle
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Resurrect":
		var e Resurrect
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Scan":
		var e Scan
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Scanned":
		var e Scanned
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Screenshot":
		var e Screenshot
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SelfDestruct":
		var e SelfDestruct
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SellDrones":
		var e SellDrones
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SellExplorationData":
		var e SellExplorationData
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SendText":
		var e SendText
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SetUserShipName":
		var e SetUserShipName
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShieldState":
		var e ShieldState
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardBuy":
		var e ShipyardBuy
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardNew":
		var e ShipyardNew
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardSell":
		var e ShipyardSell
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardSwap":
		var e ShipyardSwap
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "ShipyardTransfer":
		var e ShipyardTransfer
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "StartJump":
		var e StartJump
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SupercruiseEntry":
		var e SupercruiseEntry
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "SupercruiseExit":
		var e SupercruiseExit
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Synthesis":
		var e Synthesis
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Touchdown":
		var e Touchdown
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "Undocked":
		var e Undocked
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "USSDrop":
		var e USSDrop
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingAdd":
		var e WingAdd
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingInvite":
		var e WingInvite
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingJoin":
		var e WingJoin
		err := json.Unmarshal(bytes, &e)
		return e, err
	case "WingLeave":
		var e WingLeave
		err := json.Unmarshal(bytes, &e)
		return e, err

		///////////////////////////////////////////////////////////
	}

	return nil, &UnknownEventType{Raw: string(bytes), Type: eventType}
}

type UnknownEventType struct {
	Type string
	Raw  string
}

func (e *UnknownEventType) Error() string {
	return fmt.Sprintf("Unknown event type [%s]: %s", e.Type, e.Raw)
}

////////////////////////////////////////////////////////////////////////////////

type AfmuRepairs struct {
	FullyRepaired   bool      `json:"FullyRepaired"`
	Health          float64   `json:"Health"`
	Module          string    `json:"Module"`
	ModuleLocalised string    `json:"Module_Localised"`
	Event           string    `json:"event"`
	Timestamp       time.Time `json:"timestamp"`
}

func (e AfmuRepairs) GetEvent() string {
	return e.Event
}

func (e AfmuRepairs) GetTimestamp() time.Time {
	return e.Timestamp
}

type ApproachSettlement struct {
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ApproachSettlement) GetEvent() string {
	return e.Event
}

func (e ApproachSettlement) GetTimestamp() time.Time {
	return e.Timestamp
}

type Bounty struct {
	Rewards          []BountyFaction `json:"Rewards"`
	SharedWithOthers int64           `json:"SharedWithOthers"`
	Target           string          `json:"Target"`
	TotalReward      int64           `json:"TotalReward"`
	VictimFaction    string          `json:"VictimFaction"`
	Event            string          `json:"event"`
	Timestamp        time.Time       `json:"timestamp"`
}

type BountyFaction struct {
	Faction string `json:"Faction"`
	Reward  int64  `json:"Reward"`
}

func (e Bounty) GetEvent() string {
	return e.Event
}

func (e Bounty) GetTimestamp() time.Time {
	return e.Timestamp
}

type BuyAmmo struct {
	Cost      int64     `json:"Cost"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyAmmo) GetEvent() string {
	return e.Event
}

func (e BuyAmmo) GetTimestamp() time.Time {
	return e.Timestamp
}

type BuyDrones struct {
	BuyPrice  int64     `json:"BuyPrice"`
	Count     int64     `json:"Count"`
	TotalCost int64     `json:"TotalCost"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyDrones) GetEvent() string {
	return e.Event
}

func (e BuyDrones) GetTimestamp() time.Time {
	return e.Timestamp
}

type BuyExplorationData struct {
	Cost      int64     `json:"Cost"`
	System    string    `json:"System"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyExplorationData) GetEvent() string {
	return e.Event
}

func (e BuyExplorationData) GetTimestamp() time.Time {
	return e.Timestamp
}

type BuyTradeData struct {
	Cost      int64     `json:"Cost"`
	System    string    `json:"System"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e BuyTradeData) GetEvent() string {
	return e.Event
}

func (e BuyTradeData) GetTimestamp() time.Time {
	return e.Timestamp
}

type Cargo struct {
	Inventory []interface{} `json:"Inventory"`
	Event     string        `json:"event"`
	Timestamp time.Time     `json:"timestamp"`
}

func (e Cargo) GetEvent() string {
	return e.Event
}

func (e Cargo) GetTimestamp() time.Time {
	return e.Timestamp
}

type CockpitBreached struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CockpitBreached) GetEvent() string {
	return e.Event
}

func (e CockpitBreached) GetTimestamp() time.Time {
	return e.Timestamp
}

type CollectCargo struct {
	Stolen    bool      `json:"Stolen"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CollectCargo) GetEvent() string {
	return e.Event
}

func (e CollectCargo) GetTimestamp() time.Time {
	return e.Timestamp
}

type CommitCrime struct {
	CrimeType string    `json:"CrimeType"`
	Faction   string    `json:"Faction"`
	Fine      int64     `json:"Fine"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CommitCrime) GetEvent() string {
	return e.Event
}

func (e CommitCrime) GetTimestamp() time.Time {
	return e.Timestamp
}

type CommunityGoal struct {
	CurrentGoals []CommunityGoalDetail `json:"CurrentGoals"`
	Event        string                `json:"event"`
	Timestamp    time.Time             `json:"timestamp"`
}

type CommunityGoalDetail struct {
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
}

func (e CommunityGoal) GetEvent() string {
	return e.Event
}

func (e CommunityGoal) GetTimestamp() time.Time {
	return e.Timestamp
}

type CommunityGoalJoin struct {
	Name      string    `json:"Name"`
	System    string    `json:"System"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CommunityGoalJoin) GetEvent() string {
	return e.Event
}

func (e CommunityGoalJoin) GetTimestamp() time.Time {
	return e.Timestamp
}

type CommunityGoalReward struct {
	Name      string    `json:"Name"`
	Reward    int64     `json:"Reward"`
	System    string    `json:"System"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e CommunityGoalReward) GetEvent() string {
	return e.Event
}

func (e CommunityGoalReward) GetTimestamp() time.Time {
	return e.Timestamp
}

type DatalinkScan struct {
	Message          string    `json:"Message"`
	MessageLocalised string    `json:"Message_Localised"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e DatalinkScan) GetEvent() string {
	return e.Event
}

func (e DatalinkScan) GetTimestamp() time.Time {
	return e.Timestamp
}

type DataScanned struct {
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e DataScanned) GetEvent() string {
	return e.Event
}

func (e DataScanned) GetTimestamp() time.Time {
	return e.Timestamp
}

type Died struct {
	KillerName          string    `json:"KillerName"`
	KillerNameLocalised string    `json:"KillerName_Localised"`
	KillerRank          string    `json:"KillerRank"`
	KillerShip          string    `json:"KillerShip"`
	Event               string    `json:"event"`
	Timestamp           time.Time `json:"timestamp"`
}

func (e Died) GetEvent() string {
	return e.Event
}

func (e Died) GetTimestamp() time.Time {
	return e.Timestamp
}

type Docked struct {
	DistFromStarLS             float64   `json:"DistFromStarLS"`
	FactionState               string    `json:"FactionState"`
	StarSystem                 string    `json:"StarSystem"`
	StationEconomy             string    `json:"StationEconomy"`
	StationEconomyLocalised    string    `json:"StationEconomy_Localised"`
	StationFaction             string    `json:"StationFaction"`
	StationGovernment          string    `json:"StationGovernment"`
	StationGovernmentLocalised string    `json:"StationGovernment_Localised"`
	StationName                string    `json:"StationName"`
	StationServices            []string  `json:"StationServices"`
	StationType                string    `json:"StationType"`
	Event                      string    `json:"event"`
	Timestamp                  time.Time `json:"timestamp"`
}

func (e Docked) GetEvent() string {
	return e.Event
}

func (e Docked) GetTimestamp() time.Time {
	return e.Timestamp
}

type DockingCancelled struct {
	StationName string    `json:"StationName"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingCancelled) GetEvent() string {
	return e.Event
}

func (e DockingCancelled) GetTimestamp() time.Time {
	return e.Timestamp
}

type DockingDenied struct {
	Reason      string    `json:"Reason"`
	StationName string    `json:"StationName"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingDenied) GetEvent() string {
	return e.Event
}

func (e DockingDenied) GetTimestamp() time.Time {
	return e.Timestamp
}

type DockingGranted struct {
	LandingPad  int64     `json:"LandingPad"`
	StationName string    `json:"StationName"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingGranted) GetEvent() string {
	return e.Event
}

func (e DockingGranted) GetTimestamp() time.Time {
	return e.Timestamp
}

type DockingRequested struct {
	StationName string    `json:"StationName"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e DockingRequested) GetEvent() string {
	return e.Event
}

func (e DockingRequested) GetTimestamp() time.Time {
	return e.Timestamp
}

type DockSRV struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e DockSRV) GetEvent() string {
	return e.Event
}

func (e DockSRV) GetTimestamp() time.Time {
	return e.Timestamp
}

type EjectCargo struct {
	Abandoned bool      `json:"Abandoned"`
	Count     int64     `json:"Count"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EjectCargo) GetEvent() string {
	return e.Event
}

func (e EjectCargo) GetTimestamp() time.Time {
	return e.Timestamp
}

type EngineerApply struct {
	Blueprint string    `json:"Blueprint"`
	Engineer  string    `json:"Engineer"`
	Level     int64     `json:"Level"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EngineerApply) GetEvent() string {
	return e.Event
}

func (e EngineerApply) GetTimestamp() time.Time {
	return e.Timestamp
}

type EngineerContribution struct {
	Commodity     string    `json:"Commodity"`
	Engineer      string    `json:"Engineer"`
	Quantity      int64     `json:"Quantity"`
	TotalQuantity int64     `json:"TotalQuantity"`
	Type          string    `json:"Type"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e EngineerContribution) GetEvent() string {
	return e.Event
}

func (e EngineerContribution) GetTimestamp() time.Time {
	return e.Timestamp
}

type EngineerCraft struct {
	Blueprint   string               `json:"Blueprint"`
	Engineer    string               `json:"Engineer"`
	Ingredients []EngineerCraft_sub1 `json:"Ingredients"`
	Level       int64                `json:"Level"`
	Event       string               `json:"event"`
	Timestamp   time.Time            `json:"timestamp"`
}

type EngineerCraft_sub1 struct {
	Count int64  `json:"Count"`
	Name  string `json:"Name"`
}

func (e EngineerCraft) GetEvent() string {
	return e.Event
}

func (e EngineerCraft) GetTimestamp() time.Time {
	return e.Timestamp
}

type EngineerProgress struct {
	Engineer  string    `json:"Engineer"`
	Rank      int64     `json:"Rank"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EngineerProgress) GetEvent() string {
	return e.Event
}

func (e EngineerProgress) GetTimestamp() time.Time {
	return e.Timestamp
}

type EscapeInterdiction struct {
	Interdictor string    `json:"Interdictor"`
	IsPlayer    bool      `json:"IsPlayer"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e EscapeInterdiction) GetEvent() string {
	return e.Event
}

func (e EscapeInterdiction) GetTimestamp() time.Time {
	return e.Timestamp
}

type FactionKillBond struct {
	AwardingFaction string    `json:"AwardingFaction"`
	Reward          int64     `json:"Reward"`
	VictimFaction   string    `json:"VictimFaction"`
	Event           string    `json:"event"`
	Timestamp       time.Time `json:"timestamp"`
}

func (e FactionKillBond) GetEvent() string {
	return e.Event
}

func (e FactionKillBond) GetTimestamp() time.Time {
	return e.Timestamp
}

type FetchRemoteModule struct {
	ServerID            int64     `json:"ServerId"`
	Ship                string    `json:"Ship"`
	ShipID              int64     `json:"ShipID"`
	StorageSlot         int64     `json:"StorageSlot"`
	StoredItem          string    `json:"StoredItem"`
	StoredItemLocalised string    `json:"StoredItem_Localised"`
	TransferCost        int64     `json:"TransferCost"`
	TransferTime        int64     `json:"TransferTime"`
	Event               string    `json:"event"`
	Timestamp           time.Time `json:"timestamp"`
}

func (e FetchRemoteModule) GetEvent() string {
	return e.Event
}

func (e FetchRemoteModule) GetTimestamp() time.Time {
	return e.Timestamp
}

type Fileheader struct {
	Build       string    `json:"build"`
	Event       string    `json:"event"`
	Gameversion string    `json:"gameversion"`
	Language    string    `json:"language"`
	Part        int64     `json:"part"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e Fileheader) GetEvent() string {
	return e.Event
}

func (e Fileheader) GetTimestamp() time.Time {
	return e.Timestamp
}

type Friends struct {
	Name      string    `json:"Name"`
	Status    string    `json:"Status"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Friends) GetEvent() string {
	return e.Event
}

func (e Friends) GetTimestamp() time.Time {
	return e.Timestamp
}

type FSDJump struct {
	FuelLevel                 float64   `json:"FuelLevel"`
	FuelUsed                  float64   `json:"FuelUsed"`
	JumpDist                  float64   `json:"JumpDist"`
	Population                int64     `json:"Population"`
	StarPos                   []float64 `json:"StarPos"`
	StarSystem                string    `json:"StarSystem"`
	SystemAllegiance          string    `json:"SystemAllegiance"`
	SystemEconomy             string    `json:"SystemEconomy"`
	SystemEconomyLocalised    string    `json:"SystemEconomy_Localised"`
	SystemGovernment          string    `json:"SystemGovernment"`
	SystemGovernmentLocalised string    `json:"SystemGovernment_Localised"`
	SystemSecurity            string    `json:"SystemSecurity"`
	SystemSecurityLocalised   string    `json:"SystemSecurity_Localised"`
	Event                     string    `json:"event"`
	Timestamp                 time.Time `json:"timestamp"`
}

func (e FSDJump) GetEvent() string {
	return e.Event
}

func (e FSDJump) GetTimestamp() time.Time {
	return e.Timestamp
}

type FuelScoop struct {
	Scooped   float64   `json:"Scooped"`
	Total     float64   `json:"Total"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e FuelScoop) GetEvent() string {
	return e.Event
}

func (e FuelScoop) GetTimestamp() time.Time {
	return e.Timestamp
}

type HeatDamage struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e HeatDamage) GetEvent() string {
	return e.Event
}

func (e HeatDamage) GetTimestamp() time.Time {
	return e.Timestamp
}

type HeatWarning struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e HeatWarning) GetEvent() string {
	return e.Event
}

func (e HeatWarning) GetTimestamp() time.Time {
	return e.Timestamp
}

type HullDamage struct {
	Fighter     bool      `json:"Fighter"`
	Health      float64   `json:"Health"`
	PlayerPilot bool      `json:"PlayerPilot"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e HullDamage) GetEvent() string {
	return e.Event
}

func (e HullDamage) GetTimestamp() time.Time {
	return e.Timestamp
}

type Interdicted struct {
	Faction     string    `json:"Faction"`
	Interdictor string    `json:"Interdictor"`
	IsPlayer    bool      `json:"IsPlayer"`
	Submitted   bool      `json:"Submitted"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e Interdicted) GetEvent() string {
	return e.Event
}

func (e Interdicted) GetTimestamp() time.Time {
	return e.Timestamp
}

type JetConeBoost struct {
	BoostValue float64   `json:"BoostValue"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e JetConeBoost) GetEvent() string {
	return e.Event
}

func (e JetConeBoost) GetTimestamp() time.Time {
	return e.Timestamp
}

type LaunchSRV struct {
	Loadout          string    `json:"Loadout"`
	PlayerControlled bool      `json:"PlayerControlled"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e LaunchSRV) GetEvent() string {
	return e.Event
}

func (e LaunchSRV) GetTimestamp() time.Time {
	return e.Timestamp
}

type Liftoff struct {
	Latitude         float64   `json:"Latitude"`
	Longitude        float64   `json:"Longitude"`
	PlayerControlled bool      `json:"PlayerControlled"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e Liftoff) GetEvent() string {
	return e.Event
}

func (e Liftoff) GetTimestamp() time.Time {
	return e.Timestamp
}

type LoadGame struct {
	Commander    string    `json:"Commander"`
	Credits      int64     `json:"Credits"`
	FuelCapacity float64   `json:"FuelCapacity"`
	FuelLevel    float64   `json:"FuelLevel"`
	GameMode     string    `json:"GameMode"`
	Loan         int64     `json:"Loan"`
	Ship         string    `json:"Ship"`
	ShipID       int64     `json:"ShipID"`
	ShipIdent    string    `json:"ShipIdent"`
	ShipName     string    `json:"ShipName"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e LoadGame) GetEvent() string {
	return e.Event
}

func (e LoadGame) GetTimestamp() time.Time {
	return e.Timestamp
}

type Loadout struct {
	Modules   []Module  `json:"Modules"`
	Ship      string    `json:"Ship"`
	ShipID    int64     `json:"ShipID"`
	ShipIdent string    `json:"ShipIdent"`
	ShipName  string    `json:"ShipName"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

type Module struct {
	AmmoInClip   int64   `json:"AmmoInClip"`
	AmmoInHopper int64   `json:"AmmoInHopper"`
	Health       float64 `json:"Health"`
	Item         string  `json:"Item"`
	On           bool    `json:"On"`
	Priority     int64   `json:"Priority"`
	Slot         string  `json:"Slot"`
	Value        int64   `json:"Value"`
}

func (e Loadout) GetEvent() string {
	return e.Event
}

func (e Loadout) GetTimestamp() time.Time {
	return e.Timestamp
}

type Location struct {
	Body                      string            `json:"Body"`
	BodyType                  string            `json:"BodyType"`
	Docked                    bool              `json:"Docked"`
	FactionState              string            `json:"FactionState"`
	Factions                  []LocationFaction `json:"Factions"`
	Population                int64             `json:"Population"`
	StarPos                   []float64         `json:"StarPos"`
	StarSystem                string            `json:"StarSystem"`
	StationName               string            `json:"StationName"`
	StationType               string            `json:"StationType"`
	SystemAllegiance          string            `json:"SystemAllegiance"`
	SystemEconomy             string            `json:"SystemEconomy"`
	SystemEconomyLocalised    string            `json:"SystemEconomy_Localised"`
	SystemFaction             string            `json:"SystemFaction"`
	SystemGovernment          string            `json:"SystemGovernment"`
	SystemGovernmentLocalised string            `json:"SystemGovernment_Localised"`
	SystemSecurity            string            `json:"SystemSecurity"`
	SystemSecurityLocalised   string            `json:"SystemSecurity_Localised"`
	Event                     string            `json:"event"`
	Timestamp                 time.Time         `json:"timestamp"`
}

type LocationFaction struct {
	Allegiance    string                 `json:"Allegiance"`
	FactionState  string                 `json:"FactionState"`
	Government    string                 `json:"Government"`
	Influence     float64                `json:"Influence"`
	Name          string                 `json:"Name"`
	PendingStates []LocationFactionState `json:"PendingStates"`
}

type LocationFactionState struct {
	State string `json:"State"`
	Trend int64  `json:"Trend"`
}

func (e Location) GetEvent() string {
	return e.Event
}

func (e Location) GetTimestamp() time.Time {
	return e.Timestamp
}

type MarketBuy struct {
	BuyPrice  int64     `json:"BuyPrice"`
	Count     int64     `json:"Count"`
	TotalCost int64     `json:"TotalCost"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MarketBuy) GetEvent() string {
	return e.Event
}

func (e MarketBuy) GetTimestamp() time.Time {
	return e.Timestamp
}

type MarketSell struct {
	AvgPricePaid int64     `json:"AvgPricePaid"`
	Count        int64     `json:"Count"`
	SellPrice    int64     `json:"SellPrice"`
	TotalSale    int64     `json:"TotalSale"`
	Type         string    `json:"Type"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e MarketSell) GetEvent() string {
	return e.Event
}

func (e MarketSell) GetTimestamp() time.Time {
	return e.Timestamp
}

type MaterialCollected struct {
	Category  string    `json:"Category"`
	Count     int64     `json:"Count"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MaterialCollected) GetEvent() string {
	return e.Event
}

func (e MaterialCollected) GetTimestamp() time.Time {
	return e.Timestamp
}

type MaterialDiscarded struct {
	Category  string    `json:"Category"`
	Count     int64     `json:"Count"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MaterialDiscarded) GetEvent() string {
	return e.Event
}

func (e MaterialDiscarded) GetTimestamp() time.Time {
	return e.Timestamp
}

type MaterialDiscovered struct {
	Category        string    `json:"Category"`
	DiscoveryNumber int64     `json:"DiscoveryNumber"`
	Name            string    `json:"Name"`
	Event           string    `json:"event"`
	Timestamp       time.Time `json:"timestamp"`
}

func (e MaterialDiscovered) GetEvent() string {
	return e.Event
}

func (e MaterialDiscovered) GetTimestamp() time.Time {
	return e.Timestamp
}

type Materials struct {
	Encoded      []MaterialsCount `json:"Encoded"`
	Manufactured []MaterialsCount `json:"Manufactured"`
	Raw          []MaterialsCount `json:"Raw"`
	Event        string           `json:"event"`
	Timestamp    time.Time        `json:"timestamp"`
}

type MaterialsCount struct {
	Count int64  `json:"Count"`
	Name  string `json:"Name"`
}

func (e Materials) GetEvent() string {
	return e.Event
}

func (e Materials) GetTimestamp() time.Time {
	return e.Timestamp
}

type MiningRefined struct {
	Type          string    `json:"Type"`
	TypeLocalised string    `json:"Type_Localised"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e MiningRefined) GetEvent() string {
	return e.Event
}

func (e MiningRefined) GetTimestamp() time.Time {
	return e.Timestamp
}

type MissionAbandoned struct {
	MissionID int64     `json:"MissionID"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MissionAbandoned) GetEvent() string {
	return e.Event
}

func (e MissionAbandoned) GetTimestamp() time.Time {
	return e.Timestamp
}

type MissionAccepted struct {
	DestinationStation string    `json:"DestinationStation"`
	DestinationSystem  string    `json:"DestinationSystem"`
	Expiry             time.Time `json:"Expiry"`
	Faction            string    `json:"Faction"`
	Influence          string    `json:"Influence"`
	LocalisedName      string    `json:"LocalisedName"`
	MissionID          int       `json:"MissionID"`
	Name               string    `json:"Name"`
	PassengerCount     int       `json:"PassengerCount"`
	PassengerType      string    `json:"PassengerType"`
	PassengerVIPs      bool      `json:"PassengerVIPs"`
	PassengerWanted    bool      `json:"PassengerWanted"`
	Reputation         string    `json:"Reputation"`
	Reward             int       `json:"Reward"`
	Event              string    `json:"event"`
	Timestamp          time.Time `json:"timestamp"`
}

func (e MissionAccepted) GetEvent() string {
	return e.Event
}

func (e MissionAccepted) GetTimestamp() time.Time {
	return e.Timestamp
}

type MissionCompleted struct {
	DestinationStation string    `json:"DestinationStation"`
	DestinationSystem  string    `json:"DestinationSystem"`
	Faction            string    `json:"Faction"`
	MissionID          int64     `json:"MissionID"`
	Name               string    `json:"Name"`
	Reward             int64     `json:"Reward"`
	Event              string    `json:"event"`
	Timestamp          time.Time `json:"timestamp"`
}

func (e MissionCompleted) GetEvent() string {
	return e.Event
}

func (e MissionCompleted) GetTimestamp() time.Time {
	return e.Timestamp
}

type MissionFailed struct {
	MissionID int64     `json:"MissionID"`
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e MissionFailed) GetEvent() string {
	return e.Event
}

func (e MissionFailed) GetTimestamp() time.Time {
	return e.Timestamp
}

type MissionRedirected struct {
	MissionID             int64     `json:"MissionID"`
	Name                  string    `json:"Name"`
	NewDestinationStation string    `json:"NewDestinationStation"`
	NewDestinationSystem  string    `json:"NewDestinationSystem"`
	OldDestinationStation string    `json:"OldDestinationStation"`
	OldDestinationSystem  string    `json:"OldDestinationSystem"`
	Event                 string    `json:"event"`
	Timestamp             time.Time `json:"timestamp"`
}

func (e MissionRedirected) GetEvent() string {
	return e.Event
}

func (e MissionRedirected) GetTimestamp() time.Time {
	return e.Timestamp
}

type ModuleBuy struct {
	BuyItem          string    `json:"BuyItem"`
	BuyItemLocalised string    `json:"BuyItem_Localised"`
	BuyPrice         int64     `json:"BuyPrice"`
	Ship             string    `json:"Ship"`
	ShipID           int64     `json:"ShipID"`
	Slot             string    `json:"Slot"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e ModuleBuy) GetEvent() string {
	return e.Event
}

func (e ModuleBuy) GetTimestamp() time.Time {
	return e.Timestamp
}

type ModuleRetrieve struct {
	RetrievedItem          string    `json:"RetrievedItem"`
	RetrievedItemLocalised string    `json:"RetrievedItem_Localised"`
	Ship                   string    `json:"Ship"`
	ShipID                 int64     `json:"ShipID"`
	Slot                   string    `json:"Slot"`
	Event                  string    `json:"event"`
	Timestamp              time.Time `json:"timestamp"`
}

func (e ModuleRetrieve) GetEvent() string {
	return e.Event
}

func (e ModuleRetrieve) GetTimestamp() time.Time {
	return e.Timestamp
}

type ModuleSell struct {
	SellItem          string    `json:"SellItem"`
	SellItemLocalised string    `json:"SellItem_Localised"`
	SellPrice         int64     `json:"SellPrice"`
	Ship              string    `json:"Ship"`
	ShipID            int64     `json:"ShipID"`
	Slot              string    `json:"Slot"`
	Event             string    `json:"event"`
	Timestamp         time.Time `json:"timestamp"`
}

func (e ModuleSell) GetEvent() string {
	return e.Event
}

func (e ModuleSell) GetTimestamp() time.Time {
	return e.Timestamp
}

type ModuleSellRemote struct {
	SellItem          string    `json:"SellItem"`
	SellItemLocalised string    `json:"SellItem_Localised"`
	SellPrice         int64     `json:"SellPrice"`
	ServerID          int64     `json:"ServerId"`
	Ship              string    `json:"Ship"`
	ShipID            int64     `json:"ShipID"`
	StorageSlot       int64     `json:"StorageSlot"`
	Event             string    `json:"event"`
	Timestamp         time.Time `json:"timestamp"`
}

func (e ModuleSellRemote) GetEvent() string {
	return e.Event
}

func (e ModuleSellRemote) GetTimestamp() time.Time {
	return e.Timestamp
}

type ModuleStore struct {
	Ship                string    `json:"Ship"`
	ShipID              int64     `json:"ShipID"`
	Slot                string    `json:"Slot"`
	StoredItem          string    `json:"StoredItem"`
	StoredItemLocalised string    `json:"StoredItem_Localised"`
	Event               string    `json:"event"`
	Timestamp           time.Time `json:"timestamp"`
}

func (e ModuleStore) GetEvent() string {
	return e.Event
}

func (e ModuleStore) GetTimestamp() time.Time {
	return e.Timestamp
}

type ModuleSwap struct {
	FromItem          string    `json:"FromItem"`
	FromItemLocalised string    `json:"FromItem_Localised"`
	FromSlot          string    `json:"FromSlot"`
	Ship              string    `json:"Ship"`
	ShipID            int64     `json:"ShipID"`
	ToItem            string    `json:"ToItem"`
	ToItemLocalised   string    `json:"ToItem_Localised"`
	ToSlot            string    `json:"ToSlot"`
	Event             string    `json:"event"`
	Timestamp         time.Time `json:"timestamp"`
}

func (e ModuleSwap) GetEvent() string {
	return e.Event
}

func (e ModuleSwap) GetTimestamp() time.Time {
	return e.Timestamp
}

type Music struct {
	MusicTrack string    `json:"MusicTrack"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Music) GetEvent() string {
	return e.Event
}

func (e Music) GetTimestamp() time.Time {
	return e.Timestamp
}

type NavBeaconScan struct {
	NumBodies int64     `json:"NumBodies"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e NavBeaconScan) GetEvent() string {
	return e.Event
}

func (e NavBeaconScan) GetTimestamp() time.Time {
	return e.Timestamp
}

type Passengers struct {
	Manifest  []PassengersManifest `json:"Manifest"`
	Event     string               `json:"event"`
	Timestamp time.Time            `json:"timestamp"`
}

type PassengersManifest struct {
	Count     int64  `json:"Count"`
	MissionID int64  `json:"MissionID"`
	Type      string `json:"Type"`
	Vip       bool   `json:"VIP"`
	Wanted    bool   `json:"Wanted"`
}

func (e Passengers) GetEvent() string {
	return e.Event
}

func (e Passengers) GetTimestamp() time.Time {
	return e.Timestamp
}

type PayFines struct {
	Amount    int64     `json:"Amount"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e PayFines) GetEvent() string {
	return e.Event
}

func (e PayFines) GetTimestamp() time.Time {
	return e.Timestamp
}

type Progress struct {
	Cqc        int64     `json:"CQC"`
	Combat     int64     `json:"Combat"`
	Empire     int64     `json:"Empire"`
	Explore    int64     `json:"Explore"`
	Federation int64     `json:"Federation"`
	Trade      int64     `json:"Trade"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Progress) GetEvent() string {
	return e.Event
}

func (e Progress) GetTimestamp() time.Time {
	return e.Timestamp
}

type Promotion struct {
	Explore   int64     `json:"Explore"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Promotion) GetEvent() string {
	return e.Event
}

func (e Promotion) GetTimestamp() time.Time {
	return e.Timestamp
}

type Rank struct {
	Cqc        int64     `json:"CQC"`
	Combat     int64     `json:"Combat"`
	Empire     int64     `json:"Empire"`
	Explore    int64     `json:"Explore"`
	Federation int64     `json:"Federation"`
	Trade      int64     `json:"Trade"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e Rank) GetEvent() string {
	return e.Event
}

func (e Rank) GetTimestamp() time.Time {
	return e.Timestamp
}

type RebootRepair struct {
	Modules   []interface{} `json:"Modules"`
	Event     string        `json:"event"`
	Timestamp time.Time     `json:"timestamp"`
}

func (e RebootRepair) GetEvent() string {
	return e.Event
}

func (e RebootRepair) GetTimestamp() time.Time {
	return e.Timestamp
}

type ReceiveText struct {
	Channel          string    `json:"Channel"`
	From             string    `json:"From"`
	Message          string    `json:"Message"`
	MessageLocalised string    `json:"Message_Localised"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e ReceiveText) GetEvent() string {
	return e.Event
}

func (e ReceiveText) GetTimestamp() time.Time {
	return e.Timestamp
}

type RedeemVoucher struct {
	Amount    int64                `json:"Amount"`
	Factions  []RedeemVoucher_sub1 `json:"Factions"`
	Type      string               `json:"Type"`
	Event     string               `json:"event"`
	Timestamp time.Time            `json:"timestamp"`
}

type RedeemVoucher_sub1 struct {
	Amount  int64  `json:"Amount"`
	Faction string `json:"Faction"`
}

func (e RedeemVoucher) GetEvent() string {
	return e.Event
}

func (e RedeemVoucher) GetTimestamp() time.Time {
	return e.Timestamp
}

type RefuelAll struct {
	Amount    float64   `json:"Amount"`
	Cost      int64     `json:"Cost"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e RefuelAll) GetEvent() string {
	return e.Event
}

func (e RefuelAll) GetTimestamp() time.Time {
	return e.Timestamp
}

type Repair struct {
	Cost      int64     `json:"Cost"`
	Item      string    `json:"Item"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Repair) GetEvent() string {
	return e.Event
}

func (e Repair) GetTimestamp() time.Time {
	return e.Timestamp
}

type RepairAll struct {
	Cost      int64     `json:"Cost"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e RepairAll) GetEvent() string {
	return e.Event
}

func (e RepairAll) GetTimestamp() time.Time {
	return e.Timestamp
}

type RestockVehicle struct {
	Cost      int64     `json:"Cost"`
	Count     int64     `json:"Count"`
	Loadout   string    `json:"Loadout"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e RestockVehicle) GetEvent() string {
	return e.Event
}

func (e RestockVehicle) GetTimestamp() time.Time {
	return e.Timestamp
}

type Resurrect struct {
	Bankrupt  bool      `json:"Bankrupt"`
	Cost      int64     `json:"Cost"`
	Option    string    `json:"Option"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Resurrect) GetEvent() string {
	return e.Event
}

func (e Resurrect) GetTimestamp() time.Time {
	return e.Timestamp
}

type Scan struct {
	Atmosphere            string           `json:"Atmosphere"`
	AtmosphereComposition []ScanAtmosphere `json:"AtmosphereComposition"`
	AtmosphereType        string           `json:"AtmosphereType"`
	AxialTilt             float64          `json:"AxialTilt"`
	BodyName              string           `json:"BodyName"`
	DistanceFromArrivalLS float64          `json:"DistanceFromArrivalLS"`
	Eccentricity          float64          `json:"Eccentricity"`
	Landable              bool             `json:"Landable"`
	MassEM                float64          `json:"MassEM"`
	OrbitalInclination    float64          `json:"OrbitalInclination"`
	OrbitalPeriod         float64          `json:"OrbitalPeriod"`
	Periapsis             float64          `json:"Periapsis"`
	PlanetClass           string           `json:"PlanetClass"`
	Radius                float64          `json:"Radius"`
	RotationPeriod        float64          `json:"RotationPeriod"`
	SemiMajorAxis         float64          `json:"SemiMajorAxis"`
	SurfaceGravity        float64          `json:"SurfaceGravity"`
	SurfacePressure       float64          `json:"SurfacePressure"`
	SurfaceTemperature    float64          `json:"SurfaceTemperature"`
	TerraformState        string           `json:"TerraformState"`
	TidalLock             bool             `json:"TidalLock"`
	Volcanism             string           `json:"Volcanism"`
	Event                 string           `json:"event"`
	Timestamp             time.Time        `json:"timestamp"`
}

type ScanAtmosphere struct {
	Name    string  `json:"Name"`
	Percent float64 `json:"Percent"`
}

func (e Scan) GetEvent() string {
	return e.Event
}

func (e Scan) GetTimestamp() time.Time {
	return e.Timestamp
}

type Scanned struct {
	ScanType  string    `json:"ScanType"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Scanned) GetEvent() string {
	return e.Event
}

func (e Scanned) GetTimestamp() time.Time {
	return e.Timestamp
}

type Screenshot struct {
	Body      string    `json:"Body"`
	Filename  string    `json:"Filename"`
	Height    int64     `json:"Height"`
	System    string    `json:"System"`
	Width     int64     `json:"Width"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Screenshot) GetEvent() string {
	return e.Event
}

func (e Screenshot) GetTimestamp() time.Time {
	return e.Timestamp
}

type SelfDestruct struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e SelfDestruct) GetEvent() string {
	return e.Event
}

func (e SelfDestruct) GetTimestamp() time.Time {
	return e.Timestamp
}

type SellDrones struct {
	Count     int64     `json:"Count"`
	SellPrice int64     `json:"SellPrice"`
	TotalSale int64     `json:"TotalSale"`
	Type      string    `json:"Type"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e SellDrones) GetEvent() string {
	return e.Event
}

func (e SellDrones) GetTimestamp() time.Time {
	return e.Timestamp
}

type SellExplorationData struct {
	BaseValue  int64     `json:"BaseValue"`
	Bonus      int64     `json:"Bonus"`
	Discovered []string  `json:"Discovered"`
	Systems    []string  `json:"Systems"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e SellExplorationData) GetEvent() string {
	return e.Event
}

func (e SellExplorationData) GetTimestamp() time.Time {
	return e.Timestamp
}

type SendText struct {
	Message   string    `json:"Message"`
	To        string    `json:"To"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e SendText) GetEvent() string {
	return e.Event
}

func (e SendText) GetTimestamp() time.Time {
	return e.Timestamp
}

type SetUserShipName struct {
	Ship         string    `json:"Ship"`
	ShipID       int64     `json:"ShipID"`
	UserShipID   string    `json:"UserShipId"`
	UserShipName string    `json:"UserShipName"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e SetUserShipName) GetEvent() string {
	return e.Event
}

func (e SetUserShipName) GetTimestamp() time.Time {
	return e.Timestamp
}

type ShieldState struct {
	ShieldsUp bool      `json:"ShieldsUp"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ShieldState) GetEvent() string {
	return e.Event
}

func (e ShieldState) GetTimestamp() time.Time {
	return e.Timestamp
}

type ShipyardBuy struct {
	ShipPrice    int64     `json:"ShipPrice"`
	ShipType     string    `json:"ShipType"`
	StoreOldShip string    `json:"StoreOldShip"`
	StoreShipID  int64     `json:"StoreShipID"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e ShipyardBuy) GetEvent() string {
	return e.Event
}

func (e ShipyardBuy) GetTimestamp() time.Time {
	return e.Timestamp
}

type ShipyardNew struct {
	NewShipID int64     `json:"NewShipID"`
	ShipType  string    `json:"ShipType"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e ShipyardNew) GetEvent() string {
	return e.Event
}

func (e ShipyardNew) GetTimestamp() time.Time {
	return e.Timestamp
}

type ShipyardSell struct {
	SellShipID int64     `json:"SellShipID"`
	ShipPrice  int64     `json:"ShipPrice"`
	ShipType   string    `json:"ShipType"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e ShipyardSell) GetEvent() string {
	return e.Event
}

func (e ShipyardSell) GetTimestamp() time.Time {
	return e.Timestamp
}

type ShipyardSwap struct {
	ShipID       int64     `json:"ShipID"`
	ShipType     string    `json:"ShipType"`
	StoreOldShip string    `json:"StoreOldShip"`
	StoreShipID  int64     `json:"StoreShipID"`
	Event        string    `json:"event"`
	Timestamp    time.Time `json:"timestamp"`
}

func (e ShipyardSwap) GetEvent() string {
	return e.Event
}

func (e ShipyardSwap) GetTimestamp() time.Time {
	return e.Timestamp
}

type ShipyardTransfer struct {
	Distance      float64   `json:"Distance"`
	ShipID        int64     `json:"ShipID"`
	ShipType      string    `json:"ShipType"`
	System        string    `json:"System"`
	TransferPrice int64     `json:"TransferPrice"`
	TransferTime  int64     `json:"TransferTime"`
	Event         string    `json:"event"`
	Timestamp     time.Time `json:"timestamp"`
}

func (e ShipyardTransfer) GetEvent() string {
	return e.Event
}

func (e ShipyardTransfer) GetTimestamp() time.Time {
	return e.Timestamp
}

type StartJump struct {
	JumpType   string    `json:"JumpType"`
	StarClass  string    `json:"StarClass"`
	StarSystem string    `json:"StarSystem"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e StartJump) GetEvent() string {
	return e.Event
}

func (e StartJump) GetTimestamp() time.Time {
	return e.Timestamp
}

type SupercruiseEntry struct {
	StarSystem string    `json:"StarSystem"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e SupercruiseEntry) GetEvent() string {
	return e.Event
}

func (e SupercruiseEntry) GetTimestamp() time.Time {
	return e.Timestamp
}

type SupercruiseExit struct {
	Body       string    `json:"Body"`
	BodyType   string    `json:"BodyType"`
	StarSystem string    `json:"StarSystem"`
	Event      string    `json:"event"`
	Timestamp  time.Time `json:"timestamp"`
}

func (e SupercruiseExit) GetEvent() string {
	return e.Event
}

func (e SupercruiseExit) GetTimestamp() time.Time {
	return e.Timestamp
}

type Synthesis struct {
	Materials []Synthesis_sub1 `json:"Materials"`
	Name      string           `json:"Name"`
	Event     string           `json:"event"`
	Timestamp time.Time        `json:"timestamp"`
}

type Synthesis_sub1 struct {
	Count int64  `json:"Count"`
	Name  string `json:"Name"`
}

func (e Synthesis) GetEvent() string {
	return e.Event
}

func (e Synthesis) GetTimestamp() time.Time {
	return e.Timestamp
}

type Touchdown struct {
	Latitude         float64   `json:"Latitude"`
	Longitude        float64   `json:"Longitude"`
	PlayerControlled bool      `json:"PlayerControlled"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e Touchdown) GetEvent() string {
	return e.Event
}

func (e Touchdown) GetTimestamp() time.Time {
	return e.Timestamp
}

type Undocked struct {
	StationName string    `json:"StationName"`
	StationType string    `json:"StationType"`
	Event       string    `json:"event"`
	Timestamp   time.Time `json:"timestamp"`
}

func (e Undocked) GetEvent() string {
	return e.Event
}

func (e Undocked) GetTimestamp() time.Time {
	return e.Timestamp
}

type USSDrop struct {
	USSThreat        int64     `json:"USSThreat"`
	USSType          string    `json:"USSType"`
	USSTypeLocalised string    `json:"USSType_Localised"`
	Event            string    `json:"event"`
	Timestamp        time.Time `json:"timestamp"`
}

func (e USSDrop) GetEvent() string {
	return e.Event
}

func (e USSDrop) GetTimestamp() time.Time {
	return e.Timestamp
}

type WingAdd struct {
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingAdd) GetEvent() string {
	return e.Event
}

func (e WingAdd) GetTimestamp() time.Time {
	return e.Timestamp
}

type WingInvite struct {
	Name      string    `json:"Name"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingInvite) GetEvent() string {
	return e.Event
}

func (e WingInvite) GetTimestamp() time.Time {
	return e.Timestamp
}

type WingJoin struct {
	Others    []string  `json:"Others"`
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingJoin) GetEvent() string {
	return e.Event
}

func (e WingJoin) GetTimestamp() time.Time {
	return e.Timestamp
}

type WingLeave struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func (e WingLeave) GetEvent() string {
	return e.Event
}

func (e WingLeave) GetTimestamp() time.Time {
	return e.Timestamp
}
