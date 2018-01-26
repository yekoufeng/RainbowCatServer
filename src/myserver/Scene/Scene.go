package scene

import (
	"base/glog"
	"math"
	"myserver/consts"
	_ "myserver/interfaces"
	_ "myserver/playertaskmgr"
	"usercmd"
)

type Scene struct {
	Players map[uint32]*ScenePlayer // 玩家对象
	Cells   [][]Cell                //格子

	PlayerIdsBlue   []uint32          //蓝队
	PlayerIdsYellow []uint32          //黄队
	PlayerIdsRed    []uint32          //红队
	MaxCellColor    usercmd.ColorType //当前游戏领先队伍拥有的格子颜色
	CellColorNum    map[usercmd.ColorType]uint32
}

func (this *Scene) SceneInit() {
	glog.Error("SceneInit()...")
	this.Players = make(map[uint32]*ScenePlayer)
	//初始化格子颜色 origin
	count := int(consts.CellNum)
	for i := 0; i < count; i++ {
		var tmpArr []Cell
		for j := 0; j < count; j++ {
			tmpArr = append(tmpArr, NewCell(i, j))
		}
		this.Cells = append(this.Cells, tmpArr)
	}
	this.CellColorNum = make(map[usercmd.ColorType]uint32)
	this.CellColorNum[usercmd.ColorType_blue] = 0
	this.CellColorNum[usercmd.ColorType_red] = 0
	this.CellColorNum[usercmd.ColorType_yellow] = 0
	this.CellColorNum[usercmd.ColorType_origin] = consts.CellNum * consts.CellNum
	//默认最大颜色是原始
	this.MaxCellColor = usercmd.ColorType_origin
	glog.Error("SceneInit() success")
}

func (this *Scene) InitPlayerPosition() {
	if len(this.Players) != consts.OneGamePlayerNum {
		glog.Error("debug 人数超出两人!")
	}
	//TODO 分配位置 分配队伍
	for id, p := range this.Players {
		if len(this.PlayerIdsBlue) == 0 {
			p.SetPosition(19.5, 10, 19.5)
			p.SetRowCol(19, 19)
			p.Color = usercmd.ColorType_blue
			this.PlayerIdsBlue = append(this.PlayerIdsBlue, id)
		} else {
			p.SetPosition(0, 10, 0)
			p.SetRowCol(0, 0)
			p.Color = usercmd.ColorType_yellow
			this.PlayerIdsYellow = append(this.PlayerIdsYellow, id)
		}
	}

}

func (this *Scene) GetOnePlayerPosition(id uint32) (float32, float32, float32) {
	return this.Players[id].GetPosition()
}

func (this *Scene) GetOneCellByRowCol(tmprow uint32, tmpcol uint32) Cell {
	return this.Cells[int(tmprow)][int(tmpcol)]
}

func (this *Scene) SetCellColor(row uint32, col uint32, color usercmd.ColorType) {
	//玩家所属队伍颜色格子加一
	//被占领格子如果是有人占领的，那么那个人队伍颜色数目减一
	tmpLastColor := this.GetCellColor(row, col)
	if color != tmpLastColor {
		this.AddColorNum(color)
	}
	this.DeleteColorNum(tmpLastColor)

	//更新max
	//如果最大值颜色是原始，则覆盖最大值颜色
	if this.MaxCellColor == usercmd.ColorType_origin {
		this.MaxCellColor = color
	}
	//如果最大值颜色是本身，则不变
	//如果最大值颜色不是本身，则比较
	//如果+1的颜色队伍超过了最大值颜色的队伍，那么最大值颜色队伍就是+1的队伍
	if this.CellColorNum[color] > this.CellColorNum[this.MaxCellColor] {
		this.MaxCellColor = color
	}
	this.Cells[int(row)][int(col)].SetColor(color)
}

func (this *Scene) GetCellColor(row uint32, col uint32) usercmd.ColorType {
	return this.Cells[int(row)][int(col)].GetColor()
}

//TODO 加锁 或者 chan
func (this *Scene) AddColorNum(color usercmd.ColorType) {
	if color == usercmd.ColorType_blue {
		this.CellColorNum[usercmd.ColorType_blue]++
	} else if color == usercmd.ColorType_yellow {
		this.CellColorNum[usercmd.ColorType_yellow]++
	} else if color == usercmd.ColorType_red {
		this.CellColorNum[usercmd.ColorType_red]++
	}
}

//TODO 加锁 或者 chan
func (this *Scene) DeleteColorNum(color usercmd.ColorType) {
	if color == usercmd.ColorType_blue {
		this.CellColorNum[usercmd.ColorType_blue]--
	} else if color == usercmd.ColorType_yellow {
		this.CellColorNum[usercmd.ColorType_yellow]--
	} else if color == usercmd.ColorType_red {
		this.CellColorNum[usercmd.ColorType_red]--
	} else if color == usercmd.ColorType_origin {
		this.CellColorNum[usercmd.ColorType_origin]--
	}
}

func whichCell(px float32, py float32, pz float32) (uint32, uint32, bool) {
	//TODO -1.0  -1.0有bug  因为是-0
	col := math.Ceil(float64(px/consts.CellLength) - 0.5)
	row := math.Ceil(float64(pz/consts.CellLength) - 0.5)
	if int(row) < 0 || uint32(row) > consts.CellNum-1 || int(col) < 0 || uint32(col) > consts.CellNum-1 {
		glog.Error("[bug] error row or col ", row, " ", col)
		return 0, 0, false
	}
	return uint32(row), uint32(col), true
}
