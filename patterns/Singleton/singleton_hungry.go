package Singleton

import "time"

type Gun struct {
	name string
	time string
}

var Mygun *Gun

func init()  {
	Mygun = new(Gun)
	Mygun.name = "AK47"
	Mygun.time = time.Now().Format("20060102150405")
}

func GetGun() *Gun {
	return Mygun
}
