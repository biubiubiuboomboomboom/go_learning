package AbstractFactory

import "testing"



func TestAbstract_factory(t *testing.T )  {
	var _factory = Factory{}

	//美国
	_factory.phonefactory = new(AmericanPhoneFactory)
	phone := _factory.CreatePhone("Iphone")
	phone.PhoneType()

}
