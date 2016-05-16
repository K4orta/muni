package muni

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"time"
)

type vehicleResponse struct {
	XMLName  xml.Name   `xml:"body" json:"-"`
	Vehicles []*Vehicle `xml:"vehicle" json:"vehicles"`
	LastTime LastT      `xml:"lastTime" json:"lastTime"`
}

// LastT is a type used to unmashal the last request time from the XML response.
type LastT struct {
	Time int64 `xml:"time,attr" json:"time"`
}

// Vehicle Model
type Vehicle struct {
	ID               string    `xml:"id,attr" json:"id" db:"vehicle_id"`
	RouteTag         string    `xml:"routeTag,attr" json:"routeTag" db:"route_tag"`
	Lat              float64   `xml:"lat,attr" json:"lat" db:"lat"`
	Lng              float64   `xml:"lon,attr" json:"lng" db:"lng"`
	DirTag           string    `xml:"dirTag,attr" json:"dirTag" db:"dir_tag"`
	Heading          int       `xml:"heading,attr" json:"heading" db:"heading"`
	LeadingVehicleID string    `xml:"leadingVehicleId,attr" json:"leadingVehicleId" db:"leading_vehicle_id"`
	Predictable      bool      `xml:"predictable,attr" json:"predictable" db:"predictable"`
	SpeedKmHr        float32   `xml:"speedKmHr,attr" json:"speedKmHr" db:"speed_km_hr"`
	SecsSinceReport  int       `xml:"secsSinceReport,attr" json:"secsSinceReport" db:"secs_since_report"`
	TimeReceived     time.Time `json:"timeReceived" db:"time_received"`
}

var lastRequestTimes = map[string]int64{}

// GetMultiVehicleData takes an array of strings and runs GetVehiclesData for each.
func GetMultiVehicleData(routes []string) ([]*Vehicle, error) {
	routeData := []*Vehicle{}
	for _, r := range routes {
		vehicles, err := GetVehiclesData(r)
		if err != nil {
			return routeData, err
		}
		routeData = append(routeData, vehicles...)
	}
	return routeData, nil
}

// GetVehiclesData requests XML data from the api and parses it.
func GetVehiclesData(route string) ([]*Vehicle, error) {
	lastTime := LastRequestTime(route)
	requestQuery := "&command=vehicleLocations&r=" + route + "&t=" + strconv.FormatInt(lastTime, 10)

	resp, err := transitRequest(requestQuery)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var vr vehicleResponse

	xml.Unmarshal([]byte(b), &vr)
	lastRequestTimes[route] = vr.LastTime.Time
	for _, v := range vr.Vehicles {
		v.TimeReceived = parseTime(vr.LastTime.Time)
	}

	return vr.Vehicles, nil
}

func parseTime(input int64) time.Time {
	return time.Unix(input/1000, 0)
}

// LastRequestTime gets the last time a request was made for a specific route, or zero.
func LastRequestTime(route string) int64 {
	if val, ok := lastRequestTimes[route]; ok {
		return val
	}
	return 0
}
