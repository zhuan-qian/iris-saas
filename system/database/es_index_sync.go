package database

import (
	"context"
	"errors"
	"io/ioutil"
	"gold_hill/mine/common"
	"gold_hill/mine/service/es"
	"strings"
)

func EsIndexSync() error {

	rd, err := ioutil.ReadDir("model/es")
	if err != nil {
		return err
	}

	esClient := es.EsInit()

	for _, file := range rd {
		name := file.Name()
		prefix := strings.Split(name, ".")
		exist, err := esClient.IndexExists(prefix[0]).Do(context.Background())
		if err != nil {
			return err
		}
		if exist {
			continue
		}
		con, err := common.ReadFileToString("model/es/" + name)
		if err != nil {
			return err
		}
		createIndex, err := esClient.CreateIndex(prefix[0]).BodyString(con).Do(context.Background())
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			return errors.New("es映射创建失败 name:" + name)
		}
	}
	return nil
}
