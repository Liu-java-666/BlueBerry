package logic

import (
	"BlueBerry/database"
)

// 获取房间列表
func SentenceList(userid int, userkey string, page int, stype int) interface{} {
	// 判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		//return ck.Error
	}
	// 每页记录数
	perpage := 3
	// 开始行
	index := page * perpage
	if index < 0 {
		index = 0
	}

	var rt []*database.TSentence
	var err error

	//总记录数
	totalCnt := database.GetSentenceCountByType(stype)
	//总页数
	totalPage:=0
	if totalCnt%perpage==0 {
		totalPage = totalCnt / perpage
	}else{
		totalPage = totalCnt / perpage + 1
	}
	rt, err = database.GetSentenceList(index, perpage, stype)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]*database.TSentence
	}{
		true,
		totalPage==page+1,
		rt,
	}

	return result
}

