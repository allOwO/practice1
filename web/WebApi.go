package web

import (
	"PracticeItem/Globalvar"
	"PracticeItem/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

func CreateUser(ctx echo.Context) error {
	//body, _ := ioutil.ReadAll(ctx.Request().Body)
	//log.Printf("body:%s\n", string(body))

	tmp:= new(Globalvar.WebUserMess)
	err := ctx.Bind(tmp)
	log.Println("new user",tmp)
	if err != nil {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "输入错误",
			Data:    nil,
		})
	}
	if len(tmp.Groups) == 0 {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "未选择小组",
			Data:    nil,
		})
	}
	if tmp.UserName == "" {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "请写姓名",
			Data:    nil,
		})
	}
	if tmp.UserMail == "" {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "请写邮箱",
			Data:    nil,
		})
	}
	if EmailFormat(tmp.UserMail) == false {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "请输入正确邮箱地址",
			Data:    nil,
		})
	}
	if tmp.UserPhone == "" || MobileFormat(tmp.UserPhone) == false {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "手机号码出错",
			Data:    nil,
		})
	}
	//debug
	user := &Globalvar.User{
		Name:      tmp.UserName,
		Phone:     tmp.UserPhone,
		Mail:      tmp.UserMail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tmp.UserMail=strings.ToLower(tmp.UserMail)
	if ok := model.CreateUser(tmp.Groups, user); ok==false {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "数据库错误或已存在",
			Data:    nil,
		})
	}
	return ctx.JSON(200, SendJson{
		Code:    stSucc,
		Message: "",
		Data:    nil,
	})
}
//post /changeuser
func ChangeUser(ctx echo.Context) error {
	body, _ := ioutil.ReadAll(ctx.Request().Body)
	log.Printf("body:%s\n", string(body))
	//tmp := new(WebUserMess)
	//err := ctx.Bind(tmp)
	tmp:= new(Globalvar.WebUserMess)

	err:=json.Unmarshal(body,&tmp)
	log.Println(tmp)
	if err != nil {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "输入错误",
			Data:    nil,
		})
	}
	if len(tmp.Groups) == 0 {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "未选择小组",
			Data:    nil,
		})
	}
	if tmp.UserName == "" {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "请写姓名",
			Data:    nil,
		})
	}
	if tmp.UserMail == "" {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "请写邮箱",
			Data:    nil,
		})
	}
	if EmailFormat(tmp.UserMail) == false {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "请输入正确邮箱地址",
			Data:    nil,
		})
	}
	if tmp.UserPhone == "" || MobileFormat(tmp.UserPhone) == false {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "手机号码出错",
			Data:    nil,
		})
	}
	//debug
	user := &Globalvar.User{
		Name:      tmp.UserName,
		Phone:     tmp.UserPhone,
		Mail:      tmp.UserMail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tmp.UserMail=strings.ToLower(tmp.UserMail)
	if ok := model.UpdateUser(tmp.Groups, user); ok==false {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "数据库错误或已存在",
			Data:    nil,
		})
	}
	return ctx.JSON(200, SendJson{
		Code:    stSucc,
		Message: "",
		Data:    nil,
	})
}
//get user
func CheckUser(ctx echo.Context) error {
	tmp:=ctx.QueryParam("mail")
	log.Println(tmp)
	usr := model.GetUserInfo(tmp)
	log.Println(usr)
	//log.Println(model.GetUsertest())
	if usr == nil {
		return ctx.JSON(200, SendJson{
			Code:    stFail,
			Message: "没有此用户",
			Data:    nil,
		})
	}
	return ctx.JSON(200, SendJson{
		Code:    stSucc,
		Message: "",
		Data:    usr,
	})
}

const (
	stSucc int = 200 //正常
	stFail int = 300 //失败
)

type SendJson struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    *Globalvar.WebUserMess `json:"data"`
}

//email verify
func EmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//mobile verify
func MobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
