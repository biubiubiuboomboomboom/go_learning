package Strategy

type Operation interface {
	do(int,int)int
}

type add struct {
}

func (*add) do(digital1 int, digital2 int) int{
	return digital1+digital2
}

type reduce struct {
}

func (*reduce) do(digital1 int, digital2 int) int{
	return digital1-digital2
}

/*
	策略的执行者
 */

type Operator struct {
	operation Operation
}

func (strategy Operator) setOperation(opration Operation) {
	strategy.operation = opration
}

func (strategy Operator) Calculate(digital1 int, digital2 int) int {
	return strategy.operation.do(digital1,digital2)
}