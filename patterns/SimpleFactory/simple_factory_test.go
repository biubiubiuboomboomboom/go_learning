package SimpleFactory

import "testing"

func TestSimple(t *testing.T)  {

	phoneFactory:= new(PhoneFactory)

	phone := phoneFactory.CreatePhone("Iphone")
	phone.PhoneType()

	phone = phoneFactory.CreatePhone("Andro")
	phone.PhoneType()

	phone = phoneFactory.CreatePhone("sasasas")
	phone.PhoneType()
}
