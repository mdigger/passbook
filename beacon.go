package passbook

import (
	"encoding/json"
	"errors"
)

// Beacon Dictionary: Information about a location beacon. Available in iOS 7.0.
type Beacon struct {
	ProximityUUID string `json:"proximityUUID"`          // Unique identifier of a Bluetooth Low Energy location beacon.
	Major         uint16 `json:"major,omitempty"`        // Major identifier of a Bluetooth Low Energy location beacon.
	Minor         uint16 `json:"minor,omitempty"`        // Minor identifier of a Bluetooth Low Energy location beacon.
	RelevantText  string `json:"relevantText,omitempty"` // Text displayed on the lock screen when the pass is currently relevant.
}

func (b Beacon) Marshal() ([]byte, error) {
	if b.ProximityUUID == "" {
		return nil, errors.New("Unique identifier of a Bluetooth Low Energy location beacon must be set")
	}
	return json.Marshal(b)
}
