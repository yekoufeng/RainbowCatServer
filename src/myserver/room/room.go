package room

import (
	"base/glog"
	"common"
	"myserver/consts"
	"myserver/playertaskmgr"
	"myserver/scene"
	"time"
	"usercmd"
)

type PlayerCmd struct {
	PlayerID uint32
	Cmd      usercmd.DemoTypeCmd
	Data     []byte
	Flag     byte
}

type Room struct {
	scene.Scene //场景

	roomId         uint32          //房间的标志id
	playerNum      uint32          //房间当前人数
	isInGame       bool            //房间是否正在游戏
	startTime      int64           // 开始时间
	chan_PlayerCmd chan *PlayerCmd // 玩家输入
	playerIds      []uint32        // playerIds
	tloop          uint32          //房间当前进行时间
}

func NewRoom(rid uint32) *Room {
	room := Room{
		roomId:         rid,
		playerNum:      0,
		isInGame:       false,
		chan_PlayerCmd: make(chan *PlayerCmd, 1024),
		tloop:          0,
	}
	return &room
}

func (this Room) GetRoomId() uint32 {
	return this.roomId
}

// PostPlayerCmd 发送玩家命令到 chan_PlayerCmd. 命令在房间协程中执行。
func (this Room) PostPlayerCmd(playerID uint32, cmd usercmd.DemoTypeCmd,
	data []byte, flag byte) {
	playerCmd := &PlayerCmd{PlayerID: playerID, Cmd: cmd, Flag: flag}
	playerCmd.Data = make([]byte, len(data))
	copy(playerCmd.Data, data)
	this.chan_PlayerCmd <- playerCmd
}

func (this *Room) Start() {
	this.startTime = time.Now().Unix()
	glog.Error("游戏开始 房间号 roomId = ", this.roomId, "当前玩家人数 = ", len(this.playerIds))
	this.isInGame = true
	this.InitPlayerPosition()
	m := usercmd.GameStartS2CMsg{
		Edgenum: consts.CellNum,
	}
	for _, pId := range this.playerIds {
		m.Nums = append(m.Nums, &usercmd.GameStartS2CMsgPosition{
			PlayerId: pId,
			Col:      this.Players[pId].GetCol(),
			Row:      this.Players[pId].GetRow(),
			Color:    this.Players[pId].Color,
		})
		glog.Error("出生位置", this.Players[pId].GetRow(), "  ", this.Players[pId].GetCol())
	}

	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_GameStart), &m)
	//广播
	this.BroadCastMsg(d, f)
	//延迟三秒钟 为了 3 2 1 go
	var timer = time.NewTicker(time.Second * consts.GameStartWaitTime)

	go func() {
		for true {
			<-timer.C
			go this.Loop()
			break
		}

	}()
}

//主循环
func (this *Room) Loop() {
	glog.Error("房间loop 开始")
	var repeatedTimer = time.NewTicker(time.Second)
	for {
		select {
		case op := <-this.chan_PlayerCmd:
			if this.isInGame {
				player, ok := this.Players[op.PlayerID]
				if ok {
					//glog.Error("[CMD] cmd = ", op.Cmd, " playerId = ", op.PlayerID)
					player.OnRecvPlayerCmd(op.Cmd, op.Data, op.Flag)
				} else {
					glog.Error("PlayerCmd:no player,", op.PlayerID, " cmd:", op.Cmd)
				}
			}
		case <-repeatedTimer.C:
			if !this.isInGame {
				repeatedTimer.Stop()
				glog.Error("房间loop停止")
			}
			//glog.Error("tloop = ", this.tloop)
			//默认同步所有客户端一次时间
			this.handleSynTime()
			if this.tloop == consts.CountDownTime {
				//充能开始
				this.handleGameEnergy()
			}
			if this.tloop == consts.ItemCreateTime {
				//道具开始
				this.handleItemCreate()
			}
			if this.tloop == consts.OneGameTime {
				//一局游戏时间到
				this.HandleGameOver(usercmd.ColorType_origin)
			}
			this.tloop++
		}
	}
}

func (this *Room) handleItemCreate() {
	//scene itemmgr函数
	this.StartLoop()
}

//同步房间内时间 1秒一次
func (this *Room) handleSynTime() {
	minTmp, secTmp := getMinAndSecByLoop(this.tloop)
	m := usercmd.SynTimeS2CMsg{
		Tloop:  this.tloop,
		Minute: minTmp,
		Second: secTmp,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_GameTime), &m)
	this.BroadCastMsg(d, f)
}

func getMinAndSecByLoop(loop uint32) (uint32, uint32) {
	leftTime := consts.OneGameTime - loop
	aTmp := leftTime / 60
	bTmp := leftTime % 60
	return aTmp, bTmp
}

func (this *Room) handleGameEnergy() {
	var timer = time.NewTicker((time.Millisecond * consts.EnergyRepeatedTime))

	go func() {
		for true {
			<-timer.C
			//能量条一次
			this.AddEnergyInScene()
			if !this.isInGame {
				timer.Stop()
			}
		}
	}()
}

func (this *Room) HandleGameOver(color usercmd.ColorType) {
	if !this.isInGame {
		return
	}
	this.Scene.DeleteAllItems()
	this.isInGame = false
	m := usercmd.GameEndS2CMsg{
		WinColor: color,
	}
	//name的获取如果玩家在游戏中就断线了，可能导致这边异常
	for _, pId := range this.playerIds {
		playerTmp := playertaskmgr.GetMe().GetPlayerTask(pId)
		var nameTmp string = "玩家已下线"
		if playerTmp != nil {
			//说明玩家没有断开连接
			nameTmp = playerTmp.GetName()
		}
		m.Nums = append(m.Nums, &usercmd.GameEndS2CMsgPlayerMsg{
			PlayerId: pId,
			Name:     nameTmp,
			Cellnum:  this.Players[pId].GetOnePlayerCellNum(),
			Color:    this.Players[pId].Color,
		})
	}

	if color == usercmd.ColorType_origin {
		//没有一个队伍达到100%
		//根据max来判断谁获胜
		m.WinColor = this.MaxCellColor
	}
	m.Mvpid = this.GetMvpId(m.WinColor)
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_GameEnd), &m)
	//广播
	this.BroadCastMsg(d, f)
	glog.Error("获胜队伍颜色是", this.MaxCellColor)

	//把所有玩家踢出去房间
	//操作sceneplayer的playtask接口
	this.Scene.GameOverForEveryOne()
}

func (this *Room) AddPlayer(id []uint32) {
	if len(id) != consts.OneGamePlayerNum {
		glog.Error("room 人数不匹配 error")
	}
	this.SceneInit(this)
	m := usercmd.SearchS2CMsg{
		RoomId: this.roomId,
	}

	for i, playerid := range id {
		this.playerIds = append(this.playerIds, playerid)
		this.Players[playerid] = scene.NewScenePlayer(playerid, this)
		playTaskTmp := playertaskmgr.GetMe().GetPlayerTask(playerid)
		playTaskTmp.SetRoom(this)
		glog.Error("第", i, "个玩家初始化成功")
		m.Nums = append(m.Nums, &usercmd.SearchS2CMsgPlayer{
			PlayerId: playerid,
			Name:     playertaskmgr.GetMe().GetPlayerTask(playerid).GetName(),
		})
	}
	glog.Error("所有玩家初始化完毕")
	glog.Error("游戏开始")

	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_SearchRes), &m)
	//广播
	this.BroadCastMsg(d, f)
	this.Start()
}

func (this *Room) BroadCastMsg(data []byte, flag byte) {
	for _, player := range this.Players {
		player.SendMsg(data, flag)
	}
}

func (this *Room) IsInGame() bool {
	return this.isInGame
}
