package sms

import (
	"errors"
	"fmt"
	txcommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	txerrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711" //引入sms
	"os"
	"gold_hill/mine/common"
	"gold_hill/mine/service/cache"
	"strconv"
	"time"
)

const (
	SMS_SIGN_NAME              = "宠爱你"
	SMS_FOR_LOGIN_AND_REGISTER = "SMS_163240076"
)

type txSms struct {
	Client *sms.Client
}

func (s *txSms) newCache(phone string, code string) error {
	return cache.RedisInit().Set(s.keyBy(phone), code, time.Second*60*2).Err()
}

func (s *txSms) keyBy(phone string) string {
	return strconv.Itoa(cache.USERS_SMS_SEND_REPEAT) + "_" + phone
}

func (s *txSms) isRepeatSend(phone string) bool {
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

func NewTxSmsService() *txSms {
	credential := txcommon.NewCredential(
		os.Getenv("TX_KEY_ID"),
		os.Getenv("TX_KEY_SECRET"),
	)

	client, _ := sms.NewClient(credential, "ap-guangzhou", profile.NewClientProfile())
	return &txSms{Client: client}
}

func (s *txSms) CodeIsRightThenDel(phone string, code string) bool {
	cacheCode, _ := cache.RedisInit().Get(s.keyBy(phone)).Result()
	result := cacheCode == code
	if result {
		cache.RedisInit().Del(s.keyBy(phone)).Err()
	}
	return result
}

//发送4位随机短信
func (s *txSms) SendRandomMsg(phone string) (string, error) {
	var (
		msg        *sms.SendSmsRequest
		randomCode string
		err        error
	)

	if s.isRepeatSend(phone) {
		return "", errors.New("重复请求发送短信,请在一段时间后再尝试")
	}

	randomCode = common.GenerateRandomNum()

	msg = sms.NewSendSmsRequest()
	/* 短信应用ID: 短信SdkAppid在 [短信控制台] 添加应用后生成的实际SdkAppid，示例如1400006666 */
	msg.SmsSdkAppid = common.StringPtr(os.Getenv("TX_SMS_APP_ID"))
	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，签名信息可登录 [短信控制台] 查看 */
	msg.Sign = common.StringPtr(os.Getenv("TX_SMS_SIGN"))
	/* 国际/港澳台短信 senderid: 国内短信填空，默认未开通，如需开通请联系 [sms helper] */
	//msg.SenderId = common.StringPtr("xxx")
	/* 用户的 session 内容: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	//msg.SessionContext = common.StringPtr("xxx")
	/* 短信码号扩展号: 默认未开通，如需开通请联系 [sms helper] */
	//msg.ExtendCode = common.StringPtr("0")
	/* 模板参数: 若无模板参数，则设置为空*/
	msg.TemplateParamSet = common.StringPtrs([]string{randomCode, "2"})
	/* 模板 ID: 必须填写已审核通过的模板 ID。模板ID可登录 [短信控制台] 查看 */
	msg.TemplateID = common.StringPtr("74806")
	/* 下发手机号码，采用 e.164 标准，+[国家或地区码][手机号]
	 * 示例如：+8613711112222， 其中前面有一个+号 ，86为国家码，13711112222为手机号，最多不要超过200个手机号*/
	msg.PhoneNumberSet = common.StringPtrs([]string{fmt.Sprintf("+86%s", phone)})

	// 通过client对象调用想要访问的接口，需要传入请求对象
	_, err = s.Client.SendSms(msg)

	if _, ok := err.(*txerrors.TencentCloudSDKError); ok {
		return "", errors.New(fmt.Sprintf("An API error has returned: %s", err))
	}

	if err != nil {
		return "", err
	}

	err = s.newCache(phone, randomCode)

	return randomCode, err
}
