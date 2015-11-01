package muni

import (
	"encoding/xml"
	"io/ioutil"
)

// StopResponse is a struct representing the XML document returned by the Nextbus API
type StopResponse struct {
	XMLName xml.Name `xml:"body"`
	Routes  []*Route `xml:"route"`
}

// Route is a struct that contains the route path and stops along it
type Route struct {
	Title string  `xml:"title,attr" json:"title"`
	Stops []*Stop `xml:"stop" json:"stops"`
	Paths []*Path `xml:"path" json:"paths"`
}

// Path is an array of points
type Path struct {
	Points []*Point `xml:"point" json:"points"`
}

// Point is a Lat Lng representation of of geographical coordinate
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
	resp, err := transitRequest(stopApiUrl + route)
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
