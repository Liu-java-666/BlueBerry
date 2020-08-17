package database

import (
	"github.com/wonderivan/logger"
)

type t_pay_config struct {
	Id			int
	Money		int
	Coins		int
	Appid		string
}

type TPayConfig t_pay_config

func PayConfig_GetList() ([]*TPayConfig, error) {
	t := []*TPayConfig{}
	err := Select(&t, "SELECT * FROM pay_config")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

