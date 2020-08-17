package database

import (
	"github.com/wonderivan/logger"
)

type t_blacklist struct {
	Id					int
	User_id				int
	To_user_id			int
	Cdate				[]uint8
}

func Blacklist_Get(userid, touserid int) bool {
	t := &t_blacklist{}
	err := Get(t, "SELECT * FROM blacklist WHERE user_id = ? and to_user_id = ?",
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

func Blacklist_GetList(userid, index, maxcount int) ([]*t_blacklist, error) {
	t := []*t_blacklist{}
	err := Select(&t, "SELECT * FROM blacklist WHERE user_id = ? ORDER BY id DESC LIMIT ?,?",
		userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Blacklist_Insert(userid, touserid int) error {
	_, err := Exec("INSERT INTO blacklist(user_id, to_user_id) VALUES(?,?)",
		userid, touserid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func Blacklist_Delete(userid, touserid int) error {
	_, err := Exec("DELETE FROM blacklist WHERE user_id = ? and to_user_id = ?",
		userid, touserid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}