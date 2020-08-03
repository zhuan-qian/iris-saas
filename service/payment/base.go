package payment

import (
	"github.com/smartwalle/alipay/v3"
	"net/http"
	"sync"
)

var (
	APClient *alipay.Client
	apOnce     sync.Once
)

type Payment interface {
	PrepareToPay(userId int64,orders []string) (p string, err error)
	SyncResultForPay(r *http.Request) (err error)
}

//func GetAliPayHandle() *alipay.Client{
//	var (
//		err error
//	)
//	apOnce.Do(func() {
//
//		// 初始化支付宝客户端
//		//    appId：应用ID
//		//    privateKey：应用私钥，支持PKCS1和PKCS8
//		//    isProd：是否是正式环境
//
//		APClient, err = alipay.New(os.Getenv("ALI_APP_ID"), os.Getenv("ALI_APP_PRIVATE"),strings.ToLower(os.Getenv("APP_DEBUG")) == "false")
//		if err!=nil{
//			panic(err)
//		}
//		err=APClient.LoadAppPublicCertFromFile("config/payment/appCertPublicKey_2016090900474205.crt") // 加载应用公钥证书
//		if err != nil {
//			panic("解析支付宝证书-appCert-失败")
//		}
//
//		err=APClient.LoadAliPayRootCertFromFile("config/payment/alipayRootCert.crt") // 加载支付宝根证书
//		if err != nil {
//			panic("解析支付宝证书-rootCert-失败")
//		}
//		err=APClient.LoadAliPayPublicCertFromFile("config/payment/alipayCertPublicKey_RSA2.crt") // 加载支付宝公钥证书
//		if err != nil {
//			panic("解析支付宝证书-publicCert-失败")
//		}
//
//
//		if err != nil {
//			panic("支付宝支付参数校验出错:" + err.Error())
//		}
//	})
//	return APClient
//}
//func GetWxPayHandle() *WxPayHandle{
//	var(
//		err error
//	)
//
//	wxOnce.Do(func() {
//		// 初始化微信客户端
//		//    appId：应用ID
//		//    mchId：商户ID
//		//    apiKey：API秘钥值
//		//    isProd：是否是正式环境
//		WPHandle=&WxPayHandle{
//			Config:        WxPayConfig{
//				AppId:         os.Getenv("WX_APP_ID"),
//				MerchantId:    os.Getenv("WX_MERCHANT_ID"),
//				MerchantKeyId: os.Getenv("WX_MERCHANT_API_KEY_ID"),
//				IsProd:        os.Getenv("APP_DEBUG") == "false",
//				CertPath:      os.Getenv("WX_CERT_PATH"),
//				KeyPath:       os.Getenv("WX_KEY_PATH"),
//				Pkcs12Path:    os.Getenv("WX_PKCS12_PATH"),
//			},
//		}
//		WPHandle.DefaultClient,err=WPHandle.NewClient()
//
//		if err != nil {
//			panic("微信支付参数校验出错:" + err.Error())
//		}
//	})
//	return WPHandle
//}
