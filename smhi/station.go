package smhi

// Station defines a measurment station
type Station struct {
	Name     *string
	Owner    *string
	ID       *uint16
	From     *uint32
	To       *uint32
	Position []*Position
	Active   bool
	Key      *string
	Updated  *uint32
	Title    *string
	Summary  *string
}

// Position defines a station's position
type Position struct {
	From      *uint32
	To        *uint32
	Height    *int16
	Latitude  *float32
	Longitude *float32
}
