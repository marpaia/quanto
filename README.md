# Quanto [![CircleCI](https://circleci.com/gh/marpaia/quanto.svg?style=svg&circle-token=07906e84b293e18c892eccc2680a545094323dc4)](https://circleci.com/gh/marpaia/quanto)

Quanto is a Go library for performing a wide range of mathematical programming tasks.

See the following sections of this document for more information:

- [Introduction](#introduction)
- [Features](#features)
  - [Multivariable Expression Algebraic Operations](#multivariable-expression-algebraic-operations)
  - [Vector Arithmetic](#vector-arithmetic)
  - [Center of Curvature](#center-of-curvature)

In addition, here are some other documents and links that may be helpful:

- For development information, see the [Contributor Guide](./CONTRIBUTING.md)
- If you would like to file a bug or request a feature, please file a [GitHub Issue](https://github.com/marpaia/quanto/issues/new)

## Introduction

In the mathematical scientific programming world, we generally write code in C++, Python, or perhaps Prolog. Julia is an exciting up-and-comer, while some more statistics-heavy use-cases drive some to R. There are also proprietary languages that are popular like MATLAB and Wolfram Mathematica.

While C++ is a fast language, it is not a very productive language. While Python is a productive language, it is not a very fast language. While Prolog, Julia and R are all nice for some use-cases, none are often used as a general purpose programming language. Finally, while MATLAB and Wolfram Mathematica are excellent products, it's not always possible to use these proprietary tools.

Go, on the other hand, is both fast and productive. It has a large mindshare throughout several industries and it is an excellent general purpose programming language. Finally, it is also completely free and open source.

While Go is an excellent language to use for a number of use-cases (server software, cross-platform software, internet-of-things software, etc), the library ecosystem for scientific programming is less robust than Python or Julia or R. This library aims to fill-in part of that gap by providing a large set of well-tested, idiomatic, mathematical library code written in Go.

## Features

### Multivariable Expression Algebraic Operations

```go
package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/marpaia/quanto/algebra"
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
```

This will print:

```
5x^2 * 2y^3 = 10(x^2)(y^3)
```

### Vector Arithmetic

```go
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
```

This will print:

```
the magnitude of lhs is 19.000000
the magnitude of rhs is 4.582576
the unit vector of lhs is {I:0.05263157894736842 J:0.3157894736842105 K:0.9473684210526315}
the unit vector of rhs is {I:0.8728715609439696 J:0.4364357804719848 K:0.2182178902359924}
lhs cross rhs is {I:-30 J:71 K:-22}
lhs dot rhs is 34
the magnitude of the projection of lhs in the direction of rhs is 7.419408
the angle between lhs and rhs is 67 degrees
```

### Center of Curvature

```go
package main

import (
	"fmt"

	"github.com/marpaia/quanto/geometry"
	"github.com/marpaia/quanto/vector"
)

func main() {
	// the position of the particle in meters
	r := vector.Vector{250, 630, 430}

	// the velocity of the particle in meters per second
	v := vector.Vector{90, 125, 170}

	// the acceleration of the particle in meters per second squared
	a := vector.Vector{16, 125, 30}

	coc := geometry.CenterOfCurvature(r, v, a)
	fmt.Printf("the center of curvature is %+v\n", coc)
}

```

This will print:

```
the center of curvature is {I:95.15892343003512 J:1141.3912469446514 K:135.95171190126717}
```