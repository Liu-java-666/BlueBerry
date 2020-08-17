package database

import (
	"github.com/wonderivan/logger"
)

type t_focus_notice struct {
	Id					int
	User_id				int
	To_user_id			int
	Cdate				[]uint8
}

func FocusNotice_Get(userid, touserid int) bool {
	t := &t_focus_notice{}
	err := Get(t, "SELECT * FROM focus_notice WHERE user_id = ? and to_user_id = ?",
		userid, touserid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}

		logger.Error(err)
		return false
	}

	return true
}

func FocusNotice_GetList(userid, index, maxcount int) ([]*t_focus_notice, error) {
	t := []*t_focus_notice{}
	err := Select(&t, "SELECT * FROM focus_notice WHERE to_user_id = ? ORDER BY id DESC LIMIT ?,?",
		userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func FocusNotice_Insert(userid, touserid int) {
	_, err := Exec("INSERT INTO focus_notice(user_id, to_user_id) VALUES(?,?)",
		userid, touserid)
	if err != nil {
		logger.Error(err)
	}
}

func FocusNotice_Delete(userid, touserid int) {
	_, err := Exec("DELETE FROM focus_notice WHERE user_id = ? and to_user_id = ?",
		userid, touserid)
	if err != nil {
		logger.Error(err)
	}
}