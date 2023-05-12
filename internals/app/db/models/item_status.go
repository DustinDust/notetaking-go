package models

type ItemStatus string

const (
	STATUS_DELETED  ItemStatus = "deleted"
	STATUS_ARCHIVED ItemStatus = "archived"
	STATUS_ACTIVE   ItemStatus = "active"
)
