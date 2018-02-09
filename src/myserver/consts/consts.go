package consts

const (
	IpAddress                 = "192.168.212.192:8000" //ip地址
	CellLength        float32 = 1.0                    //方块正方形边长大小
	CellNum           uint32  = 25                     //方块正方形边长个数
	OneGameTime               = 180                    //三分钟
	GameStartWaitTime         = 4                      //开局等待时间

	CountDownTime      = 10  //充能开始时间10秒
	EnergyRepeatedTime = 500 //1秒充能一次  500ms = 0.5s
	OneGamePlayerNum   = 3   //一局游戏玩家人数
	TotalEnergyNum     = 100 //能量总数

	ItemNumOneTime = 12 //一次生成多少个道具
	ItemLiveTime   = 30 //道具刷新时间
	ItemCreateTime = 10 //道具开始生成
	AllItemsNum    = 4  //道具种类个数

	MaxItemNum = 2 //一个玩家最大道具数量

	DyeingRange = 4 //染色道具范围大小
	DyeingTime  = 6 //染色持续时间

	VirusCellNum = 5 //病毒陷阱数目

	ImprisonTime = 2 //禁锢时间2秒
	DizzyTime    = 5 //神魂颠倒时间

	SpeedTime = 5   //加速5s
	SpeedNum  = 150 //加速150%
)
