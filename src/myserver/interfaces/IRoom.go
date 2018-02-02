package interfaces

import (
	"usercmd"
)

type IRoom interface {
	PostPlayerCmd(playerID uint32, cmd usercmd.DemoTypeCmd, data []byte, flag byte)
	GetRoomId() uint32
	AddPlayer(id []uint32)
	BroadCastMsg(data []byte, flag byte)
	SetCellColor(row uint32, col uint32, color usercmd.ColorType)
	GetCellColor(row uint32, col uint32) usercmd.ColorType
	AddColorNum(color usercmd.ColorType)
	DeleteColorNum(color usercmd.ColorType)
	HandleGameOver(color usercmd.ColorType)
	MoveFromToCell(arow uint32, acol uint32, brow uint32, bcol uint32, pid uint32, itemNum uint32)
	IsInGame() bool
	DyeingFun(row uint32, col uint32, color usercmd.ColorType, pId uint32)
	SetCellVirus(row uint32, col uint32, pId uint32)
	IsCellVirus(row uint32, col uint32) bool
	RemoveCellVirus(row uint32, col uint32)
}
