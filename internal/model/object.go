package model

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"fmt"
	"strconv"
	"strings"
)

// 插入患者信息表数据（patient_info）
func InsertPatientInfo(data *global.Patient_Info) (int64, error) {
	global.Logger.Info("***插入patient_info 数据***")
	sql := `REPLACE patient_info (call_status_name,call_status_code,check_number,queue_number,
		patient_name,brdah,patient_sex,patient_age,type_name,type_name_index,green_flag,green_flag_code,
		check_body,check_items,sign_time,call_time,patient_birthday,machine_room,call_number,check_type,
		report_status,apply_department_name,apply_doctor_name,telephone,id_card,patient_number,his_sn,
		sickbed_index,society_number,clinical_manifestation,clinical_diagnosis,present_illness_history,
		note)
	VALUES (?,?,?,?,
		?,?,?,?,?,?,?,?,
		?,?,?,?,?,?,?,?,
		?,?,?,?,?,?,?,
		?,?,?,?,?,
		?);`
	res, err := global.QueueCAllDBEngine.Exec(sql, data.Call_Status_Name, data.Call_Status_Code, data.Check_Number, data.Queue_Number,
		data.Patient_Name, data.Brdah, data.Patient_Sex, data.Patient_Age, data.Type_Name, data.Type_Name_Index, data.GreenFlag, data.GreenFlagCode,
		data.Check_Body, data.Check_Items, data.Sign_Time, data.Call_Time, data.Patient_Birthday, data.Machine_Room, data.Call_Number, data.Check_Type,
		data.Report_Status, data.Apply_Department_name, data.Apply_doctor_name, data.Telephone, data.Id_Card, data.Patient_number, data.His_Sn,
		data.Sickbed_Index, data.Society_number, data.Clinical_Manifestation, data.Clinical_Diagnosis, data.Present_Illness_History,
		data.Note)
	if err != nil {
		global.Logger.Error("patient_info inser err: ", err)
		return 0, err
	}
	global.Logger.Debug("插入成功的数据： ", data.Check_Number)
	global.ArriveTime = data.Sign_Time
	return res.LastInsertId()
}

// 删除患者信息表中数据（patient_info）
func DeletePatientInfo(day int) error {
	global.Logger.Info("***删除patient_info 数据***")
	sql := `DELETE FROM patient_info pai WHERE (TIMESTAMPDIFF(DAY,pai.update_time,NOW())) < ?;`
	_, err := global.QueueCAllDBEngine.Exec(sql, day)
	if err != nil {
		return err
	}
	return err
}

// 获取患者信息表中数据（patient_info）
func QueryPatientInfo(checkNumber string) (*global.Patient_Info, error) {
	global.Logger.Info("***查询patient_info 数据***: ", checkNumber)
	sql := `SELECT pai.call_status_name,pai.queue_number,pai.patient_name,pai.brdah,
	pai.patient_sex,pai.patient_age,pai.type_name,pai.check_body,green_flag,pai.check_items,pai.check_number,
	pai.sign_time,pai.call_time,pai.patient_birthday,pai.machine_room,pai.call_status_code,
	pai.call_number,pai.check_type,pai.report_status
	FROM patient_info pai WHERE pai.check_number =?;`
	row := global.QueueCAllDBEngine.QueryRow(sql, checkNumber)
	patientinfo := global.PatientInfoData{}
	err := row.Scan(&patientinfo.Call_Status_Name, &patientinfo.Queue_Number, &patientinfo.Patient_Name, &patientinfo.Brdah,
		&patientinfo.Patient_Sex, &patientinfo.Patient_Age, &patientinfo.Type_Name, &patientinfo.Check_Body, &patientinfo.GreenFlag, &patientinfo.Check_Items, &patientinfo.Check_Number,
		&patientinfo.Sign_Time, &patientinfo.Call_Time, &patientinfo.Patient_Birthday, &patientinfo.Machine_Room, &patientinfo.Call_Status_Code,
		&patientinfo.Call_Number, &patientinfo.Check_Type, &patientinfo.Report_Status)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	data := global.Patient_Info{
		Call_Status_Name: patientinfo.Call_Status_Name.String,
		Check_Number:     patientinfo.Check_Number.String,
		Patient_Name:     patientinfo.Patient_Name.String,
		Brdah:            patientinfo.Brdah.String,
		Patient_Sex:      patientinfo.Patient_Sex.String,
		Patient_Age:      patientinfo.Patient_Age.String,
		Patient_Birthday: patientinfo.Patient_Birthday.String,
		Queue_Number:     patientinfo.Queue_Number.String,
		Machine_Room:     patientinfo.Machine_Room.String,
		Type_Name:        patientinfo.Type_Name.String,
		Call_Status_Code: int(patientinfo.Call_Status_Code.Int64),
		Check_Items:      patientinfo.Check_Items.String,
		Check_Type:       patientinfo.Check_Type.String,
		Check_Body:       patientinfo.Check_Body.String,
		GreenFlag:        patientinfo.GreenFlag.String,
		Report_Status:    patientinfo.Report_Status.String,
		Sign_Time:        patientinfo.Sign_Time.String,
		Call_Time:        patientinfo.Call_Time.String,
		Call_Number:      int(patientinfo.Call_Number.Int64),
	}
	return &data, nil
}

// 插入站点信息表数据（call_point_config）
func InsertCallPointConfig(data *global.Call_Point_Config) (int64, error) {
	sql := `INSERT INTO call_point_config (call_point,check_type,check_room,call_number,call_text,name_status)
	VALUES (?,?,?,?,?,?);`
	res, err := global.QueueCAllDBEngine.Exec(sql, data.Call_Point, data.Check_Type, data.Check_Room,
		data.Call_Number, data.Call_Text, data.Name_Status)
	if err != nil {
		global.Logger.Error("call_point_config inser err: ", err)
		return 0, err
	}
	return res.LastInsertId()
}

// 获取患者信息表中数据（call_point_config）
func QueryCallPointConfig(callPoint int) (*global.Call_Point_Config, error) {
	global.Logger.Info("***查询call_point_config 数据***: ", callPoint)
	sql := `SELECT cp.check_room,cp.call_point,cp.check_type,cp.call_number,cp.call_text,cp.name_status
	FROM call_point_config cp WHERE cp.call_point = ?;`
	row := global.QueueCAllDBEngine.QueryRow(sql, callPoint)
	callPingCfg := global.Call_Point_Config{}
	err := row.Scan(&callPingCfg.Check_Room, &callPingCfg.Call_Point, &callPingCfg.Check_Type,
		&callPingCfg.Call_Number, &callPingCfg.Call_Text, &callPingCfg.Name_Status)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return &callPingCfg, nil
}

// 获取患者信息表中数据通过检查机房（call_point_config）
func QueryCheckRoomConfig(checkroom string) (*global.Call_Point_Config, error) {
	global.Logger.Info("***查询call_point_config 数据***: ", checkroom)
	sql := `SELECT cp.check_room,cp.call_point,cp.check_type,cp.call_number,cp.call_text,cp.name_status
	FROM call_point_config cp WHERE cp.check_room = ?;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}

	row := global.QueueCAllDBEngine.QueryRow(sql, checkroom)
	callPingCfg := global.Call_Point_Config{}
	err := row.Scan(&callPingCfg.Check_Room, &callPingCfg.Call_Point, &callPingCfg.Check_Type,
		&callPingCfg.Call_Number, &callPingCfg.Call_Text, &callPingCfg.Name_Status)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return &callPingCfg, nil
}

// 获取患者数据插入数据库
func GetPatientData() {
	sql := `SELECT roi.accession_number,roi.queue_number,rp.name,rp.brdah,rp.sex_code,roi.age,fi.patient_type_code,
	roi.tsbrpb,roi.bodypart,roi.check_items,roi.arrive_time,rp.birthday,roi.room_name,roi.modality_code,
	fi.report_status,roi.apply_department_name,roi.apply_doctor_name,rp.telephone,rp.id_card,rp.patient_number,
	roi.his_sn,fi.sickbed_index,rp.society_number,fi.clinical_manifestation,fi.clinical_diagnosis,
	fi.present_illness_history,roi.note FROM 
	register_info fi
	LEFT JOIN register_order_info roi ON roi.patient_id = fi.patient_id
	LEFT JOIN register_patient rp ON rp.patient_id = fi.patient_id
	WHERE fi.status = 'Arrived' and roi.arrive_time > ? 
	order by roi.arrive_time asc LIMIT ?;`

	if global.PACSDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.PACSDBEngine.Close()
		global.PACSDBEngine, _ = NewOtherDBEngine(global.DatabaseSetting)
	}
	rows, err := global.PACSDBEngine.Query(sql, global.ArriveTime, global.GeneralSetting.MaxTasks)
	if err != nil {
		global.Logger.Error("Query error: ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		key := global.KeyData{}
		err = rows.Scan(&key.Check_Number, &key.Queue_Number, &key.Patient_Name, &key.Brdah, &key.Patient_Sex, &key.Patient_Age, &key.Type_Name,
			&key.GreenFlag, &key.Check_Body, &key.Check_Items, &key.Sign_Time, &key.Patient_Birthday, &key.Machine_Room, &key.Check_Type,
			&key.Report_Status, &key.Apply_Department_name, &key.Apply_doctor_name, &key.Telephone, &key.Id_Card, &key.Patient_number,
			&key.His_Sn, &key.Sickbed_Index, &key.Society_number, &key.Clinical_Manifestation, &key.Clinical_Diagnosis,
			&key.Present_Illness_History, &key.Note)
		if err != nil {
			global.Logger.Error("rows.Scan error: ", err)
			return
		}
		// 患者性别处理
		patient_sex := TransDict("dict_sex", global.NullStringToString(key.Patient_Sex))
		// 患者就诊来源处理
		type_name := TransDict("dict_patient_type", global.NullStringToString(key.Type_Name))
		type_index := GetPatientIndex(global.NullStringToString(key.Type_Name))

		// 排队号
		queue_number := global.NullStringToString(key.Queue_Number)
		number, _ := strconv.Atoi(queue_number)

		// 就诊类型为急诊时，排队号增加急字拼接例如：急002
		if strings.Contains(type_name, "急诊") && !strings.Contains(queue_number, "急") {
			queue_number = fmt.Sprintf("%s%03d", "急", number)
		} else {
			queue_number = fmt.Sprintf("%03d", number)
		}
		// 绿色通道
		green_code := 0
		if global.NullStringToString(key.GreenFlag) == "是" {
			green_code = 1
		}
		// 报告状态处理
		report_status := TransDict("dict_report_status", global.NullStringToString(key.Report_Status))
		data := global.Patient_Info{
			Call_Status_Name:        "未呼",
			Call_Status_Code:        0,
			Check_Number:            global.NullStringToString(key.Check_Number),
			Queue_Number:            queue_number,
			Patient_Name:            global.NullStringToString(key.Patient_Name),
			Brdah:                   global.NullStringToString(key.Brdah),
			Patient_Sex:             patient_sex,
			Patient_Age:             global.NullStringToString(key.Patient_Age),
			Type_Name:               type_name,
			Type_Name_Index:         type_index,
			GreenFlag:               global.NullStringToString(key.GreenFlag),
			GreenFlagCode:           green_code,
			Check_Body:              global.NullStringToString(key.Check_Body),
			Check_Items:             global.NullStringToString(key.Check_Items),
			Sign_Time:               global.NullStringToString(key.Sign_Time),
			Patient_Birthday:        global.NullStringToString(key.Patient_Birthday),
			Machine_Room:            global.NullStringToString(key.Machine_Room),
			Call_Number:             0,
			Check_Type:              global.NullStringToString(key.Check_Type),
			Report_Status:           report_status,
			Apply_Department_name:   global.NullStringToString(key.Apply_Department_name),
			Apply_doctor_name:       global.NullStringToString(key.Apply_doctor_name),
			Telephone:               global.NullStringToString(key.Telephone),
			Id_Card:                 global.NullStringToString(key.Id_Card),
			Patient_number:          global.NullStringToString(key.Patient_number),
			His_Sn:                  global.NullStringToString(key.His_Sn),
			Sickbed_Index:           global.NullStringToString(key.Sickbed_Index),
			Society_number:          global.NullStringToString(key.Society_number),
			Clinical_Manifestation:  global.NullStringToString(key.Clinical_Manifestation),
			Clinical_Diagnosis:      global.NullStringToString(key.Clinical_Diagnosis),
			Present_Illness_History: global.NullStringToString(key.Present_Illness_History),
			Note:                    global.NullStringToString(key.Note),
		}
		InsertPatientInfo(&data)
	}
}

// 字典值转换
func TransDict(tabName, keyCode string) string {
	if keyCode == "" {
		return ""
	}
	var sql string
	// 表名
	switch tabName {
	case "dict_patient_type":
		sql = `SELECT patient_type_name FROM dict_patient_type WHERE patient_type_code = ?;`
	case "dict_report_status":
		sql = `SELECT report_status_name FROM dict_report_status WHERE report_status_code = ?;`
	case "dict_sex":
		sql = `SELECT sex_name FROM dict_sex WHERE sex_code = ?;`
	}

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	row := global.QueueCAllDBEngine.QueryRow(sql, keyCode)
	var keyName string
	err := row.Scan(&keyName)
	if err != nil {
		global.Logger.Error(err)
		return ""
	}
	return keyName
}

// 获取患者类型编号：(排序使用)
func GetPatientIndex(keyCode string) int {
	var index int
	global.Logger.Info("***获取患者类型编号*****")
	sql := `SELECT patient_type_index FROM dict_patient_type WHERE patient_type_code = ?;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	row := global.QueueCAllDBEngine.QueryRow(sql, keyCode)
	err := row.Scan(&index)
	if err != nil {
		global.Logger.Error(err.Error())
		return 99
	}
	return index
}

// 更新患者呼叫状态
func UpdatePatientCallStatus(checkNumber string, status int) error {
	global.Logger.Info("***更新患者呼叫状态***")
	sql := `UPDATE patient_info pai SET pai.call_status = ? WHERE pai.check_number = ?;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	_, err := global.QueueCAllDBEngine.Exec(sql, status, checkNumber)
	if err != nil {
		return err
	}
	return err
}

// 获取站点呼叫信息中内容
func GetCallPointTextData(object global.Patient_Info, num int) []global.TextCfgData {
	var textCfgList []global.TextCfgData
	sql := `select pai.patient_name,pai.machine_room,pai.queue_number,pai.type_name,pai.check_type from patient_info pai 
	where pai.machine_room = ? and pai.check_number != ? order by ExtractNumber(pai.queue_number) ASC limit ?;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	rows, err := global.QueueCAllDBEngine.Query(sql, object.Machine_Room, object.Check_Number, num)
	if err != nil {
		global.Logger.Fatal("Query error: ", err)
		return textCfgList
	}
	defer rows.Close()
	for rows.Next() {
		key := global.TextCfgData{}
		err = rows.Scan(&key.Patient_Name, &key.Machine_Room, &key.Queue_Number, &key.Type_Name, &key.Check_Type)
		if err != nil {
			global.Logger.Fatal("rows.Scan error: ", err)
			return textCfgList
		}
		global.Logger.Debug("KeyData: ", key)
		textCfgList = append(textCfgList, key)
	}
	return textCfgList
}

// 获取屏幕的配置信息通过IP
func GetScreenConfig(ip string) (screenConfig global.Screen_Config) {
	global.Logger.Info("***查询屏幕的配置信息***: ", ip)
	sql := `select call_point,ip,name,title,note,department,department_code,show_status,show_size,show_type,webconfig from screen_config 
	where ip = ?;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	row := global.QueueCAllDBEngine.QueryRow(sql, ip)
	err := row.Scan(&screenConfig.Call_Point, &screenConfig.IP, &screenConfig.Name, &screenConfig.Title, &screenConfig.Note,
		&screenConfig.Department, &screenConfig.Department_Code, &screenConfig.Show_Status, &screenConfig.Show_Size, &screenConfig.Show_Type, &screenConfig.Webconfig)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	return
}

// 获取屏幕的配置信息通过 呼叫点 (获取显示的屏幕)
func GetScreenConfigByCallPoint(callpoint int) (screenConfig global.Screen_Config) {
	global.Logger.Info("***查询屏幕的配置信息***: ", callpoint)
	sql := `select call_point,ip,name,title,note,department,department_code,show_status,show_size,show_type,webconfig from screen_config 
	where call_point = ?;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	row := global.QueueCAllDBEngine.QueryRow(sql, callpoint)
	err := row.Scan(&screenConfig.Call_Point, &screenConfig.IP, &screenConfig.Name, &screenConfig.Title, &screenConfig.Note,
		&screenConfig.Department, &screenConfig.Department_Code, &screenConfig.Show_Status, &screenConfig.Show_Size, &screenConfig.Show_Type, &screenConfig.Webconfig)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	return
}

// 获取Patient_info表中已经插入数据的签到时间
func GetPatientArriveTime() string {
	var arriveTime string
	global.Logger.Info("***查询最后患者数据的签到时间*****")
	sql := `SELECT sign_time FROM patient_info  ORDER BY sign_time DESC LIMIT 1;`

	if global.QueueCAllDBEngine.Ping() != nil {
		global.Logger.Debug("数据库无效连接，重连数据库")
		global.QueueCAllDBEngine.Close()
		global.QueueCAllDBEngine, _ = NewDBEngine(global.DatabaseSetting)
	}
	row := global.QueueCAllDBEngine.QueryRow(sql)
	err := row.Scan(&arriveTime)
	if err != nil {
		global.Logger.Error(err.Error())
		return arriveTime
	}
	return arriveTime
}
