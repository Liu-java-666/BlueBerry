package config

import (
	"errors"
	"gopkg.in/gcfg.v1"
)

type tagMysqlConfig struct {
	Auth	string
	Pwd		string
	Addr	string
	Port	int
	Db		string
}

type tagEncryptConfig struct {
	Pre		string
	Key		string
	Test	int
}

type tagIMConfig struct {
	AppId int
	Key string
	Pre string
	AdminSig string
}

type tagUploadConfig struct {
	Root	string
	Path	string
}

type tagOtherConfig struct {
	Port		int
}

type tagAgentConfig struct{
	Domain string
	Test	int
}

type tagConfig struct {
	MySql			tagMysqlConfig
	Encrypt			tagEncryptConfig
	IM				tagIMConfig
	Upload			tagUploadConfig
	Other			tagOtherConfig
	Agent			tagAgentConfig
}

var config tagConfig

func Read() error {
	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		return err
	}

	if config.MySql.Port == 0 {
		return errors.New("数据库配置有误")
	}

	return nil
}

func GetMysqlConfig() tagMysqlConfig {
	return config.MySql
}

func GetListenerPort() int {
	return config.Other.Port
}

func GetUploadRoot() string {
	return config.Upload.Root
}

func GetUploadPath() string {
	return config.Upload.Path
}

func GetIMConfig() tagIMConfig {
	return config.IM
}

func GetEncryptPre() string {
	return config.Encrypt.Pre
}

//16位密钥
func GetEncryptKey() []byte {
	return []byte(config.Encrypt.Key)
}

func IsTest() bool {
	return config.Encrypt.Test > 0
}
//外网域名
func GetAgentDomainConfig() string{
	return config.Agent.Domain
}
//是否开启转发
func GetAgentTestConfig() int{
	return config.Agent.Test
}