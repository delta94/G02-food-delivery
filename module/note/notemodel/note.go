package notemodel

import (
	"fooddlv/common"
)

const EntityName = "Note"

type Note struct {
	common.SQLModel `json:",inline"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser `json:"owner" gorm:"foreignKey:Id;references:UserId;"`
	Cover           *common.Image      `json:"cover" gorm:"column:cover;"`
	Photos          *common.Images     `json:"photos" gorm:"column:photos;"`
	Title           string             `json:"title" gorm:"column:title;"`
	Content         string             `json:"content" gorm:"column:content;"`
}

func (Note) TableName() string {
	return "notes"
}

func (n *Note) Mask(isAdmin bool) {
	n.GenUID(common.DBTypeNote, 1)

	if n.User != nil {
		n.User.GenUID(common.DBTypeUser, 1)
	}
}
