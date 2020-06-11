package ChainOfResponsibility

import "testing"

func TestChain(t *testing.T)  {
	grilHandler1 := &Gril_1{}
	grilHandler2 := &Gril_2{}
	grilHandler3 := &Gril_3{}
	// 将责任链串起来
	grilHandler1.handler = grilHandler2
	grilHandler2.handler = grilHandler3

	grilHandler1.Handle("你喜欢我吗，如果不喜欢，请往前传")
}
