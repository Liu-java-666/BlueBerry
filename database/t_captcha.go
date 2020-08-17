package database

import (
	"BlueBerry/public"
	"github.com/wonderivan/logger"
)

type t_captcha struct {
	Id				int
	Captcha			string
	Phone_number	string
	Generation_time	int
	Expire_time		int
	Is_used			int
}

func Captcha_Get(phone string) (*t_captcha, error) {
	t := &t_captcha{}
	err := Get(t, "SELECT * FROM captcha WHERE phone_number = ?", phone)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}

		logger.Error(err)
		return nil, err
	}

	return t, nil
}

func Captcha_Insert(phone, captcha string) error {
	timenow := public.GetNowTimestamp()
	expiretime := timenow + 300

	_, err := Exec("INSERT INTO captcha(phone_number,`captcha`,generation_time,expire_time) VALUES(?,?,?,?)",
		phone, captcha, timenow, expiretime)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_captcha) UpdateCaptcha(captcha string) error {
	timenow := public.GetNowTimestamp()
	expiretime := timenow + 300

	_, err := Exec("UPDATE captcha SET `captcha` = ?, generation_time = ?, expire_time= ?, is_used = 0 WHERE `id` = ?",
		captcha, timenow, expiretime, t.Id)
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (t *t_captcha) SetUsed() {
	_, err := Exec("UPDATE captcha SET is_used = 1 WHERE `id` = ?", t.Id)
	if err != nil {
		logger.Error(err)
		return
	}
}