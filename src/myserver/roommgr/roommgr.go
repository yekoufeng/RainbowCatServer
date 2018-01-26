package roommgr

import (
	"base/glog"
	"math/rand"
	"myserver/consts"
	rm "myserver/room"
	"sync"
)

type RoomMgr struct {
	roomMutex     sync.RWMutex
	rooms         map[uint32]*rm.Room
	roomNum       uint32
	searchMutex   sync.RWMutex
	searchPlayers []uint32
}

//单例模式
var (
	roommgr      *RoomMgr
	roommgr_once sync.Once
)

func GetMe() *RoomMgr {
	if roommgr == nil {
		roommgr_once.Do(func() {
			roommgr = &RoomMgr{
				rooms:   make(map[uint32]*rm.Room),
				roomNum: 0,
			}
			roommgr.init()
		})
	}
	return roommgr
}

func (this *RoomMgr) init() {
	glog.Error("RoomMgr初始化成功")
}

func (this *RoomMgr) AddSearchPlayer(id uint32) {
	this.searchMutex.Lock()
	defer this.searchMutex.Unlock()
	this.searchPlayers = append(this.searchPlayers, id)
	waitingNums := len(this.searchPlayers)

	glog.Error("[匹配] 添加 [", id, "] 到匹配队列, 当前队列人数 ", waitingNums)
	glog.Error(waitingNums, " ", consts.OneGamePlayerNum)
	if waitingNums == consts.OneGamePlayerNum {
		//从队列删除前两个数据
		searchSuccessPlayerIds := this.searchPlayers[:waitingNums]
		this.searchPlayers = this.searchPlayers[waitingNums:]
		glog.Error("从匹配队列删除，当前队列人数 = ", len(this.searchPlayers))
		this.NewAndAddRoom(searchSuccessPlayerIds)
	}

}

func (this *RoomMgr) DeleteSearchPlayer(id uint32) {
	this.searchMutex.Lock()
	defer this.searchMutex.Unlock()
	//玩家下线，搜索匹配队列确实存在该玩家id
	pos, ok := -1, false
	for i, tmpId := range this.searchPlayers {
		if tmpId == id {
			ok = true
			pos = i
			break
		}
	}
	if ok {
		//移除该id
		this.searchPlayers = append(this.searchPlayers[:pos], this.searchPlayers[pos+1:]...)
	}
}

//新建一个房间，并且把匹配的人加入这个房间
func (this *RoomMgr) NewAndAddRoom(playerIds []uint32) {
	num := len(playerIds)
	glog.Error("匹配成功 匹配人数 = ", num)
	if num != consts.OneGamePlayerNum {
		glog.Error("匹配人数error")
		return
	}
	roomId := this.NewRoom()
	glog.Error("新建房间id = ", roomId)
	this.rooms[roomId].AddPlayer(playerIds)
}

func (this *RoomMgr) NewRoom() uint32 {
	this.roomMutex.Lock()
	defer this.roomMutex.Unlock()

	var roomId uint32 = rand.Uint32() % 10000
	//TODO 避免重复
	this.rooms[roomId] = rm.NewRoom(roomId)
	return roomId
}
