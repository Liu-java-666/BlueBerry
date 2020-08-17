package database

import (
	"github.com/wonderivan/logger"
)

type t_denounce struct {
	Id				int
	User_id			int
	To_user_id		int
	Post_time		[]uint8
	Type			string
	Content			string
}

func Denounce_Insert(userid, touserid int, Type, content string) error {
	_, err := Exec("INSERT INTO denounce(user_id,to_user_id,type,content) VALUES(?,?,?,?)",
		userid, touserid, Type, content)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
