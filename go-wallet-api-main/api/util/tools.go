package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	ToolsUtil    = toolsUtil{}
	allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// toolsUtil 常用工具集合类
type toolsUtil struct{}

// RandomString 返回随机字符串
func (tu toolsUtil) RandomString(length int) string {
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[rand.Intn(62)]
	}
	return string(byteList)
}

// MakeUuid 制作UUID
func (tu toolsUtil) MakeUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// MakeMd5 制作MD5
func (tu toolsUtil) MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// Contains 判断src是否包含elem元素
func (tu toolsUtil) Contains(src interface{}, elem interface{}) bool {
	srcArr := reflect.ValueOf(src)
	if srcArr.Kind() == reflect.Ptr {
		srcArr = srcArr.Elem()
	}
	if srcArr.Kind() == reflect.Slice {
		for i := 0; i < srcArr.Len(); i++ {
			if srcArr.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}

// Round float四舍五入
func (tu toolsUtil) Round(val float64, n int) float64 {
	base := math.Pow(10, float64(n))
	return math.Round(base*val) / base
}

// JsonToObj JSON转Obj
func (tu toolsUtil) JsonToObj(jsonStr string, toVal interface{}) (err error) {
	return json.Unmarshal([]byte(jsonStr), &toVal)
}

// ObjToJson Obj转JSON
func (tu toolsUtil) ObjToJson(data interface{}) (res string, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return res, err
	}
	res = string(b)
	return res, nil
}
