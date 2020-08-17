package im

import (
	"BlueBerry/config"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

func PostData(servicename, command string, sendData []byte, call func(resultData string, err error, user_data interface{}), user_data interface{}) {
	imconfig := config.GetIMConfig()
	url := fmt.Sprintf("https://console.tim.qq.com/v4/%s/%s?sdkappid=%d&identifier=admin&usersig=%s&random=%d&contenttype=json",
		servicename, command, imconfig.AppId, imconfig.AdminSig, rand.Uint32())
	r,_ := http.NewRequest("POST", url, bytes.NewBuffer(sendData))
	r.Header.Set("Content-Type", "application/json")

	if call == nil {
		call = OnIMApiCallback
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		call("", err, user_data)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		call("", errors.New(res.Status), user_data)
		return
	}

	resultData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		call("", err, user_data)
		return
	}

	call(string(resultData), nil, user_data)
	return
}

func OnIMApiCallback(resultData string, err error, user_data interface{}) {
	if err != nil {
		logger.Error("错误：", err)
		return
	}

	logger.Debug(resultData)

	revData := make(map[string]interface{})
	err = json.Unmarshal([]byte(resultData), &revData)
	if err != nil {
		logger.Error("错误：%v,resultData=%v", err, resultData)
		return
	}

	//logger.Debug(revData)


}

//导入单个帐号
func AccountImport(userid int, call func(resultData string, err error, user_data interface{})) {
	sendData := struct {
		Identifier string
	}{
		fmt.Sprintf("%s%d", config.GetIMConfig().Pre, userid),
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}
	PostData("im_open_login_svc", "account_import", b, call, userid)
}

//导入多个帐号
func MultiAccountImport(userid []int) {
	sendData := struct {
		Accounts []string
	}{}

	for _, v := range userid {
		sendData.Accounts = append(sendData.Accounts, fmt.Sprintf("%s%d", config.GetIMConfig().Pre, v))
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}
	PostData("im_open_login_svc", "multiaccount_import", b, nil, nil)
}

//增加群组成员
func AddGroupMember(group, nickname string, userid int,AvatarFile string, call func(resultData string, err error, user_data interface{})) {
	logger.Info("IM 增加群组成员 START")
	type tagMember struct {
		Member_Account string
	}
	member := tagMember{fmt.Sprintf("%s%d", config.GetIMConfig().Pre, userid)}

	sendData := struct {
		GroupId string
		MemberList []tagMember
	}{
		group,
		[]tagMember{
			member,
		},
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}

	user_data := []string{group, member.Member_Account, nickname,AvatarFile,strconv.Itoa(userid)}
	PostData("group_open_http_svc", "add_group_member", b, call, user_data)
	logger.Info("IM 增加群组成员 END")
}

//删除群组成员
func DeleteGroupMember(group string, userid int) {
	member := fmt.Sprintf("%s%d", config.GetIMConfig().Pre, userid)

	sendData := struct {
		GroupId string
		MemberToDel_Account []string
	}{
		group,
		[]string{
			member,
		},
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}

	PostData("group_open_http_svc", "delete_group_member", b, nil, nil)
}

//获取群详细资料
func GroupInfo(group string, call func(resultData string, err error, user_data interface{})) {
	type tagResponseFilter struct {
		GroupBaseInfoFilter []string
		MemberInfoFilter []string
	}
	sendData := struct {
		GroupIdList []string
		ResponseFilter tagResponseFilter
	}{
		[]string{group},
		tagResponseFilter{
			[]string{"MemberNum"},
			[]string{"Member_Account","JoinTime"},
		},
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}
	//logger.Info("获取群详细资料 group_open_http_svc")
	PostData("group_open_http_svc", "get_group_info", b, call, group)
}

//在群组中发送系统通知
func SendGroupSysNotice(group string, Type int, msgdata interface{}) {
	var msg string
	if Type == 0 {
		msg = msgdata.(string)
	} else {
		tmp, err := json.Marshal(msgdata)
		if err != nil {
			logger.Error(err)
			return
		}
		msg = string(tmp)
	}

	data := struct{
		Type int
		Msg string
	}{
		Type,
		msg,
	}

	content, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		return
	}

	sendData := struct {
		GroupId string
		Content string
	}{
		group,
		string(content),
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug(string(b))
	PostData("group_open_http_svc", "send_group_system_notification", b, nil, nil)
}

//单发单聊消息
func SendMsg(userid, touserid, Type int, msgdata interface{}) {
	var msg string
	if Type == 0 {
		msg = msgdata.(string)
	} else {
		tmp, err := json.Marshal(msgdata)
		if err != nil {
			logger.Error(err)
			return
		}
		msg = string(tmp)
	}

	type tagMsgContent struct {
		Text string
	}

	type tagMsgBody struct {
		MsgType string
		MsgContent tagMsgContent
	}

	sendData := struct {
		SyncOtherMachine int
		From_Account string
		To_Account string
		MsgRandom int
		MsgBody []tagMsgBody
	}{
		2,
		fmt.Sprintf("%s%d", config.GetIMConfig().Pre, userid),
		fmt.Sprintf("%s%d", config.GetIMConfig().Pre, touserid),
		rand.Intn(99999),
		[]tagMsgBody{
			{
				"TIMTextElem",
				tagMsgContent{
					msg,
				},
			},
		},
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug(string(b))
	PostData("openim", "sendmsg", b, nil, nil)
}

// 批量删除群组成员
/**
	@groupid
	@accountlist 账号列表
 */
func BatchDelGroupMember(groupid string,accountlist []string){
	reqestData := struct {
		GroupId string
		Silence int// 是否静默删除（选填）
		MemberToDel_Account []string
	}{
		groupid,
		1,
		accountlist,
	}
	bytes, err := json.Marshal(reqestData)
	if err != nil {
		logger.Error(err)
		return
	}
	PostData("group_open_http_svc", "delete_group_member", bytes, func(resultData string, err error, user_data interface{}) {
		log:= fmt.Sprintf("批量删除群组成员 应答%s",resultData)
		logger.Info(log)
	}, groupid)
}

// IM批量删除用户
/**
@accountlist 账号列表
*/
func BatchDelAccount(accountlist []string){
	type Account struct{
		UserID string
	}
	list := []Account{}
	for _,v := range accountlist {
		account := Account{v}
		list = append(list,account)
	}

	reqestData := struct {
		DeleteItem []Account
	}{
		list,
	}
	bytes, err := json.Marshal(reqestData)
	if err != nil {
		logger.Error(err)
		return
	}
	PostData("im_open_login_svc", "account_delete", bytes, func(resultData string, err error, user_data interface{}) {
		log:= fmt.Sprintf("批量删除IM账号 应答%s",resultData)
		logger.Info(log)
	},nil)
}