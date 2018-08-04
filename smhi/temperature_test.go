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
		fmt.Fprint(w, `{"value": [{"from": 1533254401000, "to": 1533340800000, "ref": "2018-08-03", "value": "21.8", "quality": "Y"}]}`)
	})

	temps, _, err := client.Temperatures.GetAverageDailyTemperatures(context.Background(), 12345, PeriodLatestDay, FormatJSON)
	if err != nil {
		t.Errorf("Temperatures.GetAverageDailyTemperatures returned error: %v", err)
	}

	want := &TemperatureData{Value: []TemperatureDataValue{{From: 1533254401000, To: 1533340800000, Ref: "2018-08-03", Value: "21.8", Quality: "Y"}}}
	if !reflect.DeepEqual(temps, want) {
		t.Errorf("Temperatures.GetAverageDailyTemperatures returned %+v, want %+v", temps, want)
	}
}
