package database

import (
	"github.com/wonderivan/logger"
)

type t_action_log struct {
	Id				int
	User_id			int
	To_user_id		int
	Action			string
	Type			int
	Coins			int
	Cdate			[]uint8
	Description		string
	Extra			int
}

func ActionLog_Insert(userid, touserid, Type, coins, extra int, action, description string) {
	_, err := Exec("INSERT INTO action_log(user_id,to_user_id,`action`,`type`,coins,description,extra) VALUES(?,?,?,?,?,?,?)",
		userid, touserid, action, Type, coins, description, extra)
	if err != nil {
		logger.Error(err)
		return
	}
}