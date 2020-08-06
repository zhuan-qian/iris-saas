package common

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"zhuan-qian/go-saas/model"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func LogFile() *os.File {
	filename := "public/runtime/logs/" + time.Now().Format("20060102") + ".log"
	// Open the file, this will append to the today's file if server restarted.
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func DateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func DateTimePtr(t time.Time) *string {
	return StringPtr(t.Format("2006-01-02 15:04:05"))
}

func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func AfterTimeByHourNum(HourNum int) string{
	return time.Now().Add(time.Hour*time.Duration(HourNum)).Format("2006-01-02 15:04:05")
}

func AfterTimeByHourNumPtr(HourNum int) *string{
	return common.StringPtr(time.Now().Add(time.Hour*time.Duration(HourNum)).Format("2006-01-02 15:04:05"))
}

func NowDateTimePtr() *string {
	return StringPtr(time.Now().Format("2006-01-02 15:04:05"))
}

func Date(t time.Time) string {
	return t.Format("2006-01-02")
}

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func NowDatePtr() *string {
	return StringPtr(time.Now().Format("2006-01-02"))
}

func HttpDo(method string, urlStr string, jsonData interface{}, headers *map[string]string) ([]byte, error) {
	var (
		req         = fasthttp.AcquireRequest()
		resp        = fasthttp.AcquireResponse()
		result      = []byte("")
		requestStrs []string
		requestStr  string
		requestByte []byte

		interval string

		err error
	)

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(strings.ToLower(method))
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}

	if jsonData != nil {
		if method == "GET" {
			r := reflect.ValueOf(jsonData).MapRange()
			for {
				if !r.Next() {
					break
				}
				requestStrs = append(requestStrs, r.Key().String()+"="+fmt.Sprint(r.Value().Interface()))
			}
			requestStr = strings.Join(requestStrs, "&")

			if strings.Contains(urlStr, "?") {
				interval = "&"
			} else {
				interval = "?"
			}
			urlStr += interval + requestStr
		} else {
			requestByte, _ = json.Marshal(jsonData)
			req.SetBody(requestByte)
		}
	}

	req.SetRequestURI(urlStr)

	if err = fasthttp.Do(req, resp); err != nil {
		return result, err
	}

	result = resp.Body()
	return result, nil
}

func SmartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}

func ReadFileToString(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(`%s`, bytes), nil
}

func ElasticSearchHost() string {
	return "http://"+os.Getenv("ES_HOST")+":"+os.Getenv("ES_PORT")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GenerateRandomNum() string {
	var (
		rnum int
		lpad string
	)
	rand.Seed(time.Now().UnixNano())
	rnum = rand.Intn(9999)

	if rnum < 10 {
		lpad = "000"
	} else if rnum > 9 && rnum < 100 {
		lpad = "00"
	} else if rnum > 99 && rnum < 1000 {
		lpad = "0"
	}
	return lpad + strconv.Itoa(rnum)
}

func IsDebug() bool {
	return strings.ToLower(os.Getenv("APP_DEBUG")) == "true"
}

func EsCounter(field string, num int) string {
	return `{
		"source": "ctx._source.` + field + ` += params.count",
        "lang": "painless",
        "params" : {
            "count" : ` + strconv.Itoa(num) + `
        }}`
}

func EsViewNum(num int) string {
	return `{
		"source": "ctx._source.viewNum += params.count",
        "lang": "painless",
        "params" : {
            "count" : ` + strconv.Itoa(num) + `
        }}`
}

func PageToOffset(page int, limit int) int {
	return (page - 1) * limit
}

func GenerateSha1ByFile(file *multipart.File) (r string, err error) {
	var (
		h = sha1.New()
	)
	_, err = io.Copy(h, *file)
	if err != nil {
		return
	}
	_, err = (*file).Seek(0, 0)
	if err != nil {
		return
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func IntPtr(v int) *int {
	return &v
}
func Int8Ptr(v int8) *int8 {
	return &v
}
func Int16Ptr(v int16) *int16 {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func UintPtr(v uint) *uint {
	return &v
}

func Uint64Ptr(v uint64) *uint64 {
	return &v
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func StringPtr(v string) *string {
	return &v
}

func StringValues(ptrs []*string) []string {
	values := make([]string, len(ptrs))
	for i := 0; i < len(ptrs); i++ {
		if ptrs[i] != nil {
			values[i] = *ptrs[i]
		}
	}
	return values
}

func StringPtrs(vals []string) []*string {
	ptrs := make([]*string, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

func BoolPtr(v bool) *bool {
	return &v
}

func TrimCommaAndSpace(str string) string {
	return strings.TrimSpace(strings.Trim(str, ","))
}

func Printer(args ...interface{}) (interface{}, error) {
	fmt.Print("输入的参数分别是: ")
	fmt.Println(args...)
	return true, nil
}

func IsAdminInDomain(args ...interface{}) (interface{}, error) {
	var (
		csb    = GetCasbin()
		roles  []string
		user   = args[0].(string)
		domain = args[1].(string)
	)

	roles = csb.GetRolesForUserInDomain(user, domain)
	for _, v := range roles {
		if v == model.ROLES_NAME_OF_KING {
			return true, nil
		}
	}
	return false, nil
}

func RandCode() string {
	return fmt.Sprintf("%s%06d", time.Now().Format("060102150405"), rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(999999))
}

func ParseNumToYuan(num int) (yuan int) {
	return num / 100
}

func MinusTillEnough(minuendAmount *int, discountAmount *int) (deductions int) {
	if *discountAmount > *minuendAmount {
		*discountAmount -= *minuendAmount
		deductions = *minuendAmount
		*minuendAmount = 0
	} else {
		*minuendAmount -= *discountAmount
		deductions = *discountAmount
		*discountAmount = 0
	}
	return
}
