# 项目
# ****叫号系统服务端****

# 项目描述
叫号系统服务端：
1. Web服务端
2. 数据库操作
3. 生成语音wav文件
4. 语音文件生成规则


# 设计依据

# 目录结构
configs：配置文件。
global：全局变量。
internal：内部模块。
model：数据库相关操作。
pkg：项目相关的模块包。
storage：项目生成的临时文件。

# 公共组件
配置管理
数据库连接
日志写入

# 文件配置文件读取：go get -u github.com/spf13/viper
Viper 是适用于GO 应用程序的完整配置解决方案

# 日志：go get -u gopkg.in/natefinch/lumberjack.v2
它的核心功能是将日志写入滚动文件中，该库支持设置所允许单日志文件的最大占用空间、最大生存周期、允许保留的最多旧文件数，
如果出现超出设置项的情况，就会对日志文件进行滚动处理。

# 生成接口文档
Swagger 相关的工具集会根据 OpenAPI 规范去生成各式各类的与接口相关联的内容，
常见的流程是编写注解 =》调用生成库-》生成标准描述文件 =》生成/导入到对应的 Swagger 工具
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template

@Summary	摘要
@Produce	API 可以产生的 MIME 类型的列表，MIME 类型你可以简单的理解为响应类型，例如：json、xml、html 等等
@Param	参数格式，从左到右分别为：参数名、入参类型、数据类型、是否必填、注释
@Success	响应成功，从左到右分别为：状态码、参数类型、数据类型、注释
@Failure	响应失败，从左到右分别为：状态码、参数类型、数据类型、注释
@Router	路由，从左到右分别为：路由地址，HTTP 方法

swag init

http://127.0.0.1:8000/swagger/index.html

# 国际化处理
中间件
# 邮件报警处理
go get -u gopkg.in/gomail.v2
# 接口限流控制
go get -u github.com/juju/ratelimit@v1.0.1
# 统一超时控制

# 构建命令
go build -ldflags="-H windowsgui" -o .\WowjoyQueueCall\WowjoyQueueCall.exe .\main.go .\setup.go

## 第二次相同项目提交文件到github
# git add README.md
# git commit -m "first commit"
# git push -u origin master

# 提交本地分支代码到远程分支
# git add .
# git commit -m ""
# git push origin V1.0.0.2


# 修改记录
# 2023/11/23 修改本地分支名，从V1.0.3 --> V1.1.0
* 1. 屏幕配置中,增加新字段，给前端使用，区分屏幕显示类型（科室单屏显示）
* 2. 患者信息表中增加绿色通道字段
* 3. 增加语音文件未生成时，重复生成语音文件
* 4. 修改语音播报时，最后等待患者播报
* 5. 增加就诊类型和绿色通道排序字段
* 6. 排队号中,号码前面增加"急"字区别
* 7. 获取患者信息，缺失电子申请单和临床信息
# 2023/11/21 创建新分支V1.0.3 
* 1. 修改患者数据获取,增加事件排序，解决重复数据插入报错
* 2. 增加前端log 接口
* 3. 增加访问数据库时，DB有效性验证
# 2023/11/08 增加超声内镜和门诊屏幕数据的分类
# 2023/11/07 创建V1.0.0.2版本
# 2023/11/07 语音文件通过http 获取
# 2023/10/26 V1.0.0.1 增加前端屏幕显示端相关代码
# 2022-04-14 开始创建项目
