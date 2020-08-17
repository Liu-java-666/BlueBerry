package database

import (
	"github.com/wonderivan/logger"
)

type t_match_log struct {
	Id					int
	User_id				int
	To_user_id			int
	Cdate				[]uint8
}

func MatchLog_Insert(userid, touserid int) {
	_, err := Exec("INSERT INTO match_log(user_id, to_user_id) VALUES(?,?)",
		userid, touserid)
	if err != nil {
		logger.Error(err)
	}
}
