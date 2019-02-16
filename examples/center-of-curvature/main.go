package main

import (
	"fmt"

	"github.com/marpaia/quanto/geometry"
	"github.com/marpaia/quanto/vector"
)

func main() {
	// the position of the particle in meters
	r := vector.Vector{I: 250, J: 630, K: 430}

	// the velocity of the particle in meters per second
	v := vector.Vector{I: 90, J: 125, K: 170}

	// the acceleration of the particle in meters per second squared
	a := vector.Vector{I: 16, J: 125, K: 30}

	coc := geometry.CenterOfCurvature(r, v, a)
	fmt.Printf("the center of curvature is %+v\n", coc)
}
