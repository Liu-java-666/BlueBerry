package database

import (
	"github.com/wonderivan/logger"
	"time"
)

type t_dynamic_like struct {
	Dynamic_id	int
	User_id		int
	Postuser_id int
	Like_type   string
	Create_Time int
}

func DynamicLike_Get(userid, dynamicid int) bool {
	t := &t_dynamic_like{}
	err := Get(t, "SELECT * FROM dynamic_like WHERE dynamic_id = ? and user_id = ?",
		dynamicid, userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}

		logger.Error(err)
		return false
	}

	return true
}

// 根据类型查询用户是否点赞
func DynamicLike_GetByLikeType(userid int,liketype string) bool{
	t:= &t_dynamic_like{}
	err := Get(t, "select * from dynamic_like where like_type=? and user_id=?",
		liketype, userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}

		logger.Error(err)
		return false
	}

	return true
}




func DynamicLike_Insert(userid, dynamicid, postid int,like_type string) error {
	_, err := Exec("INSERT INTO dynamic_like(dynamic_id,user_id,postuser_id,like_type,create_time) VALUES(?,?,?,?,?)",
		dynamicid, userid, postid,like_type,time.Now().Unix())
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicLike_Delete(userid, dynamicid int) error {
	_, err := Exec("DELETE FROM dynamic_like WHERE dynamic_id = ? and user_id = ?",
		dynamicid, userid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicLike_GetCount(dynamicid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_like WHERE dynamic_id = ?",
		dynamicid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func DynamicLike_GetCountByUser(userid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_like WHERE postuser_id = ?",
		userid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

//获取心动我的列表
type t_dynamic_like_heart struct{
	Dynamic_id int
	User_id int
	Create_time int
	Avatar_id int
	NickName string
	Signature string
}
func DynamicLike_GetMyHeart(index int,maxcount int,postuserid int) ([]*t_dynamic_like_heart, error){
	t := []*t_dynamic_like_heart{}
	sql := `select a.dynamic_id,a.user_id,a.create_time,b.avatar_id,b.nickname,b.signature from dynamic_like a left join user b on a.user_id=b.id 
		where like_type='work' and postuser_id=? order by create_time desc limit ?,?`
	err := Select(&t, sql,postuserid,index,maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}



func DynamicLike_GetList(dynamicid, index, maxcount int) ([]*t_dynamic_like, error) {
	t := []*t_dynamic_like{}
	err := Select(&t, "SELECT * FROM dynamic_like WHERE dynamic_id = ? ORDER BY dynamic_id DESC LIMIT ?,?",
		dynamicid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}





