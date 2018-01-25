package scene

import (
	"usercmd"
)

type Cell struct {
	color usercmd.ColorType //格子当前颜色
	row   uint32            //格子所在行
	col   uint32            //格子所在列
}

func NewCell(tmprow int, tmpcol int) Cell {
	tmp := Cell{
		color: usercmd.ColorType_red,
		row:   uint32(tmprow),
		col:   uint32(tmpcol),
	}
	return tmp
}

func (this *Cell) SetColor(color usercmd.ColorType) {
	this.color = color
}

func (this *Cell) GetColor() usercmd.ColorType {
	return this.color
}
