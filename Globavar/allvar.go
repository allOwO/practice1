package Globavar

import (
	"github.com/streadway/amqp"
	"sync"
	"time"
)

//from cmd
//消费者数量
var Recvnum int

//发送类型
var Typ string

//全局channel
var Conn *amqp.Connection

var MysqlSync sync.RWMutex

//sql User
type User struct {
	ID           int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string `gorm:"Column:user_name;not null"`
	Phone        string `gorm:"Column:user_phone;not null"`
	Mail         string `gorm:"Column:user_mail;unique;not null;index"`
	SystemUser   bool   `gorm:"Column:system_user"`
	Workers      bool   `gorm:"Column:worker"`
	ServiceStaff bool   `gorm:"Column:service_staff"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

//用来传输web
type WebUserMess struct {
	UserName  string   `json:"user_name"`
	UserPhone string   `json:"user_phone"`
	UserMail  string   `json:"user_mail" `
	Groups    []string `json:"groups"`
}

func NewWebUserMess() *WebUserMess {
	wu := new(WebUserMess)
	wu.Groups = make([]string, 0, 3)
	return wu
}

//sql message
type Message struct {
	ID          int       `gorm:"AUTO_INCREMENT;primary_key"`
	Mess        string    `gorm:"Column:mess;unique;not null;"`
	Group       string    `gorm:"Column:group"`
	Success     bool      `gorm:"Column:success"`
	SuccessUser int       `gorm:"Column:success_user"`
	FailUser    int       `gorm:"Column:fail_user"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type MessJsonBody struct {
	User    User
	Typ     string
	Message Message
}
