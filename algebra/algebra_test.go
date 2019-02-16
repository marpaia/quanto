package algebra_test

import (
	"testing"

	"github.com/marpaia/math-go/algebra"
	"github.com/marpaia/math-go/internal/assert"
)

func TestSimplifyMultivariableExpression(t *testing.T) {
	tenXSquaredYCubed := algebra.NewMultivariableExpression(
		algebra.NewSymbol("x").Power(2).Coefficient(10),
		algebra.NewSymbol("y").Power(3),
	)

	tenXSquaredYCubedSimplified := algebra.NewMultivariableExpression(
		algebra.NewSymbol("x").Power(2),
		algebra.NewSymbol("y").Power(3),
	).Coefficient(10)

	if err := assert.Equal(tenXSquaredYCubedSimplified, tenXSquaredYCubed.Simplify()); err != nil {
		t.Fatal(err)
	}
}

func TestExpressionArithmetic(t *testing.T) {
	fiveXSquared := algebra.NewSymbol("x").Power(2).Coefficient(5)
	twoYCubed := algebra.NewSymbol("y").Power(3).Coefficient(2)

	tenXSquaredYCubed := algebra.NewMultivariableExpression(
		algebra.NewSymbol("x").Power(2),
		algebra.NewSymbol("y").Power(3),
	).Coefficient(10)

	product1, err := algebra.Mul(fiveXSquared, twoYCubed)
	if err != nil {
		t.Fatal(err)
	}
	if err := assert.Equal(tenXSquaredYCubed.Simplify(), product1.Simplify()); err != nil {
		t.Fatal(err)
	}

	twentyfiveXFourth := algebra.NewSymbol("x").Power(4).Coefficient(25).MultivariableExpression()

	product2, err := algebra.Mul(fiveXSquared, fiveXSquared)
	if err != nil {
		t.Fatal(err)
	}
	if err := assert.Equal(twentyfiveXFourth.Simplify(), product2.Simplify()); err != nil {
		t.Fatal(err)
	}
}
