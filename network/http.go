package network

import (
	"BlueBerry/config"
	"BlueBerry/logic"
	"BlueBerry/public"
	"encoding/base64"
	"encoding/json"
	"github.com/wonderivan/logger"
	"io"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
)

func Encrypt(data []byte) []byte {
	body := public.AesEncryptECB(data, config.GetEncryptKey())
	bodys := base64.StdEncoding.EncodeToString(body)
	return []byte(bodys)
}

func SendResult(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		return
	}

	bodys := b
	if config.IsTest() == false {
		//加密
		bodys = Encrypt(b)
	}

	_, err = w.Write(bodys)
	if err != nil {
		logger.Error(err)
		return
	}
}

//IM第三方回调
func OnIMCallback(w http.ResponseWriter, r *http.Request) {
	//logger.Info("OnIMCallback IM回调 START")
	//转发IM回调消息到本地调试
	if config.GetAgentTestConfig()==1 {
		localIMCallBackUrl := config.GetAgentDomainConfig()
		contentType := "application/json"
		resp,err := http.Post(localIMCallBackUrl,contentType,r.Body)
		if err != nil {
			logger.Error(err)
			return
		}
		defer resp.Body.Close()
		return
	}


	resultData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(string(resultData))

	revData := make(map[string]interface{})
	err = json.Unmarshal(resultData, &revData)
	if err != nil {
		logger.Error(err)
		return
	}

	CallbackCommand := revData["CallbackCommand"].(string)
	if CallbackCommand != "Group.CallbackAfterNewMemberJoin" && CallbackCommand != "Group.CallbackAfterMemberExit" {
		return
	}

	GroupId := revData["GroupId"].(string)
	logger.Debug(GroupId)
	//logger.Info("OnIMCallback IM回调 END")
	go logic.RoomUpdateOnline(GroupId)
}

//检查版本
func OnCheckVersion(w http.ResponseWriter, r *http.Request) {
	result := struct {
		Result bool `json:"result"`
		HasUpdate int `json:"hasUpdate"`
	}{
		true,
		0,
	}
	SendResult(w, result)
}

//获取验证码
func OnGetCaptcha(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["phone"] = r.Form.Get("phone")

	retchan := make(chan interface{})
	logic.PushQue("GetCaptcha", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//登录/注册
func OnLogin(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["phone"] = r.Form.Get("phone")
	param["captcha"] = r.Form.Get("captcha")

	param["ip"] = r.RemoteAddr
	index := strings.IndexRune(r.RemoteAddr, ':')
	if index >= 0 {
		param["ip"] = r.RemoteAddr[:index]
	}

	retchan := make(chan interface{})
	logic.PushQue("Login", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

// 完善资料
func OnSetInfo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	// 昵称
	param["nickname"] = r.Form.Get("nickname")
	// 签名
	param["signature"] = r.Form.Get("signature")
	// 性别
	param["sex"], _ = strconv.Atoi(r.Form.Get("sex"))
	// 视频ID
	param["videoId"],_ = strconv.Atoi(r.Form.Get("videoId"))
	retchan := make(chan interface{})
	logic.PushQue("SetInfo", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//编辑信息
func OnEditInfo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["nickname"] = r.Form.Get("nickname")
	param["signature"] = r.Form.Get("signature")
	retchan := make(chan interface{})
	logic.PushQue("EditInfo", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//我的菜单
func OnGetMyMenu(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMyMenu", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//我的资料
func OnGetMyInfo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMyInfo", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//用户详情
func OnGetUserDetail(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("GetUserDetail", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//用户名片
func OnGetUserCard(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("GetUserCard", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//批量用户详情
func OnGetUserInfoList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	slist := r.Form["touseridlist[]"]
	ilist := []int{}
	for _, v := range slist {
		id, _ := strconv.Atoi(v)
		ilist = append(ilist, id)
	}
	param["touseridlist"] = ilist

	retchan := make(chan interface{})
	logic.PushQue("GetUserInfoList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//是否拉黑
func OnIsBlacklist(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("IsBlacklist", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取关注列表
func OnGetFocusList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetFocusList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取粉丝列表
func OnGetFansList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetFansList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取黑名单
func OnGetBlacklist(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetBlacklist", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取好友列表
func OnGetFriendList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetFriendList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取好友申请列表
func OnGetApplyList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetApplyList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取收礼列表
func OnGetReceiveGiftList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetReceiveGiftList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//我的钱包
func OnGetMyWallet(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMyWallet", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//提交充值订单
func OnPayOrder(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["money"], _ = strconv.Atoi(r.Form.Get("money"))

	retchan := make(chan interface{})
	logic.PushQue("PayOrder", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//完成充值
func OnPayFinish(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["orderid"] = r.Form.Get("orderid")
	param["status"], _ = strconv.Atoi(r.Form.Get("status"))

	retchan := make(chan interface{})
	logic.PushQue("PayFinish", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//送礼物
func OnSendGift(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["scene"] = r.Form.Get("scene")
	param["sceneid"], _ = strconv.Atoi(r.Form.Get("sceneid"))
	param["giftid"], _ = strconv.Atoi(r.Form.Get("giftid"))

	retchan := make(chan interface{})
	logic.PushQue("SendGift", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//上传图片
func OnUploadImage(w http.ResponseWriter, r *http.Request) {
	srcfile, srcfileheader, err := r.FormFile("file")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取文件数据错误"))
		return
	}
	defer srcfile.Close()

	if srcfileheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取上传文件错误：无法读取文件大小"))
		return
	} else if srcfileheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取上传文件错误：文件大小超出30M"))
		return
	}

	filetype := r.Form.Get("filetype")
	if filetype != "jpg" && filetype != "png" {
		SendResult(w, logic.ErrorResult("文件格式错误"))
		return
	}

	//检查用户
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	retchan := make(chan interface{})
	logic.PushQue("CheckUser", param, retchan)
	result := <- retchan
	ck := result.(logic.CheckUserResult)
	if !ck.Result {
		SendResult(w, ck.Error)
		return
	}

	//保存文件
	param["usetype"] = r.Form.Get("usetype")
	param["index"], _ = strconv.Atoi(r.Form.Get("index"))
	filename := public.MakeFileName(param["userid"].(int), param["index"].(int), filetype, param["usetype"].(string))
	filepath := logic.MakeImagePath(filename)
	file, err := os.Create(filepath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(file, srcfile)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	file.Close()

	//更新信息
	param["user"] = ck.User
	param["file"] = filename
	param["filetype"] = filetype

	logic.PushQue("UploadImage", param, retchan)
	result = <- retchan

	SendResult(w, result)
}


//上传声音
func OnUploadVoice(w http.ResponseWriter, r *http.Request) {
	srcfile, srcfileheader, err := r.FormFile("file")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取文件数据错误"))
		return
	}
	defer srcfile.Close()

	if srcfileheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取上传文件错误：无法读取文件大小"))
		return
	} else if srcfileheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取上传文件错误：文件大小超出30M"))
		return
	}

	//检查用户
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("CheckUser", param, retchan)
	result := <- retchan
	ck := result.(logic.CheckUserResult)
	if !ck.Result {
		SendResult(w, ck.Error)
		return
	}

	//保存文件
	filename := public.MakeFileName(param["userid"].(int), 1, "mp3", "dy")
	filepath := logic.MakeVoicePath(filename)
	file, err := os.Create(filepath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(file, srcfile)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	file.Close()

	//更新信息
	param["user"] = ck.User
	param["file"] = filename
	param["filetype"] = "mp3"
	param["second"],_ = strconv.Atoi(r.Form.Get("second"))
	logic.PushQue("UploadVoice", param, retchan)
	result = <- retchan

	SendResult(w, result)
}



//上传视频
func OnUploadVideo(w http.ResponseWriter, r *http.Request) {
	//获取封面
	srccover, srccoverheader, err := r.FormFile("cover")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取封面文件数据错误"))
		return
	}
	defer srccover.Close()

	if srccoverheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取封面文件错误：无法读取文件大小"))
		return
	} else if srccoverheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取封面文件错误：文件大小超出30M"))
		return
	}

	covertype := r.Form.Get("covertype")
	if covertype != "jpg" && covertype != "png" {
		SendResult(w, logic.ErrorResult("封面文件格式错误"))
		return
	}

	//获取视频
	srcfile, srcfileheader, err := r.FormFile("file")
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("读取视频文件数据错误"))
		return
	}
	defer srcfile.Close()

	if srcfileheader.Size <= 0 {
		SendResult(w, logic.ErrorResult("获取视频文件错误：无法读取文件大小"))
		return
	} else if srcfileheader.Size > 30*1024*1024 {
		SendResult(w, logic.ErrorResult("获取视频文件错误：文件大小超出30M"))
		return
	}

	filetype := r.Form.Get("filetype")
	if filetype != "mp4" && filetype != "rmvb" {
		SendResult(w, logic.ErrorResult("视频文件格式错误"))
		return
	}

	//检查用户
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	retchan := make(chan interface{})
	logic.PushQue("CheckUser", param, retchan)
	result := <- retchan
	ck := result.(logic.CheckUserResult)
	if !ck.Result {
		SendResult(w, ck.Error)
		return
	}

	//保存封面
	param["usetype"] = r.Form.Get("usetype")
	param["index"], _ = strconv.Atoi(r.Form.Get("index"))
	covername := public.MakeFileName(param["userid"].(int), param["index"].(int), covertype, param["usetype"].(string))
	coverpath := logic.MakeVideoPath(covername)
	cover, err := os.Create(coverpath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(cover, srccover)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	cover.Close()

	//保存视频
	filename := public.ChangeFileType(covername, filetype)
	filepath := logic.MakeVideoPath(filename)
	file, err := os.Create(filepath)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("创建文件失败"))
		return
	}

	_, err = io.Copy(file, srcfile)
	if err != nil {
		logger.Error(err)
		SendResult(w, logic.ErrorResult("写文件失败"))
		return
	}
	file.Close()

	//更新信息
	param["user"] = ck.User
	param["file"] = filename
	param["filetype"] = filetype
	param["cover"] = covername
	param["covertype"] = covertype
	param["rotation"], _ = strconv.Atoi(r.Form.Get("rotation"))
	logic.PushQue("UploadVideo", param, retchan)
	result = <- retchan

	SendResult(w, result)
}

//1V1
func OnGetMatchUser(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("GetMatchUser", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//打电话
func OnCallUp(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))

	retchan := make(chan interface{})
	logic.PushQue("CallUp", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//挂电话
func OnHangUp(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")

	retchan := make(chan interface{})
	logic.PushQue("HangUp", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//排行榜
func OnGetRankList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["tag"] = r.Form.Get("tag")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("GetRankList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//关注/取消关注
func OnFocus(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("Focus", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//拉黑/取消拉黑
func OnBlacklist(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("Blacklist", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//举报
func OnDenounce(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["type"] = r.Form.Get("type")
	param["content"] = r.Form.Get("content")

	retchan := make(chan interface{})
	logic.PushQue("Denounce", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取动态列表
func OnDynamicList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["filetype"] = r.Form.Get("filetype")
	param["tag"] = r.Form.Get("tag")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}
//获取动态详情
func OnDynamicDetail(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"],_ = strconv.Atoi(r.Form.Get("dynamicid"))
	retchan := make(chan interface{})
	logic.PushQue("OnDynamicDetail", param, retchan)
	result := <- retchan
	SendResult(w, result)
}


// 视频列表
func OnVideoList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))
	retchan := make(chan interface{})
	logic.PushQue("DynamicList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}




//点赞/取消点赞动态
func OnDynamicLike(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))
	param["like_type"] = r.Form.Get("like_type")
	retchan := make(chan interface{})
	logic.PushQue("DynamicLike", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取评论列表
func OnDynamicCommentList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicCommentList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}
//动态点赞列表
func OnDynamicLikeList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("OnDynamicLikeList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}
//评论动态
func OnDynamicComment(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	param["content"] = r.Form.Get("content")

	retchan := make(chan interface{})
	logic.PushQue("DynamicComment", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//点赞/取消点赞评论
func OnDynamicLikeComment(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["commentid"], _ = strconv.Atoi(r.Form.Get("commentid"))
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicLikeComment", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//发布动态
func OnDynamicPost(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["description"] = r.Form.Get("description")
	param["filetype"] = r.Form.Get("filetype")
	slist := r.Form["filelist[]"]
	ilist := []int{}
	for _, v := range slist {
		id, _ := strconv.Atoi(v)
		ilist = append(ilist, id)
	}
	param["filelist"] = ilist

	retchan := make(chan interface{})
	logic.PushQue("DynamicPost", param, retchan)
	result := <- retchan
	SendResult(w, result)
}
// 发布作品
func OnWorkPost(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["description"] = r.Form.Get("description")
	param["sentenceid"], _ =  strconv.Atoi(r.Form.Get("sentenceid"))
	param["filetype"] = "voice"
	param["fileid"],_ = strconv.Atoi(r.Form.Get("fileid"))
	param["second"],_ = strconv.Atoi(r.Form.Get("second"))
	retchan := make(chan interface{})
	logic.PushQue("WorkPost", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//用户动态列表
func OnDynamicUserList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	param["filetype"] = r.Form.Get("filetype")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicUserList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}
// 我的作品
func OnMyWorkList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))
	retchan := make(chan interface{})
	logic.PushQue("OnMyWorkList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

// 作品详情
func OnWorkDetail(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	retchan := make(chan interface{})
	logic.PushQue("OnWorkDetail", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//删除动态
func OnDynamicDelete(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))

	retchan := make(chan interface{})
	logic.PushQue("DynamicDelete", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//获取房间列表
func OnRoomList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))
	retchan := make(chan interface{})
	logic.PushQue("RoomList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//句子列表
func OnSentenceList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	//句子类型
	param["stype"],_ = strconv.Atoi(r.Form.Get("stype"))
	//页数
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))
	retchan := make(chan interface{})
	logic.PushQue("SentenceList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//进入房间
func OnRoomEnter(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomEnter", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//退出房间
func OnRoomLeave(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomLeave", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//申请创建房间
func OnRoomCreate(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomtype"], _ = strconv.Atoi(r.Form.Get("roomtype"))

	retchan := make(chan interface{})
	logic.PushQue("RoomCreate", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//点赞房间
func OnRoomLike(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomLike", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

//申请上座
func OnRoomSeat(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"], _ = strconv.Atoi(r.Form.Get("roomid"))

	retchan := make(chan interface{})
	logic.PushQue("RoomSeat", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

// 搜索用户
func OnSearchUser(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["nickname"] = r.Form.Get("nickname")
	retchan := make(chan interface{})
	logic.PushQue("SearchUser", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

// 收藏房间
func OnColRoom(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"],_ = strconv.Atoi(r.Form.Get("roomid"))
	param["action"],_ = strconv.Atoi(r.Form.Get("action"))
	retchan := make(chan interface{})
	logic.PushQue("OnColRoom", param, retchan)
	result := <- retchan
	SendResult(w, result)
}

// 送礼贡献榜
func OnGetPayRankList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["roomid"],_ = strconv.Atoi(r.Form.Get("roomid"))
	param["time"] = r.Form.Get("time")
	retchan := make(chan interface{})
	logic.PushQue("OnGetPayRankList", param, retchan)
	result := <- retchan
	SendResult(w, result)

}

// 心动我的
func OnMyDyLikeList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	//页数
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))
	retchan := make(chan interface{})
	logic.PushQue("OnMyDyLikeList", param, retchan)
	result := <- retchan
	SendResult(w, result)

}

//评论我的
func OnMyDyCommentList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	//页数
	param["page"], _ = strconv.Atoi(r.Form.Get("page"))
	retchan := make(chan interface{})
	logic.PushQue("OnMyDyCommentList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}



// 勿扰开关
func OnMyNoDisturb(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["action"], _ = strconv.Atoi(r.Form.Get("action"))
	retchan := make(chan interface{})
	logic.PushQue("OnMyNoDisturb", param, retchan)
	result := <- retchan
	SendResult(w, result)

}

// 设置视频为详情主页
func OnUserSetVideo(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["dynamicid"], _ = strconv.Atoi(r.Form.Get("dynamicid"))
	retchan := make(chan interface{})
	logic.PushQue("OnUserSetVideo", param, retchan)
	result := <- retchan
	SendResult(w, result)

}


//主页点赞
func OnUserDetailLike(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	param["userid"], _ = strconv.Atoi(r.Form.Get("userid"))
	param["userkey"] = r.Form.Get("userkey")
	param["touserid"], _ = strconv.Atoi(r.Form.Get("touserid"))
	retchan := make(chan interface{})
	logic.PushQue("OnUserDetailLike", param, retchan)
	result := <- retchan
	SendResult(w, result)

}



//仅供测试
func OnGetVoiceList(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	retchan := make(chan interface{})
	logic.PushQue("OnGetVoiceList", param, retchan)
	result := <- retchan
	SendResult(w, result)
}
func OnGetVoiceByUserID(w http.ResponseWriter, r *http.Request) {
	param := make(map[string]interface{})
	retchan := make(chan interface{})
	param["id"], _ = strconv.Atoi(r.Form.Get("id"))

	logic.PushQue("OnGetVoiceByUserID", param, retchan)
	result := <- retchan
	SendResult(w, result)
}


