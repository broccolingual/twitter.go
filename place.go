package main

type Place struct {
	Full_name string
	Id        string
	// TODO: contained_within
	Country      string
	Country_code string
	Geo          PlaceGeo
	Name         string
	Place_type   string
}

type PlaceGeo struct {
	Type       string
	Bbox       [4]float64
	Properties interface{}
}
