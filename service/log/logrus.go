package log

import (
	"github.com/sirupsen/logrus"
	"zhuan-qian/go-saas/service/log/hook"
	"sync"
)

var (
	instance *logrus.Logger
	once     sync.Once
)

func Get() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()
		instance.SetReportCaller(true)
		instance.SetLevel(logrus.WarnLevel)
		instance.AddHook(hook.NewFileSplitHook(73))
		//eshook, err := elogrus.NewElasticHook(es.EsInit(), os.Getenv("APP_URL"), logrus.DebugLevel, os.Getenv("APP_NAME"))
		//if err!=nil{
		//	panic(err.Error())
		//}
		//instance.AddHook(eshook)
	})
	return instance
}
