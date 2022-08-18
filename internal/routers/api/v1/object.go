package v1

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/model"
	"WowjoyProject/WowjoyQueueCall/pkg/object"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// 呼叫返回文件
func CallFile(c *gin.Context) {
	check_number := c.PostForm("check_number")
	call_point := c.PostForm("call_point")
	callpoint, _ := strconv.Atoi(call_point)
	global.Logger.Info("check_number: ", check_number, " ,call_point: ", callpoint)
	// 获取数据库数据
	patientInfo, err := model.QueryPatientInfo(check_number)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "患者信息没有查询到",
			"data":    patientInfo,
		})
		return
	}
	callPointInfo, err := model.QueryCallPointConfig(callpoint)
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

	call_text := ParsCallPointInfo(callTextList, *patientInfo)
	// 2.生成wav 文件
	if len(call_text) < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "站点呼叫内容配置为空，请配置呼叫内容",
			"data":    callPointInfo.Call_Point,
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
	fileName += patientInfo.Check_Number
	fileName += ".wav"
	object.CallExeSaveWavFile(str, fileName)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "呼叫成功",
		"data":    fileName,
	})
	return
}

// 呼叫返回数据流
func CallStream(c *gin.Context) {
	check_number := c.PostForm("check_number")
	call_point := c.PostForm("call_point")
	callpoint, _ := strconv.Atoi(call_point)
	global.Logger.Info("check_number: ", check_number, " ,call_point: ", callpoint)
	// 获取数据库数据
	patientInfo, err := model.QueryPatientInfo(check_number)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "患者信息没有查询到",
			"data":    patientInfo,
		})
		return
	}
	callPointInfo, err := model.QueryCallPointConfig(callpoint)
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

	call_text := ParsCallPointInfo(callTextList, *patientInfo)
	// 2.生成wav 文件
	if len(call_text) < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "站点呼叫内容配置为空，请配置呼叫内容",
			"data":    callPointInfo.Call_Point,
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
	fileName += patientInfo.Check_Number
	fileName += ".wav"
	object.CallExeSaveWavFile(str, fileName)
	c.File(fileName)
	return
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
	return
}

//获取结构体中字段的名称
func GetFieldName(columnName string, object global.Patient_Info) string {
	var val string
	t := reflect.TypeOf(object)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
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
func ParsCallPointInfo(callTextList []string, object global.Patient_Info) string {
	var textcfg global.TextCfg
	var textcfgData global.TextCfgData
	var textcfgDataList []global.TextCfgData

	var TextStr string

	for i := 0; i < len(callTextList); i++ {
		switch callTextList[i] {
		case "patient_name":
			textcfg.PateintNameCount++
		case "machine_room":
			textcfg.MachineRoomCount++
		case "queue_number":
			textcfg.QueueNumberCount++
		case "type_name":
			textcfg.TypeNameCount++
		case "check_type":
			textcfg.CheckTypeCount++
		case "wait_patient":
			textcfg.PateintNameCount++
		}
	}
	if textcfg.PateintNameCount > 1 {
		textcfgData.Patient_Name = object.Patient_Name
		textcfgData.Machine_Room = object.Machine_Room
		textcfgData.Queue_Number = object.Queue_Number
		textcfgData.Type_Name = object.Type_Name
		textcfgData.Check_Type = object.Check_Type
		textcfgDataList = append(textcfgDataList, textcfgData)
		dataList := model.GetCallPointTextData(object, textcfg.PateintNameCount-1)
		textcfgDataList = append(textcfgDataList, dataList...)
		ListLen := len(textcfgDataList)
		var temp global.TextCfg
		for i := 0; i < len(callTextList); i++ {
			switch callTextList[i] {
			case "patient_name":
				if temp.PateintNameCount < ListLen {
					TextStr += textcfgDataList[temp.PateintNameCount].Patient_Name
					temp.PateintNameCount++
				}
			case "machine_room":
				if temp.MachineRoomCount < ListLen {
					TextStr += textcfgDataList[temp.MachineRoomCount].Machine_Room
					temp.MachineRoomCount++
				}
			case "queue_number":
				if temp.QueueNumberCount < ListLen {
					TextStr += textcfgDataList[temp.QueueNumberCount].Queue_Number
					temp.QueueNumberCount++
				}
			case "type_name":
				if temp.TypeNameCount < ListLen {
					TextStr += textcfgDataList[temp.TypeNameCount].Type_Name
					temp.TypeNameCount++
				}
			case "check_type":
				if temp.CheckTypeCount < ListLen {
					TextStr += textcfgDataList[temp.CheckTypeCount].Check_Type
					temp.CheckTypeCount++
				}
			case "wait_patient":
				if temp.PateintNameCount < ListLen {
					TextStr += textcfgDataList[temp.PateintNameCount].Patient_Name
					temp.PateintNameCount++
				}
			default:
				TextStr += callTextList[i]
			}
		}
	} else {
		for i := 0; i < len(callTextList); i++ {
			switch callTextList[i] {
			case "patient_name":
				TextStr += object.Patient_Name
			case "machine_room":
				TextStr += object.Machine_Room
			case "queue_number":
				TextStr += object.Queue_Number
			case "type_name":
				TextStr += object.Type_Name
			case "check_type":
				TextStr += object.Check_Type
			default:
				TextStr += callTextList[i]
			}
		}
	}
	return TextStr
}

// 站点字符串解析
func ParsStrToList(callText string) (textList []string) {
	if len(callText) < 0 {
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
