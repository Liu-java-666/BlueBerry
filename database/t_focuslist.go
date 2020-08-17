package database

import (
	"github.com/wonderivan/logger"
)

type t_focuslist struct {
	Id					int
	User_id				int
	To_user_id			int
}

func FocusList_Get(userid, touserid int) bool {
	t := &t_focuslist{}
	err := Get(t, "SELECT * FROM focuslist WHERE user_id = ? and to_user_id = ?",
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

func FocusList_GetFocusList(userid, index, maxcount int) ([]*t_focuslist, error) {
	t := []*t_focuslist{}
	err := Select(&t, "SELECT * FROM focuslist WHERE user_id = ? ORDER BY id DESC LIMIT ?,?",
		userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func FocusList_GetFansList(userid, index, maxcount int) ([]*t_focuslist, error) {
	t := []*t_focuslist{}
	err := Select(&t, "SELECT * FROM focuslist WHERE to_user_id = ? ORDER BY id DESC LIMIT ?,?",
		userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func FocusList_GetFriendList(userid, index, maxcount int) ([]*t_focuslist, error) {
	t := []*t_focuslist{}
	err := Select(&t, `SELECT * FROM focuslist WHERE user_id = ? AND to_user_id IN (
			SELECT user_id FROM focuslist WHERE to_user_id = ?
		) ORDER BY id DESC LIMIT ?,?`,
		userid, userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func FocusList_GetFans(touserid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM focuslist WHERE to_user_id = ?",
		touserid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func FocusList_GetFocus(userid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM focuslist WHERE user_id = ?",
		userid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func FocusList_Insert(userid, touserid int) error {
	_, err := Exec("INSERT INTO focuslist(user_id, to_user_id) VALUES(?,?)",
		userid, touserid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func FocusList_Delete(userid, touserid int) error {
	_, err := Exec("DELETE FROM focuslist WHERE user_id = ? and to_user_id = ?",
		userid, touserid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}