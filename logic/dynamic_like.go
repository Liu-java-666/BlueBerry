package logic

import (
	"BlueBerry/database"
)

//心动我的
func OnMyDyLikeList(userid int, userkey string,page int) interface{} {
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

	tlist, err := database.DynamicLike_GetMyHeart(index, perpage,ck.User.Id)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUser struct {
		DynamicId int
		UserId int
		Nickname string
		AvatarFile string
		Signature string
		Description string
		CreateTime int
	}

	colist := []tagUser{}
	for _, v := range tlist {
		dy,err:= database.Dynamic_Get(v.Dynamic_id)
		if err!=nil{
			continue
		}
		Item := tagUser{
			DynamicId:	v.Dynamic_id,
			UserId:    v.User_id,
			Nickname:  v.NickName,
			Signature:	v.Signature,
			Description:dy.Description,
			CreateTime: v.Create_time,
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


