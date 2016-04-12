package muni

import (
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"time"
)

type vehicalResponse struct {
	XMLName  xml.Name   `xml:"body" json:"-"`
	Vehicles []*Vehicle `xml:"vehicle" json:"vehicles"`
	LastTime LastT      `xml:"lastTime" json:"lastTime"`
}

type LastT struct {
	Time int64 `xml:"time,attr" json:"time"`
}

type Vehicle struct {
	Id               string    `xml:"id,attr" json:"id" db:"vehicle_id"`
	RouteTag         string    `xml:"routeTag,attr" json:"routeTag" db:"route_tag"`
	Lat              float64   `xml:"lat,attr" json:"lat" db:"lat"`
	Lng              float64   `xml:"lon,attr" json:"lng" db:"lng"`
	DirTag           string    `xml:"dirTag,attr" json:"dirTag" db:"dir_tag"`
	Heading          int       `xml:"heading,attr" json:"heading" db:"heading"`
	LeadingVehicleId string    `xml:"leadingVehicleId,attr" json:"leadingVehicleId" db:"leading_vehicle_id"`
	Predictable      bool      `xml:"predictable,attr" json:"predictable" db:"predictable"`
	SpeedKmHr        float32   `xml:"speedKmHr,attr" json:"speedKmHr" db:"speed_km_hr"`
	SecsSinceReport  int       `xml:"secsSinceReport,attr" json:"secsSinceReport" db:"secs_since_report"`
	TimeRecieved     time.Time `json:"timeRecieved" db:"time_recieved"`
}

var (
	apiUrl = "http://webservices.nextbus.com/service/publicXMLFeed?command=vehicleLocations&a=sf-muni"
)

var lastRequestTimes = map[string]int64{}

func GetMultiVehicleData(routes []string) ([]*vehicalResponse, error) {
	routeData := []*vehicalResponse{}
	for _, r := range routes {
		vehicles, err := GetVehiclesData(r)
		if err != nil {
			return routeData, err
		}
		routeData = append(routeData, vehicles)
	}
	return routeData, nil
}

func GetVehiclesData(route string) (*vehicalResponse, error) {
	lastTime := LastRequestTime(route)
	requestURL := apiUrl + "&r=" + route + "&t=" + strconv.FormatInt(lastTime, 10)

	resp, err := transitRequest(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var vr vehicalResponse

	xml.Unmarshal([]byte(b), &vr)
	lastRequestTimes[route] = vr.LastTime.Time
	for _, v := range vr.Vehicles {
		v.TimeRecieved = time.Unix(vr.LastTime.Time, 0)
	}

	return &vr, nil
}

func LastRequestTime(route string) int64 {
	if val, ok := lastRequestTimes[route]; ok {
		return val
	}
	return 0
}
