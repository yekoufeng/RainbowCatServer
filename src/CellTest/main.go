package main

import (
	"base/glog"
	"math"
	"myserver/consts"
	"time"
)

func main() {
	//	tmprow, tmpcol := whichCell(19.4, 11, 19.4)
	//	glog.Error("row = ", tmprow, " col = ", tmpcol)
	var timer = time.NewTicker(time.Millisecond * 500)

	go func() {
		for true {
			<-timer.C
			glog.Error("0.5s")
		}

	}()
	for true {
	}
}

func whichCell(px float32, py float32, pz float32) (uint32, uint32) {
	col := math.Ceil(float64(px/consts.CellLength) - 0.5)
	row := math.Ceil(float64(pz/consts.CellLength) - 0.5)
	return uint32(row), uint32(col)
}
