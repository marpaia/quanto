package algebra

// Arithmetic is an interface which defines arithmetic values
type Arithmetic interface {
	MultivariableExpression() MultivariableExpression
}

// Mul multiplies two arithmetic values
func Mul(lhs, rhs Arithmetic) (MultivariableExpression, error) {
	l := lhs.MultivariableExpression().Simplify()
	r := rhs.MultivariableExpression().Simplify()

	expressionsMap := map[Symbol]float64{}
	for _, e := range l.expressions {
		expressionsMap[e.symbol] = e.power
	}
	for _, e := range r.expressions {
		power, ok := expressionsMap[e.symbol]
		if ok {
			expressionsMap[e.symbol] = power + e.power
		} else {
			expressionsMap[e.symbol] = e.power
		}
	}

	expressions := []Expression{}
	for symbol, power := range expressionsMap {
		expressions = append(expressions, symbol.Power(power))
	}

	product := MultivariableExpression{
		expressions: expressions,
		coefficient: l.coefficient * r.coefficient,
	}

	return product, nil
}
