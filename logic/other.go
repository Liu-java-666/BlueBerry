package logic

import "BlueBerry/database"

//提交意见反馈
func GiveFeedback(userid int, userkey string, imageid int, description, address string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//插入表
	err := database.Feedback_Insert(userid, imageid, description, address)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}




