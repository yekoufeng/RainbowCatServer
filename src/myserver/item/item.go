package item

import (
	"base/glog"
	"usercmd"
)

type Item struct {
	Row      uint32
	Col      uint32
	ItemType usercmd.ItemType
}

func NewItem(r uint32, c uint32, itype usercmd.ItemType) *Item {
	item := Item{
		Row:      r,
		Col:      c,
		ItemType: itype,
	}
	glog.Error("新道具生成 row = ", r, " col = ", c)
	return &item
}
