package global

// 屏幕显示的科室类型
const (
	Screen_Type_FS int = 0 // 放射科屏幕
	Screen_Type_US int = 1 // 超声科屏幕
	Screen_Type_ES int = 2 // 内镜科屏幕
	Screen_Type_MZ int = 3 // 门诊屏幕
)

type TextCfg struct {
	PateintNameCount int // 患者姓名计数
	MachineRoomCount int // 检查机房计数
	QueueNumberCount int // 排队号计数
	TypeNameCount    int // 就诊类型计数
	CheckTypeCount   int // 检查类型计数
	WaitPatientCount int // 等待患者计数
}

// 医技发送数据
type CallData struct {
	CurrRoom         string        `json:"currRoom"`         // 当前呼叫机房
	CurrIndex        int           `json:"currindex"`        // 当前机房位置
	CurPatientName   string        `json:"curPatientName"`   // 当前患者姓名
	CurQueueNumber   string        `json:"curQueueNumber"`   // 当前排队号
	CurCheckNumber   string        `json:"curCheckNumber"`   // 当前检查号
	CurCheckType     string        `json:"curCheckType"`     // 当前检查类型
	CurTypeName      string        `json:"curTypeName"`      // 当前就诊类型名
	CurGreenFlag     string        `json:"curGreenFlag"`     // 当前绿色通道标志（是/否）
	WaitePatienTotal int           `json:"waitePatienTotal"` // 等待患者总数
	WaitePatienList  []WaitePatien `json:"waitePatienList"`  // 等待患者列表
}

// 显示端数据
type ScreenShowData struct {
	CurShowCallPoint      int      `json:"curShowCallPoint"`      // 显示屏幕的呼叫点
	CurShowIP             string   `json:"curShowIP"`             // 当前显示屏幕的IP
	CurShowSize           int      `json:"curShowSize"`           // 当前屏幕显示数据条数
	CurShowDepartmentCode int      `json:"curShowDepartmentCode"` // 当前显示的科室类型
	CurWavFile            string   `json:"curWavFile"`            // 当前呼叫的语音文件
	RoomInfo              CallData `json:"roomInfo"`              // 房间显示信息
}

type WaitePatien struct {
	WaitQueueNumber  string `json:"waiteQueueNumber"` // 等待患者排队号
	WaitePatientName string `json:"waitePatientName"` // 等待患者姓名
	WaiteTypeName    string `json:"waiteTypeName"`    // 等待患者就诊类型
	WaiteGreenFlag   string `json:"waiteGreenFlag"`   // 等待患者绿色通道标志（是/否）
}

// 发送的科室的WebSocket数据
type WSMessage struct {
	To       string   `json:"to"`       // 接收者IP
	MsgType  int      `json:"msgType"`  // 数据类型 0:ws通讯信息，1:屏幕配置信息，2:叫号数据消息,3:ws重连数据显示信息
	DataType int      `json:"dataType"` // 数据类型 0：放射科，1：超声科，2：内镜科，3：门诊
	Data     FSWSData `json:"data"`     // 放射科数据
}
type FSWSData struct {
	CallPatient     string     `json:"callPatient"`     // 呼叫的患者
	CallRoom        string     `json:"callRoom"`        // 呼叫的科室
	CallQueueNum    string     `json:"callQueueNum"`    // 呼叫的排队号
	CallPatientType string     `json:"callPatientType"` // 呼叫患者的就诊类型
	CallWavFile     string     `json:"callWavFile"`     // 呼叫的语音文件
	CheckRoomList   []CallData `json:"checkRoomList"`   // 显示的机房数据列表
}

// 发送的门诊的WebSocket数据MZ
type WSMZMessage struct {
	To       string   `json:"to"`       // 接收者IP
	MsgType  int      `json:"msgType"`  // 数据类型 0:ws通讯信息，1:屏幕配置信息，2:叫号数据消息,3:ws重连数据显示信息
	DataType int      `json:"dataType"` // 数据类型 0：放射科，1：超声科，2：内镜科，3：门诊
	Data     CallData `json:"data"`     // 放射科数据
}

// 前端的log
type WebLog struct {
	Level int    `json:"level"` // log 等级
	Msg   string `json:"msg"`   // log 消息
}

// 增加数据签到时间
var ArriveTime string

// 屏幕中的数据 [屏幕IP][][机房]
var ScreenRoomTotalData map[string]map[string]CallData

// 超声内镜数据
var USScreenRoomTotalData map[string][]CallData
