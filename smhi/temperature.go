package smhi

import (
	"context"
	"fmt"
	"net/http"
)

// Constants used for buiding the requests
const (
	PeriodLatestHour       = "latest-hour"
	PeriodLatestDay        = "latest-day"
	PeriodLatestMonths     = "latest-months"
	PeriodCorrectedArchive = "corrected-archive"

	FormatJSON = "json"
	FormatXML  = "xml"
	FormatCSV  = "csv"
)

// TemperatureService is a service for the temperature queries
type TemperatureService service

// TemperatureData hold the returned data
type TemperatureData struct {
	Value []TemperatureDataValue `json:"value,omitempty"`
}

// TemperatureDataValue holds value data for temperatures
type TemperatureDataValue struct {
	From    uint64 `json:"from,omitempty"`
	To      uint64 `json:"to,omitempty"`
	Ref     string `json:"ref,omitempty"`
	Value   string `json:"value,omitempty"`
	Quality string `json:"quality,omitempty"`
}

// GetAverageDailyTemperatures get the average daily temperatures from a station
func (s *TemperatureService) GetAverageDailyTemperatures(ctx context.Context, station uint16, period string, format string) (*TemperatureData, *http.Response, error) {
	dataURL := fmt.Sprintf("api/version/latest/parameter/2/station/%d/period/%s/data.%s", station, period, format)
	req, err := s.client.NewRequest("GET", dataURL)
	if err != nil {
		return nil, nil, err
	}

	td := &TemperatureData{}
	resp, err := s.client.Do(ctx, req, td)
	if err != nil {
		return nil, resp, err
	}

	return td, resp, nil
}
