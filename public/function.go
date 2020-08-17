package public

import (
	"BlueBerry/config"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/wonderivan/logger"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//检查手机号格式
func ValidPhoneNumber(phone string) bool {
	ok, _ := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", phone)
	return ok
}

var fileTypeMap sync.Map

func init() {
	fileTypeMap.Store("ffd8ffe", "jpg")  //JPEG (jpg)
	fileTypeMap.Store("89504e47", "png")  //PNG (png)
	fileTypeMap.Store("47494638396126026f01", "gif")  //GIF (gif)
	fileTypeMap.Store("49492a00227105008037", "tif")  //TIFF (tif)
	fileTypeMap.Store("424d228c010000000000", "bmp")  //16色位图(bmp)
	fileTypeMap.Store("424d8240090000000000", "bmp")  //24位位图(bmp)
	fileTypeMap.Store("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	fileTypeMap.Store("41433130313500000000", "dwg")  //CAD (dwg)
	fileTypeMap.Store("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c21646f637479706520", "htm")  //HTM (htm)
	fileTypeMap.Store("48544d4c207b0d0a0942", "css")  //css
	fileTypeMap.Store("696b2e71623d696b2e71", "js")   //js
	fileTypeMap.Store("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	fileTypeMap.Store("38425053000100000000", "psd")  //Photoshop (psd)
	fileTypeMap.Store("46726f6d3a203d3f6762", "eml")  //Email [Outlook Express 6] (eml)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "doc")  //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "vsd")  //Visio 绘图
	fileTypeMap.Store("5374616E64617264204A", "mdb")  //MS Access (mdb)
	fileTypeMap.Store("252150532D41646F6265", "ps")
	fileTypeMap.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", "flv")  //flv与f4v相同
	fileTypeMap.Store("00000020667479706", "mp4")
	fileTypeMap.Store("0000001", "mp4")
	fileTypeMap.Store("49443303000000002176", "mp3")
	fileTypeMap.Store("000001ba210001000180", "mpg") //
	fileTypeMap.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", "wav") //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", "avi")
	fileTypeMap.Store("4d546864000000060001", "mid") //MIDI (mid)
	fileTypeMap.Store("504b0304140000000800", "zip")
	fileTypeMap.Store("526172211a0700cf9073", "rar")
	fileTypeMap.Store("235468697320636f6e66", "ini")
	fileTypeMap.Store("504b03040a0000000000", "jar")
	fileTypeMap.Store("4d5a9000030000000400", "exe")        //可执行文件
	fileTypeMap.Store("3c25402070616765206c", "jsp")        //jsp文件
	fileTypeMap.Store("4d616e69666573742d56", "mf")         //MF文件
	fileTypeMap.Store("3c3f786d6c2076657273", "xml")        //xml文件
	fileTypeMap.Store("494e5345525420494e54", "sql")        //xml文件
	fileTypeMap.Store("7061636b616765207765", "java")       //java文件
	fileTypeMap.Store("406563686f206f66660d", "bat")        //bat文件
	fileTypeMap.Store("1f8b0800000000000000", "gz")         //gz文件
	fileTypeMap.Store("6c6f67346a2e726f6f74", "properties") //bat文件
	fileTypeMap.Store("cafebabe0000002e0041", "class")      //bat文件
	fileTypeMap.Store("49545346030000006000", "chm")        //bat文件
	fileTypeMap.Store("04000000010000001300", "mxp")        //bat文件
	fileTypeMap.Store("504b0304140006000800", "docx")       //docx文件
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "wps")        //WPS文字wps、表格et、演示dps都是一样的
	fileTypeMap.Store("6431303a637265617465", "torrent")
	fileTypeMap.Store("6D6F6F76", "mov")         //Quicktime (mov)
	fileTypeMap.Store("FF575043", "wpd")         //WordPerfect (wpd)
	fileTypeMap.Store("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	fileTypeMap.Store("2142444E", "pst")         //Outlook (pst)
	fileTypeMap.Store("AC9EBD8F", "qdf")         //Quicken (qdf)
	fileTypeMap.Store("E3828596", "pwl")         //Windows Password (pwl)
	fileTypeMap.Store("2E7261FD", "ram")         //Real Audio (ram)
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) string {
	var fileType string
	fileCode := bytesToHexString(fSrc)

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})

	//把未知的类型记录下来，后续补充上
	if fileType == "" {
		logger.Error(fileCode)
	}
	return fileType
}

func GetFileTypeByName(name string) string {
	s := strings.Split(name, ".")
	if len(s) < 2 {
		return ""
	}
	return s[len(s)-1]
}

func GetRandString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}

	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}

	return string(result)
}

func MakeFileName(userid, index int, filetype, usetype string) string {
	filename := fmt.Sprintf("%s_%d_%d_%d_%s.%s", usetype, userid, GetNowTimestamp(), index, GetRandString(4), filetype)
	return filename
}

func ChangeFileType(filename, filetype string) string {
	s := strings.Split(filename, ".")
	if len(s) < 1 {
		return ""
	}
	newname := fmt.Sprintf("%s.%s", s[0], filetype)
	return newname
}

//获取文件列表数量
func GetFileListCnt(filelist string) int {
	strs := strings.Split(filelist, ",")
	if strs == nil {
		return 0
	}
	return len(strs)
}

//获取文件列表
func GetFileIdList(filelist string) []int {
	strs := strings.Split(filelist, ",")
	idlist := []int{}
	for _, v := range strs {
		id, _ := strconv.Atoi(v)
		idlist = append(idlist, id)
	}
	return idlist
}

//获取第一个文件
func GetFirstFileId(filelist string) int {
	strs := strings.Split(filelist, ",")
	if strs == nil {
		return 0
	}
	id, _ := strconv.Atoi(strs[0])
	return id
}

//生成文件列表
func MakeFileIdList(filelist []int) string {
	str := ""
	for _, v := range filelist {
		if str == "" {
			str = fmt.Sprintf("%d", v)
		} else {
			str += fmt.Sprintf(",%d", v)
		}
	}
	return str
}

//设置定时器
func SetMyTimer(s int, f func(), n int) chan int {
	//logger.Debug("开始设置定时器")
	time1 := time.NewTicker(time.Second*time.Duration(s))
	ch := make(chan int)
	//logger.Debug("设置定时器 ", time1)

	cnt := 0

	go func() {
		for {
			select {
			case <-time1.C:
				//logger.Debug("定时器响应1")
				f()
				//logger.Debug("定时器响应2")
				if n > 0 {
					cnt ++
					if cnt >= n {
						//logger.Debug("定时器次数到了")
						time1.Stop()
						return
					}
				}
			case <-ch:
				logger.Debug("定时器提前停止")
				time1.Stop()
				return
			}
		}
	}()

	return ch
}

//获取当前时间戳
func GetNowTimestamp() int {
	return int(time.Now().Unix())
}

//获取当前时间字符串
func GetNowTimestr() string {
	timeTemplate := "2006-01-02 15:04:05"
	return time.Now().Format(timeTemplate)
}

//字符串转时间戳
func StrToTimestamp(timestr string) int {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	return int(tm.Unix())
}

//时间戳转字符串
func TimestampToStr(tmstp int) string {
	t := int64(tmstp)
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(t, 0).Format(timeTemplate)
}

//字符串转时间
func StrToTime(timestr string) time.Time {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	return tm
}

//字符串转日期
func StrToDate(datestr string) time.Time {
	tm, _ := time.ParseInLocation("2006-01-02", datestr, time.Local)
	return tm
}

//获取IM用户的ID
func GetIMUserID(account string) int {
	s := strings.Replace(account, config.GetIMConfig().Pre, "", 1)
	userid, _ := strconv.Atoi(s)
	return userid
}