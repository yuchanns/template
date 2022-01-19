package models

type Greet struct {
	ID   string `gorm:"column:id"`
	Name string `gorm:"name"`
	Age  int    `gorm:"aget"`
}

func (Greet) TableName() string {
	return "Greet"
}
