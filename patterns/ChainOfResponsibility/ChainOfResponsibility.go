package ChainOfResponsibility

import (
	"fmt"
)

type Handler interface {
	Handle(content string)
	next(handler Handler, content string)
}

// 第一个女生
type Gril_1 struct {
	handler Handler
}

func (gril *Gril_1) Handle(content string) {
	println("女生1号不喜欢你")
	gril.next(gril.handler, content)
}

func (gril *Gril_1) next(handler Handler, content string) {
	if gril.handler != nil {
		gril.handler.Handle(content)
	}else {
		println("她把纸条给丢了")
	}
}

// 第二个女生
type Gril_2 struct {
	handler Handler
}

func (gril2 *Gril_2) Handle(content string) {
	fmt.Println("女生2号不喜欢你")
	gril2.next(gril2.handler, content)
}

func (gril2 *Gril_2) next(handler Handler, content string) {
	if gril2.handler != nil {
		gril2.handler.Handle(content)
	}else {
		println("她把纸条给丢了")
	}
}

// 第三个女生

type Gril_3 struct {
	handler Handler
}

func (gril3 *Gril_3) Handle(content string) {
	fmt.Println("女生3号不喜欢你")
	gril3.next(gril3.handler, content)
}

func (gril3 *Gril_3) next(handler Handler, content string) {
	if gril3.handler != nil {
		gril3.handler.Handle(content)
	}else {
		println("她把纸条给丢了")
	}
}