package algebra

// Symbol is a variable in a polynomial or expression like "x" or "y"
type Symbol struct {
	name string
}

// Expression is a simple mathematical expression like -1.4(x^2)
type Expression struct {
	symbol      Symbol
	power       float64
	coefficient float64
}

// MultivariableSymbol is a mathematical expression like 10(x^2)(y^3)
type MultivariableExpression struct {
	expressions []Expression
	coefficient float64
}

// NewSymbol returns a new symbol struct with the given name
func NewSymbol(name string) Symbol {
	return Symbol{
		name: name,
	}
}

// Power returns an expression where a given symbol is raised to a given power
func (s Symbol) Power(p float64) Expression {
	return Expression{
		symbol:      s,
		power:       p,
		coefficient: 1,
	}
}

// Coefficient sets the coefficient of a Expression
func (e Expression) Coefficient(c float64) Expression {
	e.coefficient = c
	return e
}

// MultivariableExpression return a MultivariableExpression type for a given
// Expression. This is often useful for operations where the type required is a
// MultivariableExpression but you have an Expression.
func (e Expression) MultivariableExpression() MultivariableExpression {
	return MultivariableExpression{
		expressions: []Expression{e},
		coefficient: 1,
	}
}

// Simplify returns a simplified version of an Expression
func (e Expression) Simplify() Expression {
	return e
}

// NewMultivariableExpression allows the user to create a multivariable
// expression from component expressions.
//
// For example, consider the following which creates (x^2)(y^3) from x^2 and y^3:
//   tenXSquaredYCubed := algebra.NewMultivariableExpression(
//     algebra.NewSymbol("x").Power(2),
//     algebra.NewSymbol("y").Power(3),
//   )
func NewMultivariableExpression(expressions ...Expression) MultivariableExpression {
	return MultivariableExpression{
		expressions: expressions,
		coefficient: 1,
	}
}

// Coefficient sets the coefficient of a MultivariableExpression
func (e MultivariableExpression) Coefficient(c float64) MultivariableExpression {
	e.coefficient = c
	return e
}

// Simplify returns a simplified version of an Expression
func (e MultivariableExpression) Simplify() MultivariableExpression {
	simplified := MultivariableExpression{}
	coefficient := e.coefficient
	for _, expression := range e.expressions {
		coefficient = coefficient * expression.coefficient
		simplified.expressions = append(simplified.expressions, Expression{
			symbol:      expression.symbol,
			power:       expression.power,
			coefficient: 1,
		})
	}
	simplified.coefficient = coefficient
	return simplified
}
