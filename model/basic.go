package model

type LatLon struct {
	Lon *float64 `json:"lon,omitempty" xorm:"null comment('经度') DECIMAL(10,7)"`
	Lat *float64 `json:"lat,omitempty" xorm:"null comment('纬度') DECIMAL(10,7)"`
}
