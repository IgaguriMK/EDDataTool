package journal
import "time"

type ModuleBuy struct {
	BuyItem             string `json:"BuyItem"`
	BuyItemLocalised    string `json:"BuyItem_Localised"`
	BuyPrice            int64  `json:"BuyPrice"`
	SellItem            string `json:"SellItem"`
	SellItemLocalised   string `json:"SellItem_Localised"`
	SellPrice           int64  `json:"SellPrice"`
	Ship                string `json:"Ship"`
	ShipID              int64  `json:"ShipID"`
	Slot                string `json:"Slot"`
	StoredItem          string `json:"StoredItem"`
	StoredItemLocalised string `json:"StoredItem_Localised"`
	Event               string `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
func (e ModuleBuy) GetEvent() string {
       return e.Event
}

func (e ModuleBuy) GetTimestamp() time.Time {
       return e.Timestamp
}
