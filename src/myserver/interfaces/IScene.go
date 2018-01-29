package interfaces

//	"usercmd"

type IScene interface {
	IsPlayerOnCell(row uint32, col uint32) bool
	BroadCastMsg(data []byte, flag byte)
	IsInGame() bool
	SetItemOnCell(row uint32, col uint32)
	DeleteItemOnCell(row uint32, col uint32)
}
