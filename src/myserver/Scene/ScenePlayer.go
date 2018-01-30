package scene

import (
	"base/glog"
	"common"
	"myserver/interfaces"
	"myserver/playertaskmgr"
	"usercmd"
)

type ScenePlayer struct {
	ScenePlayerNetMsgHelper // 房间玩家协议处理辅助类
	Sess                    interfaces.IPlayerTask
	room                    interfaces.IRoom
	PlayerId                uint32 //玩家id
	posX                    float32
	posY                    float32
	posZ                    float32
	nowrow                  uint32             //当前所在格子行
	nowcol                  uint32             //当前所在格子列
	Color                   usercmd.ColorType  //玩家颜色
	nowcellnum              uint32             //当前所拥有的格数
	items                   []usercmd.ItemType //玩家当前拥有的道具
}

func NewScenePlayer(id uint32, rm interfaces.IRoom) *ScenePlayer {
	tmp := &ScenePlayer{
		Sess:       playertaskmgr.GetMe().GetPlayerTask(id),
		room:       rm,
		posX:       0,
		posY:       0,
		posZ:       0,
		nowrow:     20,
		nowcol:     20,
		Color:      usercmd.ColorType_origin,
		nowcellnum: 0,
		PlayerId:   0,
	}
	tmp.Init(tmp)
	return tmp
}

func (this *ScenePlayer) SendMsg(data []byte, flag byte) {
	this.Sess.AsyncSend(data, flag)
}

func (this *ScenePlayer) SetPosition(px float32, py float32, pz float32) {
	this.posX = px
	this.posY = py
	this.posZ = pz
}

func (this *ScenePlayer) SetRowCol(row uint32, col uint32) {
	this.nowrow = row
	this.nowcol = col
}

func (this *ScenePlayer) GetRow() uint32 {
	return this.nowrow
}

func (this *ScenePlayer) GetCol() uint32 {
	return this.nowcol
}

func (this *ScenePlayer) GetPosition() (float32, float32, float32) {
	return this.posX, this.posY, this.posZ
}

func (this *ScenePlayer) HandleMove(px float32, py float32, pz float32) {
	this.SetPosition(px, py, pz)
	this.handleMoveColor()
	m := usercmd.MoveS2CMsg{
		PlayerId: this.Sess.GetId(),
		PosX:     this.posX,
		PosY:     this.posY,
		PosZ:     this.posZ,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_MoveRes), &m)
	this.room.BroadCastMsg(d, f)

}

func (this *ScenePlayer) handleMoveColor() {
	//获得当前坐标所在格子，判断是否进入下一个格子
	tmprow, tmpcol, ok := whichCell(this.posX, this.posY, this.posZ)
	if !ok {
		glog.Error("坐标有问题bug")
		return
	}
	if tmprow != this.nowrow || tmpcol != this.nowcol {
		//进入新的格子
		this.room.MoveFromToCell(this.nowrow, this.nowcol, tmprow, tmpcol, this.PlayerId)
		this.nowrow = tmprow
		this.nowcol = tmpcol
		tmpLastColor := this.room.GetCellColor(tmprow, tmpcol)
		if tmpLastColor == this.Color {
			//同队伍颜色，无需再发
			return
		}
		//玩家当前自己占领格子加一
		this.nowcellnum++
		this.room.SetCellColor(tmprow, tmpcol, this.Color)
		glog.Error("变色", this.Color, " 该位置所属格子为 row = ", tmprow, " col = ", tmpcol)
		m := usercmd.ChangeColorS2CMsg{
			Color: this.Color,
			Row:   tmprow,
			Col:   tmpcol,
		}
		d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_ChangeColorRes), &m)
		this.room.BroadCastMsg(d, f)

	}
}

//检查道具是否存在
func (this *ScenePlayer) checkItemOk(itype usercmd.ItemType) bool {
	for _, typeTmp := range this.items {
		if typeTmp == itype {
			return true
		}
	}
	glog.Error("[bug]玩家", this.PlayerId, " 没有该道具", itype)
	return false
}

func (this *ScenePlayer) handleUseItem(itype usercmd.ItemType) {
	//玩家使用道具
	//安全检查 玩家是否有该道具
	if !this.checkItemOk(itype) {
		return
	}
	switch itype {
	case usercmd.ItemType_virus:
		//病毒道具
		this.handleItemVirus()
	case usercmd.ItemType_dyeing:
		//染色道具
		this.handleItemDyeing()
	case usercmd.ItemType_unknown:
		//未知道具
		glog.Error("[bug]未知道具")
	default:
		glog.Error("[bug] 不明itemtype")
	}
}

func (this *ScenePlayer) handleItemVirus() {
	//发动病毒道具
}

func (this *ScenePlayer) handleItemDyeing() {
	//发动染色道具
}

func (this *ScenePlayer) GetItem(itype usercmd.ItemType) {
	//玩家获得道具
	this.items = append(this.items, itype)

	m := usercmd.GetItemS2CMsg{
		Item: itype,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_PlayerGetItem), &m)
	this.Sess.AsyncSend(d, f)
}

func (this *ScenePlayer) WinGame() {
	this.room.HandleGameOver(this.Color)
}
