package database

import (
	"github.com/wonderivan/logger"
)

type t_certification struct {
	User_id			int
	Post_time		[]uint8
	Front_img		string
	Back_img		string
	Is_audit		int
	Audit_time		[]uint8
}

func Certification_Get(userid int) (*t_certification, error) {
	t := &t_certification{}
	err := Get(t, "SELECT * FROM certification WHERE user_id = ?", userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Certification_GetMe(userid int) (*t_certification, error) {
	t := &t_certification{}
	err := Get(t, "SELECT * FROM certification WHERE user_id = ?", userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Certification_Insert(userid int, frontimg, backimg string) error {
	_, err := Exec("INSERT INTO certification(user_id,front_img,back_img) VALUES(?,?,?)",
		userid, frontimg, backimg)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func Certification_AuditList(index, maxcount int) ([]*t_certification, error) {
	t := []*t_certification{}
	err := Select(&t, "SELECT * FROM certification WHERE is_audit = 0 ORDER BY post_time DESC LIMIT ?,?",
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func (t *t_certification) SetAudit(audit int) error {
	_, err := Exec("UPDATE certification SET is_audit = ?, audit_time = CURRENT_TIMESTAMP() WHERE user_id = ?",
		audit, t.User_id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_certification) Delete() error {
	_, err := Exec("DELETE FROM certification WHERE user_id = ?",
		t.User_id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}