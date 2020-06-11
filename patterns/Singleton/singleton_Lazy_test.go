package Singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestLazy(t *testing.T) {
	myphone := GetMyPhone()
	fmt.Println("myphone 创建时间："+myphone.time)
	fmt.Printf("myphone 创建地址：%p \n",myphone)
	fmt.Println(myphone)
	time.Sleep(10000)
	myphone2 := GetMyPhone()
	fmt.Println("myphone2 创建时间: "+myphone2.time)
	fmt.Printf("myphone2 创建地址: %p",myphone2)
	fmt.Println(myphone2)
}
