package passbook

// Location Dictionary: Information about a location.
type Location struct {
	Latitude     float64 `json:"latitude"`               // Latitude, in degrees, of the location.
	Longitude    float64 `json:"longitude"`              // Longitude, in degrees, of the location.
	Altitude     float64 `json:"altitude,omitempty"`     // Altitude, in meters, of the location.
	RelevantText string  `json:"relevantText,omitempty"` // Text displayed on the lock screen when the pass is currently relevant.
}
