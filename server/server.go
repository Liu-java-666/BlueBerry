package server

import (
	"BlueBerry/config"
	"BlueBerry/database"
	"BlueBerry/logic"
	"BlueBerry/public"
	"encoding/base64"
	"fmt"
	"github.com/wonderivan/logger"
	"math/rand"
	"net/http"
	"time"
)

func Run()  {
	err := logger.SetLogger(`{"File": {"filename":"app.log","append":true,"maxlines":1000000,"maxsize":10,"daily":true,"maxdays":-1,"level":"INFO","permit":"0660","LogLevel":0}}`)
	if err != nil {
		logger.Error("日志初始化失败：", err)
		return
	}
	logger.Info("日志初始化成功")

	err = config.Read()
	if err != nil {
		logger.Error("读取配置文件失败：", err)
		return
	}

	sqlconfig := config.GetMysqlConfig()
	err = database.Open(sqlconfig.Auth, sqlconfig.Pwd, sqlconfig.Addr, sqlconfig.Db, sqlconfig.Port)
	if err != nil {
		logger.Error("连接数据库失败：", err)
		return
	}

	//设置种子
	rand.Seed(time.Now().UnixNano())

	go StartRestfulApi(config.GetListenerPort())
	go logic.Queue()

	select{}
}

func StartRestfulApi(port int) {
	SetRouter()
	SetRouter_Manager()
	SetInit()//一些必要的初始化操作
	addr := fmt.Sprintf(":%d", port)
	logger.Fatal(http.ListenAndServe(addr, nil))
}

func testAes() {
	data := "hPvfnsU3Kocggq7CFl+xkjfU0bWADHAk+0QLNrz5BvEff2XAjpMWxZdiohxV0s4NzJwJv0ogQLTMf1TLTB3PXNLZbHtenJLkhOUeN3J1V+s="
	bodys, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		logger.Error(err)
		return
	}

	param := public.AesDecryptECB(bodys, config.GetEncryptKey())
	paramstr := string(param)
	logger.Info("解密数据:", paramstr)
}