package main

import (
	"base/glog"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 5)

	go func() {
		<-timer.C
		glog.Error("debug")
	}()

	for true {

	}
}
