package smhi

type ParameterData struct {
	Key     string `json:"key,omitempty"`
	Name    string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Unit    string `json:"unit,omitempty"`
}

type StationData struct {
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Height int    `json:"height,omitempty"`
}

type PeriodData struct {
	Key      string `json:"key,omitempty"`
	From     uint64 `json:"from,omitempty"`
	To       uint64 `json:"to,omitempty"`
	Summary  string `json:"summary,omitempty"`
	Sampling string `json:"sampling,omitempty"`
}

type PositionData struct {
	From      uint64  `json:"from,omitempty"`
	To        uint64  `json:"to,omitempty"`
	Height    int     `json:"height,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
}
