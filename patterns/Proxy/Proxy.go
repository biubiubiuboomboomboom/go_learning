package Proxy

import "fmt"

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station *Station)sell(name string)  {
	if station.stock>0{
		station.stock-=1
		fmt.Printf("%s去总部买了一张票,剩余：%d \n ",name,station.stock)
	}else {
		fmt.Printf("票已售空，%s没买到票 \n",name)
	}
}

type StationProxy struct {
	station *Station
}

func (proxy *StationProxy)sell(name string)  {
	if proxy.station.stock>0{
		proxy.station.stock--
		fmt.Printf("代理点中：%s买了一张票,剩余：%d \n", name, proxy.station.stock)
	}else {
		fmt.Printf("票已售空，%s没买到票 \n",name)
	}
}


