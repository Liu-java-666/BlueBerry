package database

import (
	"BlueBerry/public"
	"github.com/wonderivan/logger"
)

type t_pay_order struct {
	Id			int
	User_id		int
	Orderid		string
	Money		int
	Coins		int
	Postdate	[]uint8
	Status		int
	Finishdate	[]uint8
}

func PayOrder_Get(orderid string) (*t_pay_order, error) {
	t := &t_pay_order{}
	err := Get(t, "SELECT * FROM pay_order WHERE orderid = ?",
		orderid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func PayOrder_Insert(userid int, orderid string, money, coins int) error {
	_, err := Exec("INSERT INTO pay_order(user_id,orderid,money,coins) VALUES(?,?,?,?)",
		userid, orderid, money, coins)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_pay_order) SetFinish(status int) error {
	_, err := Exec("UPDATE pay_order SET status = ?, finishdate = ? WHERE orderid = ?",
		status, public.GetNowTimestr(), t.Orderid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}