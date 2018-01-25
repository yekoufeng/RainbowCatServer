package scene

import (
	"base/glog"
	"myserver/consts"
	_ "myserver/interfaces"
	_ "myserver/playertaskmgr"
	"usercmd"
)

type Scene struct {
	Players        map[uint32]*ScenePlayer // 玩家对象
	Cells          [][]Cell                //格子
	PlayerIdsBlue  []uint32                //蓝队
	PlayerIdsGreen []uint32                //绿队
	CellNumBlue    uint32                  //蓝队目前的格子数
	CellNumGreen   uint32                  //绿队目前的格子数
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
			p.SetPosition(123.5, 10, 123.5)
			p.SetRowCol(19, 19)
			p.Color = usercmd.ColorType_blue
			this.PlayerIdsBlue = append(this.PlayerIdsBlue, id)
		} else {
			p.SetPosition(0, 10, 0)
			p.SetRowCol(0, 0)
			p.Color = usercmd.ColorType_green
			this.PlayerIdsGreen = append(this.PlayerIdsGreen, id)
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
	return this.Cells[int(row)][int(col)].GetColor()
}

func (this *Scene) AddColorNum(color usercmd.ColorType) {
	if color == usercmd.ColorType_blue {
		this.CellNumBlue++
	} else if color == usercmd.ColorType_green {
		this.CellNumGreen++
	}
}

func (this *Scene) DeleteColorNum(color usercmd.ColorType) {
	if color == usercmd.ColorType_blue {
		this.CellNumBlue--
	} else if color == usercmd.ColorType_green {
		this.CellNumGreen--
	}
}
