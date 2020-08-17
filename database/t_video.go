package database

import (
	"github.com/wonderivan/logger"
)

type t_video struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	File_name		string
	File_type		string
	Cover_name		string
	Cover_type		string
	Rotation		int
	Use_type		string
	Is_audit		int
	Audit_time		[]uint8
}

func Video_Get(id int) (*t_video, error) {
	t := &t_video{}
	err := Get(t, "SELECT * FROM video WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Video_Insert(userid int, file, filetype, cover, covertype, usetype string, rotation int) (int, error) {
	result, err := Exec("INSERT INTO video(user_id,file_name,file_type,cover_name,cover_type,rotation,use_type) VALUES(?,?,?,?,?,?,?)",
		userid, file, filetype, cover, covertype, rotation, usetype)
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

func (t *t_video) Delete() error {
	_, err := Exec("DELETE FROM video WHERE `id` = ?",
		t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// 更新视频用户
func (t *t_video) UpdateVideoUseType(use_type string) error{
	_, err := Exec(`update video set use_type=? where id=?`, use_type, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// 