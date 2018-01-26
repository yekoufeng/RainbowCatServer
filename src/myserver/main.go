package main

import (
	"base/glog"
	"base/gonet"
	"math/rand"
	"myserver/playertask"
	"myserver/playertaskmgr"
	"myserver/roommgr"
	"time"
)

type MyServer struct {
	gonet.Service
	myser *gonet.TcpServer
}

func myServerMain() {
	svr := MyServer{
		myser: &gonet.TcpServer{},
	}
	svr.Derived = &svr
	svr.Main()
}

func (this *MyServer) Init() bool {
	glog.Error("开始初始化")
	rand.Seed(time.Now().Unix())

	// 绑定本地端口
	var address string = "192.168.240.241:8000"
	err := this.myser.Bind(address)
	if err != nil {
		glog.Error("绑定端口失败")
		return false
	}
	glog.Error("完成初始化")
	glog.Error("监听端口 192.168.240.241:8000")
	return true
}

func (this *MyServer) MainLoop() {
	conn, err := this.myser.Accept()
	if err != nil {
		return
	}
	newtask := playertask.NewPlayerTask(conn)
	playertaskmgr.GetMe().AddPlayerTask(newtask)
	newtask.Start()
}

func (this *MyServer) Final() bool {
	glog.Error("roomser 挂了")
	return true
}

//TODO Test
func main() {
	//初始化房间管理器 玩家会话管理器
	roommgr.GetMe()
	playertaskmgr.GetMe()
	myServerMain()
}
