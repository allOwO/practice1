package PracticeItem

type SendMailService interface {
	ResendMail(string, string, string) error
	SendMail(string,string,string) error
}
