package journal
import "time"

type ModuleSellRemote struct {
	SellItem          string `json:"SellItem"`
	SellItemLocalised string `json:"SellItem_Localised"`
	SellPrice         int64  `json:"SellPrice"`
	ServerID          int64  `json:"ServerId"`
	Ship              string `json:"Ship"`
	ShipID            int64  `json:"ShipID"`
	StorageSlot       int64  `json:"StorageSlot"`
	Event             string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ModuleSellRemote) GetEvent() string {
       return e.Event
}

func (e ModuleSellRemote) GetTimestamp() time.Time {
       return e.Timestamp
}
