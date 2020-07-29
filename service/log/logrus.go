package log

import (
	"github.com/sirupsen/logrus"
	"gold_hill/mine/service/es"
	"gold_hill/mine/service/log/hook"
	"gopkg.in/sohlich/elogrus.v7"
	"os"
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
		hook, err := elogrus.NewElasticHook(es.EsInit(), os.Getenv("APP_URL"), logrus.DebugLevel, os.Getenv("APP_NAME"))
		//instance.AddHook(hook)
		instance.SetOutput(os.Stdout)
	})
	return instance
}
