package mashup

type OrgWithGroupMap map[int]*OrgWithGroup

type OrgWithGroup struct {
	GroupId int `json:"groupId"`
	Count   int `json:"count"`
}

func (m *OrgWithGroup) TableName() string {
	return "organizations"
}
