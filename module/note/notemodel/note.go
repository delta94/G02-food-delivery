package notemodel

type Note struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Title   string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
	Status  int    `json:"status"  gorm:"column:status;"`
}

func (Note) TableName() string {
	return "notes"
}
