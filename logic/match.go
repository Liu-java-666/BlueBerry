package logic

import (
	"BlueBerry/database"
)

//1V1
func GetMatchUser(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	t, err := database.User_GetMatchUser(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
		UserId int
		Nickname string
		AvatarFile string
		Sex int
		Age int
		OtherList []string
	}{
		Result: true,
	}

	//没有匹配到
	if t == nil {
		return result
	}

	//其他用户（老虎机效果）
	tlist, err := database.User_GetDestined(userid, 10)
	if err != nil{
		return ErrorResult("数据库异常")
	}

	//写日志
	database.MatchLog_Insert(userid, t.Id)

	result.UserId = t.Id
	result.Nickname = t.Nickname
	result.AvatarFile = MakeImageUrl(t.AvatarFile)
	result.Sex = t.Sex
	result.Age = t.Age
	for _, v := range tlist {
		result.OtherList = append(result.OtherList, MakeImageUrl(v.AvatarFile))
	}

	return result
}

//打电话
func CallUp(userid int, userkey string, touserid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	tu, err := database.User_GetById(touserid, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("用户不存在")
	}

	if database.Blacklist_Get(userid, touserid) {
		return ErrorResult("你已将TA拉入黑名单")
	} else if database.Blacklist_Get(touserid, userid) {
		return ErrorResult("TA已将您拉入黑名单")
	}

	if tu.Nodisturb==1 {
		return ErrorResult("对方开启勿扰")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}

//挂电话
func HangUp(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}