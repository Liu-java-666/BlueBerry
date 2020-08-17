package database

import (
	"github.com/wonderivan/logger"
)

type t_gift_log struct {
	Id				int
	User_id			int
	To_user_id		int
	Gift_id			int
	Coins			int
	Cdate			[]uint8
	Scene			string
	Scene_id		int
}



func GiftLog_Insert(userid, touserid, giftid, coins, sceneid int, scene string) {
	_, err := Exec("INSERT INTO gift_log(user_id,to_user_id,gift_id,coins,scene,scene_id) VALUES(?,?,?,?,?,?)",
		userid, touserid, giftid, coins, scene, sceneid)
	if err != nil {
		logger.Error(err)
		return
	}
}

func GiftLog_GetReceiveCnt(touserid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COUNT(*) FROM gift_log WHERE to_user_id = ?",
		touserid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}

func GiftLog_LastList(touserid, maxcount int) ([]*t_gift_log, error) {
	t := []*t_gift_log{}
	err := Select(&t, `SELECT * FROM gift_log WHERE to_user_id = ? ORDER BY id DESC LIMIT ?`,
		touserid, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}


//查询用户收礼列表
type tagGiftUser struct {
	User_id int
	Coins int
}
func GiftLog_UserList(touserid, index, maxcount int) ([]*tagGiftUser, error) {
	t := []*tagGiftUser{}
	err := Select(&t, `SELECT user_id, coins FROM (
			SELECT MIN(id) AS id, user_id, SUM(coins) AS coins FROM gift_log 
			WHERE to_user_id = ? GROUP BY user_id ORDER BY id DESC LIMIT ?,?
		) AS a`,
		touserid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

type tagRecGift struct {
	t_gift_log
	Gift_Name string
}
func GiftLog_List(touserid, index, maxcount int) ([]*tagRecGift, error) {
	t := []*tagRecGift{}
	err := Select(&t, `SELECT a.*,b.name as gift_name 
		FROM gift_log a,gift b WHERE 
		a.gift_id=b.id and  
		to_user_id = ? 
		ORDER BY id DESC LIMIT ?,?`,
		touserid, index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func GiftLog_GetValue(scene string, sceneid int) int {
	cnt := 0
	err := Get(&cnt, "SELECT COALESCE(SUM(coins),0) FROM gift_log WHERE scene = ? AND scene_id = ?",
		scene, sceneid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}
