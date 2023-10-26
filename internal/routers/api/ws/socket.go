package ws

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/model"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 客户端 Client
type Client struct {
	Conn *websocket.Conn
	IP   string
}

type Message struct {
	To       string               `json:"to"`       // 接收者IP
	MsgType  int                  `json:"msgType"`  // 数据类型 0:ws通讯信息，1:屏幕配置信息，2:叫号数据消息
	DataType int                  `json:"dataType"` // 数据类型 0：放射科，1：超声科，2：内镜科，3：门诊
	Data     global.Screen_Config `json:"data"`     // 配置信息
}

var Clients = make(map[*websocket.Conn]*Client)

// 用来升级HTTP请求
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, //指定读缓存大小
	WriteBufferSize: 1024, //指定写缓存大小
	CheckOrigin:     checkOrigin,
}

// 检测请求来源  该函数用于拦截或放行跨域请求
func checkOrigin(r *http.Request) bool {
	return true
}

// 定义客户端结构体的read方法
func (c *Client) ReadMsg() {
	for {
		//读取消息
		_, _, err := c.Conn.ReadMessage()
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				// 连接正常关闭或正在关闭
				global.Logger.Error("连接关闭:", err, "IP: ", c.IP)
			} else {
				// 连接异常关闭
				global.Logger.Error("连接异常关闭:", err, "IP: ", c.IP)
			}
			c.Conn.Close()
			delete(Clients, c.Conn)
			break
		}
	}
}

func HandleWebSocket(c *gin.Context) {
	// 升级HTTP连接为WebSocket连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("WebSocket Upgrade err ", err.Error())
		return
	}
	clientip := c.ClientIP()
	global.Logger.Debug("WebSocket连接的 IP :", clientip)
	client := &Client{
		Conn: conn,
		IP:   clientip,
	}
	Clients[conn] = client
	// 返回配置信息
	data := model.GetScreenConfig(clientip)
	msg := Message{
		To:       clientip,
		MsgType:  1,
		DataType: 0,
		Data:     data,
	}
	conn.WriteJSON(msg)
	// 连接时，sleep(5s) 然后发送数据
	time.Sleep(time.Second * 5)
	switch data.Department_Code {
	case global.Screen_Type_FS:
		var roomlist []global.CallData
		for _, v := range global.ScreenRoomTotalData[clientip] {
			roomlist = append(roomlist, v)
		}
		// 通过机房的位置排序
		sort.Slice(roomlist, func(i, j int) bool {
			return roomlist[i].CurrIndex < roomlist[j].CurrIndex
		})
		initdata := global.FSWSData{
			CallPatient:     "",
			CallRoom:        "",
			CallQueueNum:    "",
			CallPatientType: "",
			CallWavFile:     "",
			CheckRoomList:   roomlist,
		}
		initmsg := global.WSMessage{
			To:       clientip,
			MsgType:  3,
			DataType: global.Screen_Type_FS,
			Data:     initdata,
		}
		SendMsg(initmsg)
	case global.Screen_Type_US:
	case global.Screen_Type_ES:
	case global.Screen_Type_MZ:
	default:
	}
	//启动协程收web端传过来的消息
	go client.ReadMsg()
}

// 发送消息
func SendMsg(msg global.WSMessage) {
	for _, value := range Clients {
		if value.IP == msg.To {
			err := value.Conn.WriteJSON(msg)
			if err != nil {
				global.Logger.Error("websocket send msg err :", err)
				break
			}
			global.Logger.Debug("websocket send msg Successful: ", msg)
			break
		}
	}
}
