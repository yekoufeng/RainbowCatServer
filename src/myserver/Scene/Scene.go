package scene

import (
	"base/glog"
	"common"
	"math"
	"myserver/consts"
	_ "myserver/interfaces"
	_ "myserver/playertaskmgr"
	"usercmd"
)

//TODO 抽出队伍抽象struct
type Scene struct {
	Players map[uint32]*ScenePlayer // 玩家对象
	Cells   [][]Cell                //格子

	PlayerIdsBlue   []uint32          //蓝队
	PlayerIdsYellow []uint32          //黄队
	PlayerIdsRed    []uint32          //红队
	MaxCellColor    usercmd.ColorType //当前游戏领先队伍拥有的格子颜色
	CellColorNum    map[usercmd.ColorType]uint32
	EnergyRed       uint32 //红队能量
	EnergyBlue      uint32 //蓝队能量
	EnergyYellow    uint32 //黄队能量

	isInGame bool //room通知是否还在游戏中
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
		glog.Error("debug 人数超出设定人数!")
		return
	}
	//TODO 分配位置 分配队伍
	for id, p := range this.Players {
		if len(this.PlayerIdsBlue) == 0 {
			p.SetPosition(19.5, 10, 19.5)
			p.SetRowCol(19, 19)
			p.Color = usercmd.ColorType_blue
			this.PlayerIdsBlue = append(this.PlayerIdsBlue, id)
		} else if len(this.PlayerIdsRed) == 0 {
			p.SetPosition(0, 10, 0)
			p.SetRowCol(0, 0)
			p.Color = usercmd.ColorType_red
			this.PlayerIdsRed = append(this.PlayerIdsRed, id)
		} else {
			p.SetPosition(0, 10, 0)
			p.SetRowCol(10, 10)
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

func (this *Scene) AddEnergyInScene() {
	if this.MaxCellColor == usercmd.ColorType_origin {
		return
	}
	var tmpStatue uint32 = 0
	var tmpPlayer *ScenePlayer
	//根据当前领先队伍颜色增加能量条
	switch this.MaxCellColor {
	case usercmd.ColorType_red:
		this.EnergyRed++
		tmpStatue = this.EnergyRed
		glog.Error("红队能量条+1")
		tmpPlayer = this.Players[this.PlayerIdsRed[0]]
		if this.EnergyRed == consts.TotalEnergyNum {
			if len(this.PlayerIdsRed) < 1 {
				glog.Error("[bug] 红色获胜队伍没有成员")
			}
			this.Players[this.PlayerIdsRed[0]].WinGame()
		}
	case usercmd.ColorType_blue:
		this.EnergyBlue++
		tmpStatue = this.EnergyBlue
		glog.Error("蓝队能量条+1")
		tmpPlayer = this.Players[this.PlayerIdsBlue[0]]
		if this.EnergyBlue == consts.TotalEnergyNum {
			if len(this.PlayerIdsBlue) < 1 {
				glog.Error("[bug] 蓝色获胜队伍没有成员")
			}
			this.Players[this.PlayerIdsBlue[0]].WinGame()
		}
	case usercmd.ColorType_yellow:
		this.EnergyYellow++
		tmpStatue = this.EnergyYellow
		glog.Error("黄队能量条+1")
		tmpPlayer = this.Players[this.PlayerIdsYellow[0]]
		if this.EnergyYellow == consts.TotalEnergyNum {
			if len(this.PlayerIdsYellow) < 1 {
				glog.Error("[bug] 黄色获胜队伍没有成员")
			}
			this.Players[this.PlayerIdsYellow[0]].WinGame()
		}
	}

	m := usercmd.GameEnergyS2CMsg{
		Color:  this.MaxCellColor,
		Status: tmpStatue,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_GameEnergy), &m)
	tmpPlayer.room.BroadCastMsg(d, f)
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
