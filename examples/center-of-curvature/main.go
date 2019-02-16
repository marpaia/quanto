package main

import (
	"fmt"

	"github.com/marpaia/math-go/vector"
)

func main() {
	// the position of the particle in meters
	r := vector.Vector{250, 630, 430}

	// the velocity of the particle in meters per second
	v := vector.Vector{90, 125, 170}

	// the acceleration of the particle in meters per second squared
	a := vector.Vector{16, 125, 30}

	coc := vector.CenterOfCurvature(r, v, a)
	fmt.Printf("the center of curvature is %+v\n", coc)
}
