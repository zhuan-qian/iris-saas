package dao

import "xorm.io/xorm"

type Handle interface {
	Write (session *xorm.Session) Handle
	Reade () *xorm.Session
}