package Strategy

import (
	"testing"
)

func TestOperator_Calculate(t *testing.T) {
	operation := new(Operator)
	operation.operation = new(add)
	println(operation.Calculate(1,2))
}
