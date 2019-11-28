package PracticeItem


type WebUserMess struct {
	UserName  string   `json:"user_name"`
	UserPhone string   `json:"user_phone"`
	UserMail  string   `json:"user_mail" `
	Groups    []string `json:"groups"`
}

type WebUserService interface {
	GetUserInfo(string) *WebUserMess
}


func NewWebUserMess() *WebUserMess {
	wu := new(WebUserMess)
	wu.Groups = make([]string, 0, 3)
	return wu
}