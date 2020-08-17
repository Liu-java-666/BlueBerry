package manage

import (
	"BlueBerry/database"
	"BlueBerry/im"
	"BlueBerry/logic"
	"BlueBerry/public"
	"fmt"
	"os"
)

func IM_Import(min, max int) string {
	if max == 0 {
		max = min
	}

	if min <= 0 || min > max {
		return "参数错误"
	}

	userid := []int{}
	for i:=min; i<=max; i++ {
		userid = append(userid, i)
	}

	im.MultiAccountImport(userid)

	return "操作完成"
}

func Avatar_AuditList(page, count int) string {
	if page < 0 {
		page = 0
	}

	if count <= 0 {
		count = 10
	} else if count > 50 {
		count = 50
	}

	tlist, err := database.Image_AuditList(page*count, count, "avatar")
	if err != nil {
		return err.Error()
	}

	List := ""
	for _, v := range tlist {
		List += fmt.Sprintf("Id:%d, File:%s\n",
			v.Id, v.File_name)
	}

	return List
}

func Avatar_SetAudit(id, audit int) string {
	if id <= 0 {
		return "头像id错误"
	}

	if audit == 0 {
		audit = 1
	} else if audit != 1 && audit != -1 {
		return "审核参数错误"
	}

	tavatar, err := database.Image_Get(id)
	if err != nil {
		return err.Error()
	}

	if tavatar == nil {
		return "头像文件不存在"
	}

	if tavatar.Is_audit != 0 {
		return "非审核中状态"
	}

	err = tavatar.SetAudit(audit)
	if err != nil {
		return err.Error()
	}

	if audit == 1 {
		database.User_SetAvatar(tavatar.User_id, id)

		//拒绝该用户其他头像
		tlist, _ := database.Image_UserList(tavatar.User_id, id, "avatar")
		for _, v := range tlist {
			ta, _ := database.Image_Get(v.Id)
			if ta != nil {
				ta.SetAudit(-1)
				filepath := logic.MakeImageUrl(ta.File_name)
				os.Remove(filepath)
			}
		}
	} else {
		filepath := logic.MakeImageUrl(tavatar.File_name)
		os.Remove(filepath)
	}

	return "设置成功"
}


func Dynamic_AuditList(page, count int) string {
	if page < 0 {
		page = 0
	}

	if count <= 0 {
		count = 10
	} else if count > 50 {
		count = 50
	}

	tlist, err := database.Dynamic_AuditList(page*count, count)
	if err != nil {
		return err.Error()
	}

	List := ""
	for _, v := range tlist {
		List += fmt.Sprintf("Id:%d, Description:%s, Topic:%s, FileType:%s, FileList:%s\n",
			v.Id, v.Description, v.Topic, v.Filetype, v.Filelist)
	}

	return List
}

func Dynamic_SetAudit(id, audit int) string {
	if id <= 0 {
		return "动态id错误"
	}

	if audit == 0 {
		audit = 1
	} else if audit != 1 && audit != -1 {
		return "审核参数错误"
	}

	t, err := database.Dynamic_Get(id)
	if err != nil {
		return err.Error()
	}

	if t == nil {
		return "动态不存在"
	}

	if t.Is_audit != 0 {
		return "非审核中状态"
	}

	err = t.SetAudit(audit)
	if err != nil {
		return err.Error()
	}

	if audit == -1 {
		fidlist := public.GetFileIdList(t.Filelist)
		for _, v := range fidlist {
			if t.Filetype == "video" {
				tv, _ := database.Video_Get(v)
				if tv != nil {
					tv.Delete()
					filepath := logic.MakeVideoPath(tv.File_name)
					os.Remove(filepath)
					filepath = logic.MakeVideoPath(tv.Cover_name)
					os.Remove(filepath)
				}
			} else if t.Filetype == "image" {
				ti, _ := database.Image_Get(v)
				if ti != nil {
					ti.Delete()
					filepath := logic.MakeImagePath(ti.File_name)
					os.Remove(filepath)
				}
			}else if t.Filetype == "voice"{
				vi,_ := database.Voice_Get(v)
				if vi != nil{
					vi.Delete()
					filepath := logic.MakeVoicePath(vi.File_name)
					os.Remove(filepath)
				}
			}
		}
	}

	return "设置成功"
}

func Photo_AuditList(page, count int) string {
	if page < 0 {
		page = 0
	}

	if count <= 0 {
		count = 10
	} else if count > 50 {
		count = 50
	}

	tlist, err := database.PhotoList_AuditList(page*count, count)
	if err != nil {
		return err.Error()
	}

	List := ""
	for _, v := range tlist {
		List += fmt.Sprintf("Id:%d, UserId:%d, Photolist:%s\n",
			v.Id, v.User_id, v.Photolist)
	}

	return List
}

func Photo_SetAudit(id, audit int) string {
	if id <= 0 {
		return "官方认证id错误"
	}

	if audit == 0 {
		audit = 1
	} else if audit != 1 && audit != -1 {
		return "审核参数错误"
	}

	t, err := database.PhotoList_Get(id)
	if err != nil {
		return err.Error()
	}

	if t == nil {
		return "认证数据不存在"
	}

	if t.Is_audit != 0 {
		return "非审核中状态"
	}

	err = t.SetAudit(audit)
	if err != nil {
		return err.Error()
	}

	if audit == 1 {
		database.User_SetPhotoList(t.User_id, id)

		pidlist := public.GetFileIdList(t.Photolist)
		for _, v := range pidlist {
			tp, _ := database.Image_Get(v)
			if tp != nil {
				tp.SetAudit(1)
			}
		}

		//拒绝该用户其他认证
		tlist, _ := database.PhotoList_UserList(t.User_id, id)
		for _, v := range tlist {
			to, _ := database.PhotoList_Get(v.Id)
			if to != nil {
				to.SetAudit(-1)
			}
		}

		//删除该用户的其他照片
		tplist, _ := database.Image_UnusedList(t.User_id, t.Photolist, "photo")
		for _, v := range tplist {
			tp, _ := database.Image_Get(v.Id)
			if tp != nil {
				tp.SetAudit(-1)
				filepath := logic.MakeImagePath(tp.File_name)
				os.Remove(filepath)
			}
		}
	} else {
		//删除该用户所有待审照片
		tplist, _ := database.Image_UserAuditList(t.User_id, "photo")
		for _, v := range tplist {
			tp, _ := database.Image_Get(v.Id)
			if tp != nil {
				tp.SetAudit(-1)
				filepath := logic.MakeImagePath(tp.File_name)
				os.Remove(filepath)
			}
		}
	}

	return "设置成功"
}