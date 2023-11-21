package object

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/routers/api/ws"
	"os/exec"
	"sort"
)

// 调用exe生成wav语音文件
func CallExeSaveWavFile(str, fileName string) bool {
	exepath := global.ObjectSetting.TTSPath

	arg := make([]string, 0)
	arg = append(arg, str)
	arg = append(arg, fileName)
	global.Logger.Info(exepath, " 参数是：", arg)

	cmd := exec.Command(exepath, arg...)
	if err := cmd.Run(); err != nil {
		global.Logger.Error("生成Wav文件失败")
		return false
	}
	global.Logger.Info("Wav语音文件生成成功")
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
