package smhi

// ParameterData holds information on the parameters
type ParameterData struct {
	Key     string `json:"key,omitempty"`
	Name    string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Unit    string `json:"unit,omitempty"`
}

// StationData holds information on the station
type StationData struct {
	Key    string  `json:"key,omitempty"`
	Name   string  `json:"name,omitempty"`
	Owner  string  `json:"owner,omitempty"`
	Height float32 `json:"height,omitempty"`
}

// PeriodData holds information on the period
type PeriodData struct {
	Key      string `json:"key,omitempty"`
	From     uint64 `json:"from,omitempty"`
	To       uint64 `json:"to,omitempty"`
	Summary  string `json:"summary,omitempty"`
	Sampling string `json:"sampling,omitempty"`
}

// PositionData holds information on the position
type PositionData struct {
	From      uint64  `json:"from,omitempty"`
	To        uint64  `json:"to,omitempty"`
	Height    float32 `json:"height,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}
