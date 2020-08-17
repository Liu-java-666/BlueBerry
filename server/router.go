package server

import (
	"BlueBerry/config"
	"BlueBerry/database"
	"BlueBerry/im"
	"BlueBerry/logic"
	"BlueBerry/manage"
	"BlueBerry/network"
	"BlueBerry/public"
	"encoding/base64"
	"fmt"

	"github.com/wonderivan/logger"
	"net/http"
	"strings"
)

//解密
func Decode(r *http.Request) {
	data := r.Form.Get(config.GetEncryptPre())

	logger.Info(data)

	if data != "" {
		bodys, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			logger.Error(err)
			return
		}

		param := public.AesDecryptECB(bodys, config.GetEncryptKey())

		paramstr := string(param)
		logger.Info("Decrypted data:", paramstr)
		paramlist := strings.Split(paramstr, "&")
		for _, v := range paramlist {
			kv := strings.Split(v, "=")
			if len(kv) == 2 {
				r.Form.Add(kv[0], kv[1])
			}
		}
	}
}

type MyHandler struct{}
func (mh MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	if r.Method != "POST" {
		network.SendResult(w, logic.ErrorResult("不支持的方法"))
		return
	}

	//body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))

	logger.Info(r.URL.Path, r.RemoteAddr)

	r.ParseMultipartForm(100*1024*1024)

	//解密
	if config.IsTest() == false {
		Decode(r)
	}

	switch r.URL.Path {
	case "/": network.OnIMCallback(w, r)//开启内网穿透后 测试服收到IM回调消息 会转发回本地 用于本地调试
	case "/imcallback": network.OnIMCallback(w, r)
	// 获取配置
	case "/blueberry/config": network.OnCheckVersion(w, r)
	// 获取验证码
	case "/blueberry/verity": network.OnGetCaptcha(w, r)
	// 登录/注册
	case "/blueberry/login-in": network.OnLogin(w, r)
	// 完善资料
	case "/blueberry/comp-data": network.OnSetInfo(w, r)
	// 编辑资料
	case "/blueberry/edit-data": network.OnEditInfo(w, r)
	// 我的菜单
	case "/blueberry/my-menu": network.OnGetMyMenu(w, r)
	// 我的信息
	case "/blueberry/my-info": network.OnGetMyInfo(w, r)
	// 用户详情
	case "/blueberry/xqing": network.OnGetUserDetail(w, r)
	// 用户名片
	case "/blueberry/ucard": network.OnGetUserCard(w, r)
	// 批量用户详情
	case "/blueberry/batch-xqing": network.OnGetUserInfoList(w, r)
	// 是否拉黑
	case "/blueberry/querybk": network.OnIsBlacklist(w, r)
	// 关注列表
	case "/blueberry/gzlist": network.OnGetFocusList(w, r)
	// 粉丝列表
	case "/blueberry/fslist": network.OnGetFansList(w, r)
	// 黑名单列表
	case "/blueberry/bklist": network.OnGetBlacklist(w, r)
	// 好友列表
	case "/blueberry/fdlist": network.OnGetFriendList(w, r)
	// 申请列表
	case "/blueberry/aplist": network.OnGetApplyList(w, r)
	// 收礼列表
	case "/blueberry/rec-gift": network.OnGetReceiveGiftList(w, r)
	// 我的钱包
	case "/blueberry/my-qianbao": network.OnGetMyWallet(w, r)
	// 创建订单
	case "/blueberry/create-pay": network.OnPayOrder(w, r)
	// 充值完成
	case "/blueberry/finish-pay": network.OnPayFinish(w, r)
	// 送礼物
	case "/blueberry/send-gift": network.OnSendGift(w, r)
	// 上传图片
	case "/blueberry/up-img": network.OnUploadImage(w, r)
	// 上传视频
	case "/blueberry/up-vdo": network.OnUploadVideo(w, r)
	// 上传声音
	case "/blueberry/up-voice": network.OnUploadVoice(w, r)

	// 1V1匹配
	case "/blueberry/1v1": network.OnGetMatchUser(w, r)
	// 拨打电话
	case "/blueberry/tocall" : network.OnCallUp(w, r)
	// 挂断电话
	case "/blueberry/tohung": network.OnHangUp(w, r)
	// 排行榜
	case "/blueberry/ranklist": network.OnGetRankList(w, r)
	// 关注 取消
	case "/blueberry/gz": network.OnFocus(w, r)
	// 拉黑 取消
	case "/blueberry/bk": network.OnBlacklist(w, r)
	// 举报
	case "/blueberry/jb": network.OnDenounce(w, r)
	// 动态列表
	case "/blueberry/dtlist": network.OnDynamicList(w, r)
	// 动态点赞
	case "/blueberry/dt-like": network.OnDynamicLike(w, r)
	// 动态评论列表
	case "/blueberry/dt-commentlist": network.OnDynamicCommentList(w, r)
	// 动态评论
	case "/blueberry/dt-comment": network.OnDynamicComment(w, r)
	// 动态评论点赞
	case "/blueberry/dt-likecomment": network.OnDynamicLikeComment(w, r)
	// 发布动态
	case "/blueberry/dt-post": network.OnDynamicPost(w, r)
	// 用户动态列表
	case "/blueberry/dt-ulist": network.OnDynamicUserList(w, r)

	// 删除动态
	case "/blueberry/dt-del": network.OnDynamicDelete(w, r)

	//发布作品
	case "/blueberry/wk-post": network.OnWorkPost(w, r)
	//视频列表
	case "/blueberry/sp-list":network.OnVideoList(w, r)
	// ==============================================
	//房间
	// ==============================================
	// 房间列表
	case "/blueberry/getRoomList": network.OnRoomList(w, r)
	// 进入房间
	case "/blueberry/toEnterRoom": network.OnRoomEnter(w, r)
	// 离开房间
	case "/blueberry/toLeaveRoom": network.OnRoomLeave(w, r)
	// 申请创建房间
	case "/blueberry/toCreateRoom": network.OnRoomCreate(w, r)
	// 房间点赞
	case "/blueberry/ToLikeRoom": network.OnRoomLike(w, r)
	// 申请上座
	case "/blueberry/ToSeatRoom": network.OnRoomSeat(w, r)
	// 句子列表
	case "/blueberry/sentence-list":network.OnSentenceList(w,r)

	// 搜索用户
	case "/blueberry/search-user":network.OnSearchUser(w,r)
	// 收藏房间
	case "/blueberry/doColRoom":network.OnColRoom(w,r)
	// 贡献榜
	case "/blueberry/getPayRankList":network.OnGetPayRankList(w,r)
	// 心动我的
	case "/blueberry/mdy-likelist":network.OnMyDyLikeList(w,r)
	// 评论我的
	case "/blueberry/mdy-commentlist":network.OnMyDyCommentList(w,r)
	// 勿扰开关
	case "/blueberry/nodisturb":network.OnMyNoDisturb(w,r)
	// 我的作品
	case "/blueberry/my-worklist":network.OnMyWorkList(w,r)
	// 作品详情
	case "/blueberry/work-detail":network.OnWorkDetail(w,r)
	// 作品心动列表
	case "/blueberry/work-likelist":network.OnDynamicLikeList(w,r)
	// 设置视频为详情主页
	case "/blueberry/xq-setVideo":network.OnUserSetVideo(w,r)
	// 主页点赞
	case "/blueberry/zy-like":network.OnUserDetailLike(w,r)
	//动态详情
	case "/blueberry/dt-detail":network.OnDynamicDetail(w,r)

	//声音文件列表 //该路由只是为了测试 无实际用途
	case "/blueberry/voicelist":network.OnGetVoiceList(w,r)
	//根据UserID查询声音 //该路由只是为了测试 无实际用途
	case "/blueberry/voicebyid":network.OnGetVoiceByUserID(w,r)

	default: network.SendResult(w, logic.ErrorResult("未知路径"))
	}
}

func SetRouter() {
	http.Handle("/", MyHandler{})
}

func SetRouter_Manager() {
	http.HandleFunc("/vpartymanager/im/import", manage.OnIMImport)
	http.HandleFunc("/vpartymanager/avatar/list", manage.OnAvatarList)
	http.HandleFunc("/vpartymanager/avatar/audit", manage.OnAvatarAudit)
	http.HandleFunc("/vpartymanager/dynamic/list", manage.OnDynamicList)
	http.HandleFunc("/vpartymanager/dynamic/audit", manage.OnDynamicAudit)
	http.HandleFunc("/vpartymanager/photo/list", manage.OnPhotoList)
	http.HandleFunc("/vpartymanager/photo/audit", manage.OnPhotoAudit)
}

//初始化操作
func SetInit(){
	// ================================================
	// 将房间中的用户全部退群
	// ================================================
	// 获取房间用户列表
	logger.Error("开始清理房间中的用户")
	roomuserlist,err := database.RoomUser_GetUsers()
	if err != nil {
		logger.Error("获取房间列表 数据库异常")
		return
	}
	//房间中的用户强制退群
	for _,v := range roomuserlist {
		account := fmt.Sprintf("%s%d", config.GetIMConfig().Pre, v.User_id )

		im.BatchDelGroupMember(v.Im_group,[]string{account})
	}
	//清空房间用户表
	database.RoomUser_DelAll()


	// ================================================
	// 删除IM账号 专业版无法删除 该接口仅支持体验版
	// ================================================
	//
	//imconfig := config.GetIMConfig()
	//logger.Error("开始清理IM用户")
	//account:=[]string{}
	////每次最多请求100个
	//for i:=1;i<100;i++ {
	//	account = append(account,fmt.Sprintf("%s%d", imconfig.Pre, i))
	//}
	//im.BatchDelAccount(account)

}
