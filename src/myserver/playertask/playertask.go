package playertask

import (
	"base/glog"
	"base/gonet"
	"common"
	"myserver/interfaces"
	"myserver/playertaskmgr"
	"myserver/roommgr"
	"net"
	"usercmd"
)

type PlayerTask struct {
	gonet.TcpTask
	id          uint32
	name        string
	room        interfaces.IRoom
	isInRoom    bool
	isSearching bool
}

func NewPlayerTask(conn net.Conn) *PlayerTask {
	s := &PlayerTask{
		TcpTask:     *gonet.NewTcpTask(conn),
		name:        "",
		isInRoom:    false,
		isSearching: false,
	}
	s.Derived = s
	// 设置发送缓冲区限制
	s.SetSendBuffSizeLimt(256 * 1024)
	return s
}

func (this *PlayerTask) SetId(id uint32) {
	this.id = id
}

func (this *PlayerTask) GetId() uint32 {
	return this.id
}

func (this *PlayerTask) GetName() string {
	return this.name
}

func (this *PlayerTask) ParseMsg(data []byte, flag byte) bool {
	cmd := usercmd.DemoTypeCmd(common.GetCmd(data))
	//glog.Error("cmd = ", cmd)
	switch cmd {
	//把所有不需要和房间有关的cmd放在外面处理
	case usercmd.DemoTypeCmd_LoginReq:
		this.handleLogin(data, flag)

	case usercmd.DemoTypeCmd_SearchReq:
		this.handleSearch(data, flag)
	//所有和房间有关的cmd放进各自用户的room里处理
	default:
		this.room.PostPlayerCmd(this.id, cmd, data, flag)
	}
	return true
}

func (this *PlayerTask) handleLogin(data []byte, flag byte) {
	if this.name != "" {
		glog.Error("[bug] login twice!")
		return
	}
	recvCmd, ok := common.DecodeCmd(data, flag, &usercmd.LoginC2SMsg{}).(*usercmd.LoginC2SMsg)
	glog.Error("[Login req]")
	if !ok {
		glog.Error("Login req error")
		return
	}
	glog.Error("login name = ", recvCmd.GetName())
	this.name = recvCmd.GetName()
	m := usercmd.LoginS2CMsg{
		PlayerId: this.id,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_LoginRes), &m)
	this.AsyncSend(d, f)
	glog.Error("[login success] name = ", recvCmd.GetName(), " id = ", this.id)
}

func (this *PlayerTask) handleSearch(data []byte, flag byte) {
	if this.isSearching || this.isInRoom {
		glog.Error("[bug] is in searching")
		return
	}
	recvCmd, ok := common.DecodeCmd(data, flag, &usercmd.SearchC2SMsg{}).(*usercmd.SearchC2SMsg)
	glog.Error("[Search req]")
	if !ok {
		glog.Error("Search req error")
		return
	}
	glog.Error("Search id = ", recvCmd.GetPlayerId())
	if this.id != recvCmd.GetPlayerId() {
		glog.Error("[bug] 客户端请求id与本身id不匹配 客户端id = ", recvCmd.GetPlayerId(), " 本身id = ", this.id)
		return
	}
	this.isSearching = true
	roommgr.GetMe().AddSearchPlayer(recvCmd.GetPlayerId())
}

func (this *PlayerTask) OnClose() {
	playertaskmgr.GetMe().DeletePlayerTask(this)
	if this.isSearching {
		roommgr.GetMe().DeleteSearchPlayer(this.id)
	}
	glog.Error("disconnect id = ", this.id, " name = ", this.name, " isInRoom = ", this.isInRoom, " isSearching = ", this.isSearching)
}

func (this *PlayerTask) SetRoom(room interfaces.IRoom) {
	this.room = room
	this.isSearching = false
	this.isInRoom = true
}
