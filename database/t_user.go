package database

import (
	"BlueBerry/public"
	"fmt"
	"github.com/wonderivan/logger"
	"time"
)

type t_user struct {
	Id					int
	Phone_number		string
	Registration_time	[]uint8
	Nickname			string
	Sex					int
	Birthday			[]uint8
	User_key			string
	Lastlogon_time		[]uint8
	Lastlogon_ip		string
	Avatar_id			int
	Video_id			int
	Photolist_id		int
	Certification		int
	Signature			string
	Relationship_status string
	Friends_purpose		string
	Hobbies				string
	Coins				int
	Coins_used			int
	Nodisturb			int
}

type TUser struct {
	t_user
	AvatarFile	string
	AvatarAudit int
	Age			int
}

func GetAge(birthday []uint8) int {
	tm := public.StrToDate(string(birthday))
	if tm.Year() <= 1900 {
		return 0
	}
	now := time.Now()

	age := now.Year() - tm.Year()
	if now.Month() < tm.Month() || (now.Month() == tm.Month() && now.Day() < tm.Day()) {
		age --
	}

	return age
}

func User_GetByPhone(phone string) (*TUser, error) {
	t := TUser{}
	err := Get(&t, "SELECT * FROM user WHERE phone_number = ?", phone)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	t.AvatarFile, t.AvatarAudit = Image_GetMyAvatar(t.Id)
	t.Age = GetAge(t.Birthday)

	return &t, nil
}

func User_GetById(id int, bMe bool) (*TUser, error) {
	t := TUser{}
	err := Get(&t, "SELECT * FROM user WHERE `id` = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	if bMe {
		t.AvatarFile, t.AvatarAudit = Image_GetMyAvatar(t.Id)
	} else {
		t.AvatarFile = Image_GetOtherAvatar(t.Avatar_id)
		t.AvatarAudit = 1
	}
	t.Age = GetAge(t.Birthday)

	return &t, nil
}

func User_Insert(phone, ip, userkey string,nodisturb int) (*TUser, error) {
	result, err := Exec("INSERT INTO user(phone_number,user_key,lastlogon_ip,nodisturb) VALUES(?,?,?,?)",
		phone, userkey, ip,nodisturb)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return User_GetById(int(id), true)
}
// 用户获赞数
func User_GetLikeCount(userid int) int {
	cnt := 0
	err := Get(&cnt, `SELECT COALESCE(SUM(num),0) FROM (
			SELECT COUNT(*) AS num FROM dynamic_like WHERE postuser_id = ?
			UNION ALL
			SELECT COUNT(*) AS num FROM dynamic_comment_like WHERE commentuser_id = ?
			UNION ALL
			SELECT like_num AS num FROM room WHERE user_id = ?
		) AS c `,
		userid, userid, userid)
	if err != nil {
		logger.Error(err)
		return 0
	}

	return cnt
}
// 用户心动数 (作品或赞)
func User_GetHeartCount(userid int) int{
	cnt:=0
	err:=Get(&cnt,"select COUNT(*) from dynamic_like where postuser_id=? and like_type='work'",userid)
	if err != nil {
		logger.Error(err)
		return 0
	}
	return cnt
}


func User_SetAvatar(id, avatar int) {
	_, err := Exec("UPDATE user SET avatar_id = ? WHERE `id` = ?",
		avatar, id)
	if err != nil {
		logger.Error(err)
	}
}

func User_SetCertification(id int) {
	_, err := Exec("UPDATE user SET certification = 1 WHERE `id` = ?",
		id)
	if err != nil {
		logger.Error(err)
	}
}

func User_SetPhotoList(id, photolistid int) {
	_, err := Exec("UPDATE user SET photolist_id = ? WHERE `id` = ?",
		photolistid, id)
	if err != nil {
		logger.Error(err)
	}
}

func User_GetMatchUser(userid int) (*TUser, error) {
	t := TUser{}
	sqlstr := fmt.Sprintf(`SELECT * FROM user WHERE nickname != '' AND id > 10 AND id != %d
		AND id NOT IN (SELECT to_user_id FROM blacklist WHERE user_id = %d)
		AND id NOT IN (SELECT user_id FROM blacklist WHERE to_user_id = %d)
		AND id NOT IN (SELECT to_user_id FROM match_log WHERE user_id = %d AND cdate >= CAST(SYSDATE()AS DATE))
		GROUP BY id ORDER BY RAND() LIMIT 1`,
		userid, userid, userid, userid)
	err := Get(&t, sqlstr)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}

	t.AvatarFile = Image_GetOtherAvatar(t.Avatar_id)
	t.AvatarAudit = 1
	t.Age = GetAge(t.Birthday)

	return &t, nil
}

func User_GetDestined(id, cnt int) ([]*TUser, error) {
	t := []t_user{}
	err := Select(&t, "SELECT * FROM user WHERE nickname != '' AND id > 10 AND id != ? ORDER BY RAND() LIMIT ?",
		id, cnt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	T := []*TUser{}
	for _, v := range t {
		item := &TUser{
			t_user:      v,
			AvatarFile:  Image_GetOtherAvatar(v.Avatar_id),
			AvatarAudit: 1,
			Age:		 GetAge(v.Birthday),
		}
		T = append(T, item)
	}

	return T, nil
}

func User_Search(id int, keyword string, index, maxcount int) ([]*TUser, error) {
	t := []t_user{}
	sqlstr := fmt.Sprintf("SELECT * FROM user WHERE nickname like '%%%s%%' AND id > 10 AND id != %d ORDER BY id DESC LIMIT %d,%d",
		keyword, id, index, maxcount)
	err := Select(&t, sqlstr)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	T := []*TUser{}
	for _, v := range t {
		item := &TUser{
			t_user:      v,
			AvatarFile:  Image_GetOtherAvatar(v.Avatar_id),
			AvatarAudit: 1,
			Age:		 GetAge(v.Birthday),
		}
		T = append(T, item)
	}

	return T, nil
}

type RankData struct {
	TUser
	Num int
}

func User_RichList(index, maxcount int) ([]*RankData, error) {
	t := []*RankData{}
	err := Select(&t, `SELECT a.*, COALESCE(b.num,0) AS num FROM user AS a LEFT JOIN (
			SELECT SUM(coins) AS num, user_id FROM gift_log GROUP BY user_id
		) AS b ON a.id = b.user_id WHERE a.id > 10 AND b.num > 0 ORDER BY b.num DESC LIMIT ?,?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return t, nil
}

func User_StarList(index, maxcount int) ([]*RankData, error) {
	t := []*RankData{}
	err := Select(&t, `SELECT a.*, COALESCE(b.num,0) AS num FROM user AS a LEFT JOIN (
			SELECT SUM(coins) AS num, to_user_id FROM gift_log GROUP BY to_user_id
		) AS b ON a.id = b.to_user_id WHERE a.id > 10 AND b.num > 0 ORDER BY b.num DESC LIMIT ?,?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return t, nil
}

func User_CharmList(index, maxcount int) ([]*RankData, error) {
	t := []*RankData{}
	err := Select(&t, `SELECT a.*, COALESCE(b.num,0) AS num FROM user AS a LEFT JOIN (
			SELECT SUM(num) AS num, user_id FROM (
				SELECT SUM(like_num) AS num, user_id FROM room GROUP BY user_id
				UNION ALL
				SELECT COUNT(*) AS num, postuser_id AS user_id FROM dynamic_like GROUP BY postuser_id
				UNION ALL
				SELECT COUNT(*) AS num, commentuser_id AS user_id FROM dynamic_comment_like GROUP BY commentuser_id
			) AS c GROUP BY user_id
		) AS b ON a.id = b.user_id 
		WHERE a.id > 10 AND b.num > 0 ORDER BY b.num DESC LIMIT ?,?`,
		index, maxcount)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	for _, v := range t {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return t, nil
}

func (t *TUser) UpdateLogin(ip, userkey string) {
	_, err := Exec("UPDATE user SET user_key = ?, lastlogon_time = ?, lastlogon_ip= ? WHERE `id` = ?",
		userkey, public.GetNowTimestr(), ip, t.Id)
	if err != nil {
		logger.Error(err)
	}
}

// 完善资料
func (t *TUser) SetInfo(nickname string, sex int,signature string,video_id int) error {
	_, err := Exec("UPDATE user SET nickname = ?, sex = ?, signature= ?,video_id=? WHERE `id` = ?",
		nickname, sex, signature,video_id,t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
//更新用户信息
func (t *TUser) UpdateInfo(nickname, signature string) error {
	_, err := Exec(`UPDATE user SET nickname = ?, signature = ? WHERE id = ?`, nickname, signature, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *TUser) UseCoins(coins int) (error, int) {
	_, err := Exec("UPDATE user SET coins = coins - ?, coins_used = coins_used + ? WHERE `id` = ?",
		coins, coins, t.Id)
	if err != nil {
		logger.Error(err)
		return err, 0
	}

	cnt := 0
	_ = Get(&cnt, "SELECT coins FROM user WHERE `id` = ?", t.Id)

	return nil, cnt
}

func (t *TUser) AddCoins(coins int) (error, int) {
	_, err := Exec("UPDATE user SET coins = coins + ? WHERE `id` = ?",
		coins, t.Id)
	if err != nil {
		logger.Error(err)
		return err, 0
	}

	cnt := 0
	_ = Get(&cnt, "SELECT coins FROM user WHERE `id` = ?", t.Id)

	return nil, cnt
}



// 根据昵称查询用户
func GetUserByNickName(nickname string) ([]*TUser, error) {
	t := []t_user{}
	sql := fmt.Sprintf("SELECT * FROM user where nickname like '%%%s%%' and id > 10 ORDER BY id desc", nickname)
	err := Select(&t, sql)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	T := []*TUser{}
	for _, v := range t {
		item := &TUser{
			t_user:      v,
			AvatarFile:  Image_GetOtherAvatar(v.Avatar_id),
			AvatarAudit: 1,
			Age:		 GetAge(v.Birthday),
		}
		T = append(T, item)
	}

	return T, nil
}



type RankUser struct{
	TUser
	Num int
}
// 房间贡献榜
func GetSendUserList(roomid int,time string) ([]*RankUser, error) {
	list := []*RankUser{}

	sql := ""
	if time == "day" {
		// 日榜
		sql = "select SUM(a.coins) num,b.* from gift_log a,user b where a.user_id = b.id and date(cdate) = CURDATE() and scene_id=? and a.coins>0  group by user_id order by num desc"
	}else{
		// 本周
		sql = "select SUM(a.coins) num,b.* from gift_log a,user b where a.user_id = b.id and YEARWEEK(DATE_FORMAT(cdate,'%Y-%m-%d')) = YEARWEEK(NOW()) and scene_id=? and a.coins>0 group by user_id order by num desc"
	}

	err := Select(&list, sql,roomid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	for _, v := range list {
		v.AvatarFile = Image_GetOtherAvatar(v.Avatar_id)
		v.AvatarAudit = 1
		v.Age = GetAge(v.Birthday)
	}

	return list, nil
}

// 勿扰开关
func (t *TUser) SetNoDisturb(action int) error {
	_, err := Exec("update user set nodisturb=? where id=?",action,t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}



// 更新用户视频ID
func (t *TUser) UpdateVideoId(videoId int) error {
	_, err := Exec(`update user set video_id=? where id=?`, videoId, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}