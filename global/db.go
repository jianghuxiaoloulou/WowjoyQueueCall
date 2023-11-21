package global

import (
	"database/sql"
)

var (
	QueueCAllDBEngine *sql.DB
	PACSDBEngine      *sql.DB
)

// 屏幕配置
type Screen_Config struct {
	Call_Point      int    `db:"callPoint"`      // 呼叫点 （唯一值）
	IP              string `db:"ip"`             // 屏幕IP
	Name            string `db:"name"`           // 屏幕名字
	Title           string `db:"title"`          // 屏幕标题
	Note            string `db:"note"`           // 屏幕注意事项
	Department      string `db:"department"`     // 屏幕现实的科室
	Department_Code int    `db:"departmentCode"` // 屏幕显示的科室类型 0：放射科，1：超声科，2：内镜科，3：门诊
	Show_Status     int    `db:"showStatus"`     // 屏幕姓名显示状态 0：加密，1：不加密
	Show_Size       int    `db:"showSize"`       // 屏幕显示的数据size(超声内镜指定屏幕显示的数据条目)
	Webconfig       string `db:"webConfig"`      // 前端的配置字符串
}

type Call_Point_Config struct {
	Call_Point  string `db:"call_point"`  // 呼叫点多个值通过|*|分割
	Check_Type  string `db:"check_type"`  // 检查类型
	Check_Room  string `db:"check_room"`  // 检查机房
	Call_Number int    `db:"call_number"` // 呼叫次数
	Call_Text   string `db:"call_text"`   // 呼叫内容
	Name_Status int    `db:"name_status"` // 姓名加密状态
	CreatedTime string `db:"create_time"` // 创建时间
	UpdatedTime string `db:"update_time"` // 更新时间
}

type Patient_Info struct {
	Call_Status_Name string `db:"call_status_name"` // 呼叫状态名
	Check_Number     string `db:"check_number"`     // 检查号
	Patient_Name     string `db:"patient_name"`     // 患者姓名
	Brdah            string `db:"brdah"`            // 病历号
	Patient_Sex      string `db:"patient_sex"`      // 患者性别
	Patient_Age      string `db:"patient_age"`      // 患者年龄
	Patient_Birthday string `db:"patient_birthday"` // 患者出生日期
	Queue_Number     string `db:"queue_number"`     // 排队序号
	Machine_Room     string `db:"machine_room"`     // 机房
	Type_Name        string `db:"type_name"`        // 就诊类型
	Call_Status_Code int    `db:"call_status_code"` // 呼叫状态code
	Check_Items      string `db:"check_items"`      // 检查项目
	Check_Type       string `db:"check_type"`       // 检查类型
	Check_Body       string `db:"check_body"`       // 检查部位
	Report_Status    string `db:"report_status"`    // 报告状态int
	Sign_Time        string `db:"sign_time"`        // 报到时间
	Call_Time        string `db:"call_time"`        // 呼叫时间
	Call_Number      int    `db:"call_number"`      // 呼叫次数
	CreatedTime      string `db:"create_time"`      // 创建时间
	UpdatedTime      string `db:"update_time"`      // 更新时间
}

type PatientInfoData struct {
	Call_Status_Name sql.NullString // 呼叫状态名
	Check_Number     sql.NullString // 检查号
	Patient_Name     sql.NullString // 患者姓名
	Brdah            sql.NullString // 病历号
	Patient_Sex      sql.NullString // 患者性别
	Patient_Age      sql.NullString // 患者年龄
	Patient_Birthday sql.NullString // 患者出生日期
	Queue_Number     sql.NullString // 排队序号
	Machine_Room     sql.NullString // 机房
	Type_Name        sql.NullString // 就诊类型
	Call_Status_Code sql.NullInt64  // 呼叫状态code
	Check_Items      sql.NullString // 检查项目
	Check_Type       sql.NullString // 检查类型
	Check_Body       sql.NullString // 检查部位
	Report_Status    sql.NullString // 报告状态int
	Sign_Time        sql.NullString // 报到时间
	Call_Time        sql.NullString // 呼叫时间
	Call_Number      sql.NullInt64  // 呼叫次数
	CreatedTime      sql.NullString // 创建时间
	UpdatedTime      sql.NullString // 更新时间
}

type KeyData struct {
	Call_Status_Name sql.NullString // 呼叫状态名
	Check_Number     sql.NullString // 检查号
	Patient_Name     sql.NullString // 患者姓名
	Brdah            sql.NullString // 病历号
	Patient_Sex      sql.NullString // 患者性别
	Patient_Age      sql.NullString // 患者年龄
	Patient_Birthday sql.NullString // 患者出生日期
	Queue_Number     sql.NullString // 排队序号
	Machine_Room     sql.NullString // 机房
	Type_Name        sql.NullString // 就诊类型
	Call_Status_Code sql.NullInt64  // 呼叫状态code
	Check_Items      sql.NullString // 检查项目
	Check_Type       sql.NullString // 检查类型
	Check_Body       sql.NullString // 检查部位
	Report_Status    sql.NullString // 报告状态
	Sign_Time        sql.NullString // 报到时间
	Call_Time        sql.NullString // 呼叫时间
	Call_Number      sql.NullInt64  // 呼叫次数
}

type PatientData struct {
	Call_Status_Name string `json:"call_status_name"` // 呼叫状态名
	Check_Number     string `json:"check_number"`     // 检查号
	Patient_Name     string `json:"patient_name"`     // 患者姓名
	Brdah            string `json:"brdah"`            // 病历号
	Patient_Sex      string `json:"patient_sex"`      // 患者性别
	Patient_Age      string `json:"patient_age"`      // 患者年龄
	Patient_Birthday string `json:"patient_birthday"` // 患者出生日期
	Queue_Number     string `json:"queue_number"`     // 排队序号
	Machine_Room     string `json:"machine_room"`     // 机房
	Type_Name        string `json:"type_name"`        // 就诊类型
	Call_Status_Code int    `json:"call_status_code"` // 呼叫状态code
	Check_Items      string `json:"check_items"`      // 检查项目
	Check_Type       string `json:"check_type"`       // 检查类型
	Check_Body       string `json:"check_body"`       // 检查部位
	Report_Status    string `json:"report_status"`    // 报告状态
	Sign_Time        string `json:"sign_time"`        // 报到时间
	Call_Time        string `json:"call_Time"`        // 呼叫时间
	Call_Number      string `json:"call_number"`      // 呼叫次数
}

type TextCfgData struct {
	Patient_Name string `db:"patient_name"` // 患者姓名
	Queue_Number string `db:"queue_number"` // 排队序号
	Machine_Room string `db:"machine_room"` // 机房
	Type_Name    string `db:"type_name"`    // 就诊类型
	Check_Type   string `db:"check_type"`   // 检查类型
}

func NullStringToString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return ""
}
