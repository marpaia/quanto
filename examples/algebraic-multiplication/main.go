package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/marpaia/math-go/algebra"
)

func main() {
	fiveXSquared := algebra.NewSymbol("x").Power(2).Coefficient(5)
	twoYCubed := algebra.NewSymbol("y").Power(3).Coefficient(2)
	product, err := algebra.Mul(fiveXSquared, twoYCubed)
	if err != nil {
		fmt.Printf("error: %q\n", err)
		os.Exit(1)
	}

	tenXSquaredYCubed := algebra.NewMultivariableExpression(
		algebra.NewSymbol("x").Power(2),
		algebra.NewSymbol("y").Power(3),
	).Coefficient(10)

	if reflect.DeepEqual(product, tenXSquaredYCubed) {
		fmt.Println("5x^2 * 2y^3 = 10(x^2)(y^3)")
	} else {
		fmt.Println("An error occurred")
		os.Exit(1)
	}
}
