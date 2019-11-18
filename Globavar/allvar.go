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
type User struct {
	ID        int       `gorm:"primary_key,AUTO_INCREMENT"`
	Name      string    `gorm:"user_name"`
	Phone     string    `gorm:"user_phone"`
	Mail      string    `gorm:"user_mail"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

//type Group struct {
//	ID        int    `gorm:"AUTO_INCREMENT"`
//	GroupName string `gorm:"group_name"`
//	CreatedAt time.Time `gorm:"created_at"`
//	UpdatedAt time.Time `gorm:"updated_at"`
//}
type Message struct {
	ID          int       `gorm:"AUTO_INCREMENT"`
	Mess        string    `gorm:"mess"`
	Group       string    `gorm:"group"`
	Success     bool      `gorm:"success"`
	SuccessUser int       `gorm:"success_user"`
	FailUser    int       `gorm:"fail_user"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

type MessJsonBody struct {
	User User
	Typ     string
	Message Message
}
