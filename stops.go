package transit

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type StopResponse struct {
	XMLName xml.Name `xml:"body"`
	Routes  []*Route `xml:"route"`
}

type Route struct {
	Title string  `xml:"title,attr" json:"title"`
	Stops []*Stop `xml:"stop" json:"stops"`
	Paths []*Path `xml:"path" json:"paths"`
}

type Path struct {
	Points []*Point `xml:"point" json:"points"`
}

type Point struct {
	Lat float64 `xml:"lat,attr" json:"lat"`
	Lng float64 `xml:"lon,attr" json:"lng"`
}

type Stop struct {
	Title  string  `xml:"title,attr" json:"title"`
	Tag    string  `xml:"tag,attr" json:"tag"`
	StopId string  `xml:"stopId,attr" json:"stopId"`
	Lat    float64 `xml:"lat,attr" json:"lat"`
	Lng    float64 `xml:"lon,attr" json:"lng"`
}

var (
	stopApiUrl = "http://webservices.nextbus.com/service/publicXMLFeed?command=routeConfig&a=sf-muni&r="
)

func GetStopData(route string) (*StopResponse, error) {
	resp, err := http.Get(stopApiUrl + route)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var vr StopResponse
	xml.Unmarshal([]byte(b), &vr)
	return &vr, nil
}
