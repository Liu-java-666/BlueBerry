package database

import (
	"github.com/wonderivan/logger"
)

type t_dynamic_comment struct {
	Id				int
	Dynamic_id		int
	Postuser_id		int
	User_id			int
	Content			string
	Cdate			[]uint8
}

func DynamicComment_Insert(dynamicid, postid, userid int, content string) error {
	_, err := Exec("INSERT INTO dynamic_comment(dynamic_id,postuser_id,user_id,content) VALUES(?,?,?,?)",
		dynamicid, postid, userid, content)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicComment_Get(id int) (*t_dynamic_comment, error) {
	t := &t_dynamic_comment{}
	err := Get(t, "SELECT * FROM dynamic_comment WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func DynamicComment_GetList(dynamicid, index, maxcount int) ([]*t_dynamic_comment, error) {
	t := []*t_dynamic_comment{}
	err := Select(&t, "SELECT * FROM dynamic_comment WHERE dynamic_id = ? ORDER BY id DESC LIMIT ?,?",
		dynamicid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func DynamicComment_GetCount(dynamicid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_comment WHERE dynamic_id = ?",
		dynamicid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

type tagDynamic struct{
	t_dynamic_comment
	Avatar_id int
	Nickname string
}
//获取用户的评论列表
func DynamicComment_GetByUserID(userid int, index int, maxcount int)([]*tagDynamic,error){
	t := []*tagDynamic{}
	sql:=`
		select a.*,b.avatar_id,b.nickname from dynamic_comment a,user b 
		where a.user_id=b.id and postuser_id=? 
		order by cdate desc limit ?,?
		`
	err := Select(&t, sql, userid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}