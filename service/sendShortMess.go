package service

type sendShortMessService struct {
	phone string
	total string
	messagebody string
}
func NewsendShortMessService()*sendShortMessService{
	return &sendShortMessService{}
}