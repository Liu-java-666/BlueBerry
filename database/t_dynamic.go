package database

import (
	"fmt"
	"github.com/wonderivan/logger"
)

type t_dynamic struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	Description		string
	Sentence_id		int
	Topic			string
	Filetype		string
	Filelist		string
	Is_audit		int
	Audit_time		[]uint8
	Voice_Second 	int
}

type TDynamic t_dynamic

func Dynamic_Get(id int) (*t_dynamic, error) {
	t := &t_dynamic{}
	err := Get(t, "SELECT * FROM dynamic WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}
func Dynamic_GetByID(id int) (*TDynamic, error) {
	t := &TDynamic{}
	err := Get(t, "SELECT * FROM dynamic WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_Insert(userid int, description, topic, filetype, filelist string) error {
	_, err := Exec("INSERT INTO dynamic(user_id,description,topic,filetype,filelist) VALUES(?,?,?,?,?)",
		userid, description, topic, filetype, filelist)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

//发布声音动态
func Work_Insert(userid int, description, filetype, filelist string,second int,sentenceid int) error {
	_, err := Exec("INSERT INTO dynamic(user_id,description,sentence_id,filetype,filelist,is_audit,voice_second) VALUES(?,?,?,?,?,1,?)",
		userid, description, sentenceid, filetype, filelist,second)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func Dynamic_GetCount(userid int, filetype string) int {
	cnt := 0

	var sqlstr string
	if filetype == "" {
		sqlstr = fmt.Sprintf("SELECT COUNT(*) FROM dynamic WHERE user_id = %d AND is_audit = 1",
			userid)
	} else {
		sqlstr = fmt.Sprintf("SELECT COUNT(*) FROM dynamic WHERE user_id = %d AND filetype = '%s' AND is_audit = 1",
			userid, filetype)
	}

	err := Get(&cnt, sqlstr)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func GetFileTypeString(filetype string) string {
	str := ""
	if filetype == "" {
		str = "('image','video')"
	} else {
		str = fmt.Sprintf("('%s')", filetype)
	}
	return str
}

func Dynamic_AllList(index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND is_audit = 1 ORDER BY audit_time DESC, id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_TopicList(topic string, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND topic = %s AND is_audit = 1 ORDER BY audit_time DESC, id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), topic, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_FocusList(userid, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf(`SELECT * FROM dynamic WHERE filetype IN %s AND is_audit = 1 
		AND user_id IN (SELECT to_user_id FROM focuslist WHERE user_id = %d)
		ORDER BY audit_time DESC, id DESC LIMIT %d,%d`,
		GetFileTypeString(filetype), userid, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_HotList(index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf(`SELECT a.* FROM dynamic AS a LEFT JOIN (
			SELECT COUNT(*) AS likenum, postuser_id AS likeuser FROM dynamic_like GROUP BY likeuser
		) AS b ON a.user_id = b.likeuser 
		WHERE filetype IN %s AND is_audit = 1 ORDER BY likenum DESC, audit_time DESC LIMIT %d,%d`,
		GetFileTypeString(filetype), index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_MyList(userid, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND user_id = %d ORDER BY id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), userid, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_UserList(userid, index, maxcount int, filetype string) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf("SELECT * FROM dynamic WHERE filetype IN %s AND user_id = %d AND is_audit = 1 ORDER BY audit_time DESC, id DESC LIMIT %d,%d",
		GetFileTypeString(filetype), userid, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Dynamic_AuditList(index, maxcount int) ([]*t_dynamic, error) {
	t := []*t_dynamic{}
	err := Select(&t, "SELECT * FROM dynamic WHERE is_audit = 0 ORDER BY id DESC LIMIT ?,?",
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func (t *t_dynamic) SetAudit(audit int) error {
	_, err := Exec("UPDATE dynamic SET is_audit = ?, audit_time = CURRENT_TIMESTAMP() WHERE `id` = ?",
		audit, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_dynamic) Delete() error {
	_, err := Exec("DELETE FROM dynamic WHERE `id` = ?",
		t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

// 用户详情 动态列表
type DetailDynamic struct{
	t_dynamic
	LikeNum int `db:"like_num"`
}

// 详情动态列表
func Dynamic_GetByUserID(userid int)([]*DetailDynamic, error) {
	t := []*DetailDynamic{}
	err := Select(&t, "select a.*,(select COUNT(*) from dynamic_like where dynamic_id=a.id) like_num from dynamic a where filetype in ('image','video') and user_id=? and is_audit=1 order by id desc",
		userid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}
//获取自己的动态
func Dynamic_GetMyByUserID(userid int)([]*DetailDynamic, error) {
	t := []*DetailDynamic{}
	err := Select(&t, "select a.*,(select COUNT(*) from dynamic_like where dynamic_id=a.id) like_num from dynamic a where filetype in ('image','video') and user_id=? order by id desc",
		userid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}


type DetailWork struct{
	TDynamic
	LikeNum int `db:"like_num"`
	CommentNum int `db:"comment_num"`
}
//详情作品列表
func Work_GetByUserID(userid int)([]*DetailWork, error) {
	t := []*DetailWork{}
	sql := "select a.*,(select COUNT(*) from dynamic_like where dynamic_id=a.id) like_num,(select COUNT(*) from dynamic_comment where dynamic_id=a.id) comment_num from dynamic a where filetype in ('voice') and user_id=? and is_audit=1 order by id desc"
	err := Select(&t, sql,
		userid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}
func Work_GetMyByUserID(userid int)([]*DetailWork, error) {
	t := []*DetailWork{}
	sql := "select a.*,(select COUNT(*) from dynamic_like where dynamic_id=a.id) like_num,(select COUNT(*) from dynamic_comment where dynamic_id=a.id) comment_num from dynamic a where filetype in ('voice') and user_id=? order by id desc "
	err := Select(&t, sql,
		userid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}





//小视频列表
func Dynamic_VideoList(index, maxcount int) ([]*TDynamic, error) {
	t := []*TDynamic{}
	sqlstr := fmt.Sprintf(`SELECT a.* FROM dynamic AS a LEFT JOIN (
			SELECT COUNT(*) AS likenum, postuser_id AS likeuser FROM dynamic_like GROUP BY likeuser
		) AS b ON a.user_id = b.likeuser 
		WHERE filetype = 'video' AND is_audit = 1 ORDER BY likenum DESC, audit_time DESC LIMIT %d,%d`,index,maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}