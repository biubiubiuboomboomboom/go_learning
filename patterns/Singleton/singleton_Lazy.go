package Singleton

import (
	"sync"
	"time"
)

type Phone struct {
	name string
	time string
}

var (
	MyPhone *Phone
	once sync.Once
)

/*
 	懒汉模式 
	事先绝不实现实例，只有第一次需要时才实例化。

 */

func GetMyPhone() *Phone{
	once.Do(func() {
		MyPhone = &Phone{name: "小米",time: time.Now().Format("20060102150405")}
	})
	return MyPhone
}






