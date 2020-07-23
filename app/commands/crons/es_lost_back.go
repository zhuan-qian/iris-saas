package crons

import (
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"scaffold/dao"
	"scaffold/model"
	"reflect"
)

func EsLostBack() {
	err := esLostIndex()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = esLostUpdate()
	if err != nil {
		fmt.Println(err.Error())
	}

}

func esLostIndex() error {
	var (
		//dbEngine = common.GetDB()

		listUndo     []*model.EsLostIndex
		indexMapUndo = make(map[string][]int64)
		err          error
		//articlesList []*model.Articles
		//dynamicList  []*model.Dynamic
	)
	listUndo, err = dao.NewEsLostIndexDao().WithSession(nil).List(0)
	if err != nil {
		return err
	}

	//按索引建立map
	for _, v := range listUndo {
		if _, ok := indexMapUndo[v.Index]; ok {
			indexMapUndo[v.Index] = append(indexMapUndo[v.Index], v.ObjId)
		} else {
			indexMapUndo[v.Index] = []int64{v.ObjId}
		}
	}

	//按索引查询需要更新内容
	//for index, ids := range indexMapUndo {
	//	session := dbEngine.NewSession().In("id", ids)
	//	switch index {
	//	case dao.ES_INDEX_OF_ARTICLE:
	//		err = session.Find(&articlesList)
	//		if err != nil {
	//			return err
	//		}
	//		err = esIndexHandle(articlesList, dao.NewArticlesDao(), index)
	//		if err != nil {
	//			return err
	//		}
	//
	//	case dao.ES_INDEX_OF_DYNAMIC:
	//		err = session.Find(dynamicList)
	//		if err != nil {
	//			return err
	//		}
	//		err = esIndexHandle(dynamicList, dao.NewDynamicDao(), index)
	//		if err != nil {
	//			return err
	//		}
	//
	//	default:
	//		continue
	//	}
	//}

	return nil
}

func esIndexHandle(list interface{}, d interface{}, index string) (err error) {
	var (
		finished  []string
		daoValue  = reflect.ValueOf(d)
		daoFunc   = daoValue.MethodByName("NewEsIndex")
		listValue = reflect.ValueOf(list)
	)

	for i := 0; i <= listValue.Len(); i++ {
		v := listValue.Index(i).Elem()
		result := daoFunc.Call([]reflect.Value{v})
		if len(result) == 1 && !result[0].IsNil() {
			continue
			//return errors.New(fmt.Sprintf("%v", result[0].Interface()))
		}
		finished = append(finished, fmt.Sprintf("%d", v.FieldByName("id").Int()))
	}

	if len(finished) > 0 {
		err = dao.NewEsLostIndexDao().Delete(index, finished)
	}

	return
}

func esLostUpdate() error {
	var (
		//dbEngine     = common.GetDB()
		listUndo     []*model.EsLostUpdate
		indexMapUndo = make(map[string][]int64)
		err          error
		//articlesList []*model.Articles
		//dynamicList  []*model.Dynamic
	)
	listUndo, err = dao.NewEsLostUpdateDao().WithSession(nil).List(0)
	if err != nil {
		return err
	}

	//按索引建立map
	for _, v := range listUndo {
		if _, ok := indexMapUndo[v.Index]; ok {
			indexMapUndo[v.Index] = append(indexMapUndo[v.Index], v.ObjId)
		} else {
			indexMapUndo[v.Index] = []int64{v.ObjId}
		}
	}

	//按索引查询需要更新内容
	//for index, ids := range indexMapUndo {
	//	session := dbEngine.NewSession().In("id", ids)
	//	switch index {
	//	case dao.ES_INDEX_OF_ARTICLE:
	//		err = session.Find(articlesList)
	//		if err != nil {
	//			return err
	//		}
	//		err = esUpdateHandle(articlesList, dao.NewArticlesDao(), index)
	//		if err != nil {
	//			return err
	//		}
	//
	//	case dao.ES_INDEX_OF_DYNAMIC:
	//		err = session.Find(dynamicList)
	//		if err != nil {
	//			return err
	//		}
	//		err = esUpdateHandle(dynamicList, dao.NewDynamicDao(), index)
	//		if err != nil {
	//			return err
	//		}
	//	default:
	//		continue
	//	}
	//
	//}

	return nil
}

func esUpdateHandle(list interface{}, d interface{}, index string) (err error) {
	var (
		finished   []string
		daoValue   = reflect.ValueOf(d)
		daoFunc    = daoValue.MethodByName("UpdateEsIndex")
		daoNewFunc = daoValue.MethodByName("NewEsIndex")
		listValue  = reflect.ValueOf(list)
	)

	for i := 0; i <= listValue.Len(); i++ {
		v := listValue.Index(i).Elem()
		id := v.FieldByName("id").Int()

		result := daoFunc.Call([]reflect.Value{reflect.ValueOf(id), v})
		if len(result) == 1 && !result[0].IsNil() && !elastic.IsNotFound(result[0].Interface()) {
			continue
			//return errors.New(fmt.Sprintf("%v", result[0].Interface()))
		}

		if elastic.IsNotFound(result[0].Interface()) {
			iresult := daoNewFunc.Call([]reflect.Value{v})
			if len(iresult) == 1 && !iresult[0].IsNil() {
				return errors.New(fmt.Sprintf("%v", iresult[0].Interface()))
			}
		}

		finished = append(finished, fmt.Sprintf("%d", v.FieldByName("id").Int()))
	}

	if len(finished) > 0 {
		err = dao.NewEsLostIndexDao().Delete(index, finished)
	}
	return nil
}
