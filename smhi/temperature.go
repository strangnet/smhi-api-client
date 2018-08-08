package smhi

import (
	"context"
	"fmt"
	"net/http"
)

// Temperature parameter definitions
const (
	TemperatureParameterHourly            = 1
	TemperatureParameterAverageDaily      = 2
	TemperatureParameterMinimumDaily      = 19
	TemperatureParameterMaximumDaily      = 20
	TemperatureParameterAverageMonthly    = 22
	TemperatureParameterMinimumTwiceDaily = 26 // Not implemented
	TemperatureParameterMaximumTwiceDaily = 27 // Not implemented
)

// TemperatureService is a service for the temperature queries
type TemperatureService service

// TemperatureData hold the returned data
type TemperatureData struct {
	Value     []TemperatureDataValue `json:"value,omitempty"`
	Updated   uint64                 `json:"updated,omitempty"`
	Parameter ParameterData          `json:"parameter,omitempty"`
	Station   StationData            `json:"station,omitempty"`
	Period    PeriodData             `json:"period,omitempty"`
	Position  []PositionData         `json:"position,omitempty"`
}

// TemperatureDataValue holds value data for temperatures
type TemperatureDataValue struct {
	From    uint64 `json:"from,omitempty"`
	To      uint64 `json:"to,omitempty"`
	Ref     string `json:"ref,omitempty"`
	Value   string `json:"value,omitempty"`
	Quality string `json:"quality,omitempty"`
}

func getTemperatureData(ctx context.Context, client *Client, parameter int, station uint32, period string) (*TemperatureData, *http.Response, error) {
	dataURL := fmt.Sprintf("api/version/latest/parameter/%d/station/%d/period/%s/data.json", parameter, station, period)
	req, err := client.NewRequest("GET", dataURL)
	if err != nil {
		return nil, nil, err
	}

	td := &TemperatureData{}
	resp, err := client.Do(ctx, req, td)
	if err != nil {
		return nil, resp, err
	}

	return td, resp, nil
}

func getParameterData(ctx context.Context, client *Client, parameter int, includeInactive bool) (*Parameter, *http.Response, error) {
	dataURL := fmt.Sprintf("api/version/latest/parameter/%d.json", parameter)
	req, err := client.NewRequest("GET", dataURL)
	if err != nil {
		return nil, nil, err
	}

	p := &Parameter{}
	resp, err := client.Do(ctx, req, p)
	if err != nil {
		return nil, resp, err
	}

	// Filter out the inactive stations
	if !includeInactive {
		newStation := make([]Station, 0)
		for _, s := range p.Station {
			if s.Active {
				newStation = append(newStation, s)
			}
		}
		p.Station = newStation
	}

	return p, resp, nil
}

// GetHourlyTemperatures retrieves hourly temperatures from a station
func (s *TemperatureService) GetHourlyTemperatures(ctx context.Context, station uint32, period string) (*TemperatureData, *http.Response, error) {
	return getTemperatureData(ctx, s.client, TemperatureParameterHourly, station, period)
}

// GetStationsWithHourlyTemperatures retrives all stations with hourly temperatures
func (s *TemperatureService) GetStationsWithHourlyTemperatures(ctx context.Context, includeInactive bool) (*Parameter, *http.Response, error) {
	return getParameterData(ctx, s.client, TemperatureParameterHourly, includeInactive)
}

// GetAverageDailyTemperatures retrieves the average daily temperatures from a station
func (s *TemperatureService) GetAverageDailyTemperatures(ctx context.Context, station uint32, period string) (*TemperatureData, *http.Response, error) {
	return getTemperatureData(ctx, s.client, TemperatureParameterAverageDaily, station, period)
}

// GetStationsWithAverageDailyTemperatures retrieves all stations with average daily temperatures
func (s *TemperatureService) GetStationsWithAverageDailyTemperatures(ctx context.Context, includeInactive bool) (*Parameter, *http.Response, error) {
	return getParameterData(ctx, s.client, TemperatureParameterAverageDaily, includeInactive)
}

// GetAverageMonthlyTemperatures retrieves the average monthly temperatures from a station
func (s *TemperatureService) GetAverageMonthlyTemperatures(ctx context.Context, station uint32, period string) (*TemperatureData, *http.Response, error) {
	return getTemperatureData(ctx, s.client, TemperatureParameterAverageMonthly, station, period)
}

// GetStationsWithAverageMonthlyTemperatures retrieves all stations with average daily temperatures
func (s *TemperatureService) GetStationsWithAverageMonthlyTemperatures(ctx context.Context, includeInactive bool) (*Parameter, *http.Response, error) {
	return getParameterData(ctx, s.client, TemperatureParameterAverageMonthly, includeInactive)
}

// GetMinimumDailyTemperatures retrieves the minimum daily temperatures from a station
func (s *TemperatureService) GetMinimumDailyTemperatures(ctx context.Context, station uint32, period string) (*TemperatureData, *http.Response, error) {
	return getTemperatureData(ctx, s.client, TemperatureParameterMinimumDaily, station, period)
}

// GetStationsWithMinimumDailyTemperatures retrieves all stations with minimum daily temperatures
func (s *TemperatureService) GetStationsWithMinimumDailyTemperatures(ctx context.Context, includeInactive bool) (*Parameter, *http.Response, error) {
	return getParameterData(ctx, s.client, TemperatureParameterMinimumDaily, includeInactive)
}

// GetMaximumDailyTemperatures retrieves the maximum daily temperatures from a station
func (s *TemperatureService) GetMaximumDailyTemperatures(ctx context.Context, station uint32, period string) (*TemperatureData, *http.Response, error) {
	return getTemperatureData(ctx, s.client, TemperatureParameterMaximumDaily, station, period)
}

// GetStationsWithMaximumDailyTemperatures retrieves all stations with maximum daily temperatures
func (s *TemperatureService) GetStationsWithMaximumDailyTemperatures(ctx context.Context, includeInactive bool) (*Parameter, *http.Response, error) {
	return getParameterData(ctx, s.client, TemperatureParameterMaximumDaily, includeInactive)
}
