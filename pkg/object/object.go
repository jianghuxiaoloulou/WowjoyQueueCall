package object

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/routers/api/ws"
	"os"
	"os/exec"
	"sort"
	"syscall"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// 退出进程
func QuitExe(exeName string) {
	cmd := exec.Command("taskkill.exe", "/F", "/im", exeName)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		global.Logger.Error("退出程序失败", err)
		return
	}
}

// 调用exe生成wav语音文件
func CallExeSaveWavFile(str, fileName string) bool {
	// 如果文件没有生成，重试三次
	count := 1
	exepath := global.ObjectSetting.TTSPath
	arg := make([]string, 0)
	arg = append(arg, str)
	arg = append(arg, fileName)
	global.Logger.Info(exepath, " 参数是：", arg)

LOOP:
	cmd := exec.Command(exepath, arg...)
	if err := cmd.Run(); err != nil {
		global.Logger.Error("生成Wav文件失败")
		return false
	}
	// 判断生成的语音文件是否存在
	_, err := os.Stat(fileName)
	if err != nil {
		QuitExe("TTSCfg.exe")
		count++
		if count < 4 {
			global.Logger.Error("语音文件不存在,重新生成第: ", count, " 次")
			goto LOOP
		}
		global.Logger.Error("Wav语音文件不存在")
		return false
	}
	global.Logger.Info("Wav语音文件生成成功")
	return true
}

// 通过go-ole库生成语音文件
func GetWavFile(str, path string) bool {
	// var err error
	global.Logger.Debug("呼叫内容：", str)
	count := 1
LOOP:
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("SAPI.SpVoice")
	voice, _ := unknown.QueryInterface(ole.IID_IDispatch)
	saveFile, _ := oleutil.CreateObject("SAPI.SpFileStream")
	ff, _ := saveFile.QueryInterface(ole.IID_IDispatch)
	// 打开wav文件
	oleutil.CallMethod(ff, "Open", path, 3, true)
	// 设置voice的AudioOutputStream属性，必须是PutPropertyRef，如果是PutProperty就无法生效
	oleutil.PutPropertyRef(voice, "AudioOutputStream", ff)
	// 设置语速
	oleutil.PutProperty(voice, "Rate", global.ObjectSetting.Rate)
	// 设置音量
	oleutil.PutProperty(voice, "Volume", global.ObjectSetting.Volume)
	// 说话
	oleutil.CallMethod(voice, "Speak", str)
	// oleutil.CallMethod(voice, "Speak", "bb", 1)
	// 停止说话
	//oleutil.CallMethod(voice, "Pause")
	// 恢复说话
	//oleutil.CallMethod(voice, "Resume")
	// 等待结束
	oleutil.CallMethod(voice, "WaitUntilDone", 1000000)
	// 关闭文件
	oleutil.CallMethod(ff, "Close")
	ff.Release()
	voice.Release()
	ole.CoUninitialize()
	_, err := os.Stat(path)
	if err != nil {
		count++
		if count < 4 {
			global.Logger.Error("语音文件不存在,重新生成第: ", count, " 次")
			goto LOOP
		}
		global.Logger.Error("Wav语音文件不存在: ", path)
		return false
	}
	global.Logger.Info("Wav语音文件生成成功: ", path)
	return true
}

// 分发显示信息到显示屏
func SendShowScreenInfo(info global.ScreenShowData) {
	// for _, value := range info {
	switch info.CurShowDepartmentCode {
	case global.Screen_Type_FS:
		ShowFSScreenInfo(info)
	case global.Screen_Type_US:
		ShowUSScreenInfo(info)
	case global.Screen_Type_ES:
		ShowESScreenInfo(info)
	case global.Screen_Type_MZ:
		ShowMZScreenInfo(info)
	default:
	}
}

// 显示放射科屏幕信息
func ShowFSScreenInfo(info global.ScreenShowData) {
	global.Logger.Debug("显示放射科屏幕信息")
	// 放射科数据需要保存
	if global.ScreenRoomTotalData[info.CurShowIP] != nil {
		// 存在显示屏
		// 1. 屏幕的已经存在的机房位置+1，新呼叫的机房位置在最上面
		for key, v := range global.ScreenRoomTotalData[info.CurShowIP] {
			v.CurrIndex++
			global.ScreenRoomTotalData[info.CurShowIP][key] = v
		}
		global.ScreenRoomTotalData[info.CurShowIP][info.RoomInfo.CurrRoom] = info.RoomInfo
	} else {
		// 不存在显示屏
		global.ScreenRoomTotalData[info.CurShowIP] = make(map[string]global.CallData)
		global.ScreenRoomTotalData[info.CurShowIP][info.RoomInfo.CurrRoom] = info.RoomInfo
	}
	// 获取所有机房的数据
	var roomlist []global.CallData
	for _, v := range global.ScreenRoomTotalData[info.CurShowIP] {
		roomlist = append(roomlist, v)
	}
	// 通过机房的位置排序
	sort.Slice(roomlist, func(i, j int) bool {
		return roomlist[i].CurrIndex < roomlist[j].CurrIndex
	})

	data := global.FSWSData{
		CallPatient:     info.RoomInfo.CurPatientName,
		CallRoom:        info.RoomInfo.CurrRoom,
		CallQueueNum:    info.RoomInfo.CurQueueNumber,
		CallPatientType: info.RoomInfo.CurTypeName,
		CallWavFile:     info.CurWavFile,
		CheckRoomList:   roomlist,
	}
	// 发送数据给websocket
	msg := global.WSMessage{
		To:       info.CurShowIP,
		MsgType:  2,
		DataType: global.Screen_Type_FS,
		Data:     data,
	}
	ws.SendMsg(msg)
}

// 显示超声科屏幕信息
func ShowUSScreenInfo(info global.ScreenShowData) {
	global.Logger.Debug("显示超声科屏幕信息")
	// 超声科数据需要保存
	var roomlist []global.CallData
	// 1. 屏幕的已经存在的机房位置+1，新呼叫的机房位置在最上面
	for _, v := range global.USScreenRoomTotalData[info.CurShowIP] {
		if v.CurCheckNumber != info.RoomInfo.CurCheckNumber || v.CurPatientName != info.RoomInfo.CurPatientName || v.CurQueueNumber != info.RoomInfo.CurQueueNumber {
			v.CurrIndex++
			roomlist = append(roomlist, v)
		}
	}
	roomlist = append(roomlist, info.RoomInfo)

	// 通过机房的位置排序
	sort.Slice(roomlist, func(i, j int) bool {
		return roomlist[i].CurrIndex < roomlist[j].CurrIndex
	})

	if len(roomlist) > info.CurShowSize {
		roomlist = roomlist[:info.CurShowSize]
	}

	global.USScreenRoomTotalData[info.CurShowIP] = roomlist

	data := global.FSWSData{
		CallPatient:     info.RoomInfo.CurPatientName,
		CallRoom:        info.RoomInfo.CurrRoom,
		CallQueueNum:    info.RoomInfo.CurQueueNumber,
		CallPatientType: info.RoomInfo.CurTypeName,
		CallWavFile:     info.CurWavFile,
		CheckRoomList:   roomlist,
	}
	// 发送数据给websocket
	msg := global.WSMessage{
		To:       info.CurShowIP,
		MsgType:  2,
		DataType: global.Screen_Type_US,
		Data:     data,
	}
	ws.SendMsg(msg)
}

// 显示内镜科屏幕信息
func ShowESScreenInfo(info global.ScreenShowData) {
	global.Logger.Debug("显示内镜科屏幕信息")
	// 内镜科数据需要保存
	var roomlist []global.CallData
	// 1. 屏幕的已经存在的机房位置+1，新呼叫的机房位置在最上面
	for _, v := range global.USScreenRoomTotalData[info.CurShowIP] {
		if v.CurCheckNumber != info.RoomInfo.CurCheckNumber || v.CurPatientName != info.RoomInfo.CurPatientName || v.CurQueueNumber != info.RoomInfo.CurQueueNumber {
			v.CurrIndex++
			roomlist = append(roomlist, v)
		}
	}
	roomlist = append(roomlist, info.RoomInfo)

	// 通过机房的位置排序
	sort.Slice(roomlist, func(i, j int) bool {
		return roomlist[i].CurrIndex < roomlist[j].CurrIndex
	})

	if len(roomlist) > info.CurShowSize {
		roomlist = roomlist[:info.CurShowSize]
	}

	global.USScreenRoomTotalData[info.CurShowIP] = roomlist

	data := global.FSWSData{
		CallPatient:     info.RoomInfo.CurPatientName,
		CallRoom:        info.RoomInfo.CurrRoom,
		CallQueueNum:    info.RoomInfo.CurQueueNumber,
		CallPatientType: info.RoomInfo.CurTypeName,
		CallWavFile:     info.CurWavFile,
		CheckRoomList:   roomlist,
	}
	// 发送数据给websocket
	msg := global.WSMessage{
		To:       info.CurShowIP,
		MsgType:  2,
		DataType: global.Screen_Type_ES,
		Data:     data,
	}
	ws.SendMsg(msg)
}

// 显示门诊屏幕信息
func ShowMZScreenInfo(info global.ScreenShowData) {
	global.Logger.Debug("显示门诊屏幕信息")
	if global.ScreenRoomTotalData[info.CurShowIP] != nil {
		// 存在显示屏
		global.ScreenRoomTotalData[info.CurShowIP][info.RoomInfo.CurrRoom] = info.RoomInfo
	} else {
		// 不存在显示屏
		global.ScreenRoomTotalData[info.CurShowIP] = make(map[string]global.CallData)
		global.ScreenRoomTotalData[info.CurShowIP][info.RoomInfo.CurrRoom] = info.RoomInfo
	}

	// 发送数据给websocket
	msg := global.WSMZMessage{
		To:       info.CurShowIP,
		MsgType:  2,
		DataType: global.Screen_Type_MZ,
		Data:     info.RoomInfo,
	}
	ws.SendMZMsg(msg)
}
