package logic

import (
	"BlueBerry/config"
	"BlueBerry/database"
	"BlueBerry/im"
	"BlueBerry/public"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
	"sort"
	"time"
)

//获取房间列表
func RoomList(userid int, userkey string, page int) interface{} {
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

	type HomeRoomInfo struct {
		UserId			int
		RoomId			int
		NickName		string
		Sex				int
		AvatarFile		string
		Signature		string
		RoomCover 		string
	}


	list := []HomeRoomInfo{}
	var err error

	var tlist []*database.HomeRoom
	tlist, err = database.Room_GetAllRoom(index, perpage)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	for _,v := range tlist {
		item := HomeRoomInfo{}
		item.UserId = v.User_Id
		item.RoomId = v.RoomId
		item.NickName = v.NickName
		item.RoomCover =  config.GetUploadPath()+v.Room_Cover
		item.Sex = v.Sex
		img,_ := database.Image_Get(v.Avatar_Id)
		if img==nil {
			continue
		}
		item.AvatarFile = MakeImageUrl(img.File_name)
		item.Signature = v.Signature
		list = append(list,item)
	}

	result := struct{
		Result	bool
		IsEnd	bool
		Data	[]HomeRoomInfo
	}{
		true,
		len(list) < perpage,
		list,
	}

	return result
}


// 收藏房间
// action 0 取消 1 收藏
func OnColRoom(userid int, userkey string, roomid int,action int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//查询房间

	Room,err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if Room == nil {
		return ErrorResult("房间不存在")
	}
	UserColRoom,err := database.GetUserColRoomByUserIDAndRoomID(userid,roomid)
	// 取消
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if action == 0 {
		if UserColRoom == nil {
			return ErrorResult("未收藏不能取消")
		}
		//取消收藏
		database.DelUserColRoom(UserColRoom.ID)
		result := struct {
			Result bool
		}{
			true,
		}
		return result
	}else if action == 1 {
		// 收藏
		if UserColRoom!=nil{
			return ErrorResult("不能重复收藏")
		}
		database.AddUserColRoom(ck.User.Id,Room.Id, time.Now().Unix())
		result := struct {
			Result bool
		}{
			true,
		}
		return result
	}
	result := struct {
		Result bool
	}{
		true,
	}
	return result
}




//进入房间
func RoomEnter(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if r == nil {
		return ErrorResult("房间不存在")
	}

	tu, err := database.User_GetById(r.User_id, false)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("房主不存在")
	}

	tru, err := database.RoomUser_Get(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	if tru == nil {
		//保存房间号
		database.RoomUser_Insert(userid, roomid, r.Im_group)
	} else if tru.Room_id != roomid {
		//退出上个房间群
		if tru.Room_id>0 {
			im.DeleteGroupMember(tru.Im_group, userid)
		}
		//保存现在的房间号
		tru.SetRoom(roomid, r.Im_group)
	}




	userColRoom,_ := database.GetUserColRoomByUserIDAndRoomID(userid,roomid)
	isCol := false
	if userColRoom!= nil {
		isCol = true
	}




	//加IM群
	avatarfile := ""
	if ck.User.AvatarAudit == 1 {
		//已审核
		avatarfile = MakeImageUrl(ck.User.AvatarFile)
	}else{
		avatarfile = MakeImageUrl(database.Image_GetDefaultAvatar())
	}
	go im.AddGroupMember(r.Im_group, ck.User.Nickname, userid,avatarfile, OnAddGroupMember)


	//返回成功消息
	result := struct{
		Result 			bool
		ImGroup 		string
		UserId			int
		Sex 			int
		Nickname		string
		AvatarFile		string
		GiftValue 		int
		LikeNum 		int
		IsCol			bool
		IsFocus			bool
	}{
		true,
		r.Im_group,
		r.User_id,
		tu.Sex,
		tu.Nickname,
		MakeImageUrl(tu.AvatarFile),
		database.GiftLog_GetValue("room", roomid),
		r.Like_num,
		isCol,
		database.FocusList_Get(ck.User.Id,r.User_id),
	}

	return result
}

//增加群组成员结果
func OnAddGroupMember(resultData string, err error, user_data interface{}) {
	logger.Info("IM 增加群组成员结果 START")
	userdata := user_data.([]string)

	if err != nil {
		logger.Error("增加群组成员失败,group=%s,account=%s,err=%v", userdata[0], userdata[1], err)
		return
	}

	//logger.Debug(resultData)

	revData := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultData), &revData)
	if err != nil {
		logger.Error("增加群组成员失败,group=%s,account=%s,err=%v,resultData=%v", userdata[0], userdata[1], err, resultData)
		return
	}

	//logger.Debug(revData)

	ActionStatus := revData["ActionStatus"].(string)
	if ActionStatus != "OK" {
		ErrorCode := int(revData["ErrorCode"].(float64))
		ErrorInfo := revData["ErrorInfo"].(string)
		logger.Error("增加群组成员失败,group=%s,account=%s,errcode=%d,errinfo=%s", userdata[0], userdata[1], ErrorCode, ErrorInfo)
		return
	}

	//发送欢迎消息
	WelMsg := struct {
		Content string
		AvatarFile string
		UserId string
	}{
		fmt.Sprintf("%s来了，欢迎！", userdata[2]),
		userdata[3],
		userdata[4],
	}

	bytes,err := json.Marshal(WelMsg)
	if err == nil {
		msg := string(bytes)
		im.SendGroupSysNotice(userdata[0], 0, msg)
	}

	//go RoomUpdateOnline(userdata[0])
	logger.Info("IM 增加群组成员结果 END")
}

//退出房间
func RoomLeave(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, _ := database.Room_Get(roomid)
	if r != nil {
		//退IM群
		im.DeleteGroupMember(r.Im_group, userid)
	}else{
		return ErrorResult("参数错误 房间不存在")
	}

	tru, _ := database.RoomUser_Get(userid)
	if tru != nil {
		//清除房间号
		tru.SetRoom(0, "")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		Result: true,
	}

	return result
}

//申请创建房间
func RoomCreate(userid int, userkey string, roomtype int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if roomtype < 0 || roomtype > 1 {
		return ErrorResult("房间类型错误")
	}

	err := database.Room_Insert(userid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
	}{
		Result: true,
	}

	return result
}

//点赞房间
func RoomLike(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if r == nil {
		return ErrorResult("房间不存在")
	}

	err = database.Room_Like(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	RoomUpdateData(r)

	//返回成功消息
	result := struct{
		Result bool
	}{
		Result: true,
	}

	return result
}

//申请上座
func RoomSeat(userid int, userkey string, roomid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	r, err := database.Room_Get(roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}
	if r == nil {
		return ErrorResult("房间不存在")
	}

	result := struct{
		Result bool
		IsRepeat bool
	}{
		true,
		false,
	}

	if database.RoomSeat_Get(userid, roomid) {
		result.IsRepeat = true
		return result
	}

	err = database.RoomSeat_Insert(userid, roomid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	return result
}

//房间在线人数更新
func RoomUpdateOnline(imGroup string) {
	//logger.Info("房间在线人数更新，imGroup=", imGroup)
	r, _ := database.Room_Find(imGroup)
	if r == nil {
		logger.Error("查询房间失败，imGroup=", imGroup)
		return
	}

	im.GroupInfo(imGroup, OnGroupInfo)
}

type tagMember struct {
	Account string
	JoinTime int
}
type memberListSlice []tagMember
func (s memberListSlice) Len() int {return len(s)}
func (s memberListSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s memberListSlice) Less(i, j int) bool { return s[i].JoinTime < s[j].JoinTime }

//获取群详细资料结果
func OnGroupInfo(resultData string, err error, user_data interface{}) {
	//logger.Info("获取群详细资料结果")
	imGroup := user_data.(string)

	if err != nil {
		logger.Error("获取群详细资料失败,group=%s,err=%v", imGroup, err)
		return
	}

	logger.Debug(resultData)

	revData := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultData), &revData)
	if err != nil {
		logger.Error("获取群详细资料失败,group=%s,err=%v,resultData=%v", imGroup, err, resultData)
		return
	}

	//logger.Debug(revData)

	ActionStatus := revData["ActionStatus"].(string)
	if ActionStatus != "OK" {
		ErrorCode := int(revData["ErrorCode"].(float64))
		ErrorInfo := revData["ErrorInfo"].(string)
		logger.Error("获取群详细资料失败,group=%s,errcode=%d,errinfo=%s", imGroup, ErrorCode, ErrorInfo)
		return
	}

	r, _ := database.Room_Find(imGroup)
	if r == nil {
		logger.Error("获取群详细资料结果错误：房间不存在，group=%s", imGroup)
		return
	}

	GroupInfoList, ok := revData["GroupInfo"].([]interface{})
	if !ok {
		logger.Error("获取群详细资料结果错误：返回GroupInfoList数据有误,group=%s", imGroup)
		return
	}
	if len(GroupInfoList) <= 0 {
		logger.Error("获取群详细资料结果错误：返回GroupInfo数据为空,group=%s", imGroup)
		return
	}
	GroupInfo, ok := GroupInfoList[0].(map[string]interface{})
	if !ok {
		logger.Error("获取群详细资料结果错误：返回GroupInfo数据有误,group=%s", imGroup)
		return
	}

	MemberNum := int(GroupInfo["MemberNum"].(float64))
	logger.Debug("MemberNum=", MemberNum)

	MemberList, ok := GroupInfo["MemberList"].([]interface{})
	if !ok {
		logger.Error("获取群详细资料结果错误：返回MemberList数据有误,group=%s", imGroup)
		return
	}

	memberList := memberListSlice{}
	for _, v := range MemberList {
		item, ok := v.(map[string]interface{})
		if !ok {
			logger.Error("获取群详细资料结果错误：返回Member数据有误,group=%s", imGroup)
			continue
		}
		memberList = append(memberList, tagMember{
			Account:  item["Member_Account"].(string),
			JoinTime: int(item["JoinTime"].(float64)),
		})
	}
	//sort.Sort(sort.Reverse(memberList))
	sort.Sort(memberList)

	type tagOnlineUser struct {
		UserId int
		AvatarFile string
		Sex int
		Nickname string
	}
	onlineList := []tagOnlineUser{}
	for _, v := range memberList {
		userid := public.GetIMUserID(v.Account)
		if userid <= 0 {
			continue
		}
		tu, _ := database.User_GetById(userid, false)
		if tu == nil {
			continue
		}
		onlineList = append(onlineList, tagOnlineUser{
			UserId:     userid,
			AvatarFile: MakeImageUrl(tu.AvatarFile),
			Sex:tu.Sex,
			Nickname:tu.Nickname,
		})
		//
		//if len(onlineList) >= 5 {
		//	break
		//}
		//
	}

	//发送在线列表消息
	msgdata := struct{
		RoomID int
		OnlineUserCnt int
		OnlineUserList []tagOnlineUser
	}{
		r.Id,
		len(onlineList),
		onlineList,
	}
	//logger.Info("发送系统消息 在线列表")
	go im.SendGroupSysNotice(imGroup, 2, msgdata)
}

//更新房间数据
func RoomUpdateData(r *database.TRoom) {
	//发送房间更新消息
	msgdata := struct{
		RoomID int
		GiftValue int
		LikeNum int
	}{
		r.Id,
		database.GiftLog_GetValue("room", r.Id),
		r.Like_num,
	}
	go im.SendGroupSysNotice(r.Im_group, 3, msgdata)
}


