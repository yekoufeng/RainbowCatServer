package scene

import (
	"base/glog"
	"common"
	"math"
	"myserver/consts"
	"myserver/interfaces"
	"myserver/playertaskmgr"
	"usercmd"
)

type ScenePlayer struct {
	ScenePlayerNetMsgHelper // 房间玩家协议处理辅助类
	Sess                    interfaces.IPlayerTask
	room                    interfaces.IRoom
	posX                    float32
	posY                    float32
	posZ                    float32
	nowrow                  uint32
	nowcol                  uint32
	Color                   usercmd.ColorType
	nowcellnum              uint32 //当前所拥有的格数
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
	tmprow, tmpcol := whichCell(this.posX, this.posY, this.posZ)
	if tmprow != this.nowrow || tmpcol != this.nowcol {
		//进入新的格子
		tmpLastColor := this.room.GetCellColor(tmprow, tmpcol)
		if tmpLastColor == this.Color {
			//同队伍颜色，无需再发
			//TODO
			return
		}
		//玩家当前自己占领格子加一
		this.nowcellnum++

		this.room.SetCellColor(tmprow, tmpcol, this.Color)
		glog.Error("变色 该位置所属格子为 row = ", tmprow, " col = ", tmpcol)
		m := usercmd.ChangeColorS2CMsg{
			Color: this.Color,
			Row:   tmprow,
			Col:   tmpcol,
		}
		d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_ChangeColorRes), &m)
		this.room.BroadCastMsg(d, f)
		this.nowrow = tmprow
		this.nowcol = tmpcol
	}
}

func whichCell(px float32, py float32, pz float32) (uint32, uint32) {
	//TODO -3.5 -3.5有bug  因为是-0
	col := math.Ceil(float64(px/consts.CellLength) - 0.5)
	row := math.Ceil(float64(pz/consts.CellLength) - 0.5)
	return uint32(row), uint32(col)
}
