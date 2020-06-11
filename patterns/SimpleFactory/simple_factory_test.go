package SimpleFactory

import "testing"


func TestSimple(t *testing.T)  {

	var _phoneFactory = PhoneFactory{}

	phone := _phoneFactory.CreatePhone("Iphone")
	phone.PhoneType()

	phone = _phoneFactory.CreatePhone("Andro")
	phone.PhoneType()

	phone = _phoneFactory.CreatePhone("sasasas")
	phone.PhoneType()
}
