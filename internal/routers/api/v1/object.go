package v1

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/model"
	"WowjoyProject/WowjoyQueueCall/pkg/object"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 呼叫返回文件
func CallFile(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var calldata global.CallData
	c.ShouldBind(&calldata)
	global.Logger.Debug(calldata)
	// 查询科室配置信息
	callPointInfo, err := model.QueryCheckRoomConfig(calldata.CurrRoom)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "呼叫站点信息不存在",
			"data":    callPointInfo,
		})
		return
	}
	// 解析站点呼叫内容信息
	callTextList := ParsStrToList(callPointInfo.Call_Text)

	call_text := ParsCallPointInfo(callTextList, calldata)
	var baseWav, fileName string
	// 2.生成wav 文件
	if len(call_text) > 0 {
		global.Logger.Debug("开始生成wav 文件")
		var str string
		// 判定呼叫次数
		for i := 0; i < callPointInfo.Call_Number; i++ {
			str += call_text
			str += "   "
		}
		fileName = global.ObjectSetting.WAVFilePath
		fileName += "\\"
		fileName += calldata.CurCheckNumber
		fileName += ".wav"
		object.CallExeSaveWavFile(str, fileName)
		baseWav = global.ObjectSetting.WAVURL
		baseWav += calldata.CurCheckNumber
		baseWav += ".wav"
		// 语音文件转Base64编码
		// baseWav = general.File2Base64(fileName)
		// baseWav = "data: audio/wav: base64," + baseWav
	}
	// 1.机房可以分配多个呼叫点
	// 分发消息到显示端
	// 查询机房配置的呼叫点（多呼叫点通过|*|分割符分隔）
	callpoints := strings.Split(callPointInfo.Call_Point, "|*|")
	// var showinfolist []global.ScreenShowData
	for _, value := range callpoints {
		callpoint, _ := strconv.Atoi(value)
		screenInfo := model.GetScreenConfigByCallPoint(callpoint)
		// 通过屏幕显示类容分发消息
		showinfo := global.ScreenShowData{
			CurShowCallPoint:      callpoint,
			CurShowIP:             screenInfo.IP,
			CurShowDepartmentCode: screenInfo.Department_Code,
			CurWavFile:            baseWav,
			RoomInfo:              calldata,
		}
		go object.SendShowScreenInfo(showinfo)
		// showinfolist = append(showinfolist, showinfo)
	}
	// go object.SendShowScreenInfo(showinfolist)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "呼叫成功",
		"data":    fileName,
	})
}

// 呼叫返回数据流
func CallStream(c *gin.Context) {
	reqIP := c.ClientIP()
	global.Logger.Debug("请求的主机IP: ", reqIP)
	var calldata global.CallData
	c.ShouldBind(&calldata)
	global.Logger.Debug(calldata)
	callPointInfo, err := model.QueryCheckRoomConfig(calldata.CurrRoom)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "呼叫站点信息不存在",
			"data":    callPointInfo,
		})
		return
	}

	// 解析站点呼叫内容信息
	callTextList := ParsStrToList(callPointInfo.Call_Text)

	call_text := ParsCallPointInfo(callTextList, calldata)
	// 2.生成wav 文件
	if len(call_text) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "站点呼叫内容配置为空，请配置呼叫内容",
			"data":    callPointInfo.Check_Room,
		})
		return
	}

	var str string
	// 判定呼叫次数
	for i := 0; i < callPointInfo.Call_Number; i++ {
		str += call_text
		str += "   "
	}
	fileName := global.ObjectSetting.WAVFilePath
	fileName += "\\"
	fileName += calldata.CurCheckNumber
	fileName += ".wav"
	object.CallExeSaveWavFile(str, fileName)
	c.File(fileName)
}

// 插入患者数据
func InsPatientData(c *gin.Context) {
	var data global.PatientData
	// 绑定数据
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	// 插入数据库
	patientdata := global.Patient_Info{
		Check_Number:     data.Check_Number,
		Patient_Name:     data.Patient_Name,
		Patient_Sex:      data.Patient_Sex,
		Patient_Age:      data.Patient_Age,
		Patient_Birthday: data.Patient_Birthday,
		Queue_Number:     data.Queue_Number,
		Machine_Room:     data.Machine_Room,
		Type_Name:        data.Type_Name,
		Call_Status_Code: data.Call_Status_Code,
		Check_Items:      data.Check_Items,
		Check_Type:       data.Check_Type,
		Check_Body:       data.Check_Body,
		Report_Status:    data.Report_Status,
		Sign_Time:        data.Sign_Time,
	}
	model.InsertPatientInfo(&patientdata)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "患者数据插入成功",
		"data":    "",
	})
}

// // 获取结构体中字段的名称
func GetFieldName(columnName string, object global.Patient_Info) string {
	var val string
	t := reflect.TypeOf(object)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		global.Logger.Debug("Check type error not Struct")
		return ""
	}
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if strings.ToUpper(t.Field(i).Name) == strings.ToUpper(columnName) {
			v := reflect.ValueOf(object)
			val := v.FieldByName(t.Field(i).Name).String()
			return val
		}
	}
	return val
}

// 解析站点信息
func ParsCallPointInfo(callTextList []string, object global.CallData) string {
	var waitename []string
	for _, v := range object.WaitePatienList {
		waitename = append(waitename, v.WaitePatientName)
	}
	var TextStr string
	var waitenum int
	for i := 0; i < len(callTextList); i++ {
		switch callTextList[i] {
		case "patient_name":
			TextStr += object.CurPatientName
		case "machine_room":
			TextStr += object.CurrRoom
		case "queue_number":
			TextStr += object.CurQueueNumber
		case "type_name":
			TextStr += object.CurTypeName
		case "check_type":
			TextStr += object.CurCheckType
		case "wait_patient":
			if waitenum < len(waitename) {
				TextStr += waitename[waitenum]
				waitenum++
			}
		default:
			TextStr += callTextList[i]
		}
	}
	return TextStr
}

// 站点字符串解析
func ParsStrToList(callText string) (textList []string) {
	if len(callText) <= 0 {
		global.Logger.Error("呼叫内容长度小于0：", callText)
		return textList
	}
	for {
		start := strings.Index(callText, "<")
		if start > 0 {
			textList = append(textList, callText[:start])
		} else if start < 0 {
			textList = append(textList, callText)
			break
		}
		end := strings.Index(callText, ">")
		if end >= 0 {
			textList = append(textList, callText[start+1:end])
		}
		callText = callText[end+1:]
	}
	return textList
}

// 手动获取患者数据
func HandGetPatientData(c *gin.Context) {
	model.GetPatientData()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "手动获取患者数据成功",
		"data":    "",
	})
}
