package Singleton

import (
	"testing"
)

func TestHungry(t *testing.T) {



	mygun := GetGun()
	//fmt.Println("myphone 创建时间："+mygun.time)
	//fmt.Printf("myphone 创建地址：%p \n",&mygun)
	////time.Sleep(10000)
	mygun2 := GetGun()

	println(mygun == mygun2)

	//fmt.Println("myphone2 创建时间: "+mygun2.time)
	//fmt.Printf("myphone2 创建地址: %p",&mygun2)
}
