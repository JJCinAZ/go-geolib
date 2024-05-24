package geolib

import (
	"math"
)

// Semi-axes of WGS-84 geoidal reference
const (
	// Major semiaxis (meters)
	WGS84A = 6378137.0
	// Minor semiaxis (meters)
	WGS84B = 6356752.3
)

// BoundingBox represents the geo-polygon that encompasses the given point and radius
type BoundingBox struct {
	LatMin float64
	LatMax float64
	LonMin float64
	LonMax float64
}

// Get the Earth's radius in meters at a given latitude based on the WGS84 ellipsoid
func getWgs84EarthRadius(lat float64) float64 {
	an := WGS84A * WGS84A * math.Cos(lat)
	bn := WGS84B * WGS84B * math.Sin(lat)

	ad := WGS84A * math.Cos(lat)
	bd := WGS84B * math.Sin(lat)

	return math.Sqrt((an*an + bn*bn) / (ad*ad + bd*bd))
}

// GetBoundingBox calculates a target point from a starting point, distance, and bearing
// Parameters:
//   - φ1: latitude of the starting point in degrees
//   - λ1: longitude of the starting point in degrees
//   - distance: distance from the starting point to the target point in kilometers
//   - bearing: bearing from the starting point to the target point in degrees (0-360)
//
// Result: new point
//   - φ2: latitude of the target point in degrees
//   - λ2: longitude of the target point in degrees
//
// GetBoundingBox returns a BoundingBox encompassing the given lat/long point and radius
func GetBoundingBox(φ1 float64, λ1 float64, radiusKm float64) BoundingBox {
	φ1 = Deg2Rad(φ1)
	λ1 = Deg2Rad(λ1)
	halfSide := 1000 * radiusKm

	// Radius of Earth at given latitude
	radius := getWgs84EarthRadius(φ1)

	pradius := radius * math.Cos(φ1)

	latMin := φ1 - halfSide/radius
	latMax := φ1 + halfSide/radius
	lonMin := λ1 - halfSide/pradius
	lonMax := λ1 + halfSide/pradius

	return BoundingBox{
		LatMin: Rad2Deg(latMin),
		LatMax: Rad2Deg(latMax),
		LonMin: Rad2Deg(lonMin),
		LonMax: Rad2Deg(lonMax),
	}
}
