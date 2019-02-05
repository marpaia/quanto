# Mathematics in Go [![CircleCI](https://circleci.com/gh/marpaia/math-go/tree/master.svg?style=svg&circle-token=07906e84b293e18c892eccc2680a545094323dc4)](https://circleci.com/gh/marpaia/math-go/tree/master)

This repository is a Go library for performing a wide range of mathematical programming tasks.

See the following sections of this document for more information:

- [Introduction](#introduction)
- [Features](#features)
  - [Multivariable Expression Algebraic Operations](#multivariable-expression-algebraic-operations)

In addition, here are some other documents that may be helpful:

- [Documentation](./docs/README.md)
- [Developer Guide](./docs/developer-guide.md)

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
```

This will print:

```
5x^2 * 2y^3 = 10(x^2)(y^3)
```
