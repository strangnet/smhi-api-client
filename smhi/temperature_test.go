package smhi

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTemperatureService_GetDailyAverageTemperatures(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/2/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"value": [{"from": 1533254401000, "to": 1533340800000, "ref": "2018-08-03", "value": "21.8", "quality": "Y"}],
		"updated": 1533470400000,
		"parameter": {
		"key": "2",
		"name": "Lufttemperatur",
		"summary": "medelvärde 1 dygn, 1 gång/dygn, kl 00",
		"unit": "degree celsius"
		},
		"station": {
		"key": "97100",
		"name": "Tullinge A",
		"owner": "SMHI",
		"height": 2
		},
		"period": {
		"key": "latest-day",
		"from": 1533380401000,
		"to": 1533470400000,
		"summary": "Data från senaste dygnet",
		"sampling": "24 timmar"
		},
		"position": [
		{
		"from": 818985600000,
		"to": 1533470400000,
		"height": 45,
		"latitude": 59.1789,
		"longitude": 17.9125
		}
		]}`)
	})

	temps, _, err := client.Temperatures.GetAverageDailyTemperatures(context.Background(), 12345, PeriodLatestDay, FormatJSON)
	if err != nil {
		t.Errorf("Temperatures.GetAverageDailyTemperatures returned error: %v", err)
	}

	want := &TemperatureData{
		Value: []TemperatureDataValue{
			{
				From:    1533254401000,
				To:      1533340800000,
				Ref:     "2018-08-03",
				Value:   "21.8",
				Quality: "Y",
			},
		},
		Updated: 1533470400000,
		Parameter: ParameterData{
			Key:     "2",
			Name:    "Lufttemperatur",
			Summary: "medelvärde 1 dygn, 1 gång/dygn, kl 00",
			Unit:    "degree celsius",
		},
		Station: StationData{
			Key:    "97100",
			Name:   "Tullinge A",
			Owner:  "SMHI",
			Height: 2,
		},
		Period: PeriodData{
			Key:      "latest-day",
			From:     1533380401000,
			To:       1533470400000,
			Summary:  "Data från senaste dygnet",
			Sampling: "24 timmar",
		},
		Position: []PositionData{
			{
				From:      818985600000,
				To:        1533470400000,
				Height:    45,
				Latitude:  59.1789,
				Longitude: 17.9125,
			},
		},
	}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetAverageDailyTemperatures returned %+v, want %+v", temps, want)
	}
}
