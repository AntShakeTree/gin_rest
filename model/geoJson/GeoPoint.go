package geoJson

import "fmt"

type GeoPoint struct {
	Type        string                 `json:"type" bson:"type"`
	Coordinates Coordinate             `json:"coordinates" bson:"coordinates"`
	Properties  map[string]interface{} `json:"properties" bson:"properties"`
	Crs         *CRS                   `json:"crs,omitempty" bson:"crs,omitempty"`
}

type GeoPointCollection struct {
	GeoPoints []*GeoPoint `json:"geoPoints" bson:"geo_points"`
	Type      string      `json:"type" bson:"type"`
}

func (t *GeoPointCollection) AddGeoPoint(f ...*GeoPoint) {
	if f == nil {
		t.GeoPoints = make([]*GeoPoint, 0, 100)
	}
	t.GeoPoints = append(t.GeoPoints, f...)
}
func (t *GeoPointCollection) Appends(c GeoPointCollection) {
	if c.GeoPoints == nil {
		c.GeoPoints = make([]*GeoPoint, 0, 100)
	}
	for _, v := range c.GeoPoints {
		t.GeoPoints = append(t.GeoPoints, v)
	}
}

func NewGeoPointCollection(gs []*GeoPoint) *GeoPointCollection {
	return &GeoPointCollection{Type: "GeoPoints", GeoPoints: gs}
}

// AddGeometry adds geometry to coordinates.
// New value will replace existing
func (t *GeoPoint) AddGeometry(g interface{}) error {
	if c, ok := g.(Coordinate); ok {
		t.Coordinates = c
	} else {
		return fmt.Errorf("AssertionError: %v to %v", g, "Coordinate")
	}
	return nil
}

// GetType return type of geometry
func (t GeoPoint) GetType() string {
	return t.Type
}

//GetCoordinates return coordinates of geometry
func (t GeoPoint) GetCoordinates() interface{} {
	return t.Coordinates
}

//Factory function to create new object
func NewGeoPoint(c Coordinate) *GeoPoint {
	return &GeoPoint{Type: "Point", Coordinates: c}
}

func (t *GeoPoint) PutProperties(key string, v interface{}) {
	if t.Properties == nil {
		t.Properties = make(map[string]interface{})
	}
	t.Properties[key] = v
}
func (t *GeoPoint) SetProperties(ps map[string]interface{}) error {
	if ps == nil {
		return fmt.Errorf("AssertionError: %v to %v", ps, "Map")
	}
	if t.Properties == nil && ps != nil {
		t.Properties = ps
	}
	for k, v := range ps {
		t.Properties[k] = v
	}
	return nil
}
