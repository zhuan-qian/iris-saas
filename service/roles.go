package service

import (
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"strings"
)

type Roles struct {
	d *dao.Roles
}

func NewRolesService() *Roles {
	return &Roles{d: dao.NewRolesDao().WithSession(nil)}
}

func (s *Roles) Create(m *model.Roles) (id int64, err error) {
	return s.d.InsertOne(m)
}

func (s *Roles) GetListByKey(orgId int, limit *int, page *int, keyword *string) ([]model.Roles, int64, error) {
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
	*keyword = strings.TrimSpace(*keyword)
	return s.d.GetAll(orgId, limit, common.IntPtr(common.PageToOffset(*page, *limit)), keyword)
}

func (s *Roles) Modify(orgId int, id int, p *params.RolesPut) (num int64, err error) {
	var (
		cols   []string
		m      = &model.Roles{}
		isKing bool
	)

	isKing, err = dao.NewRolesDao().WithSession(nil).IsKingBy(orgId, id)
	if err != nil {
		return 0, err
	}
	if isKing {
		return 0, common.NewRequireError("admin无法修改")
	}

	if p.Name != nil {
		cols = append(cols, "name")
		m.Name = *p.Name
	}

	if p.Status != nil {
		cols = append(cols, "status")
		m.Status = *p.Status
	}

	return s.d.ModifyBy(cols, int64(id), m)
}
