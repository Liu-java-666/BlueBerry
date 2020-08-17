package database

import (
	"github.com/wonderivan/logger"
)

type t_room_seat struct {
	Id			int
	User_id		int
	Room_id 	int
}

func RoomSeat_Get(userid, roomid int) bool {
	t := &t_room_seat{}
	err := Get(t, "SELECT * FROM room_seat WHERE room_id = ? and user_id = ?",
		roomid, userid)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false
		}

		logger.Error(err)
		return false
	}

	return true
}

func RoomSeat_Insert(userid, roomid int) error {
	_, err := Exec("INSERT INTO room_seat(user_id,room_id) VALUES(?,?)",
		userid, roomid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func RoomSeat_Delete(userid, roomid int) error {
	_, err := Exec("DELETE FROM room_seat WHERE room_id = ? and user_id = ?",
		roomid, userid)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}