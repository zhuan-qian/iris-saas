package service

import (
	"scaffold/app/controllers/params"
	"scaffold/common"
	"scaffold/dao"
	"scaffold/model"
	"strconv"
	"strings"
)

type organizations struct {
	d *dao.Organizations
}

func NewOrganizationsService() *organizations {
	return &organizations{d: dao.NewOrganizationsDao().WithSession(nil)}
}

//通过关键词查询学校列表
func (s *organizations) GetListByKey(limit *int, page *int, keyword *string, groupId *int, sort *string, sortType *string) ([]*model.Organizations, int64, error) {
	if limit == nil {
		limit = new(int)
		*limit = 0
	}
	if page == nil {
		page = new(int)
		*page = 0
	}
	if keyword == nil {
		keyword = new(string)
		*keyword = ""
	}
	if groupId == nil {
		groupId = new(int)
		*groupId = 0
	}
	if sort == nil {
		sort = new(string)
		*sort = ""
	}
	if sortType == nil {
		sortType = new(string)
		*sortType = ""
	}
	*keyword = strings.TrimSpace(*keyword)
	return s.d.GetAll(limit, common.IntPtr(common.PageToOffset(*page, *limit)), keyword, keyword, groupId, sort, sortType)
}

func (s *organizations) Create(p *params.OrganizationsPost) (id int, err error) {
	var m = &model.Organizations{}

	if p.Owner != nil {
		m.Owner = *p.Owner
	}

	if p.Name != nil {
		m.Name = *p.Name
	}

	if p.Phone != nil {
		m.Phone = *p.Phone
	}

	if p.Area != nil {
		m.Area = *p.Area
	}

	if p.Address != nil {
		m.Address = *p.Address
	}

	if p.GroupId != nil {
		m.GroupId = *p.GroupId
	}

	if p.Type != nil {
		m.Type = *p.Type
	}

	if p.Status != nil {
		m.Status = *p.Status
	} else {
		m.Status = 0
	}

	m.Code = common.RandCode()

	_, err = s.d.InsertOne(m)
	return m.Id, err
}

func (s *organizations) CreateShop(p *params.OrganizationsPost) error {
	var (
		err   error
		orgId int
	)

	orgId, err = s.Create(p)
	if err != nil {
		return err
	}

	u := &params.UsersModify{}
	oid := new(int)
	*oid = orgId
	u.OrgId = oid
	saleable := new(int8)
	*saleable = 1
	u.Saleable = saleable
	_, err = NewUsersService().Modify(*p.Owner, u)
	if err != nil {
		return err
	}

	m := &model.Roles{
		Name:   model.ROLES_NAME_OF_KING,
		OrgId:  orgId,
		Status: 1,
	}
	_, err = NewRolesService().Create(m)
	if err != nil {
		return err
	}
	_, err = common.GetCasbin().AddPolicy(model.ROLES_NAME_OF_KING, strconv.Itoa(orgId), "all-objs", "all-acts")
	if err != nil {
		return err
	}
	_, err = NewAdminsRolesService().Modify(*p.Owner, []string{strconv.Itoa(m.Id)}, orgId)
	if err != nil {
		return err
	}

	return err
}

func (s *organizations) Modify(id int, p *params.OrganizationsPut) (num int64, err error) {
	var (
		cols []string
		m    = &model.Organizations{}
	)
	if p.Name != nil {
		cols = append(cols, "name")
		m.Name = *p.Name
	}

	if p.Phone != nil {
		cols = append(cols, "phone")
		m.Phone = *p.Phone
	}

	//if p.ExpireAt != nil {
	//	cols = append(cols, "expireAt")
	//	m.ExpireAt = p.ExpireAt
	//}

	if p.GroupId != nil {
		cols = append(cols, "groupId")
		m.GroupId = *p.GroupId
	}

	if p.Area != nil {
		cols = append(cols, "area")
		m.Area = *p.Area
	}

	if p.Address != nil {
		cols = append(cols, "address")
		m.Address = *p.Address
	}

	if p.Status != nil {
		cols = append(cols, "status")
		m.Status = *p.Status
	}

	return s.d.ModifyBy(cols, int64(id), m)
}

func (s *organizations) ModifyBy(id int) error {
	m := &model.Organizations{
		GroupId: 0,
	}
	_, err := s.d.ModifyByField([]string{"groupId"}, "groupId", id, m)
	return err
}

func (s *organizations) InfoBy(id int64) (*model.Organizations, error) {
	var (
		m   = &model.Organizations{}
		ok  bool
		err error
	)
	ok, err = s.d.InfoBy(nil, id, m)
	if err != nil {
		return nil, err
	}
	if ok == false {
		return nil, nil
	}

	return m, err
}
