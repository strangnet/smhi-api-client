package smhi

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTemperatureService_GetAverageDailyTemperatures_returnsOK(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/2/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{"value": [{"from": 1533254401000, "to": 1533340800000, "ref": "2018-08-03", "value": "21.8", "quality": "Y"}],
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
		"height": 2.0
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
		"height": 45.0,
		"latitude": 59.1789,
		"longitude": 17.9125
		}
		]}`)
	})

	temps, _, err := client.Temperatures.GetAverageDailyTemperatures(context.Background(), 12345, PeriodLatestDay)
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
			Height: 2.0,
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
				Height:    45.0,
				Latitude:  59.1789,
				Longitude: 17.9125,
			},
		},
	}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetAverageDailyTemperatures returned %+v, want %+v", temps, want)
	}
}

func TestTemperatureService_GetHourlyTemperatures_returnsOK(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/1/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"value": [{"from": 1533254401000, "to": 1533340800000, "ref": "2018-08-03", "value": "21.8", "quality": "Y"}],
		"updated": 1533470400000,
		"parameter": {
		"key": "1",
		"name": "Lufttemperatur",
		"summary": "momentanvärde, 1 gång/tim",
		"unit": "degree celsius"
		},
		"station": {
		"key": "97100",
		"name": "Tullinge A",
		"owner": "SMHI",
		"height": 2.0
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
		"height": 45.0,
		"latitude": 59.1789,
		"longitude": 17.9125
		}
		]}`)
	})

	temps, _, err := client.Temperatures.GetHourlyTemperatures(context.Background(), 12345, PeriodLatestDay)
	if err != nil {
		t.Errorf("Temperatures.GetHourlyTemperatures returned error: %v", err)
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
			Key:     "1",
			Name:    "Lufttemperatur",
			Summary: "momentanvärde, 1 gång/tim",
			Unit:    "degree celsius",
		},
		Station: StationData{
			Key:    "97100",
			Name:   "Tullinge A",
			Owner:  "SMHI",
			Height: 2.0,
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
				Height:    45.0,
				Latitude:  59.1789,
				Longitude: 17.9125,
			},
		},
	}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetHourlyTemperatures returned %+v, want %+v", temps, want)
	}
}

func TestTemperatureService_GetAverageMonthlyTemperatures_returnsOK(t *testing.T) {

	data := `
	{"value": [
		{
			"from": 1519862401000,
			"to": 1522540800000,
			"ref": "2018-03",
			"value": "-3.0",
			"quality": "Y"
			},
			{
			"from": 1522540801000,
			"to": 1525132800000,
			"ref": "2018-04",
			"value": "5.6",
			"quality": "Y"
			},
			{
			"from": 1525132801000,
			"to": 1527811200000,
			"ref": "2018-05",
			"value": "14.2",
			"quality": "Y"
			},
			{
			"from": 1527811201000,
			"to": 1530403200000,
			"ref": "2018-06",
			"value": "15.6",
			"quality": "Y"
			},
			{
			"from": 1530403201000,
			"to": 1533081600000,
			"ref": "2018-07",
			"value": "20.5",
			"quality": "Y"
			}
		],
		"updated": 1533470400000,
		"parameter": {
			"key": "22",
			"name": "Lufttemperatur",
			"summary": "medel, 1 gång per månad",
			"unit": "degree celsius"
		},
		"station": {
			"key": "97100",
			"name": "Tullinge A",
			"owner": "SMHI",
			"height": 2.0
		},
		"period": {
			"key": "latest-months",
			"from": 1522454401000,
			"to": 1533690000000,
			"summary": "Data från senaste fyra månaderna",
			"sampling": "1 månad"
		},
		"position": [
		{
		"from": 818985600000,
		"to": 1533470400000,
		"height": 45.0,
		"latitude": 59.1789,
		"longitude": 17.9125
		}
		]}
	`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/22/station/12345/period/latest-months/data.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	temps, _, err := client.Temperatures.GetAverageMonthlyTemperatures(context.Background(), 12345, PeriodLatestMonths)
	if err != nil {
		t.Errorf("Temperatures.GetAverageMonthlyTemperatures returned error: %v", err)
	}

	want := &TemperatureData{
		Value: []TemperatureDataValue{
			{
				From:    1519862401000,
				To:      1522540800000,
				Ref:     "2018-03",
				Value:   "-3.0",
				Quality: "Y",
			},
			{
				From:    1522540801000,
				To:      1525132800000,
				Ref:     "2018-04",
				Value:   "5.6",
				Quality: "Y",
			},
			{
				From:    1525132801000,
				To:      1527811200000,
				Ref:     "2018-05",
				Value:   "14.2",
				Quality: "Y",
			},
			{
				From:    1527811201000,
				To:      1530403200000,
				Ref:     "2018-06",
				Value:   "15.6",
				Quality: "Y",
			},
			{
				From:    1530403201000,
				To:      1533081600000,
				Ref:     "2018-07",
				Value:   "20.5",
				Quality: "Y",
			},
		},
		Updated: 1533470400000,
		Parameter: ParameterData{
			Key:     "22",
			Name:    "Lufttemperatur",
			Summary: "medel, 1 gång per månad",
			Unit:    "degree celsius",
		},
		Station: StationData{
			Key:    "97100",
			Name:   "Tullinge A",
			Owner:  "SMHI",
			Height: 2.0,
		},
		Period: PeriodData{
			Key:      "latest-months",
			From:     1522454401000,
			To:       1533690000000,
			Summary:  "Data från senaste fyra månaderna",
			Sampling: "1 månad",
		},
		Position: []PositionData{
			{
				From:      818985600000,
				To:        1533470400000,
				Height:    45.0,
				Latitude:  59.1789,
				Longitude: 17.9125,
			},
		},
	}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetAverageMonthlyTemperatures returned %+v, want %+v", temps, want)
	}
}

func TestTemperatureService_GetMinimumDailyTemperatures_returnsOK(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/19/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"value": [{"from": 1533254401000, "to": 1533340800000, "ref": "2018-08-03", "value": "21.8", "quality": "Y"}],
		"updated": 1533470400000,
		"parameter": {
		"key": "19",
		"name": "Lufttemperatur",
		"summary": "min, 1 gång per dygn",
		"unit": "degree celsius"
		},
		"station": {
		"key": "97100",
		"name": "Tullinge A",
		"owner": "SMHI",
		"height": 2.0
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
		"height": 45.0,
		"latitude": 59.1789,
		"longitude": 17.9125
		}
		]}`)
	})

	temps, _, err := client.Temperatures.GetMinimumDailyTemperatures(context.Background(), 12345, PeriodLatestDay)
	if err != nil {
		t.Errorf("Temperatures.GetMinimumDailyTemperatures returned error: %v", err)
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
			Key:     "19",
			Name:    "Lufttemperatur",
			Summary: "min, 1 gång per dygn",
			Unit:    "degree celsius",
		},
		Station: StationData{
			Key:    "97100",
			Name:   "Tullinge A",
			Owner:  "SMHI",
			Height: 2.0,
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
				Height:    45.0,
				Latitude:  59.1789,
				Longitude: 17.9125,
			},
		},
	}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetMinimumDailyTemperatures returned %+v, want %+v", temps, want)
	}
}

func TestTemperatureService_GetMaximumDailyTemperatures_returnsOK(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/20/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"value": [{"from": 1533254401000, "to": 1533340800000, "ref": "2018-08-03", "value": "21.8", "quality": "Y"}],
		"updated": 1533470400000,
		"parameter": {
		"key": "20",
		"name": "Lufttemperatur",
		"summary": "max, 1 gång per dygn",
		"unit": "degree celsius"
		},
		"station": {
		"key": "97100",
		"name": "Tullinge A",
		"owner": "SMHI",
		"height": 2.0
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
		"height": 45.0,
		"latitude": 59.1789,
		"longitude": 17.9125
		}
		]}`)
	})

	temps, _, err := client.Temperatures.GetMaximumDailyTemperatures(context.Background(), 12345, PeriodLatestDay)
	if err != nil {
		t.Errorf("Temperatures.GetMaximumDailyTemperatures returned error: %v", err)
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
			Key:     "20",
			Name:    "Lufttemperatur",
			Summary: "max, 1 gång per dygn",
			Unit:    "degree celsius",
		},
		Station: StationData{
			Key:    "97100",
			Name:   "Tullinge A",
			Owner:  "SMHI",
			Height: 2.0,
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
				Height:    45.0,
				Latitude:  59.1789,
				Longitude: 17.9125,
			},
		},
	}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetMaximumDailyTemperatures returned %+v, want %+v", temps, want)
	}
}

func TestTemperatureService_GetAverageDailyTemperatures_returns404(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/2/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {})

	_, resp, err := client.Temperatures.GetAverageDailyTemperatures(context.Background(), 12346, PeriodLatestDay)
	if err != nil {
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Temperatures.GetAverageDailyTemperatures returned error code %d, expected %d", resp.StatusCode, http.StatusNotFound)
		}

	}
}

func TestTemperatureService_GetStationsWithAverageDailyTemperatures_returnsOK(t *testing.T) {
	data := `{
		"key": "2",
		"updated": 1533495600000,
		"title": "Lufttemperatur: Välj station (sedan tidsutsnitt)",
		"summary": "medelvärde 1 dygn, 1 gång/dygn, kl 00",
		"valueType": "INTERVAL",
		"station": [
		{
			"name": "Abelvattnet Aut",
			"owner": "SMHI",
			"id": 154860,
			"height": 665.0,
			"latitude": 65.53,
			"longitude": 14.97,
			"active": true,
			"key": "154860",
			"updated": 841535999000,
			"title": "Lufttemperatur - Abelvattnet Aut",
			"summary": "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0"
		},
		{
			"name": "Abisko",
			"owner": "SMHI",
			"id": 188800,
			"height": 388.0,
			"latitude": 68.3557,
			"longitude": 18.8206,
			"active": false,
			"key": "188800",
			"updated": 1533081599000,
			"title": "Lufttemperatur - Abisko",
			"summary": "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0"
		}
		]
	}`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/2.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	t.Run("includes all stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithAverageDailyTemperatures(context.Background(), true)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithAverageDailyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "2",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "medelvärde 1 dygn, 1 gång/dygn, kl 00",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
				{
					Name:      "Abisko",
					Owner:     "SMHI",
					ID:        188800,
					Height:    388.0,
					Latitude:  68.3557,
					Longitude: 18.8206,
					Active:    false,
					Key:       "188800",
					Updated:   1533081599000,
					Title:     "Lufttemperatur - Abisko",
					Summary:   "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithAverageDailyTemperatures returned %+v, want %+v", p, want)
		}
	})
	t.Run("includes only active stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithAverageDailyTemperatures(context.Background(), false)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithAverageDailyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "2",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "medelvärde 1 dygn, 1 gång/dygn, kl 00",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithAverageDailyTemperatures returned %+v, want %+v", p, want)
		}
	})
}

func TestTemperatureService_GetMinimumDailyTemperatures_returns404(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/19/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {})

	_, resp, err := client.Temperatures.GetMinimumDailyTemperatures(context.Background(), 12346, PeriodLatestDay)
	if err != nil {
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Temperatures.GetMinimumDailyTemperatures returned error code %d, expected %d", resp.StatusCode, http.StatusNotFound)
		}

	}
}

func TestTemperatureService_GetStationsWithMinimumDailyTemperatures_returnsOK(t *testing.T) {
	data := `{
		"key": "19",
		"updated": 1533495600000,
		"title": "Lufttemperatur: Välj station (sedan tidsutsnitt)",
		"summary": "min, 1 gång per dygn",
		"valueType": "INTERVAL",
		"station": [
		{
			"name": "Abelvattnet Aut",
			"owner": "SMHI",
			"id": 154860,
			"height": 665.0,
			"latitude": 65.53,
			"longitude": 14.97,
			"active": true,
			"key": "154860",
			"updated": 841535999000,
			"title": "Lufttemperatur - Abelvattnet Aut",
			"summary": "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0"
		},
		{
			"name": "Abisko",
			"owner": "SMHI",
			"id": 188800,
			"height": 388.0,
			"latitude": 68.3557,
			"longitude": 18.8206,
			"active": false,
			"key": "188800",
			"updated": 1533081599000,
			"title": "Lufttemperatur - Abisko",
			"summary": "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0"
		}
		]
	}`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/19.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	t.Run("includes all stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithMinimumDailyTemperatures(context.Background(), true)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithMinimumDailyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "19",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "min, 1 gång per dygn",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
				{
					Name:      "Abisko",
					Owner:     "SMHI",
					ID:        188800,
					Height:    388.0,
					Latitude:  68.3557,
					Longitude: 18.8206,
					Active:    false,
					Key:       "188800",
					Updated:   1533081599000,
					Title:     "Lufttemperatur - Abisko",
					Summary:   "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithMinimumDailyTemperatures returned %+v, want %+v", p, want)
		}
	})
	t.Run("includes only active stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithMinimumDailyTemperatures(context.Background(), false)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithMinimumDailyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "19",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "min, 1 gång per dygn",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithMinimumDailyTemperatures returned %+v, want %+v", p, want)
		}
	})
}

func TestTemperatureService_GetMaximumTemperatures_returns404(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/20/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {})

	_, resp, err := client.Temperatures.GetMaximumDailyTemperatures(context.Background(), 12346, PeriodLatestDay)
	if err != nil {
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Temperatures.GetMaximumDailyTemperatures returned error code %d, expected %d", resp.StatusCode, http.StatusNotFound)
		}
	}
}

func TestTemperatureService_GetAverageMonthlyTemperatures_returns404(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/22/station/12345/period/latest-months/data.json", func(w http.ResponseWriter, r *http.Request) {})

	_, resp, err := client.Temperatures.GetAverageMonthlyTemperatures(context.Background(), 12346, PeriodLatestDay)
	if err != nil {
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Temperatures.GetAverageMonthlyTemperatures returned error code %d, expected %d", resp.StatusCode, http.StatusNotFound)
		}
	}
}

func TestTemperatureService_GetHourlyTemperatures_returns404(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/1/station/12345/period/latest-day/data.json", func(w http.ResponseWriter, r *http.Request) {})

	_, resp, err := client.Temperatures.GetHourlyTemperatures(context.Background(), 12346, PeriodLatestDay)
	if err != nil {
		if resp != nil && resp.StatusCode != http.StatusNotFound {
			t.Errorf("Temperatures.GetHourlyTemperatures returned error code %d, expected %d", resp.StatusCode, http.StatusNotFound)
		}
	}
}

func TestTemperatureService_GetStationsWithMaximumDailyTemperatures_returnsOK(t *testing.T) {
	data := `{
		"key": "20",
		"updated": 1533495600000,
		"title": "Lufttemperatur: Välj station (sedan tidsutsnitt)",
		"summary": "max, 1 gång per dygn",
		"valueType": "INTERVAL",
		"station": [
		{
			"name": "Abelvattnet Aut",
			"owner": "SMHI",
			"id": 154860,
			"height": 665.0,
			"latitude": 65.53,
			"longitude": 14.97,
			"active": true,
			"key": "154860",
			"updated": 841535999000,
			"title": "Lufttemperatur - Abelvattnet Aut",
			"summary": "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0"
		},
		{
			"name": "Abisko",
			"owner": "SMHI",
			"id": 188800,
			"height": 388.0,
			"latitude": 68.3557,
			"longitude": 18.8206,
			"active": false,
			"key": "188800",
			"updated": 1533081599000,
			"title": "Lufttemperatur - Abisko",
			"summary": "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0"
		}
		]
	}`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/20.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	t.Run("includes all stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithMaximumDailyTemperatures(context.Background(), true)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithMaximumDailyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "20",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "max, 1 gång per dygn",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
				{
					Name:      "Abisko",
					Owner:     "SMHI",
					ID:        188800,
					Height:    388.0,
					Latitude:  68.3557,
					Longitude: 18.8206,
					Active:    false,
					Key:       "188800",
					Updated:   1533081599000,
					Title:     "Lufttemperatur - Abisko",
					Summary:   "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithMaximumDailyTemperatures returned %+v, want %+v", p, want)
		}
	})
	t.Run("includes only active stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithMaximumDailyTemperatures(context.Background(), false)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithMaximumDailyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "20",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "max, 1 gång per dygn",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithMaximumDailyTemperatures returned %+v, want %+v", p, want)
		}
	})

}

func TestTemperatureService_GetStationsWithHourlyTemperatures_returnsOK(t *testing.T) {

	data := `{
		"key": "1",
		"updated": 1533495600000,
		"title": "Lufttemperatur: Välj station (sedan tidsutsnitt)",
		"summary": "momentanvärde, 1 gång/tim",
		"valueType": "SAMPLING",
		"station": [
		{
			"name": "Abelvattnet Aut",
			"owner": "SMHI",
			"id": 154860,
			"height": 665.0,
			"latitude": 65.53,
			"longitude": 14.97,
			"active": true,
			"key": "154860",
			"updated": 841535999000,
			"title": "Lufttemperatur - Abelvattnet Aut",
			"summary": "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0"
		},
		{
			"name": "Abisko",
			"owner": "SMHI",
			"id": 188800,
			"height": 388.0,
			"latitude": 68.3557,
			"longitude": 18.8206,
			"active": false,
			"key": "188800",
			"updated": 1533081599000,
			"title": "Lufttemperatur - Abisko",
			"summary": "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0"
		}
		]
	}`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/1.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	t.Run("includes all stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithHourlyTemperatures(context.Background(), true)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithHourlyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "1",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "momentanvärde, 1 gång/tim",
			ValueType: "SAMPLING",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
				{
					Name:      "Abisko",
					Owner:     "SMHI",
					ID:        188800,
					Height:    388.0,
					Latitude:  68.3557,
					Longitude: 18.8206,
					Active:    false,
					Key:       "188800",
					Updated:   1533081599000,
					Title:     "Lufttemperatur - Abisko",
					Summary:   "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithHourlyTemperatures returned %+v, want %+v", p, want)
		}
	})
	t.Run("includes only active stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithHourlyTemperatures(context.Background(), false)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithHourlyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "1",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "momentanvärde, 1 gång/tim",
			ValueType: "SAMPLING",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithHourlyTemperatures returned %+v, want %+v", p, want)
		}
	})
}

func TestTemperatureService_GetStationsWithAverageMonthlyTemperatures_returnsOK(t *testing.T) {

	data := `{
		"key": "22",
		"updated": 1533495600000,
		"title": "Lufttemperatur: Välj station (sedan tidsutsnitt)",
		"summary": "medel, 1 gång per månad",
		"valueType": "INTERVAL",
		"station": [
		{
			"name": "Abelvattnet Aut",
			"owner": "SMHI",
			"id": 154860,
			"height": 665.0,
			"latitude": 65.53,
			"longitude": 14.97,
			"active": true,
			"key": "154860",
			"updated": 841535999000,
			"title": "Lufttemperatur - Abelvattnet Aut",
			"summary": "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0"
		},
		{
			"name": "Abisko",
			"owner": "SMHI",
			"id": 188800,
			"height": 388.0,
			"latitude": 68.3557,
			"longitude": 18.8206,
			"active": false,
			"key": "188800",
			"updated": 1533081599000,
			"title": "Lufttemperatur - Abisko",
			"summary": "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0"
		}
		]
	}`

	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/version/latest/parameter/22.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, data)
	})

	t.Run("includes all stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithAverageMonthlyTemperatures(context.Background(), true)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithAverageMonthlyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "22",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "medel, 1 gång per månad",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
				{
					Name:      "Abisko",
					Owner:     "SMHI",
					ID:        188800,
					Height:    388.0,
					Latitude:  68.3557,
					Longitude: 18.8206,
					Active:    false,
					Key:       "188800",
					Updated:   1533081599000,
					Title:     "Lufttemperatur - Abisko",
					Summary:   "Latitud: 68.3557 Longitud: 18.8206 Höjd: 388.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithAverageMonthlyTemperatures returned %+v, want %+v", p, want)
		}
	})
	t.Run("includes only active stations", func(t *testing.T) {
		p, _, err := client.Temperatures.GetStationsWithAverageMonthlyTemperatures(context.Background(), false)
		if err != nil {
			t.Errorf("Temperatures.GetStationsWithAverageMonthlyTemperatures returned error: %v", err)
		}

		want := &Parameter{
			Key:       "22",
			Updated:   1533495600000,
			Title:     "Lufttemperatur: Välj station (sedan tidsutsnitt)",
			Summary:   "medel, 1 gång per månad",
			ValueType: "INTERVAL",
			Station: []Station{
				{
					Name:      "Abelvattnet Aut",
					Owner:     "SMHI",
					ID:        154860,
					Height:    665.0,
					Latitude:  65.53,
					Longitude: 14.97,
					Active:    true,
					Key:       "154860",
					Updated:   841535999000,
					Title:     "Lufttemperatur - Abelvattnet Aut",
					Summary:   "Latitud: 65.5300 Longitud: 14.9700 Höjd: 665.0",
				},
			},
		}

		if !reflect.DeepEqual(p, want) {
			t.Errorf("Temperatures.GetStationsWithAverageMonthlyTemperatures returned %+v, want %+v", p, want)
		}
	})
}
