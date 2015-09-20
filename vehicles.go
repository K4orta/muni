package transit

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type VehicalResponse struct {
	XMLName  xml.Name   `xml:"body"`
	Vehicles []*Vehicle `xml:"vehicle" json:"vehicles"`
	LastTime LastT      `xml:"lastTime" json:"lastTime"`
}

type LastT struct {
	Time int64 `xml:"time,attr" json:"time"`
}

type Vehicle struct {
	Id               string    `xml:"id,attr" json:"id"`
	RouteTag         string    `xml:"routeTag,attr" json:"routeTag"`
	Lat              float64   `xml:"lat,attr" json:"lat"`
	Lng              float64   `xml:"lon,attr" json:"lng"`
	DirTag           string    `xml:"dirTag,attr" json:"dirTag"`
	Heading          int       `xml:"heading,attr" json:"heading"`
	LeadingVehicleId string    `xml:"leadingVehicleId,attr" json:"leadingVehicleId"`
	Predictable      bool      `xml:"predictable,attr" json:"predictalbe"`
	SpeedKmHr        float32   `xml:"speedKmHr,attr" json:"speedKmHr"`
	SecsSinceReport  int       `xml:"secsSinceReport,attr" json:"secsSinceReport"`
	TimeLogged       time.Time `json:"timeLogged"`
}

var (
	apiUrl = "http://webservices.nextbus.com/service/publicXMLFeed?command=vehicleLocations&a=sf-muni"
)

var lastRequestTimes = map[string]int64{}

func GetVehiclesData(route string) (*VehicalResponse, error) {
	lastTime := LastRequestTime(route)
	client := &http.Client{}
	resp, err := http.Get(apiUrl + "&r=" + route + "&t=" + strconv.FormatInt(lastTime, 10))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var vr VehicalResponse

	xml.Unmarshal([]byte(b), &vr)
	lastRequestTimes[route] = vr.LastTime.Time
	return &vr, nil
}

func LastRequestTime(route string) int64 {
	if val, ok := lastRequestTimes[route]; ok {
		return val
	}
	return 0
}
