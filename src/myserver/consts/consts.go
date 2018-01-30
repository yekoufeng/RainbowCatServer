package consts

const (
	IpAddress           = "192.168.240.241:8000" //ip地址
	CellLength  float32 = 1.0                    //方块正方形边长大小
	CellNum     uint32  = 20                     //方块正方形边长个数
	OneGameTime         = 60                     //一分钟

	CountDownTime      = 10  //充能开始时间10秒
	EnergyRepeatedTime = 500 //0.5秒充能一次  500ms = 0.5s
	OneGamePlayerNum   = 3   //一局游戏玩家人数
	TotalEnergyNum     = 100 //能量总数

	ItemNumOneTime = 10 //一次生成多少个道具
	ItemLiveTime   = 10 //道具刷新时间
	ItemCreateTime = 10 //道具开始生成
	AllItemsNum    = 2  //道具种类个数
)
