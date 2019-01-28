package calculus

import (
	"errors"
)

// Differentiable is an interface that defines whether or not the derivative can
// be taken of a given type. A differentiable function of one real variable is a
// function whose derivative exists at each point in its domain.
type Differentiable interface {
}

func Derive(f Differentiable) error {
	return errors.New("Derive is unimplemented")
}
