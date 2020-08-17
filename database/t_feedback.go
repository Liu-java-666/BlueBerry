package database

import (
	"github.com/wonderivan/logger"
)

type t_feedback struct {
	Id				int
	User_id			int
	Post_time		[]uint8
	Image_id		int
	Description		string
	Address			string
}

func Feedback_Insert(userid, imageid int, description, address string) error {
	_, err := Exec("INSERT INTO feedback(user_id,image_id,description,address) VALUES(?,?,?,?)",
		userid, imageid, description, address)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
