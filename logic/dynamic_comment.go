package logic

import (
	"BlueBerry/database"
	"BlueBerry/public"
)




//评论我的
func OnMyDyCommentList(userid int, userkey string,page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 20
	index := page * perpage
	if index < 0 {
		index = 0
	}

	tlist, err := database.DynamicComment_GetByUserID(ck.User.Id,index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUser struct {
		DynamicId int
		UserId int
		Nickname string
		AvatarFile string
		CreateTime int
		Content string
	}

	colist := []tagUser{}
	for _, v := range tlist {
		Item := tagUser{
			DynamicId:	v.Dynamic_id,
			UserId:    v.User_id,
			Nickname:  v.Nickname,
			AvatarFile:"",
			CreateTime:public.StrToTimestamp(string(v.Cdate)) ,
			Content:v.Content,
		}

		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
		}

		colist = append(colist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUser
	}{
		true,
		len(colist) < perpage,
		colist,
	}

	return result
}