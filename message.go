package PracticeItem

import "time"

type Message struct {
	ID          int    `gorm:"AUTO_INCREMENT;primary_key"`
	Mess        string `gorm:"Column:mess;unique;not null;"`
	Type        string `gorm:"Column:type;not null;"`
	Group       string `gorm:"Column:group"`
	Success     bool   `gorm:"Column:success"`
	//SuccessUser int    `gorm:"Column:success_user"`
	//FailUser    int    `gorm:"Column:fail_user"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type MessService interface {
	CreateMessage(*Message)bool
	//InsertMessSucess(*Message)bool
	//InsertMessFail(*Message)bool
}