package database

import (
	"github.com/wonderivan/logger"
)

type t_dynamic_comment_like struct {
	Comment_id		int
	User_id			int
	Commentuser_id	int
	Dynamic_id		int
}

func DynamicCommentLike_Get(userid, commentid int) bool {
	t := &t_dynamic_comment_like{}
	err := Get(t, "SELECT * FROM dynamic_comment_like WHERE comment_id = ? and user_id = ?",
		commentid, userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}

		logger.Error(err)
		return false
	}

	return true
}

func DynamicCommentLike_Insert(userid, commentid, commentuserid, dynamicid int) error {
	_, err := Exec("INSERT INTO dynamic_comment_like(comment_id,user_id,commentuser_id,dynamic_id) VALUES(?,?,?,?)",
		commentid, userid, commentuserid, dynamicid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicCommentLike_Delete(userid, commentid int) error {
	_, err := Exec("DELETE FROM dynamic_comment_like WHERE comment_id = ? and user_id = ?",
		commentid, userid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func DynamicCommentLike_GetCount(commentid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_comment_like WHERE comment_id = ?",
		commentid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func DynamicCommentLike_GetCountByUser(userid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM dynamic_comment_like WHERE commentuser_id = ?",
		userid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}