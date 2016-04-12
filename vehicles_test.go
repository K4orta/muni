package muni

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLastRequestTimeInit(t *testing.T) {
	time := LastRequestTime("N")
	if time != 0 {
		t.Error("Expected a time of zero, got: ", time)
	}
}

func TestLastRequestTime(t *testing.T) {
	fakeServer := makeFakeServer()
	SetConfig(TransitConfig{fakeServer.URL})
	startTime := LastRequestTime("71")
	GetVehiclesData("71")
	afterTime := LastRequestTime("71")

	if startTime == afterTime {
		t.Error("Last request time was the same after calling GetVehiclesData")
	}
}

func TestLastRequestTimeMultiRoute(t *testing.T) {
	fakeServer := makeFakeServer()
	SetConfig(TransitConfig{fakeServer.URL})
	GetVehiclesData("71")
	startTime := LastRequestTime("71")
	GetVehiclesData("N")
	afterTime := LastRequestTime("71")

	if startTime != afterTime {
		t.Error("GetVehiclesData is polluting unrelated route times")
	}
}

func TestGetVehicles(t *testing.T) {
	fakeServer := makeFakeServer()
	SetConfig(TransitConfig{fakeServer.URL})

	vd, err := GetVehiclesData("N")
	if err != nil {
		t.Error("Test failed", err)
	}

	if len(vd) != 19 {
		t.Error("Failed to unmarshal vehicles")
	}

	if vd[0].TimeRecieved != time.Unix(1420919252102/1000, 0) {
		t.Error("Failed to unmarshal lastTime field")
	}
}

func TestMultipleVehicleData(t *testing.T) {
	fakeServer := makeFakeServer()
	SetConfig(TransitConfig{fakeServer.URL})

	vd, err := GetMultiVehicleData([]string{"N", "L", "J"})
	if err != nil {
		t.Error("Test failed", err)
	}

	if len(vd) != 57 {
		t.Error("MultiVehicleData did not return three sets of data")
	}
}

func TestSetTimeRecieved(t *testing.T) {
	fakeServer := makeFakeServer()
	SetConfig(TransitConfig{fakeServer.URL})
	vd, _ := GetVehiclesData("N")

	for _, v := range vd {
		// if v.TimeRecieved != time.Unix(1420919252102, 0) {
		// 	t.Error("Did not set time for each vehicle")
		// }

		if v.TimeRecieved.Year() != 2015 || v.TimeRecieved.Month() != time.January || v.TimeRecieved.Day() != 10 {
			t.Error("Did not convert year correctly. Expected 2015, got:", v.TimeRecieved.Year(), v.TimeRecieved.Month(), v.TimeRecieved.Day())
		}
	}

}

func makeFakeServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/xml")
		fmt.Fprint(w, `
			<?xml version="1.0" encoding="utf-8" ?>
			<body copyright="All data copyright San Francisco Muni 2015.">
				<vehicle id="1495" routeTag="N" dirTag="N__IB1" lat="37.7836" lon="-122.38814" secsSinceReport="25" predictable="true" heading="183" speedKmHr="7"/>
				<vehicle id="1473" routeTag="N" dirTag="N__OB1" lat="37.7733599" lon="-122.39765" secsSinceReport="62" predictable="true" heading="218" speedKmHr="0" leadingVehicleId="1433"/>
				<vehicle id="1454" routeTag="N" lat="37.75221" lon="-122.38423" secsSinceReport="25" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1420" routeTag="N" lat="37.72186" lon="-122.44708" secsSinceReport="69" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1485" routeTag="N" lat="37.75184" lon="-122.38424" secsSinceReport="14" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1534" routeTag="N" lat="37.75135" lon="-122.38422" secsSinceReport="25" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1403" routeTag="N" lat="37.76943" lon="-122.42999" secsSinceReport="134" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1419" routeTag="N" dirTag="N__IB1" lat="37.79321" lon="-122.39211" secsSinceReport="21" predictable="true" heading="136" speedKmHr="45" leadingVehicleId="1536"/>
				<vehicle id="1423" routeTag="N" lat="37.72195" lon="-122.44716" secsSinceReport="71" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1436" routeTag="N" lat="37.75198" lon="-122.38447" secsSinceReport="37" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1533" routeTag="N" lat="37.75156" lon="-122.38438" secsSinceReport="77" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1475" routeTag="N" dirTag="N__OB1" lat="37.78911" lon="-122.38812" secsSinceReport="9" predictable="true" heading="342" speedKmHr="20" leadingVehicleId="1497"/>
				<vehicle id="1477" routeTag="N" dirTag="N__IB1" lat="37.78366" lon="-122.38815" secsSinceReport="25" predictable="true" heading="183" speedKmHr="37" leadingVehicleId="1495"/>
				<vehicle id="1441" routeTag="N" lat="37.76943" lon="-122.43085" secsSinceReport="83" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1536" routeTag="N" dirTag="N__IB1" lat="37.79321" lon="-122.39211" secsSinceReport="21" predictable="true" heading="136" speedKmHr="45"/>
				<vehicle id="1482" routeTag="N" lat="37.72171" lon="-122.44717" secsSinceReport="33" predictable="false" heading="-4" speedKmHr="0"/>
				<vehicle id="1497" routeTag="N" dirTag="N__OB1" lat="37.78776" lon="-122.38781" secsSinceReport="29" predictable="true" heading="356" speedKmHr="48"/>
				<vehicle id="1523" routeTag="N" lat="37.7694" lon="-122.43055" secsSinceReport="47" predictable="false" heading="-4" speedKmHr="0" leadingVehicleId="1441"/>
				<vehicle id="1433" routeTag="N" dirTag="N__OB1" lat="37.77352" lon="-122.39754" secsSinceReport="8" predictable="true" heading="218" speedKmHr="0"/>
				<lastTime time="1420919252102"/>
			</body>
		 `)
	}))
}
