package service

type sendOfficialAccountService struct {
	wechatid    string
	total       string
	messagebody string
}

func NewSendOfficialAccountService()*sendOfficialAccountService{
	return &sendOfficialAccountService{}
}
