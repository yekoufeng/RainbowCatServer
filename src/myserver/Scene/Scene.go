package scene

import (
	"base/glog"
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
	CellNumBlue     uint32            //蓝队目前的格子数
	CellNumYellow   uint32            //黄队目前的格子数
	CellNumRed      uint32            //红队目前的格子数
	MaxCellNum      uint32            //当前游戏领先队伍拥有的格子数目
	MaxCellColor    usercmd.ColorType //当前游戏领先队伍拥有的格子颜色
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
	if tmpLastColor != usercmd.ColorType_origin {
		this.DeleteColorNum(tmpLastColor)
	}
	this.Cells[int(row)][int(col)].SetColor(color)
}

func (this *Scene) GetCellColor(row uint32, col uint32) usercmd.ColorType {
	//TODO 貌似有bug
	if int(row) < 0 || row > consts.CellNum || int(col) < 0 || col > consts.CellNum {
		glog.Error("[bug] error row or col ", row, " ", col)
	}
	return this.Cells[int(row)][int(col)].GetColor()
}

//TODO 加锁 或者 chan
func (this *Scene) AddColorNum(color usercmd.ColorType) {
	if color == usercmd.ColorType_blue {
		this.CellNumBlue++
	} else if color == usercmd.ColorType_yellow {
		this.CellNumYellow++
	} else if color == usercmd.ColorType_red {
		this.CellNumRed++
	}
}

//TODO 加锁 或者 chan
func (this *Scene) DeleteColorNum(color usercmd.ColorType) {
	if color == usercmd.ColorType_blue {
		this.CellNumBlue--
	} else if color == usercmd.ColorType_yellow {
		this.CellNumYellow--
	} else if color == usercmd.ColorType_red {
		this.CellNumRed--
	}
}
