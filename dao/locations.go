package dao

import (
	"errors"
	"gold_hill/scaffold/model"
	"sort"
	"strings"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Locations struct {
	Base
}

func NewLocationsDao() *Locations {
	return &Locations{}
}

func (d *Locations) WithSession(s *xorm.Session) *Locations {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Locations) ListByPath(path string) (list []*model.Locations, err error) {
	err = d.session.Where(builder.Like{"path", "%" + path + "%"}).Find(&list)
	return
}

func (d *Locations) ListBy(cols []string, ids []string) (list []*model.Locations, err error) {
	err = d.session.Cols(cols...).In("id", ids).Asc("path").Find(&list)
	return
}

func (d *Locations) ParseLocationsBy(link string) (province *int, provinceName *string, city *int, cityName *string,
	region *int, regionName *string, err error) {

	var (
		links = strings.Split(strings.Trim(link, ","), ",")
		list  []*model.Locations
	)

	if len(links) != 4 {
		goto ERR
	}

	list, err = d.ListBy([]string{"id", "name"}, links)

	if len(list) != 4 {
		goto ERR
	}

	province = &list[2].Id
	provinceName = &list[2].Name
	city = &list[3].Id
	cityName = &list[3].Name
	//region = &list[1].Id
	//regionName = &list[1].Name
	return

ERR:
	err = errors.New("地址路径信息错误")
	return

}

func (d *Locations) GetFatherBy(path string) string {
	return path[:strings.LastIndex(path[0:len(path)-1], ",")+1]
}

func (d *Locations) BuildByTree(list []*model.Locations) []*model.Locations {
	for i := len(list) - 1; i >= 0; i-- {
		fatherPath := d.GetFatherBy(list[i].Path)
		for j := len(list) - 1; j >= 0; j-- {
			if list[j].Path != fatherPath {
				continue
			}
			list[j].Subs = append(list[j].Subs, list[i])
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	return list
}

func (d *Locations) Sort(isSmallToLarge bool, list []*model.Locations) {
	var (
		compare func(one int, two int) bool
	)
	if isSmallToLarge {
		compare = func(one int, two int) bool {
			return one < two
		}
	} else {
		compare = func(one int, two int) bool {
			return one > two
		}
	}

	sort.Slice(list, func(i, j int) bool {
		i = list[i].Id
		j = list[j].Id
		return compare(i, j)
	})

	for _, v := range list {
		if len(v.Subs) > 0 {
			d.Sort(true, v.Subs)
		}
	}
}
