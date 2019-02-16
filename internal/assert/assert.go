package assert

import (
	"fmt"
	"reflect"
)

// Equal asserts that two objects are equal.
func Equal(expected, actual interface{}) error {
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("\nexpected: %+v\nactual: %+v", expected, actual)
	}
	return nil
}
