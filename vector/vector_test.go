package vector_test

import (
	"testing"

	"github.com/marpaia/quanto/vector"
)

func TestMagnitude(t *testing.T) {
	v := vector.Vector{1, -4, 8}
	if m := vector.Magnitude(v); m != 9 {
		t.Fatalf("magnitude should be 9, got %f", m)
	}
}

func TestUnitVector(t *testing.T) {
	v := vector.Vector{1, -4, 8}
	u := vector.UnitVector(v)
	if u.I != (float64(1) / float64(9)) {
		t.Fatalf("u.I should be 1/9 but it is %f", u.I)
	}
	if u.J != (float64(-4) / float64(9)) {
		t.Fatalf("u.J should be -4/9 but it is %f", u.J)
	}
	if u.K != (float64(8) / float64(9)) {
		t.Fatalf("u.K should be 8/9 but it is %f", u.K)
	}
}

func TestDirection(t *testing.T) {
	v := vector.Vector{1, -4, 8}
	d := vector.DirectionDegrees(v)
	if int(d.X) != 83 {
		t.Fatalf("d.X should be 83 but it is %d", int(d.X))
	}
	if int(d.Y) != 116 {
		t.Fatalf("d.Y should be 116 but it is %d", int(d.Y))
	}
	if int(d.Z) != 27 {
		t.Fatalf("d.Z should be 27 but it is %d", int(d.Z))
	}
}

func TestDotProduct(t *testing.T) {
	lhs := vector.Vector{1, 6, 18}
	rhs := vector.Vector{42, -69, 98}
	dot := vector.DotProduct(lhs, rhs)
	if dot != 1392 {
		t.Fatalf("expected dot product to be 1392 but it is %f", dot)
	}
}

func TestAngle(t *testing.T) {
	lhs := vector.Vector{1, 6, 18}
	rhs := vector.Vector{42, -69, 98}
	angle := vector.AngleDegrees(lhs, rhs)
	if int(angle) != 54 {
		t.Fatalf("expected angle to be 54 but it is %d", int(angle))
	}
}

func TestProjection(t *testing.T) {
	lhs := vector.Vector{1, 6, 18}
	rhs := vector.Vector{42, -69, 98}
	projection := vector.Project(lhs, rhs)
	if int(projection) != 10 {
		t.Fatalf("expected projection to be 10 but it is %d", int(projection))
	}
}

func TestCrossProduct(t *testing.T) {
	lhs := vector.Vector{1, 6, 18}
	rhs := vector.Vector{4, 2, 1}
	cross := vector.CrossProduct(lhs, rhs)
	if cross.I != -30 {
		t.Fatalf("expected cross.I to be -30 but it is %f", cross.I)
	}
	if cross.J != 71 {
		t.Fatalf("expected cross.J to be 71 but it is %f", cross.J)
	}
	if cross.K != -22 {
		t.Fatalf("expected cross.K to be -22 but it is %f", cross.K)
	}
}

func TestCenterOfCurvature(t *testing.T) {
	r := vector.Vector{250, 630, 430}
	v := vector.Vector{90, 125, 170}
	a := vector.Vector{16, 125, 30}

	coc := vector.CenterOfCurvature(r, v, a)

	if int(coc.I) != 95 {
		t.Fatalf("coc.I should be 95 but it is %d", int(coc.I))
	}
	if int(coc.J) != 1141 {
		t.Fatalf("coc.J should be 1141 but it is %d", int(coc.J))
	}
	if int(coc.K) != 135 {
		t.Fatalf("coc.K should be 135 but it is %d", int(coc.K))
	}
}
