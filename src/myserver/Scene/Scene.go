package scene

import (
	"base/glog"
	"common"
	"math"
	"myserver/cell"
	"myserver/consts"
	_ "myserver/interfaces"
	"myserver/itemmgr"
	_ "myserver/playertaskmgr"
	"time"
	"usercmd"
)

//TODO 抽出队伍抽象struct
type Scene struct {
	itemmgr.ItemMgr //道具管理

	Players         map[uint32]*ScenePlayer // 玩家对象
	Cells           [][]cell.Cell           //格子
	PlayerIdsBlue   []uint32                //蓝队
	PlayerIdsYellow []uint32                //黄队
	PlayerIdsRed    []uint32                //红队
	MaxCellColor    usercmd.ColorType       //当前游戏领先队伍拥有的格子颜色
	CellColorNum    map[usercmd.ColorType]uint32
	EnergyRed       uint32 //红队能量
	EnergyBlue      uint32 //蓝队能量
	EnergyYellow    uint32 //黄队能量
	isInGame        bool   //room通知是否还在游戏中
}

func (this *Scene) SceneInit() {
	glog.Error("[init] SceneInit()...")
	this.ItemMgr.Scene = this
	this.Players = make(map[uint32]*ScenePlayer)
	//初始化格子颜色 origin
	count := int(consts.CellNum)
	for i := 0; i < count; i++ {
		var tmpArr []cell.Cell
		for j := 0; j < count; j++ {
			tmpArr = append(tmpArr, cell.NewCell(i, j))
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
	glog.Error("[init] SceneInit() success")
}

func (this *Scene) InitPlayerPosition() {
	if len(this.Players) != consts.OneGamePlayerNum {
		glog.Error("debug 人数超出设定人数!")
		return
	}
	//TODO 分配位置 分配队伍
	for id, p := range this.Players {
		p.PlayerId = id
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

func (this *Scene) GetOneCellByRowCol(tmprow uint32, tmpcol uint32) cell.Cell {
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
		//glog.Error("红队能量条+1")
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
		//glog.Error("蓝队能量条+1")
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
		//glog.Error("黄队能量条+1")
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

//从 a 格子 移动到 b 格子  对格子isPlayerOnMe参数修改
func (this *Scene) MoveFromToCell(arow uint32, acol uint32, brow uint32, bcol uint32, pid uint32, itemNum uint32) {
	this.Cells[int(arow)][int(acol)].PlayerLeaveMe()
	this.Cells[int(brow)][int(bcol)].PlayerOnMe()
	//判断格子上是否有道具
	//glog.Error("玩家从", arow, " ", acol, "运动到", brow, " ", bcol)
	if itemNum >= consts.MaxItemNum {
		//玩家道具数量达到上限
		//glog.Error("玩家道具包已满")
		return
	}
	if this.Cells[int(brow)][int(bcol)].GetItemOnMe() {
		//格子有道具属性变更
		//道具管理那边也要去除道具
		//glog.Error("道具被玩家捡了 row = ", brow, " col = ", bcol)
		this.Cells[int(brow)][int(bcol)].ItemLeaveMe()
		this.ItemMgr.DeleteOneItem(brow, bcol)
		this.GetItemToPlayer(pid, this.ItemMgr.GetItemByRowCol(brow, bcol))
	}
}

func (this *Scene) GetItemToPlayer(pid uint32, itype usercmd.ItemType) {
	//玩家获得道具type
	playerTmp, ok := this.Players[pid]
	if !ok {
		glog.Error("[bug] wrong playerid")
	}
	playerTmp.GetItem(itype)
}

func (this *Scene) SetItemOnCell(row uint32, col uint32) {
	this.Cells[int(row)][int(col)].ItemOnMe()
}

func (this *Scene) DeleteItemOnCell(row uint32, col uint32) {
	this.Cells[int(row)][int(col)].ItemLeaveMe()
}

func (this *Scene) IsPlayerOnCell(row uint32, col uint32) bool {
	return this.Cells[int(row)][int(col)].GetPlayerOnMe()
}

func (this *Scene) BroadCastMsg(data []byte, flag byte) {
	//TODO 默认取红队来发送消息 整个逻辑架构有问题
	tmpPlayer := this.Players[this.PlayerIdsRed[0]]
	tmpPlayer.room.BroadCastMsg(data, flag)
}

func whichCell(px float32, py float32, pz float32) (uint32, uint32, bool) {
	//TODO -1.0  -1.0有bug  因为是-0
	col := math.Ceil(float64(px/consts.CellLength) - 0.5)
	row := math.Ceil(float64(pz/consts.CellLength) - 0.5)
	if int(row) < 0 || uint32(row) > consts.CellNum-1 || int(col) < 0 || uint32(col) > consts.CellNum-1 {
		//glog.Error("[bug] error row or col ", row, " ", col)
		return 0, 0, false
	}
	return uint32(row), uint32(col), true
}

func (this *Scene) IsInGame() bool {
	//TODO 默认取红队来发送消息 整个逻辑架构有问题
	tmpPlayer := this.Players[this.PlayerIdsRed[0]]
	return tmpPlayer.room.IsInGame()
}

func (this *Scene) SetCellVirus(row uint32, col uint32, pId uint32) {
	this.Cells[int(row)][int(col)].SetVirus()
	//广播格子病毒陷阱
	m := usercmd.VirusCreateS2CMsg{
		PlayerId: pId,
		Row:      row,
		Col:      col,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_VirusCreate), &m)
	//TODO 默认取红队来发送消息 整个逻辑架构有问题
	tmpPlayer := this.Players[this.PlayerIdsRed[0]]
	tmpPlayer.room.BroadCastMsg(d, f)
}

func (this *Scene) IsCellVirus(row uint32, col uint32) bool {
	//判断格子如果有病毒陷阱，就返回true并且移除病毒陷阱
	tmp := this.Cells[int(row)][int(col)]
	if tmp.GetVirus() {
		glog.Error("移除病毒陷阱")
		tmp.RemoveVirus()
		m := usercmd.VirusDestroyS2CMsg{
			Row: row,
			Col: col,
		}
		d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_VirusDestroy), &m)
		//TODO 默认取红队来发送消息 整个逻辑架构有问题
		tmpPlayer := this.Players[this.PlayerIdsRed[0]]
		tmpPlayer.room.BroadCastMsg(d, f)
		return true
	}
	return false
}

func (this *Scene) AbsFun(a uint32, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func (this *Scene) DyeingFun(row uint32, col uint32, color usercmd.ColorType, pId uint32) {
	var num uint32 = 0
	var arrays []*ScenePlayer
	//判断以row col为中心 行列差之和在n范围内 染上特别的颜色
	for _, playerTmp := range this.Players {
		if playerTmp.PlayerId == pId {
			//自己跳过
			continue
		}
		rowDiffer := this.AbsFun(playerTmp.nowrow, row)
		colDiffer := this.AbsFun(playerTmp.nowcol, col)
		glog.Error("rowDiffer = ", rowDiffer)
		glog.Error("colDiffer = ", colDiffer)
		if rowDiffer+colDiffer < consts.DyeingRange || rowDiffer+colDiffer == consts.DyeingRange {
			playerTmp.TurnDyeing(color, pId)
			num++
			arrays = append(arrays, playerTmp)
		}
	}
	//glog.Error("染色道具发动 一共", num, "个玩家受到影响")
	timer := time.NewTimer(time.Second * consts.DyeingTime)
	go func() {
		for true {
			<-timer.C
			//染色状态结束
			//广播客户端  并且 将玩家 player属性修改
			for _, pTmp := range arrays {
				pTmp.TurnNoDyeing()
			}
			return
		}
	}()
}
