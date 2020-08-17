package logic

import (
	"BlueBerry/database"
)

//关注/取消关注
func Focus(userid int, userkey string, touserid, action int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if userid == touserid {
		return ErrorResult("不能关注自己哦")
	}

	tu, err := database.User_GetById(touserid, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("用户不存在")
	}

	result := struct{
		Result bool
	}{
		true,
	}

	bFocus := database.FocusList_Get(userid, touserid)
	if (bFocus && action > 0) || (!bFocus && action <= 0) {
		//重复关注或者重复取消，直接返回成功消息
		return result
	}

	if action > 0 {
		//拉黑用户不能关注
		if database.Blacklist_Get(userid, touserid) {
			return ErrorResult("你已将TA拉入黑名单")
		} else if database.Blacklist_Get(touserid, userid) {
			return ErrorResult("TA已将您拉入黑名单")
		}

		err = database.FocusList_Insert(userid, touserid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, touserid, 1, 0, 0, "focus", "关注了你")

		//好友申请
		//我关注了对方，就把对方的好友申请删掉
		database.FocusNotice_Delete(touserid, userid)
		//如果对方没有关注我，则向对方添加一条好友申请
		if database.FocusList_Get(touserid, userid) == false {
			//删掉我的历史好友申请
			database.FocusNotice_Delete(userid, touserid)
			//插入最新的好友申请
			database.FocusNotice_Insert(userid, touserid)
		}
	} else {
		err = database.FocusList_Delete(userid, touserid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, touserid, 0, 0, 0, "focus", "对你取消了关注")
	}

	//返回成功消息
	return result
}

//拉黑/取消拉黑
func Blacklist(userid int, userkey string, touserid, action int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if touserid == 1 {
		return ErrorResult("不能拉黑小客服")
	}

	if userid == touserid {
		return ErrorResult("不能拉黑自己哦")
	}

	tu, err := database.User_GetById(touserid, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("用户不存在")
	}

	bBlacklist := database.Blacklist_Get(userid, touserid)
	if bBlacklist && action > 0 {
		return ErrorResult("已在黑名单")
	} else if !bBlacklist && action <= 0 {
		return ErrorResult("不在黑名单")
	}

	if action > 0 {
		//取消关注关系
		if database.FocusList_Get(userid, touserid) {
			err = database.FocusList_Delete(userid, touserid)
			if err != nil {
				return ErrorResult("数据库异常")
			}
		}
		if database.FocusList_Get(touserid, userid) {
			err = database.FocusList_Delete(touserid, userid)
			if err != nil {
				return ErrorResult("数据库异常")
			}
		}

		//加入黑名单
		err = database.Blacklist_Insert(userid, touserid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, touserid, 1, 0, 0, "blacklist", "拉黑了你")
	} else {
		//从黑名单移除
		err = database.Blacklist_Delete(userid, touserid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		//操作日志
		database.ActionLog_Insert(userid, touserid, 0, 0, 0, "blacklist", "将你移除了黑名单")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		true,
	}

	return result
}

//举报
func Denounce(userid int, userkey string, touserid int, Type, content string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if userid == touserid {
		return ErrorResult("不能举报自己哦")
	}

	tu, err := database.User_GetById(touserid, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("用户不存在")
	}

	err = database.Denounce_Insert(userid, touserid, Type, content)
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