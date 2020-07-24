package sms

import (
	"encoding/json"
	"errors"
	"github.com/denverdino/aliyungo/sms"
	"os"
	"gold_hill/scaffold/common"
	"gold_hill/scaffold/service/cache"
	"strconv"
	"time"
)

const (
	SMS_Ali_SIGN_NAME              = "宠爱你"
	SMS_Ali_FOR_LOGIN_AND_REGISTER = "SMS_163240076"
)

type aliSms struct {
	Client *sms.DYSmsClient
}

func (s *aliSms) newCache(phone string, code string) error {
	return cache.RedisInit().Set(s.keyBy(phone), code, time.Second*60*2).Err()
}

func (s *aliSms) keyBy(phone string) string {
	return strconv.Itoa(cache.USERS_SMS_SEND_REPEAT) + "_" + phone
}

func (s *aliSms) isRepeatSend(phone string) bool {
	var (
		cacheKey = s.keyBy(phone)
		i        int64
	)
	i, _ = cache.RedisInit().Exists(cacheKey).Result()
	if i > 0 {
		return true
	}
	return false

}

func NewAliSmsService() *aliSms {
	return &aliSms{Client: sms.NewDYSmsClient(os.Getenv("ALI_KEY_ID"), os.Getenv("ALI_KEY_SECRET"))}
}

func (s *aliSms) CodeIsRightThenDel(phone string, code string) bool {
	cacheCode, _ := cache.RedisInit().Get(s.keyBy(phone)).Result()
	result := cacheCode == code
	if result {
		cache.RedisInit().Del(s.keyBy(phone)).Err()
	}
	return result
}

//发送4位随机短信
func (s *aliSms) SendRandomMsg(phone string) (string, error) {
	var (
		msg   *sms.SendSmsArgs
		jsonm = make(map[string]string)
		jsonb []byte
		err   error
	)

	if s.isRepeatSend(phone) {
		return "", errors.New("重复请求发送短信,请在一段时间后再尝试")
	}

	jsonm["code"] = common.GenerateRandomNum()
	jsonb, err = json.Marshal(jsonm)
	if err != nil {
		return "", err
	}

	msg = &sms.SendSmsArgs{
		PhoneNumbers:    phone,
		SignName:        SMS_Ali_SIGN_NAME,
		TemplateCode:    SMS_Ali_FOR_LOGIN_AND_REGISTER,
		TemplateParam:   string(jsonb),
		SmsUpExtendCode: "",
		OutId:           "",
	}

	_, err = s.Client.SendSms(msg)
	if err != nil {
		return "", err
	}

	err = s.newCache(phone, jsonm["code"])

	return jsonm["code"], err

}
