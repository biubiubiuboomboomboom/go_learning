package Adaptor

import "fmt"

type EnglishGirl interface {
	SayEngilish(Type string ,name string)
}

type Me struct {
}

func (*Me) SayChinese(name string){
	fmt.Println("你好！",name)
}

type MiddleMan struct {
	me Me
}

func (middle *MiddleMan)SayEngilish(Type string ,name string)  {
	if Type == "english"{
		middle.me.SayChinese(name)
	}else {
		println("这个语种我也不熟")
	}
}