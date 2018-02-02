package cell

import (
	"base/glog"
	"usercmd"
)

type Cell struct {
	color        usercmd.ColorType //格子当前颜色
	row          uint32            //格子所在行
	col          uint32            //格子所在列
	isPlayerOnMe bool              //当前格子是否有玩家存在
	isItemOnMe   bool              //当前格子是否有道具
	isVirus      bool              //是否病毒陷阱
}

func NewCell(tmprow int, tmpcol int) Cell {
	tmp := Cell{
		color:        usercmd.ColorType_origin,
		row:          uint32(tmprow),
		col:          uint32(tmpcol),
		isPlayerOnMe: false,
	}
	return tmp
}

func (this *Cell) SetColor(color usercmd.ColorType) {
	this.color = color
}

func (this *Cell) GetColor() usercmd.ColorType {
	return this.color
}

func (this *Cell) PlayerOnMe() {
	this.isPlayerOnMe = true
}

func (this *Cell) PlayerLeaveMe() {
	this.isPlayerOnMe = false
}

func (this *Cell) GetPlayerOnMe() bool {
	return this.isPlayerOnMe
}

func (this *Cell) ItemOnMe() {
	this.isItemOnMe = true
}

func (this *Cell) ItemLeaveMe() {
	this.isItemOnMe = false
}

func (this *Cell) GetItemOnMe() bool {
	return this.isItemOnMe
}

func (this *Cell) SetVirus() {
	this.isVirus = true
}

func (this *Cell) RemoveVirus() {
	this.isVirus = false
}

func (this *Cell) GetVirus() bool {
	return this.isVirus
}
