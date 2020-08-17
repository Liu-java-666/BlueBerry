package database

import "github.com/wonderivan/logger"
//用户收藏房间
type TUserColRoom struct{
	ID int
	User_ID int
	Room_ID int
	Create_Time int64
}

// 新增
func AddUserColRoom(userid int,roomid int,createTime int64) error{
	sql := "insert into user_col_room(user_id,room_id,create_time) values(?,?,?)"
	_, err := Exec(sql,userid, roomid, createTime)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

//删除
func DelUserColRoom(id int) error{
	sql := "delete from user_col_room where id = ?"
	_, err := Exec(sql,id)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

//查询
func QueryUserColRoomByID(id int) (*TUserColRoom, error){
	sql := "select * from user_col_room where id=?"
	t := TUserColRoom{}
	err := Get(&t, sql,id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &t, nil
}
//根据UserID查询
func GetUserColRoomByUserID(userid int)(*TUserColRoom, error){
	sql := "select * from user_col_room where user_id=?"
	t := TUserColRoom{}
	err := Get(&t, sql,userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &t, nil
}

// 根据用户ID查询列表
func QueryUserColRoomListByUserID(userid int) ([]*TUserColRoom, error){
	sql := "select * from user_col_room where user_id=? order by create_time desc"
	t := []*TUserColRoom{}
	err := Select(&t, sql,userid)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return t, nil
}
// 根据用户ID 房间ID查询收藏记录
func GetUserColRoomByUserIDAndRoomID(userid int , roomid int)(*TUserColRoom, error){
	sql := "select * from user_col_room where user_id=? and room_id = ?"
	t := TUserColRoom{}
	err := Get(&t, sql,userid,roomid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		logger.Error(err)
		return nil, err
	}
	return &t, nil
}

