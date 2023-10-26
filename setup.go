package main

import (
	"WowjoyProject/WowjoyQueueCall/global"
	"WowjoyProject/WowjoyQueueCall/internal/model"
	"WowjoyProject/WowjoyQueueCall/pkg/logger"
	"WowjoyProject/WowjoyQueueCall/pkg/setting"
	"io"
	"log"
	"os"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	readSetup()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("General", &global.GeneralSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Object", &global.ObjectSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	lunberLogger := &lumberjack.Logger{
		Filename:  global.GeneralSetting.LogSavePath + "/" + global.GeneralSetting.LogFileName + global.GeneralSetting.LogFileExt,
		MaxSize:   global.GeneralSetting.LogMaxSize,
		MaxAge:    global.GeneralSetting.LogMaxAge,
		LocalTime: true,
	}
	global.Logger = logger.NewLogger(io.MultiWriter(lunberLogger, os.Stdout), "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupReadDBEngine() error {
	var err error
	global.QueueCAllDBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 增加其它系统的数据库连接
func setupWriteDBEngine() error {
	var err error
	global.PACSDBEngine, err = model.NewOtherDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func readSetup() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupReadDBEngine()
	if err != nil {
		log.Fatalf("init.setupReadDBEngine err: %v", err)
	}
	err = setupWriteDBEngine()
	if err != nil {
		log.Fatalf("init.setupWriteDBEngine err: %v", err)
	}
}
