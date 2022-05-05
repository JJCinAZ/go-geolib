go-geolib
=========

## About

Geographical functions for Golang apps!

Copied from https://github.com/alouche/go-geolib

## installation

	go get github.com/jjcinaz/go-geolib

## Symbolism

    φ = latitude
    λ = longitude

## Example Usage

	package main																																												

	import (
		"fmt"
		"github.com/jjcinaz/go-geolib"
	)

	func main() {
		haversine_distance := fmt.Sprintf("%2.f", geolib.Haversine_Distance(50.116667, 8.683333, 52.516667, 13.3833))
		println("Havesine Distance:", haversine_distance, "km")
	}
