package scene

// 房间玩家协议处理 辅助类

import (
	"base/glog"
	"common"
	"usercmd"
)

type ScenePlayerNetMsgHelper struct {
	msgHandlerMap MsgHandlerMap // 玩家协议处理器
	selfPlayer    *ScenePlayer  // 玩家自身的引用
}

func (this *ScenePlayerNetMsgHelper) Init(selfPlayer *ScenePlayer) {
	this.selfPlayer = selfPlayer
	this.msgHandlerMap.Init()
	this.RegCmds()
}

//注册网络消息
func (this *ScenePlayerNetMsgHelper) RegCmds() {
	this.msgHandlerMap.RegisterHandler(usercmd.DemoTypeCmd_MoveReq, this.OnNetMove)
}

//收到玩家消息
func (this *ScenePlayerNetMsgHelper) OnRecvPlayerCmd(cmd usercmd.DemoTypeCmd, data []byte, flag byte) {
	this.msgHandlerMap.Call(cmd, data, flag)
}

func (this *ScenePlayerNetMsgHelper) OnNetMove(data []byte, flag byte) {
	op, ok := common.DecodeCmd(data, flag, &usercmd.MoveC2SMsg{}).(*usercmd.MoveC2SMsg)
	if !ok {
		glog.Error("DecodeCmd error : OnNetMove")
		return
	}
	this.selfPlayer.HandleMove(op.GetPosX(), op.GetPosY(), op.GetPosZ())
}
