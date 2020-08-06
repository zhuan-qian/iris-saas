package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"zhuan-qian/go-saas/common"
	"sync"
	"time"
)

var (
	instance *elastic.Client
	once     sync.Once
)

func EsInit() *elastic.Client {
	once.Do(func() {
		host := common.ElasticSearchHost()
		//es 配置
		errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
		var err error

		var config []elastic.ClientOptionFunc
		config = append(config, elastic.SetErrorLog(errorlog))
		config = append(config, elastic.SetURL(host))
		config = append(config, elastic.SetScheme("http"))
		config = append(config, elastic.SetSniff(false))
		if !common.IsDebug() {
			config = append(config, elastic.SetBasicAuth(os.Getenv("ES_USER"), os.Getenv("ES_PASS")))
		}

		instance, err = elastic.NewClient(config...)
		if err != nil {
			panic(err)
		}
		ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
		info, code, err := instance.Ping(host).Do(ctx)
		if err != nil {
			panic(err)
		}
		cancelFn()
		fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

		esversion, err := instance.ElasticsearchVersion(host)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Elasticsearch version %s\n", esversion)
	})
	return instance
}

type Basic interface {
	NewEsIndex(m interface{}) error
	UpdateEsIndex(id int64, m interface{}) error
}

func IntPtrSingle(v interface{}) *int {
	if values, ok := v.([]interface{}); ok {
		container := int(values[0].(float64))
		return &container
	}
	return nil
}

func Int16PtrSingle(v interface{}) *int16 {
	if values, ok := v.([]interface{}); ok {
		container := int16(values[0].(float64))
		return &container
	}
	return nil
}

func StringPtrSingle(v interface{}) *string {
	if values, ok := v.([]interface{}); ok {
		container := values[0].(string)
		return &container
	}
	return nil
}
