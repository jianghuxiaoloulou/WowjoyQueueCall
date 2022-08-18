package global

type TextCfg struct {
	PateintNameCount int // 患者姓名计数
	MachineRoomCount int // 检查机房计数
	QueueNumberCount int // 排队号计数
	TypeNameCount    int // 就诊类型计数
	CheckTypeCount   int // 检查类型计数
	WaitPatientCount int // 等待患者计数
}

var (
	ObjectDataChan chan Patient_Info
)
