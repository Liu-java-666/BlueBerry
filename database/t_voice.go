package database

import (
	"fmt"
	"github.com/wonderivan/logger"
)
type t_voice struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	File_name		string
	File_type		string
	File_second     int
}

//插入声音
func Voice_Insert(userid int, filename string, filetype string,second int) (int, error) {
	result, err := Exec("insert into voice(user_id,file_name,file_type,file_second) values(?,?,?,?)",
		userid, filename, filetype,second)
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
//查询声音
func Voice_Get(id int) (*t_voice, error) {
	t := &t_voice{}
	err := Get(t, `SELECT * FROM voice WHERE id = ?`,
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



//该函数只是为了测试 无实际用途
func Voice_GetAll()(DataTable,error){
	return QueryListBySql("select * from voice")
}

//该函数只是为了测试 无实际用途
func Voice_GetByUserID(userid int)(DataMap,error){
	sql := fmt.Sprintf("select * from voice where user_id=%d",userid)
	return QueryBySql(sql)
}



func (t *t_voice) Delete() error {
	_, err := Exec("DELETE FROM voice WHERE `id` = ?",
		t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}