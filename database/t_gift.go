package database

import (
	"github.com/wonderivan/logger"
)

type t_gift struct {
	Id			int
	Price		int
	Name		string
}

type TGift t_gift

func Gift_GetList() ([]*TGift, error) {
	t := []*TGift{}
	err := Select(&t, "SELECT * FROM gift")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

