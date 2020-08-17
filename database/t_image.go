package database

import (
	"fmt"
	"github.com/wonderivan/logger"
)

type t_image struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	File_name		string
	File_type		string
	Use_type		string
	Is_audit		int
	Audit_time		[]uint8
}

func Image_Get(id int) (*t_image, error) {
	t := &t_image{}
	err := Get(t, `SELECT * FROM image WHERE id = ?`,
		id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Image_Insert(userid int, file, filetype, usetype string) (int, error) {
	result, err := Exec("INSERT INTO image(user_id,file_name,file_type,use_type) VALUES(?,?,?,?)",
		userid, file, filetype, usetype)
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

var default_avatar string

func Image_GetDefaultAvatar() string {
	if default_avatar != "" {
		return default_avatar
	}

	t, err := Image_Get(1)
	if err != nil {
		logger.Error("读取默认头像失败，err=", err)
		return ""
	}

	if t == nil {
		logger.Error("默认头像文件不存在")
		return ""
	}

	default_avatar = t.File_name
	return default_avatar
}

func Image_GetMyAvatar(userid int) (string, int) {
	t := &t_image{}
	err := Get(t, "SELECT * FROM image WHERE user_id = ? AND use_type = 'avatar' AND is_audit >= 0 ORDER BY `id` DESC LIMIT 1",
		userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return Image_GetDefaultAvatar(), 1
		}

		logger.Error(err)
		return Image_GetDefaultAvatar(), 1
	}

	return t.File_name, t.Is_audit
}

func Image_GetOtherAvatar(id int) string {
	if id <= 1 {
		return Image_GetDefaultAvatar()
	}

	t, err := Image_Get(id)
	if err != nil {
		logger.Error(err)
		return Image_GetDefaultAvatar()
	}

	if t == nil {
		logger.Error("未找到头像文件,id=", id)
		return Image_GetDefaultAvatar()
	}

	if t.Is_audit <= 0 {
		return Image_GetDefaultAvatar()
	}

	return t.File_name
}

func Image_AuditList(index, maxcount int, usetype string) ([]*t_image, error) {
	t := []*t_image{}
	err := Select(&t, "SELECT * FROM image WHERE use_type = ? AND is_audit = 0 ORDER BY id DESC LIMIT ?,?",
		usetype, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Image_UserList(userid, excludeid int, usetype string) ([]*t_image, error) {
	t := []*t_image{}
	err := Select(&t, "SELECT * FROM image WHERE user_id = ? AND use_type = ? AND `id` != ?",
		userid, usetype, excludeid)
	if err != nil {
		logger.Error(err)
		return t, err
	}

	return t, nil
}

func Image_UnusedList(userid int, usedlist, usetype string) ([]*t_image, error) {
	t := []*t_image{}
	sqlstr := fmt.Sprintf("SELECT * FROM image WHERE user_id = %d AND use_type = '%s' AND is_audit != -1 AND `id` NOT IN (%s)",
		userid, usetype, usedlist)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return t, err
	}

	return t, nil
}

func Image_UserAuditList(userid int, usetype string) ([]*t_image, error) {
	t := []*t_image{}
	sqlstr := fmt.Sprintf("SELECT * FROM image WHERE user_id = %d AND use_type = %s AND is_audit = 0",
		userid, usetype)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return t, err
	}

	return t, nil
}

func (t *t_image) SetAudit(audit int) error {
	_, err := Exec("UPDATE image SET is_audit = ?, audit_time = CURRENT_TIMESTAMP() WHERE `id` = ?",
		audit, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_image) Delete() error {
	_, err := Exec("DELETE FROM image WHERE `id` = ?",
		t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}