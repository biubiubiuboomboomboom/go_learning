package AbstractFactory

import "fmt"

/*
	一个统一的 Phone 手机接口
 */


type Phone interface {
	PhoneType()
}

/*
	Iphone 手机
	根据工厂分为美国产与国产
 */

type ChineseIphone struct {
}

func (ChineseIphone)PhoneType() {
	fmt.Println("This is Iphone from China")
}



type AmericanIphone struct {
}

func (AmericanIphone)PhoneType() {
	fmt.Println("This is Iphone from American")
}


/*
	安卓 手机
	根据工厂 分为美国产与中国产
*/

type ChineseAndro struct {
}

func (ChineseAndro)PhoneType()  {
	fmt.Println("this is Andro from China")
}


type AmericanAndro struct {
}

func (AmericanAndro)PhoneType()  {
	fmt.Println("this is Andro from American")
}


/*
	别的手机统一称为其他
	分为美国产和中国产
*/

type ChineseOther struct {
}

func (ChineseOther)PhoneType()  {
	fmt.Println("this is other phone from China ,but I dont know ")
}



type AmericanOther struct {
}

func (AmericanOther)PhoneType()  {
	fmt.Println("this is other phone from American,but I dont know ")
}



type PhoneFactory interface {
	CreatePhone(name string) Phone
}

type ChinesePhoneFactory struct {}

func (* ChinesePhoneFactory) CreatePhone(name string) Phone {
	if name == "Iphone"{
		return &ChineseIphone{}
	} else if name == "Andro" {
		return &ChineseAndro{}
	}else {
		return ChineseOther{}
	}
}

type AmericanPhoneFactory struct {}

func (* AmericanPhoneFactory) CreatePhone(name string) Phone {
	if name == "Iphone"{
		return &AmericanIphone{}
	} else if name == "Andro" {
		return &AmericanAndro{}
	}else {
		return AmericanOther{}
	}
}

type Factory struct {
	phonefactory PhoneFactory
}

func (factory *Factory) CreatePhone(name string) Phone {
	return factory.phonefactory.CreatePhone(name)
}