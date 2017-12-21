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
