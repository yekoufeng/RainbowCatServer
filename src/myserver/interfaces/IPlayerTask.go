package interfaces

//	"usercmd"

type IPlayerTask interface {
	//OnRecvPlayerCmd(cmd usercmd.DemoTypeCmd, data []byte, flag byte)
	AsyncSend(buffer []byte, flag byte) bool
	SetId(id uint32)
	GetId() uint32
	GetName() string
	SetRoom(room IRoom)
}
