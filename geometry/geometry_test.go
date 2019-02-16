package geometry_test

import (
	"testing"

	"github.com/marpaia/quanto/geometry"
	"github.com/marpaia/quanto/vector"
)

func TestCenterOfCurvature(t *testing.T) {
	r := vector.Vector{250, 630, 430}
	v := vector.Vector{90, 125, 170}
	a := vector.Vector{16, 125, 30}

	coc := geometry.CenterOfCurvature(r, v, a)

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
