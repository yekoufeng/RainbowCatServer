package itemmgr

import (
	"base/glog"
	"common"
	"math/rand"
	"myserver/consts"
	"myserver/interfaces"
	"myserver/item"
	"time"
	"usercmd"
)

type ItemMgr struct {
	items []*item.Item //当前场景内道具
	Scene interfaces.IScene
}

func (this *ItemMgr) StartLoop() {
	//道具循环开始入口
	glog.Error("[start] Item loop")
	timer := time.NewTicker(time.Second * consts.ItemLiveTime)
	for n := 0; n < consts.ItemNumOneTime; n++ {
		this.CreateItem()
	}
	glog.Error("当前场景道具总数: ", len(this.items))
	go func() {
		for true {
			<-timer.C
			glog.Error("道具刷新")
			if !this.Scene.IsInGame() {
				timer.Stop()
				glog.Error("道具管理结束loop")
			}
			//this.RefreshItem()
		}
	}()
}

func (this *ItemMgr) RandItemtype() usercmd.ItemType {
	//随机一个数 在 1-n范围内  n = 道具个数
	//先取 0 - n-1  再加1
	return usercmd.ItemType(rand.Int()%consts.AllItemsNum + 1)
}

func (this *ItemMgr) CreateItem() {
	var itemRow uint32 = rand.Uint32() % consts.CellNum
	var itemCol uint32 = rand.Uint32() % consts.CellNum
	//如果有玩家在格子上，重新随机生成
	if this.Scene.IsPlayerOnCell(itemRow, itemCol) {
		this.CreateItem()
		return
	}
	//道具种类随机
	//TODO 方便测试 全是染色道具
	//itemTmp := item.NewItem(itemRow, itemCol, this.RandItemtype())
	itemTmp := item.NewItem(itemRow, itemCol, usercmd.ItemType_dyeing)
	this.Scene.SetItemOnCell(itemRow, itemCol)
	//广播道具生成信息
	m := usercmd.CreateItemsS2CMsg{
		Row:  itemRow,
		Col:  itemCol,
		Item: usercmd.ItemType_unknown,
	}
	d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_ItemCreate), &m)
	this.Scene.BroadCastMsg(d, f)
	this.items = append(this.items, itemTmp)
}

func (this *ItemMgr) DeleteAllItems() {
	//释放所有当前的道具
	for _, itemTmp := range this.items {
		this.Scene.DeleteItemOnCell(itemTmp.Row, itemTmp.Col)
		//广播道具销毁信息
		m := usercmd.DestroyItemsS2CMsg{
			Row: itemTmp.Row,
			Col: itemTmp.Col,
		}
		d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_ItemDestroy), &m)
		this.Scene.BroadCastMsg(d, f)
	}
	//清空道具切片
	this.items = this.items[:0]
}

func (this *ItemMgr) RefreshItem() {
	//检测上次残留道具
	if len(this.items) != 0 {
		this.DeleteAllItems()
	}
	//生成固定个数道具
	for n := 0; n < consts.ItemNumOneTime; n++ {
		this.CreateItem()
	}
	glog.Error("当前场景道具总数: ", len(this.items))
}

func (this *ItemMgr) DeleteOneItem(row uint32, col uint32) {
	//释放所有当前的道具
	for _, itemTmp := range this.items {
		//如果对应
		if itemTmp.Row == row && itemTmp.Col == col {
			this.Scene.DeleteItemOnCell(row, col)
			//广播道具销毁信息
			m := usercmd.DestroyItemsS2CMsg{
				Row: row,
				Col: col,
			}
			d, f, _ := common.EncodeGoCmd(uint16(usercmd.DemoTypeCmd_ItemDestroy), &m)
			this.Scene.BroadCastMsg(d, f)
			return
		}
	}
}

func (this *ItemMgr) GetItemByRowCol(row uint32, col uint32) usercmd.ItemType {
	for _, itemTmp := range this.items {
		if itemTmp.Row == row && itemTmp.Col == col {
			return itemTmp.ItemType
		}
	}
	glog.Error("[bug]  应该存在对应的道具")
	return usercmd.ItemType_unknown
}
