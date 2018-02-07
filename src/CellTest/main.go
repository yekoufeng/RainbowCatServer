package main

import (
	"base/glog"
	"math"
	"myserver/consts"
)

func main() {
	aTmp, bTmp := getMinAndSecByLoop(179)
	glog.Error("minute = ", aTmp)
	glog.Error("second = ", bTmp)
}

func getMinAndSecByLoop(loop uint32) (uint32, uint32) {
	leftTime := consts.OneGameTime - loop
	glog.Error("leftTime = ", leftTime)
	aTmp := leftTime / 60
	bTmp := leftTime % 60
	return aTmp, bTmp
}

func whichCell(px float32, py float32, pz float32) (uint32, uint32) {
	col := math.Ceil(float64(px/consts.CellLength) - 0.5)
	row := math.Ceil(float64(pz/consts.CellLength) - 0.5)
	return uint32(row), uint32(col)
}
