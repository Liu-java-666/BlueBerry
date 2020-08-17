package manage

import (
	"net/http"
	"strconv"
)

func OnIMImport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	min,_ := strconv.Atoi(r.Form.Get("min"))
	max,_ := strconv.Atoi(r.Form.Get("max"))

	result := IM_Import(min, max)
	w.Write([]byte(result))
}

func OnAvatarList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	page,_ := strconv.Atoi(r.Form.Get("page"))
	count,_ := strconv.Atoi(r.Form.Get("count"))

	result := Avatar_AuditList(page, count)
	w.Write([]byte(result))
}

func OnAvatarAudit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	id,_ := strconv.Atoi(r.Form.Get("id"))
	action,_ := strconv.Atoi(r.Form.Get("action"))

	result := Avatar_SetAudit(id, action)
	w.Write([]byte(result))
}

func OnDynamicList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	page,_ := strconv.Atoi(r.Form.Get("page"))
	count,_ := strconv.Atoi(r.Form.Get("count"))

	result := Dynamic_AuditList(page, count)
	w.Write([]byte(result))
}

func OnDynamicAudit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	id,_ := strconv.Atoi(r.Form.Get("id"))
	action,_ := strconv.Atoi(r.Form.Get("action"))

	result := Dynamic_SetAudit(id, action)
	w.Write([]byte(result))
}

func OnPhotoList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	page,_ := strconv.Atoi(r.Form.Get("page"))
	count,_ := strconv.Atoi(r.Form.Get("count"))

	result := Photo_AuditList(page, count)
	w.Write([]byte(result))
}

func OnPhotoAudit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json")             //返回数据格式是json

	r.ParseMultipartForm(1048576)
	id,_ := strconv.Atoi(r.Form.Get("id"))
	action,_ := strconv.Atoi(r.Form.Get("action"))

	result := Photo_SetAudit(id, action)
	w.Write([]byte(result))
}