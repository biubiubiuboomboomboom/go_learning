package SimpleFactory

import "fmt"

/*
	定义一个手机的工厂
 */
type Phone interface {
	PhoneType()
}

/*
	第一种是 Iphone
 */

type Iphone struct {
}

func (*Iphone)PhoneType()  {
	fmt.Println("this is iphone")
}

/*
	第二种是安卓
 */

type Andro struct {
}

func (*Andro)PhoneType()  {
	fmt.Println("this is Andro")
}


/*
	别的手机统一称为其他
 */

type Other struct {
}

func (*Other)PhoneType()  {
	fmt.Println("this is other phone ,but I dont know ")
}




type PhoneFactory struct {
}


/*
	根据用户的的需要，通过工厂模式，实现了各个型号的手机
 */

func (*PhoneFactory) CreatePhone(name string) Phone {
	if name == "Iphone"{
		return &Iphone{}
	}else if name == "Andro"{
		return &Andro{}
	}else {
		return &Other{}
	}
}

