package transit

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUnmarshalStops(t *testing.T) {
	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/xml")
		fmt.Fprint(w, `
			<?xml version="1.0" encoding="utf-8" ?> 
			<body copyright="All data copyright San Francisco Muni 2015.">
				<route tag="N" title="N-Judah" color="003399" oppositeColor="ffffff" latMin="37.7601699" latMax="37.7932299" lonMin="-122.5092" lonMax="-122.38798">
					<stop tag="5240" title="King St &amp; 4th St" lat="37.7760599" lon="-122.39436" stopId="15240"/>
					<stop tag="5237" title="King St &amp; 2nd St" lat="37.7796199" lon="-122.38982" stopId="15237"/>
					<stop tag="7145" title="The Embarcadero &amp; Brannan St" lat="37.7846299" lon="-122.38798" stopId="17145"/>
					<stop tag="4510" title="Embarcadero Folsom St" lat="37.7907499" lon="-122.3898399" stopId="14510"/>
					<stop tag="7217" title="Embarcadero Station Outbound" lat="37.7932299" lon="-122.39654" stopId="17217"/>
					<direction tag="N__IB1" title="Inbound to Caltrain via Downtown" name="Inbound" useForUI="true">
						<stop tag="5223" />
						<stop tag="5216" />
						<stop tag="5214" />
						<stop tag="5212" />
					</direction>
					<direction tag="N__OB1" title="Outbound to Ocean Beach via Downtown" name="Outbound" useForUI="true">
						<stop tag="5240" />
						<stop tag="5237" />
						<stop tag="7145" />
						<stop tag="4510" />
					</direction>
					<path>
						<point lat="37.76017" lon="-122.50878"/>
						<point lat="37.7603" lon="-122.50812"/>
						<point lat="37.76039" lon="-122.50606"/>
						<point lat="37.76052" lon="-122.50284"/>
						<point lat="37.76068" lon="-122.49915"/>
						<point lat="37.76083" lon="-122.49596"/>
					</path>
				</route>
			</body>
		 `)
	}))
	stopApiUrl = fakeServer.URL + "/"

	sd, err := GetStopData("N")
	if err != nil {
		t.Error("Test failed", err)
	}

	if sd == nil {
		t.Error("Failed to unmashal route", sd)
	}

	if sd.Routes[0].Paths == nil {
		t.Error("Failed to unmashal paths", sd.Routes)
	}

	if len(sd.Routes[0].Stops) != 5 {
		t.Error("Failed to unmashal routes", sd.Routes)
	}

}
