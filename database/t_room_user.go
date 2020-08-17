package database

import (
	"github.com/wonderivan/logger"
)

type t_room_user struct {
	User_id		int
	Room_id 	int
	Im_group	string
}

func RoomUser_Get(userid int) (*t_room_user, error) {
	t := &t_room_user{}
	err := Get(t, "SELECT * FROM room_user WHERE user_id = ?",
		userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func RoomUser_Insert(userid, roomid int, imgroup string) {
	_, err := Exec("INSERT INTO room_user(user_id,room_id,im_group) VALUES(?,?,?)",
		userid, roomid, imgroup)
	if err != nil {
		logger.Error(err)
	}
}

func (t *t_room_user) SetRoom(roomid int, imgroup string) {
	_, err := Exec("UPDATE room_user SET room_id = ?, im_group = ? WHERE user_id = ?",
		roomid, imgroup, t.User_id)
	if err != nil {
		logger.Error(err)
	}
}


func RoomUser_GetUsers() ([]*t_room_user, error) {
	t := []*t_room_user{}
	err := Select(&t, "SELECT * FROM room_user WHERE room_id>0")
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func RoomUser_DelAll(){
	_, err := Exec("delete from room_user")
	if err != nil {
		logger.Error(err)
	}
}
