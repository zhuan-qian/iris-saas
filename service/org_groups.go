package service

import (
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/model/mashup"
)

type orgGroups struct {
	d *dao.OrgGroups
}

func NewOrgGroupsService() *orgGroups {
	return &orgGroups{d: dao.NewOrgGroupsDao().WithSession(nil)}
}

//通过关键词查询学校分组列表
func (s *orgGroups) GetListByKey(limit *int, page *int) (list []*model.OrgGroups, count int64, err error) {
	var (
		ids       []int
		orgDao    = dao.NewOrganizationsDao().WithSession(nil)
		CountList []*mashup.OrgWithGroup
		CountMap  = make(mashup.OrgWithGroupMap)
	)

	if limit == nil {
		limit = new(int)
		*limit = 0
	}
	if page == nil {
		page = new(int)
		*page = 0
	}
	list, count, err = s.d.GetAll(limit, common.IntPtr(common.PageToOffset(*page, *limit)))
	if err != nil {
		return nil, 0, err
	}

	ids = s.d.CollectIdBy(list)
	CountList, err = orgDao.CountBy(ids)
	if err != nil {
		return nil, 0, err
	}
	CountMap = orgDao.MapOfGroupIdBy(CountList)
	s.d.FillCountBy(list, CountMap)
	return
}

func (s *orgGroups) Create(m *model.OrgGroups) (result int64, err error) {
	return s.d.CreateOne(m)
}

func (s *orgGroups) Modify(id int, p *params.OrgGroupsPut) (num int64, err error) {
	var (
		cols []string
		m    = &model.OrgGroups{}
	)
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
