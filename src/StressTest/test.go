package main

import (
	"base/glog"
	"base/gonet"
	"common"
	"net"
	"os"
	"strconv"
	"usercmd"
)

type playertask struct {
	gonet.TcpTask
	id uint32
}

func (this *playertask) ParseMsg(data []byte, flag byte) bool {
	cmd := usercmd.DemoTypeCmd(common.GetCmd(data))
	glog.Error("cmd = ", cmd)
	switch cmd {
	case usercmd.DemoTypeCmd_LoginRes:
		recvCmd, ok := common.DecodeCmd(data, flag, &usercmd.LoginS2CMsg{}).(*usercmd.LoginS2CMsg)
		if !ok {
			glog.Error("login res error")
		}
		glog.Error("[login res] id = ", recvCmd.GetPlayerId())
		this.id = recvCmd.GetPlayerId()
		this.searchReq()

	case usercmd.DemoTypeCmd_SearchRes:
		recvCmd, ok := common.DecodeCmd(data, flag, &usercmd.SearchS2CMsg{}).(*usercmd.SearchS2CMsg)
		if !ok {
			glog.Error("search res error")
		}
		glog.Error("[search res] roomid = ", recvCmd.GetRoomId())
		this.moveReq()

	case usercmd.DemoTypeCmd_MoveRes:
		recvCmd, ok := common.DecodeCmd(data, flag, &usercmd.MoveS2CMsg{}).(*usercmd.MoveS2CMsg)
		if !ok {
			glog.Error("move res error")
		}
		glog.Error("[move res] playerId = ", recvCmd.GetPlayerId())

	}

	return true
}

func (this *playertask) loginReq() {
	glog.Error("login req")
	m := usercmd.LoginC2SMsg{
		Name: "玩家123",
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_LoginReq), &m)
	this.AsyncSend(d, f)
}

func (this *playertask) searchReq() {
	glog.Error("search req")
	m := usercmd.SearchC2SMsg{
		PlayerId: this.id,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_SearchReq), &m)
	this.AsyncSend(d, f)
}

func (this *playertask) moveReq() {
	glog.Error("move req")
	m := usercmd.MoveC2SMsg{
		PosX: 10,
		PosY: 9,
		PosZ: 8,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_MoveReq), &m)
	this.AsyncSend(d, f)
}

func (this *playertask) OnClose() {
	glog.Error("断开连接")
}

func NewTcpCon(tmp string) *playertask {
	conn, _ := net.Dial("tcp", tmp)
	s := &playertask{
		TcpTask: *gonet.NewTcpTask(conn),
	}
	s.Derived = s
	// 设置发送缓冲区限制
	s.SetSendBuffSizeLimt(256 * 1024)
	s.Start()
	return s
}

func main() {
	glog.Error("压力测试开始... ")

	ipTmp := make([]byte, 25)
	os.Stdin.Read(ipTmp)
	glog.Error("读取ip成功 ", string(ipTmp))

	var connNum int = 0
	for {
		input := make([]byte, 1)
		os.Stdin.Read(input)
		cmdStr := string(input[0:1])
		cmdInt, _ := strconv.Atoi(cmdStr)

		switch cmdInt {
		case 1:
			glog.Error("读取指令成功  生成50个新连接")
			for i := 0; i < 50; i++ {
				connNum++
				glog.Error("测试连接个数 : ", connNum)
				conn, _ := net.Dial("tcp", string(ipTmp))
				s := &playertask{
					TcpTask: *gonet.NewTcpTask(conn),
				}
				s.Derived = s
				// 设置发送缓冲区限制
				s.SetSendBuffSizeLimt(256 * 1024)
				s.Start()
				s.loginReq()
			}
		case 2:
			glog.Error("读取指令成功  生成500个新连接")
			for i := 0; i < 500; i++ {
				connNum++
				glog.Error("测试连接个数 : ", connNum)
				conn, _ := net.Dial("tcp", string(ipTmp))
				s := &playertask{
					TcpTask: *gonet.NewTcpTask(conn),
				}
				s.Derived = s
				// 设置发送缓冲区限制
				s.SetSendBuffSizeLimt(256 * 1024)
				s.Start()
				s.loginReq()
			}
		case 3:
			glog.Error("读取指令成功  生成1000个新连接")
			connNum++
			glog.Error("测试连接个数 : ", connNum)
			conn, _ := net.Dial("tcp", string(ipTmp))
			s := &playertask{
				TcpTask: *gonet.NewTcpTask(conn),
			}
			s.Derived = s
			// 设置发送缓冲区限制
			s.SetSendBuffSizeLimt(256 * 1024)
			s.Start()
			s.loginReq()
		}
	}
}
