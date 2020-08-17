package logic

import (
	"BlueBerry/database"
	"BlueBerry/im"
	"encoding/json"
	"fmt"
	"github.com/wonderivan/logger"
)

var giftList []*database.TGift

//加载礼物列表
func LoadGiftList() {
	giftList = []*database.TGift{}

	glist, err := database.Gift_GetList()
	if err != nil {
		logger.Error("读取礼物列表失败：", err)
		return
	}
	giftList = glist
}

//获取礼物
func GetGift(giftid int) *database.TGift {
	if len(giftList) == 0 {
		LoadGiftList()
	}

	for _, v := range giftList {
		if v.Id == giftid {
			return v
		}
	}

	return nil
}

//送礼物
func SendGift(userid int, userkey, scene string, sceneid, giftid int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	var tu *database.TUser
	var err error
	var r *database.TRoom
	if scene == "dynamic" {
		t, _ := database.Dynamic_Get(sceneid)
		if t == nil {
			return ErrorResult("动态不存在")
		}

		tu, err = database.User_GetById(t.User_id, false)
	}else if scene == "room" {
		r, err = database.Room_Get(sceneid)
		if err != nil {
			return ErrorResult("数据库异常")
		}
		if r == nil {
			return ErrorResult("房间不存在")
		}

		tu, err = database.User_GetById(r.User_id, false)
	} else {
		tu, err = database.User_GetById(sceneid, false)
	}

	if userid == tu.Id {
		return ErrorResult("不能给自己送礼哦")
	}

	if err != nil {
		return ErrorResult("数据库异常")
	}
	if tu == nil {
		return ErrorResult("房主不存在")
	}
	if tu.Id == userid {
		return ErrorResult("不能给自己送礼")
	}

	gift := GetGift(giftid)
	if gift == nil {
		return ErrorResult("礼物不存在")
	}

	if ck.User.Coins < gift.Price {
		return ErrorResultSpecial(2, "钻石不足！")
	}

	err, coins := ck.User.UseCoins(gift.Price)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	database.GiftLog_Insert(userid, tu.Id, giftid, gift.Price, sceneid, scene)
	database.ActionLog_Insert(userid, tu.Id, 1, gift.Price, gift.Id, "gift", "给你赠送了礼物")

	if scene == "room" {
		//发送文字消息
		Msg := struct{
			Content string
			AvatarFile string
		}{
			fmt.Sprintf("%s送出了1个%s", ck.User.Nickname, gift.Name),
			MakeImageUrl(ck.User.AvatarFile),
		}
		bytes,err := json.Marshal(Msg)
		if err == nil {
			msgStr := string(bytes)
			go im.SendGroupSysNotice(r.Im_group, 0, msgStr)
		}

		//发送滚屏消息
		msgdata := struct{
			SendID int
			SendName string
			ReceiveID int
			ReceiveName string
			GiftID int
		}{
			userid,
			ck.User.Nickname,
			tu.Id,
			tu.Nickname,
			giftid,
		}
		go im.SendGroupSysNotice(r.Im_group, 1, msgdata)
		RoomUpdateData(r)
	}

	//返回成功消息
	result := struct{
		Result bool
		Coins int
	}{
		true,
		coins,
	}

	return result
}