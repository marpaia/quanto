package main

import (
	"fmt"

	"github.com/marpaia/quanto/vector"
)

func main() {
	lhs := vector.Vector{1, 6, 18}
	rhs := vector.Vector{4, 2, 1}

	fmt.Printf("the magnitude of lhs is %f\n", vector.Magnitude(lhs))
	fmt.Printf("the magnitude of rhs is %f\n", vector.Magnitude(rhs))

	fmt.Printf("the unit vector of lhs is %+v\n", vector.UnitVector(lhs))
	fmt.Printf("the unit vector of rhs is %+v\n", vector.UnitVector(rhs))

	cross := vector.CrossProduct(lhs, rhs)
	fmt.Printf("lhs cross rhs is %+v\n", cross)

	dot := vector.DotProduct(lhs, rhs)
	fmt.Printf("lhs dot rhs is %+v\n", dot)

	projection := vector.Project(lhs, rhs)
	fmt.Printf("the magnitude of the projection of lhs in the direction of rhs is %f\n", projection)

	angle := vector.AngleDegrees(lhs, rhs)
	fmt.Printf("the angle between lhs and rhs is %d degrees\n", int(angle))
}
