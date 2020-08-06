package dao

import (
	"fmt"
	"zhuan-qian/go-saas/model"
	"sort"
	"strconv"
	"strings"
	"xorm.io/xorm"
)

type Menus struct {
	Base
}

func NewMenusDao() *Menus {
	return &Menus{}
}

func (d *Menus) WithSession(s *xorm.Session) *Menus {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Menus) GetFatherBy(path string) string {
	if strings.HasSuffix(path, model.SUFFIX_OF_LEVEL_1) {
		return model.SUFFIX_OF_LEVEL_0
	}
	if strings.HasSuffix(path, model.SUFFIX_OF_LEVEL_2) {
		return path[0:2] + model.SUFFIX_OF_LEVEL_1
	}
	if strings.HasSuffix(path, model.SUFFIX_OF_LEVEL_3) {
		return path[0:4] + model.SUFFIX_OF_LEVEL_2
	} else {
		return path[0:6] + model.SUFFIX_OF_LEVEL_3
	}
}

func (d *Menus) ListAndDescBy(genre int8, in []string) (list []*model.Menus, err error) {
	if in != nil {
		d.session = d.session.In("path", in)
	}
	err = d.session.Where("genre=?", genre).Desc("path").Find(&list)
	return
}

func (d *Menus) ListBy(genre int8, in []string) (list []*model.Menus, err error) {
	if in != nil {
		d.session = d.session.In("path", in)
	}
	err = d.session.Where("genre=?", genre).Find(&list)
	return
}

func (d *Menus) DeleteBy(path string, genre int8) error {
	sql := fmt.Sprintf("delete from `%s` where `path`='%s' and `genre`=%d", new(model.Menus).TableName(), path, genre)
	_, err := d.session.Exec(sql)
	return err
}

func (d *Menus) BuildTreeBy(list []*model.Menus) []*model.Menus {
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

func (d *Menus) Sort(isSmallToLarge bool, list []*model.Menus) {
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
		i, _ = strconv.Atoi(list[i].Path)
		j, _ = strconv.Atoi(list[j].Path)
		return compare(i, j)
	})
}

func (d *Menus) CollectPathBy(list []*model.Menus) (paths []string) {
	for _, v := range list {
		paths = append(paths, v.Path)
	}
	return
}
