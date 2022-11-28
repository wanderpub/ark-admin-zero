package utils

import (
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/enums/channelType"
	"ark-admin-zero/common/enums/messageType"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/jsonx"
)

const PhoneRegex = "^((13[0-9])|(14[5,7,9])|(15[0-3,5-9])|(166)|(17[0-9])|(18[0-9])|(19[1,8,9]))\\d{8}$"
const PARAM = "?"

func GenerateBusinessId(templateId, templateType int64) int64 {
	str := fmt.Sprintf("%d%s", int64(templateType*1000000)+templateId, time.Now().Format("20060102"))
	return cast.ToInt64(str)
}
func GenerateUrl(url string, templateId int64, templateType int64) string {
	businessId := GenerateBusinessId(templateId, templateType)
	if strings.Contains(url, "?") {
		return fmt.Sprintf("%s?track_code_bid=%d", url, businessId)
	}
	return fmt.Sprintf("%s&track_code_bid=%d", url, businessId)
}

// ReplaceByMap returns a copy of `origin`,
// which is replaced by a map in unordered way, case-sensitively.
func ReplaceByMap(origin string, replaces map[string]string) string {
	for k, v := range replaces {
		origin = strings.Replace(origin, "{$"+k+"}", v, -1)
	}
	return origin
}

func GetAllGroupIds() []string {
	list := make([]string, 0)
	for _, ct := range channelType.TypeCodeEn {
		for _, mt := range messageType.TypeCodeEn {
			list = append(list, ct+"."+mt)
		}
	}
	return list
}
func GetGroupIdByTaskInfo(info model.TaskInfo) string {
	channelCodeEn := channelType.TypeCodeEn[int(info.SendChannel)]
	msgCodeEn := messageType.TypeCodeEn[info.MsgType]
	return channelCodeEn + "." + msgCodeEn
}

func GetMqKey(channel, msgType string) string {
	return fmt.Sprintf("austin.biz.%s.%s", channel, msgType)
}

func ArrayStringUniq(arr []string) []string {
	set := make(map[string]struct{}, 0)
	for _, s := range arr {
		set[s] = struct{}{}
	}
	var newStr = make([]string, 0)
	for k := range set {
		newStr = append(newStr, k)
	}
	return newStr
}

func GetIntKeysByMap(data map[int]string) []int {
	var list []int
	for key := range data {
		list = append(list, key)
	}
	return list
}
func GetStringValuesByMap(data map[int]string) []string {
	var list []string
	for _, value := range data {
		list = append(list, value)
	}
	return list
}

func ArrayStringToInt64(ids []string) []int64 {
	var newIds []int64
	for _, id := range ids {
		newIds = append(newIds, cast.ToInt64(id))
	}
	return newIds
}

func ArrayInt64ToString(ids []int64) []string {
	var newIds []string
	for _, id := range ids {
		newIds = append(newIds, cast.ToString(id))
	}
	return newIds
}

func GetContentModel(in interface{}, out interface{}) error {
	marshal, _ := jsonx.Marshal(in)
	return jsonx.Unmarshal(marshal, &out)
}

func ArrayStringIn(list []string, found string) bool {
	for _, s := range list {
		if found == s {
			return true
		}
	}
	return false
}
func ArrayInt64In(list []int64, found int64) bool {
	for _, s := range list {
		if found == s {
			return true
		}
	}
	return false
}
