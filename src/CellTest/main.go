package main

import (
	"base/glog"
	"math"
	"myserver/consts"
)

func main() {
	tmprow, tmpcol := whichCell(19.4, 11, 19.4)
	glog.Error("row = ", tmprow, " col = ", tmpcol)
}

func whichCell(px float32, py float32, pz float32) (uint32, uint32) {
	//TODO -1.0  -1.0有bug  因为是-0
	col := math.Ceil(float64(px/consts.CellLength) - 0.5)
	row := math.Ceil(float64(pz/consts.CellLength) - 0.5)
	return uint32(row), uint32(col)
}
