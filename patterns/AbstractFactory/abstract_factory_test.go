package AbstractFactory

import "testing"

func TestAbstract_factory(t *testing.T )  {
	factory := new(Factory)

	//美国
	factory.phonefactory = new(AmericanPhoneFactory)
	phone := factory.CreatePhone("Iphone")
	phone.PhoneType()

}
