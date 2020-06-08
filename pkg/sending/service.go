package sending

type Service interface{
	SendPush(...Notification) error
}

type Repository interface{
	SendPush(Notification) error
}

type service struct{
	r Repository
}

func NewService() Service {
	return &service{}
}

func (s *service) SendPush(n ...Notification) error{
	return nil
}

