package geolib

import "math"

// Terminal calculates a target point from a starting point, distance, and bearing
// Result: new point
//  - Type: float64
// distance is in kilometers and bearing in degrees
func Terminal(φ1, λ1, distance, bearing float64) (φ2, λ2 float64) {
	R := 6378.1 // Radius of the Earth
	φ1 = Deg2Rad(φ1)
	λ1 = Deg2Rad(λ1)
	bearing = Deg2Rad(bearing)
	φ2 = math.Asin(math.Sin(φ1)*math.Cos(distance/R) + math.Cos(φ1)*
		math.Sin(distance/GreatEarthCircleRadiusKM)*math.Cos(bearing))
	λ2 = λ1 + math.Atan2(math.Sin(bearing)*math.Sin(distance/GreatEarthCircleRadiusKM)*math.Cos(φ1),
		math.Cos(distance/GreatEarthCircleRadiusKM)-math.Sin(φ1)*math.Sin(φ2))
	return Rad2Deg(φ2), Rad2Deg(λ2)
}
