package algebra_test

import (
	"testing"

	"github.com/marpaia/quanto/algebra"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, tenXSquaredYCubedSimplified, tenXSquaredYCubed.Simplify())
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
	assert.Equal(t, tenXSquaredYCubed.Simplify(), product1.Simplify())

	twentyfiveXFourth := algebra.NewSymbol("x").Power(4).Coefficient(25).MultivariableExpression()

	product2, err := algebra.Mul(fiveXSquared, fiveXSquared)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, twentyfiveXFourth.Simplify(), product2.Simplify())
}
