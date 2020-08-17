package logic

import (
	"BlueBerry/database"
	"BlueBerry/public"
	"strconv"
)

//该函数只是为了测试 无实际用途
//声音列表

func  OnGetVoiceList() interface{}{
	list,err := database.Voice_GetAll()
	if err != nil{
		return ErrorResult("数据库异常")
	}

	type Voice struct{
		ID	int
		UserID string
		PostTime string
		FileName string
	}

	var r []Voice
	for _,v := range list {
		intID,_ :=  strconv.Atoi(v["id"].(string)) //转INT
		t := Voice{
			ID:intID,
			UserID:   v["user_id"].(string),
			PostTime: v["post_time"].(string),
			FileName: v["file_name"].(string),
		}
		r = append(r,t)
	}


	result := struct {
		Data []Voice
	}{
		r,
	}
	return result
}


func  OnGetVoiceByUserID(id int) interface{}{
	v,err := database.Voice_GetByUserID(id)
	if err != nil{
		return ErrorResult("数据库异常")
	}
	if v==nil{
		return ErrorResult("没有数据")
	}
	type Voice struct{
		ID	int
		UserID int
		PostTime int
		FileName string
	}
	intID,_ :=  strconv.Atoi(v["id"].(string)) //转INT
	intUserID,_:=  strconv.Atoi(v["user_id"].(string))
	t := Voice{
		ID:intID,
		UserID: intUserID,
		PostTime: public.StrToTimestamp(v["post_time"].(string)),
		FileName: v["file_name"].(string),
	}

	result := struct {
		Data Voice
	}{
		t,
	}
	return result
}


