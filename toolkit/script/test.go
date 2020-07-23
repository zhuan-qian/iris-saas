package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"scaffold/model"
	"strconv"
	"strings"
	"xorm.io/core"
	"xorm.io/xorm"
)

type Province struct {
	Name string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	Code string `json:"code" xorm:"not null comment('城市代码') VARCHAR(16) index"`
}

func (m *Province) TableName() string {
	return "province"
}

type City struct {
	Name         string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	Code         string `json:"code" xorm:"not null comment('城市代码') VARCHAR(16) index"`
	ProvinceCode string `json:"provinceCode" xorm:"not null comment('城市代码') VARCHAR(16) index"`
}

func (m *City) TableName() string {
	return "city"
}

type Area struct {
	Name     string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	Code     string `json:"code" xorm:"not null comment('城市代码') VARCHAR(16) index"`
	CityCode string `json:"cityCode" xorm:"not null comment('城市代码') VARCHAR(16) index"`
}

func (m *Area) TableName() string {
	return "area"
}

type Street struct {
	Name     string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	Code     string `json:"code" xorm:"not null comment('城市代码') VARCHAR(16) index"`
	AreaCode string `json:"areaCode" xorm:"not null comment('城市代码') VARCHAR(16) index"`
}

func (m *Street) TableName() string {
	return "street"
}

type Village struct {
	Name       string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	Code       string `json:"code" xorm:"not null comment('城市代码') VARCHAR(16) index"`
	StreetCode string `json:"streetCode" xorm:"not null comment('城市代码') VARCHAR(16) index"`
}

func (m *Village) TableName() string {
	return "village"
}

func main() {
	var (
		provinces           []*Province
		cities              []*City
		citiesMapByProvince = make(map[string][]*model.Locations)
		areas               []*Area
		areasMapByCity      = make(map[string][]*model.Locations)
		streets             []*Street
		streetsMapByArea    = make(map[string][]*model.Locations)
		villages            []*Village
		villagesMapByStreet = make(map[string][]*model.Locations)
		cnLocations         []*model.Locations

		locations    []*model.Locations
		locationsMap = make(map[string]*model.Locations)

		topLocations []*model.Locations

		engine  *xorm.Engine
		sengine *xorm.Engine
		err     error
	)

	//环境配置获取
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	engine, err = xorm.NewEngine("sqlite3", "data.sqlite")
	if err != nil {
		goto HASERR
	}

	sengine = common.GetDB()

	engine.SetMapper(common.CamelMapper{})
	engine.SetTableMapper(core.SnakeMapper{})
	sengine.SetMapper(common.CamelMapper{})
	sengine.SetTableMapper(core.SnakeMapper{})

	//engine.SQL("select p.code as pcode,p.name as pname,c.code as ccode,c.name as cname," +
	//	"a.code as acode,a.name as aname,s.code as scode,s.name as sname" +
	//	" from province p" +
	//	" join city c on c.provinceCode = p.code" +
	//	" join area a on a.cityCode = c.code" +
	//	" join street s on s.areaCode = a.code").Find(&provinces)
	//for _, v := range provinces {
	//	fmt.Println(v.Pname, v.Cname, v.Aname, v.Sname)
	//}

	//TODO 移除中国除港澳台的所有地址信息
	engine.Exec("delete from locations where path like ',1,7,%' and path not like ',1,7,' and path not like ',1,7,278,%' and path not like ',1,7,279,%' and path not like ',1,7,280%'")

	sengine.Omit("id").Desc("path").Find(&locations)
	for i, v := range locations {
		locations[i].Parent = v.Path[0 : strings.LastIndex(strings.Trim(v.Path, ","), ",")+2]
		locations[i].Level = len(strings.Split(strings.Trim(v.Path, ","), ","))
		locationsMap[v.Path] = locations[i]
	}

	for _, v := range locations {
		if _, ok := locationsMap[v.Parent]; ok {
			locationsMap[v.Parent].Subs = append(locationsMap[v.Parent].Subs, locationsMap[v.Path])
		}
		if v.Level == 1 {
			topLocations = append(topLocations, locationsMap[v.Path])
		}
	}

	engine.SQL("select code,name from province").Find(&provinces)
	engine.SQL("select code,name,provinceCode from city").Find(&cities)
	engine.SQL("select code,name,cityCode from area").Find(&areas)
	engine.SQL("select code,name,areaCode from street").Find(&streets)
	engine.SQL("select code,name,streetCode from village").Find(&villages)

	for _, v := range villages {
		l := &model.Locations{
			Path:       v.Code,
			Level:      7,
			Name:       v.Name,
			NameEn:     "",
			NamePinyin: "",
			Code:       v.Code,
			Subs:       nil,
		}

		if _, ok := villagesMapByStreet[v.StreetCode]; ok {
			villagesMapByStreet[v.StreetCode] = append(villagesMapByStreet[v.StreetCode], l)
		} else {
			villagesMapByStreet[v.StreetCode] = []*model.Locations{l}
		}
	}

	for _, v := range streets {
		l := &model.Locations{
			Path:       v.Code,
			Level:      6,
			Name:       v.Name,
			NameEn:     "",
			NamePinyin: "",
			Code:       v.Code,
			Subs:       nil,
		}

		if _, ok := villagesMapByStreet[v.Code]; ok {
			l.Subs = villagesMapByStreet[v.Code]
		}
		if _, ok := streetsMapByArea[v.AreaCode]; ok {
			streetsMapByArea[v.AreaCode] = append(streetsMapByArea[v.AreaCode], l)
		} else {
			streetsMapByArea[v.AreaCode] = []*model.Locations{l}
		}
	}

	for _, v := range areas {
		l := &model.Locations{
			Path:       v.Code,
			Level:      5,
			Name:       v.Name,
			NameEn:     "",
			NamePinyin: "",
			Code:       v.Code,
			Subs:       nil,
		}

		if _, ok := streetsMapByArea[v.Code]; ok {
			l.Subs = streetsMapByArea[v.Code]
		}

		if _, ok := areasMapByCity[v.CityCode]; ok {
			areasMapByCity[v.CityCode] = append(areasMapByCity[v.CityCode], l)
		} else {
			areasMapByCity[v.CityCode] = []*model.Locations{l}
		}
	}

	for _, v := range cities {
		l := &model.Locations{
			Path:       v.Code,
			Level:      4,
			Name:       v.Name,
			NameEn:     "",
			NamePinyin: "",
			Code:       v.Code,
			Subs:       nil,
		}

		if _, ok := areasMapByCity[v.Code]; ok {
			l.Subs = areasMapByCity[v.Code]
		}

		if _, ok := citiesMapByProvince[v.ProvinceCode]; ok {
			citiesMapByProvince[v.ProvinceCode] = append(citiesMapByProvince[v.ProvinceCode], l)
		} else {
			citiesMapByProvince[v.ProvinceCode] = []*model.Locations{l}
		}
	}

	for _, v := range provinces {
		l := &model.Locations{
			Path:       v.Code,
			Level:      3,
			Name:       v.Name,
			NameEn:     "",
			NamePinyin: "",
			Code:       v.Code,
			Subs:       nil,
		}

		if _, ok := citiesMapByProvince[v.Code]; ok {
			l.Subs = citiesMapByProvince[v.Code]
		}

		cnLocations = append(cnLocations, l)
	}

	locationsMap[",1,7,"].Subs = cnLocations

	engine.Exec("truncate table locations")

	err = intoLocations(topLocations, 0, ",")
	if err != nil {
		fmt.Println(err.Error())
	}

	return

HASERR:
	fmt.Println(err.Error())
	return

}

func intoLocations(list []*model.Locations, parentLevel int, parentPath string) (err error) {
	level := parentLevel + 1
	d := dao.NewLocationsDao().WithSession(nil)
	for _, v := range list {
		v.Level = level
		_, err = d.InsertOne(v)
		if err != nil {
			return err
		}
		v.Path = parentPath + strconv.Itoa(v.Id) + ","
		_, err = d.ModifyBy([]string{"path"}, int64(v.Id), v)
		if err != nil {
			return err
		}
		if v.Subs != nil {
			err = intoLocations(v.Subs, level, v.Path)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
