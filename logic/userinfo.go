package logic

import (
	"BlueBerry/database"
	"BlueBerry/public"
	"github.com/wonderivan/logger"
)

//检查生日参数
func CheckBirthday(year, month, day int) bool {
	if year <= 1900 || year >= 2020 || month < 1 || month > 12 || day < 1 || day > 31 {
		return false
	}

	if (month == 4 || month == 6 || month == 9 || month == 11) && day > 30 {
		return false
	}

	if month == 2 && day > 29 {
		return false
	}

	if month == 2 && (year % 4 != 0 || year % 100 == 0) && day > 28 {
		return false
	}

	return true
}

//设置信息
func SetInfo(userid int, userkey string, nickname string, sex int, signature string, videoId int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//昵称已存在就不能修改
	if ck.User.Nickname != "" {
		return ErrorResult("不能修改性别和生日")
	}

	//检查参数
	if nickname == "" {
		return ErrorResult("昵称不能为空")
	}
	if sex != 0 && sex != 1 {
		return ErrorResult("性别参数错误")
	}
	if signature == ""{
		return ErrorResult("签名不能为空")
	}

	tvideo, _ := database.Video_Get(videoId)
	if tvideo == nil {
		return ErrorResult("视频ID有误")
	}


	// 更新用户信息
	err := ck.User.SetInfo(nickname, sex , signature,tvideo.Id)
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

//编辑信息
func EditInfo(userid int, userkey string, nickname string, signature string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//检查参数
	if nickname == "" {
		return ErrorResult("昵称不能为空")
	}


	//没有修改
	if ck.User.Nickname == nickname && ck.User.Signature == signature {
		return ErrorResult("没有修改")
	}

	//保存资料
	err := ck.User.UpdateInfo(nickname, signature)
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

//我的菜单
func GetMyMenu(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Coins int
		FocusNum int
		FansNum int
		LikeNum int
		GiftList []int
		IsGiftMore bool
	}{
		true,
		ck.User.Nickname,
		MakeImageUrl(ck.User.AvatarFile),
		ck.User.AvatarAudit,
		ck.User.Sex,
		ck.User.Age,
		ck.User.Coins,
		database.FocusList_GetFocus(userid),
		database.FocusList_GetFans(userid),
		database.User_GetLikeCount(userid),
		[]int{},
		false,
	}

	tlist, _ := database.GiftLog_LastList(userid, 6)
	for _, v := range tlist {
		result.GiftList = append(result.GiftList, v.Gift_id)
		if len(result.GiftList) >= 5 {
			break
		}
	}
	result.IsGiftMore = len(tlist) > 5

	return result
}

//我的资料
func GetMyInfo(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	type tagPhotoItem struct {
		Id int
		Url string
	}

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Birthday string
		Signature string
		Status string
		Purpose	string
		Hobbies	string
		PhotoList []tagPhotoItem
	}{
		true,
		ck.User.Nickname,
		MakeImageUrl(ck.User.AvatarFile),
		ck.User.AvatarAudit,
		ck.User.Sex,
		ck.User.Age,
		string(ck.User.Birthday),
		ck.User.Signature,
		ck.User.Relationship_status,
		ck.User.Friends_purpose,
		ck.User.Hobbies,
		[]tagPhotoItem{},
	}

	to, _ := database.PhotoList_MyList(userid)
	if to != nil {
		pidlist := public.GetFileIdList(to.Photolist)
		for _, v := range pidlist {
			tp, _ := database.Image_Get(v)
			if tp != nil {
				Item := tagPhotoItem{
					Id:  v,
					Url: MakeImageUrl(tp.File_name),
				}
				result.PhotoList = append(result.PhotoList, Item)
			}
		}
	}

	return result
}

//用户详情
func GetUserDetail(userid int, userkey string, touserid int) interface{} {
	defer TryCatch()
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	var tu *database.TUser
	var err error
	// 查询自己
	if userid == touserid {
		tu = ck.User
	// 查询他人
	} else {
		tu, err = database.User_GetById(touserid, false)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		if tu == nil {
			return ErrorResult("用户不存在")
		}
	}

	relation:=0 //关系 0 没有关系 1 我关注他 2 他关注我 3 相互关注
	isFocus := database.FocusList_Get(userid, touserid)
	isBeFocus:= database.FocusList_Get(touserid,userid )
	if isFocus==true && isBeFocus== true {
		relation = 3
	}else if isFocus==false && isBeFocus==false {
		relation = 0
	}else if  isFocus {
		relation = 1
	}else{
		relation = 2
	}



	//动态列表
	type dy struct{
		DynamicId int		//动态ID
		Description string  //描述
		FileType string		//文件类型
		LikeNum int			//点赞数
		PostTime int		//发布时间
		VideoCover string 	//图片
		VideoUrl string 	//视频地址
		IsAudit	bool		//是否审核 true 审核 false 未审核
	}
	//作品列表
	type zp struct{
		DynamicId int		//动态ID
		SentenceID		int
		SentenceText	string
		SentenceType 	string
		Description string  //描述
		VoiceUrl string //声音
		Second int //秒
		HeartNum int // 心动数
		CommentNum int // 评论数
	}


	video,_ := database.Video_Get(tu.Video_id)
	VideoFile := ""
	VideoCoverFile :=""
	VideoId := 0
	if video != nil {
		VideoFile =  MakeVideoUrl(video.File_name)
		VideoCoverFile= MakeVideoUrl(video.Cover_name)
		VideoId = video.Id
	}

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		VideoId int
		VideoFile string
		VideoCoverFile string
		Sex int
		Signature string
		Relation int
		FocusNum int
		FansNum int
		LikeNum int
		HeartNum int
		Dylist []dy
		Zplist []zp
	}{
		true,
		tu.Nickname,
		MakeImageUrl(tu.AvatarFile),
		tu.AvatarAudit,
		VideoId,
		VideoFile,
		VideoCoverFile,
		tu.Sex,
		tu.Signature,
		relation,
		database.FocusList_GetFocus(touserid),
		database.FocusList_GetFans(touserid),
		database.User_GetLikeCount(touserid),
		database.User_GetHeartCount(touserid),

		[]dy{},
		[]zp{},

	}

	var dylist []*database.DetailDynamic
	if userid==touserid {
		//自己
		dylist,_ = database.Dynamic_GetMyByUserID(touserid)
	}else{
		dylist,_ = database.Dynamic_GetByUserID(touserid)
	}
	//动态列表
	for _,v := range dylist{
		dynamic := dy{}
		dynamic.DynamicId = v.Id
		dynamic.Description = v.Description
		dynamic.FileType = v.Filetype
		dynamic.LikeNum= v.LikeNum
		dynamic.IsAudit = v.Is_audit > 0
		dynamic.PostTime=public.StrToTimestamp(string(v.Post_time))
		pidlist := public.GetFileIdList(v.Filelist)
		if v.Filetype == "video" {
			//视频
			tv, _ := database.Video_Get(pidlist[0])
			if tv != nil {
				dynamic.VideoCover = MakeVideoUrl(tv.Cover_name)
				dynamic.VideoUrl = MakeVideoUrl(tv.File_name)
				result.Dylist = append(result.Dylist, dynamic)
			}

		}else if v.Filetype == "image"{
			//图片
			tv, _ := database.Image_Get(pidlist[0])
			if tv != nil {
				dynamic.VideoCover = MakeImageUrl(tv.File_name)
				dynamic.VideoUrl = ""//图片 视频地址为空
				result.Dylist = append(result.Dylist, dynamic)
			}
		}
	}


	//作品列表
	var zplist []*database.DetailWork
	if userid==touserid {
		//自己
		zplist,_ = database.Work_GetMyByUserID(touserid)
	}else{
		zplist,_ = database.Work_GetByUserID(touserid)
	}

	for _,v := range zplist{
		sentenceid := 0
		sentencetype:="3"
		sentencetext:=""
		if v.Sentence_id > 0 {
			sentence,err := database.GetSentenceByID(v.Sentence_id)
			if err != nil {
				logger.Info("数据库异常 查询句子出错")
				continue
			}
			if sentence == nil{
				logger.Info("数据库异常 查询句子为空")
				continue
			}
			sentenceid = sentence.Id
			sentencetype=sentence.Sentence_type
			sentencetext=sentence.Sentence_text
		}
		dynamic := zp{}
		dynamic.DynamicId=v.Id
		dynamic.SentenceID = sentenceid
		dynamic.SentenceType= sentencetype
		dynamic.SentenceText = sentencetext
		dynamic.Description = v.Description
		dynamic.HeartNum= v.LikeNum
		dynamic.CommentNum=v.CommentNum
		dynamic.Second=v.Voice_Second
		pidlist := public.GetFileIdList(v.Filelist)
		tv, _ := database.Voice_Get(pidlist[0])
		if tv != nil {
			dynamic.VoiceUrl = MakeVoiceUrl(tv.File_name)
			result.Zplist = append(result.Zplist, dynamic)
		}
	}

	return result
}

//用户名片
func GetUserCard(userid int, userkey string, touserid int) interface{} {
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

	//返回成功消息
	result := struct{
		Result bool
		Nickname string
		AvatarFile string
		AvatarAudit int
		Sex int
		Age int
		Signature string
		IsFocus bool
		FocusNum int
		FansNum int
		LikeNum int
		CoinsUsed int
	}{
		true,
		tu.Nickname,
		MakeImageUrl(tu.AvatarFile),
		tu.AvatarAudit,
		tu.Sex,
		tu.Age,
		tu.Signature,
		database.FocusList_Get(userid, touserid),
		database.FocusList_GetFocus(touserid),
		database.FocusList_GetFans(touserid),
		database.User_GetLikeCount(touserid),
		tu.Coins_used,
	}

	return result
}

//批量用户详情
func GetUserInfoList(userid int, userkey string, touseridlist []int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	type tagUserData struct {
		UserId int
		Nickname string
		AvatarFile string
		IsFocus bool
	}

	UserInfoList := []tagUserData{}
	for _, v := range touseridlist {
		Item := tagUserData{
			UserId:     v,
		}
		tu, err := database.User_GetById(v, false)
		if err == nil && tu != nil {
			Item.Nickname = tu.Nickname
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
			Item.IsFocus = database.FocusList_Get(userid, v)
		}
		UserInfoList = append(UserInfoList, Item)
	}

	//返回成功消息
	result := struct{
		Result bool
		UserInfoList []tagUserData
	}{
		true,
		UserInfoList,
	}

	return result
}

//是否拉黑
func IsBlacklist(userid int, userkey string, touserid int) interface{} {
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

	//返回成功消息
	result := struct{
		Result bool
		IsBlacklist bool
		IsBeBlacklist bool
	}{
		true,
		database.Blacklist_Get(userid, touserid),
		database.Blacklist_Get(touserid, userid),
	}

	return result
}

//获取关注列表
func GetFocusList(userid int, userkey string, page int) interface{} {
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
	tlist, err := database.FocusList_GetFocusList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		Signature   string
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.To_user_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Sex,
				tu.Age,
				tu.Signature,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取粉丝列表
func GetFansList(userid int, userkey string, page int) interface{} {
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
	tlist, err := database.FocusList_GetFansList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		IsFocus		bool
		Signature 	string
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Sex,
				tu.Age,
				database.FocusList_Get(userid, tu.Id),
				tu.Signature,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取黑名单
func GetBlacklist(userid int, userkey string, page int) interface{} {
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
	tlist, err := database.Blacklist_GetList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		Signature   string
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.To_user_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Sex,
				tu.Age,
				tu.Signature,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取好友列表
func GetFriendList(userid int, userkey string, page int) interface{} {
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
	tlist, err := database.FocusList_GetFriendList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Signature	string
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.To_user_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				tu.Signature,
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取好友申请列表
func GetApplyList(userid int, userkey string, page int) interface{} {
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
	tlist, err := database.FocusNotice_GetList(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Cdate		int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item := tagUserData{
				tu.Id,
				tu.Nickname,
				MakeImageUrl(tu.AvatarFile),
				public.StrToTimestamp(string(v.Cdate)),
			}
			userlist = append(userlist, Item)
		}
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取收礼列表
func GetReceiveGiftList(userid int, userkey string, page int) interface{} {
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
	tlist, err := database.GiftLog_List(userid, index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		GiftId		int
		GiftName	string
		Cdate		int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		Item := tagUserData{
			UserId: v.User_id,
			GiftId: v.Gift_id,
			GiftName:v.Gift_Name,
			Cdate:  public.StrToTimestamp(string(v.Cdate)),
		}

		tu, _ := database.User_GetById(v.User_id, false)
		if tu != nil {
			Item.Nickname = tu.Nickname
			Item.AvatarFile = MakeImageUrl(tu.AvatarFile)
			Item.Sex = tu.Sex
			Item.Age = tu.Age
		}

		userlist = append(userlist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}

//获取排行榜
func GetRankList(userid int, userkey, tag string, page int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	perpage := 10
	index := page * perpage
	if index < 0 {
		index = 0
	}

	var tlist []*database.RankData
	var err error

	switch tag {
	case "star":
		tlist, err = database.User_StarList(index, perpage)
	case "charm":
		tlist, err = database.User_CharmList(index, perpage)
	default:
		tlist, err = database.User_RichList(index, perpage)
	}

	if err != nil {
		return ErrorResult("数据库异常")
	}


	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Age			int
		Amount		int
	}

	userlist := []tagUserData{}
	for _, v := range tlist {
		Item := tagUserData{
			v.Id,
			v.Nickname,
			MakeImageUrl(v.AvatarFile),
			v.Sex,
			v.Age,
			v.Num,
		}
		userlist = append(userlist, Item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]tagUserData
	}{
		true,
		len(userlist) < perpage,
		userlist,
	}

	return result
}


// 搜索用户
func SearchUser(userid int, userkey string, nickname string) interface{} {
	// 判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	var rt []*database.TUser
	var err error
	rt, err = database.GetUserByNickName(nickname)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	type tagUserData struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Signature string
	}

	userlist := []tagUserData{}
	for _, v := range rt {
		Item := tagUserData{
			v.Id,
			v.Nickname,
			MakeImageUrl(v.AvatarFile),
			v.Sex,
			v.Signature,
		}
		userlist = append(userlist, Item)
	}


	result := struct{
		Result	bool
		Data	 []tagUserData
	}{
		true,
		userlist,
	}

	return result
}

// 房间送礼排行榜
type SendGiftUser struct{
	UserId int
	AvatarFile string
	Sex int
	Num int
}
func OnGetPayRankList(userid int, userkey string,roomid int,time string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	var list []*database.RankUser
	var err error
	// 获取贡献榜
	list, err = database.GetSendUserList(roomid,time)

	if err != nil {
		return ErrorResult("数据库异常")
	}


	type model struct {
		UserId		int
		Nickname	string
		AvatarFile	string
		Sex			int
		Num			int
	}

	myrank := -1 //自己名次
	rtlist := []model{}
	for i, v := range list {
		if v.Id == ck.User.Id {
			//自己 保存名次
			myrank = i+1
		}
		Item := model{
			v.Id,
			v.Nickname,
			MakeImageUrl(v.AvatarFile),
			v.Sex,
			v.Num,
		}
		rtlist = append(rtlist, Item)
	}



	result := struct{
		Result	bool
		MyRank  int //自己名次
		Data	[]model
	}{
		true,
		myrank,
		rtlist,
	}

	return result
}

// 勿扰开关
func OnMyNoDisturb(userid int, userkey string, action int) interface{} {
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	err := ck.User.SetNoDisturb(action)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	return struct{Result bool}{true}
}



// 设置视频为详情主页
func OnUserSetVideo(userid int, userkey string, dynamicid int) interface{} {
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}


	dy,err := database.Dynamic_Get(dynamicid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if dy == nil {
		return ErrorResult("动态不存在")
	}

	filelist := public.GetFileIdList(dy.Filelist)
	if len(filelist)==0 {
		return ErrorResult("视频文件不存在")
	}

	video,err := database.Video_Get(filelist[0])
	if err != nil{
		return ErrorResult("数据库异常")
	}
	if video== nil{
		return ErrorResult("视频不存在")
	}
	//更新用户视频ID
	err = ck.User.UpdateVideoId(filelist[0])
	if err != nil{
		return ErrorResult("数据库异常 更新用户视频ID")
	}

	//更新视频用途
	err = video.UpdateVideoUseType("detail")
	if err != nil{
		return ErrorResult("数据库异常 更新视频用途")
	}

	//删除这条动态信息
	err = dy.Delete()
	if err != nil {
		return ErrorResult("数据库异常 删除动态出错")
	}

	return struct{Result bool}{true}



}

