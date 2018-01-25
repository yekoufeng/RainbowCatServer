package playertaskmgr

import (
	"base/glog"
	"math/rand"
	"myserver/interfaces"
	"sync"
)

type PlayerTaskMgr struct {
	mutex sync.RWMutex
	tasks map[uint32]interfaces.IPlayerTask
}

//单例模式
var (
	playtaskmgr      *PlayerTaskMgr
	playtaskmgr_once sync.Once
)

func GetMe() *PlayerTaskMgr {
	if playtaskmgr == nil {
		playtaskmgr_once.Do(func() {
			playtaskmgr = &PlayerTaskMgr{
				tasks: make(map[uint32]interfaces.IPlayerTask),
			}
			playtaskmgr.init()
		})
	}
	return playtaskmgr
}

func (this *PlayerTaskMgr) init() {
	glog.Error("PlayerTaskMgr初始化成功")
}

func (this *PlayerTaskMgr) AddPlayerTask(playertask interfaces.IPlayerTask) {
	//TODO 避免重复
	tmp := rand.Uint32() % 10000
	glog.Error("[new task] id = ", tmp)
	playertask.SetId(tmp)
	this.mutex.Lock()
	this.tasks[tmp] = playertask
	this.mutex.Unlock()
}

func (this *PlayerTaskMgr) DeletePlayerTask(playertask interfaces.IPlayerTask) {
	tmp := playertask.GetId()
	glog.Error("[delete task] id = ", tmp)
	this.mutex.Lock()
	delete(this.tasks, tmp)
	this.mutex.Unlock()
}

func (this *PlayerTaskMgr) GetPlayerTask(id uint32) interfaces.IPlayerTask {
	tmp, ok := this.tasks[id]
	if !ok {
		glog.Error("[错误] id = ", id)
	}
	return tmp
}
