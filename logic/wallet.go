package logic

import (
	"BlueBerry/database"
	tsgutils "github.com/typa01/go-utils"
	"github.com/wonderivan/logger"
)

type tagPayItem struct {
	Money		int
	Coins		int
	AppId		string
}

var payList []tagPayItem

//加载充值列表
func LoadPayList() {
	tlist, _ := database.PayConfig_GetList()
	if tlist != nil {
		for _, v := range tlist {
			Item := tagPayItem{
				v.Money,
				v.Coins,
				v.Appid,
			}
			payList = append(payList, Item)
		}
	}
}

//获取充值项
func GetPayItem(money int) *tagPayItem {
	//读取充值选项
	if len(payList) == 0 {
		LoadPayList()
	}

	for _, v := range payList {
		if v.Money == money {
			return &v
		}
	}

	return nil
}

//我的钱包
func GetMyWallet(userid int, userkey string) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	//读取充值选项
	if len(payList) == 0 {
		LoadPayList()
	}

	result := struct{
		Result		bool
		Coins		int
		PayList		[]tagPayItem
	}{
		true,
		ck.User.Coins,
		payList,
	}

	return result
}

//提交充值订单
func PayOrder(userid int, userkey string, money int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	payitem := GetPayItem(money)
	if payitem == nil {
		return ErrorResult("充值金额有误")
	}

	orderid := tsgutils.GUID()
	err := database.PayOrder_Insert(userid, orderid, money, payitem.Coins)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	result := struct{
		Result		bool
		OrderId		string
	}{
		true,
		orderid,
	}

	return result
}

//完成充值
func PayFinish(userid int, userkey, orderid string, status int) interface{} {
	//判断用户
	ck := CheckUser(userid, userkey)
	if ck.Result == false {
		return ck.Error
	}

	if status != 1 && status != -1 {
		return ErrorResult("订单状态有误")
	}

	t, err := database.PayOrder_Get(orderid)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	if t == nil {
		logger.Error("订单不存在,", orderid)
		return ErrorResult("订单不存在")
	}

	if t.User_id != userid {
		return ErrorResult("不是自己的订单")
	}

	if t.Status != 0 {
		return ErrorResult("该订单已完成")
	}

	err = t.SetFinish(status)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	coins := ck.User.Coins
	if status == 1 {
		err, coins = ck.User.AddCoins(t.Coins)
		if err != nil {
			return ErrorResult("数据库异常")
		}
	}

	result := struct{
		Result		bool
		Coins		int
	}{
		true,
		coins,
	}

	return result
}