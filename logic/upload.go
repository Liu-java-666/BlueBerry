package logic

import (
	"BlueBerry/database"
)

//上传图片
func UploadImage(user *database.TUser, file, filetype, usetype string, index int) interface{} {
	//插入图片表
	id, err := database.Image_Insert(user.Id, file, filetype, usetype)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
		UseType string
		Index int
		Id int
	}{
		true,
		usetype,
		index,
		id,
	}

	return result
}

// 上传声音
func UploadVoice(user *database.TUser, filename string, filetype string,second int) interface{} {
	//插入声音表
	id, err := database.Voice_Insert(user.Id, filename, filetype,second)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
		Id int
	}{
		true,
		id,
	}

	return result
}
//上传视频
func UploadVideo(user *database.TUser, file, filetype, cover, covertype, usetype string, rotation, index int) interface{} {
	//插入视频表
	id, err := database.Video_Insert(user.Id, file, filetype, cover, covertype, usetype, rotation)
	if err != nil {
		return ErrorResult("数据库异常")
	}

	//返回成功消息
	result := struct{
		Result bool
		UseType string
		Index int
		Id int
	}{
		true,
		usetype,
		index,
		id,
	}

	return result
}
