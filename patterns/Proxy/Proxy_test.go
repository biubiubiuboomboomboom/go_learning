package Proxy

import "testing"

func TestProxy(t *testing.T)  {
	station := &Station{3}
	proxy := &StationProxy{station: station}

	proxy.sell("派大星")
	proxy.sell("小明")
	proxy.sell("小兰")
	station.sell("海绵宝宝")
}
