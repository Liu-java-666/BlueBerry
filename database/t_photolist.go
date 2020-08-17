package database

import (
	"github.com/wonderivan/logger"
)

type t_photolist struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	Photolist		string
	Is_audit		int
	Audit_time		[]uint8
}

func PhotoList_Get(id int) (*t_photolist, error) {
	t := &t_photolist{}
	err := Get(t, "SELECT * FROM photolist WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func PhotoList_Insert(userid int, photolist string) (int, error) {
	result, err := Exec(`INSERT INTO photolist(user_id,photolist) VALUES(?,?)`,
		userid, photolist)
	if err != nil {
		logger.Error(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err)
		return 0, err
	}

	return int(id), nil
}

func PhotoList_MyList(userid int) (*t_photolist, error) {
	t := &t_photolist{}
	err := Get(t, "SELECT * FROM photolist WHERE user_id = ? ORDER BY `id` DESC LIMIT 1",
		userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func PhotoList_AuditList(index, maxcount int) ([]*t_photolist, error) {
	t := []*t_photolist{}
	err := Select(&t, `SELECT * FROM photolist WHERE id IN (
			SELECT MAX(id) AS id FROM photolist WHERE is_audit = 0 GROUP BY (user_id)
		) ORDER BY id DESC LIMIT ?, ?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func PhotoList_UserList(userid, excludeid int) ([]*t_photolist, error) {
	t := []*t_photolist{}
	err := Select(&t, "SELECT * FROM photolist WHERE user_id = ? AND `id` != ?",
		userid, excludeid)
	if err != nil {
		logger.Error(err)
		return t, err
	}

	return t, nil
}

func PhotoList_UserAuditList(userid, excludeid int) ([]*t_photolist, error) {
	t := []*t_photolist{}
	err := Select(&t, "SELECT * FROM photolist WHERE user_id = ? AND is_audit = 0 AND `id` != ?",
		userid, excludeid)
	if err != nil {
		logger.Error(err)
		return t, err
	}

	return t, nil
}

func (t *t_photolist) SetAudit(audit int) error {
	_, err := Exec("UPDATE photolist SET is_audit = ?, audit_time = CURRENT_TIMESTAMP() WHERE id = ?",
		audit, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}