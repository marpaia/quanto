package vector

import (
	"math"
)

// Vector is an element of a vector space.
type Vector struct {
	I float64
	J float64
	K float64
}

// Direction of a vector. May be in radians or degrees, depending on what you
// use to populate it.
type Direction struct {
	X float64
	Y float64
	Z float64
}

// Magnitude represents the vector displacement of point from the origin.
func Magnitude(v Vector) float64 {
	return math.Sqrt(math.Pow(v.I, 2) + math.Pow(v.J, 2) + math.Pow(v.K, 2))
}

// UnitVector returns a vector in the same direction as v, but with a magnitude
// of 1.
func UnitVector(v Vector) Vector {
	magnitude := Magnitude(v)
	return Vector{
		I: v.I * (1 / magnitude),
		J: v.J * (1 / magnitude),
		K: v.K * (1 / magnitude),
	}
}

// DirectionRadians calculates the angles of a given vector in radians.
func DirectionRadians(v Vector) Direction {
	magnitude := Magnitude(v)
	return Direction{
		X: math.Acos(v.I / magnitude),
		Y: math.Acos(v.J / magnitude),
		Z: math.Acos(v.K / magnitude),
	}
}

// DirectionDegrees calculates the angles of a given vector in degrees.
func DirectionDegrees(v Vector) Direction {
	directionDegrees := DirectionRadians(v)
	return Direction{
		X: directionDegrees.X * (180 / math.Pi),
		Y: directionDegrees.Y * (180 / math.Pi),
		Z: directionDegrees.Z * (180 / math.Pi),
	}
}

// AngleDegrees calculuates the andgle between the two given vectors in degrees.
func AngleDegrees(lhs Vector, rhs Vector) float64 {
	return AngleRadians(lhs, rhs) * (180 / math.Pi)
}

// AngleRadians calculuates the andgle between the two given vectors in radiaans.
func AngleRadians(lhs Vector, rhs Vector) float64 {
	return math.Acos(DotProduct(lhs, rhs) / (Magnitude(lhs) * Magnitude(rhs)))
}

// DotProduct is is an algebraic operation that takes two equal-length sequences
// of numbers and returns a single number.
func DotProduct(lhs Vector, rhs Vector) float64 {
	return (lhs.I * rhs.I) + (lhs.J * rhs.J) + (lhs.K * rhs.K)
}

// Project calculates the magnitude of the projection of lhs in the direction of
// rhs.
func Project(lhs Vector, rhs Vector) float64 {
	return DotProduct(lhs, rhs) / Magnitude(rhs)
}

// CrossProduct is a binary operation on two vectors in three-dimensional space.
// Given two linearly independent vectors, the cross product is a vector that is
// perpendicular to both and and thus normal to the plane containing them.
func CrossProduct(lhs Vector, rhs Vector) Vector {
	return Vector{
		I: lhs.J*rhs.K - lhs.K*rhs.J,
		J: lhs.K*rhs.I - lhs.I*rhs.K,
		K: lhs.I*rhs.J - lhs.J*rhs.I,
	}
}

// CenterOfCurvature calculates the coordinates of the center of curvature given
// the position, velocity and acceleration of a kinematic point.
//
// r is the position of a particle in meters
// v is the velocity of a particle in meters per second
// a is the acceleration of a particle in meters per second squared
func CenterOfCurvature(r Vector, v Vector, a Vector) Vector {
	vMagnitude := Magnitude(v)
	vCrossA := CrossProduct(v, a)
	vCrossAMagnitude := Magnitude(vCrossA)

	unitTangentVector := Vector{
		I: v.I / vMagnitude,
		J: v.J / vMagnitude,
		K: v.K / vMagnitude,
	}

	unitBinormalVector := Vector{
		I: vCrossA.I / vCrossAMagnitude,
		J: vCrossA.J / vCrossAMagnitude,
		K: vCrossA.K / vCrossAMagnitude,
	}

	unitPrincipalNormalVector := CrossProduct(unitBinormalVector, unitTangentVector)

	radiusOfCurvature := math.Pow(vMagnitude, 2) / Project(a, unitPrincipalNormalVector)

	return Vector{
		I: r.I + radiusOfCurvature*unitPrincipalNormalVector.I,
		J: r.J + radiusOfCurvature*unitPrincipalNormalVector.J,
		K: r.K + radiusOfCurvature*unitPrincipalNormalVector.K,
	}
}
