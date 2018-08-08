package smhi

// Parameter holds information on all datapoints for that parameter
// Parameter is a type of weather observation
type Parameter struct {
	Key        string     `json:"key,omitempty"`
	Updated    uint64     `json:"updated,omitempty"`
	Title      string     `json:"title,omitempty"`
	Summary    string     `json:"summary,omitempty"`
	ValueType  string     `json:"valueType,omitempty"`
	StationSet StationSet `json:"stationSet,omitempty"`
	Station    []Station  `json:"station,omitempty"`
}

// StationSet is a set of stations
type StationSet struct{}

// Station defines a measurment station
type Station struct {
	Name      string  `json:"name,omitempty"`
	Owner     string  `json:"owner,omitempty"`
	ID        uint32  `json:"id,omitempty"`
	Height    float32 `json:"height,omitempty"`
	Latitude  float32 `json:"latitude,omitempty"`
	Longitude float32 `json:"longitude,omitempty"`
	Active    bool    `json:"active,omitempty"`
	Key       string  `json:"key,omitempty"`
	Updated   uint64  `json:"updated,omitempty"`
	Title     string  `json:"title,omitempty"`
	Summary   string  `json:"summary,omitempty"`
}
