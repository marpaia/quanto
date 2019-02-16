package geometry

import (
	"math"

	"github.com/marpaia/quanto/vector"
)

// CenterOfCurvature calculates the coordinates of the center of curvature given
// the position, velocity and acceleration of a kinematic point.
//
// r is the position of a particle in meters
// v is the velocity of a particle in meters per second
// a is the acceleration of a particle in meters per second squared
func CenterOfCurvature(r vector.Vector, v vector.Vector, a vector.Vector) vector.Vector {
	vMagnitude := vector.Magnitude(v)
	vCrossA := vector.CrossProduct(v, a)
	vCrossAMagnitude := vector.Magnitude(vCrossA)

	unitTangentVector := vector.Vector{
		I: v.I / vMagnitude,
		J: v.J / vMagnitude,
		K: v.K / vMagnitude,
	}

	unitBinormalVector := vector.Vector{
		I: vCrossA.I / vCrossAMagnitude,
		J: vCrossA.J / vCrossAMagnitude,
		K: vCrossA.K / vCrossAMagnitude,
	}

	unitPrincipalNormalVector := vector.CrossProduct(unitBinormalVector, unitTangentVector)

	radiusOfCurvature := math.Pow(vMagnitude, 2) / vector.Project(a, unitPrincipalNormalVector)

	return vector.Vector{
		I: r.I + radiusOfCurvature*unitPrincipalNormalVector.I,
		J: r.J + radiusOfCurvature*unitPrincipalNormalVector.J,
		K: r.K + radiusOfCurvature*unitPrincipalNormalVector.K,
	}
}
